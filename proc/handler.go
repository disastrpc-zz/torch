package proc

import (
	"os/exec"
	"torch/utils"
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

	var args []string = utils.Unpack(conf)

	cmd = &exec.Cmd{
		Path: conf.JavPath,
		Args: args,
		Dir:  conf.WorkDir,
	}

	return cmd
}

// Hook starts JVM instance
func Hook(conf *utils.Config) {

	cmd := build(conf)
	Listen(cmd, conf.Interval)
	err := cmd.Start()
	execErr(err)
	cmd.Wait()
}
