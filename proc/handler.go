package proc

import (
	"bufio"
	"fmt"
	"os/exec"
	"time"
	"torch/utils"

	"github.com/rivo/tview"
)

const refreshInterval = 1 * time.Second

// Status Represents server restart status
type Status struct {
	T    int
	Flag int
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

// build conf struct by unmarshaling json config
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

// Hook starts JVM instance
func Hook(conf *utils.Config) {

	Con := make(chan string, 10)
	Stat := make(chan int, conf.WarnCount)

	for {

		cmd := build(conf)

		err := cmd.Start()
		execErr(err)

		stdout, err := cmd.StdoutPipe()
		execErr(err)

		stdin, err := cmd.StdinPipe()
		execErr(err)

		reader := bufio.NewReader(stdout)

		Listen(cmd, conf.Interval, conf.WarnCount, Con, Stat, reader)

		for data := range Con {
			fmt.Println(data)
			select {
			case S := <-Stat:
				println(S)
				if S == 1 {
					msg := utils.FormatWarn(T)
					stdin.Write(msg)
					println("here1", S)
				} else if S == 2 {
					stdin.Write([]byte("\x73\x74\x6f\x70\n"))
				}
			default:
				continue
			}
		}
		cmd.Wait()
	}
}
