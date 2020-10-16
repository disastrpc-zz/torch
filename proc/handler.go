package proc

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
	"torch/utils"

	"github.com/rivo/tview"
)

const refreshInterval = 1 * time.Second

type procHook struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	reader io.Reader
	conf   utils.Config
}

var (
	view *tview.Modal
	proc *tview.Application
)

func execErr(err error) {
	if err != nil {
		panic(err)
	}
}

func build(conf *utils.Config) (cmd *exec.Cmd) {

	if conf.JavPath == "java" {
		conf.JavPath, _ = exec.LookPath("java")
	}

	workDir, err := utils.ParseWorkingDir(conf.JarFile)
	args := utils.Unpack(conf)
	if err != nil {
		panic(err)
	}

	cmd = &exec.Cmd{
		Path: conf.JavPath,
		Args: args,
		Dir:  workDir,
	}

	return cmd
}

// Executes command sent to sever
func execute(s string, stdin io.WriteCloser) {
	stdin.Write([]byte(s))
}

func recSout(Sout chan<- string,
	Stat chan int,
	Rem chan int,
	hook *procHook) {

	scanner := bufio.NewScanner(hook.reader)
	for scanner.Scan() {

		// TODO - Move detect strings to own func
		if strings.Contains(scanner.Text(), "For help, type") {
			msg := utils.FormatLog("Started Torch ticker")
			fmt.Print(msg)
			InitTicker(hook.conf.Interval, hook.conf.WarnCount, Stat, Rem)
			Stat <- 0
		}
		Sout <- scanner.Text()
	}
}

// Listen takes pointer to process struct and listens on standard pipes, and on channels comming from ticker
func Listen(hook *procHook,
	Sout chan string,
	Stat chan int,
	Rem chan int) {

	//Brk := make(chan bool)

	go recSout(Sout, Stat, Rem, hook)

	go func(Stat, Rem chan int, stdin io.WriteCloser) {
		for {
			select {
			case S := <-Stat:
				if S == 1 {
					R := <-Rem
					msg := utils.FormatWarn(R)
					hook.stdin.Write(msg)
				} else if S == 2 {
					hook.stdin.Write([]byte("\x73\x74\x6f\x70\n"))
					//Brk <- true
					break
				}
			case out := <-Sout:
				fmt.Println(out)
			}

		}
	}(Stat, Rem, hook.stdin)

	// for s := range Sout {
	// 	fmt.Println(s)
	// 	select {
	// 	case <-Brk:
	// 		return
	// 	default:
	// 		continue
	// 	}
	// }
	// Brk <- false
}

// Hook starts JVM instance and listener/ticker
func Hook(conf *utils.Config) {

	var hook procHook
	var err error

	Sout := make(chan string, 10)
	Stat := make(chan int, 2)
	Rem := make(chan int, 2)

	hook.conf = *conf

	hook.cmd = build(&hook.conf)

	hook.stdout, err = hook.cmd.StdoutPipe()
	execErr(err)

	hook.stdin, err = hook.cmd.StdinPipe()
	execErr(err)

	hook.reader = bufio.NewReader(hook.stdout)
	err = hook.cmd.Start()
	execErr(err)

	Listen(&hook, Sout, Stat, Rem)
	hook.cmd.Wait()
	time.Sleep(5 * time.Second)
	Stat <- 0

}
