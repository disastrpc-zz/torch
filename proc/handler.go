package proc

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"torch/tui"
	"torch/utils"

	"github.com/gdamore/tcell/v2"
)

// Represents a hooked process including its Cmd struct, UI elements, pipes, configuration settings and reader object
type procHook struct {
	cmd    *exec.Cmd
	view   *tui.Tui
	stdin  io.WriteCloser
	stdout io.ReadCloser
	reader io.Reader
	conf   utils.Config
}

// Executes command sent to sever
func execute(s string, stdin io.WriteCloser) {
	stdin.Write([]byte(s))
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

// Set all hook properties
func initHook(hook *procHook, conf *utils.Config, Stat chan int) {
	var err error

	// Setup hook struct values

	hook.cmd = build(&hook.conf)

	// Setup pipes and reader, start server and listen on channels
	hook.stdout, err = hook.cmd.StdoutPipe()
	execErr(err)
	hook.stdin, err = hook.cmd.StdinPipe()
	execErr(err)
	hook.reader = bufio.NewReader(hook.stdout)
}

// Hook starts JVM instance and listener/ticker
func Hook(conf *utils.Config) {

	var hook procHook
	var err error

	Sout := make(chan string, 10)
	Stat := make(chan int, 2)
	Rem := make(chan int, 2)
	stop := make(chan bool, 2)

	hook.view = tui.Init()

	// Capture input from InputField and write to stdin.
	// If command executed is stop program will enter shutdown routine.
	hook.view.In.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			cmd := hook.view.In.GetText()

			if cmd == "stop" {
				hook.stdin.Write([]byte(cmd + "\n"))
				hook.view.In.SetText("")
				fmt.Fprintf(hook.view.Tv, "%v", string(utils.FormatLog("Stopping Torch")))
				stop <- true
			}

			hook.stdin.Write([]byte(cmd + "\n"))
			hook.view.In.SetText("")

		}
	})

	// Set Flex as App root
	go hook.view.App.SetRoot(hook.view.Flx, true).Run()
	hook.conf = *conf

	if hook.conf.ShowBanner {
		fmt.Fprintf(hook.view.Tv, "%v\n", utils.Banner(&hook.conf))
	}

	for {
		initHook(&hook, conf, Stat)
		select {
		case s := <-stop:
			if s {
				hook.cmd.Wait()
				hook.view.App.Stop()
				return
			}
		default:
			err = hook.cmd.Start()
			execErr(err)

			hook.Listen(Sout, Stat, Rem)

			msg := utils.FormatLog("Starting JVM")

			fmt.Fprintf(hook.view.Tv, "%v", msg)

			hook.cmd.Wait()
		}
	}

}
