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

	for {
		cmd := build(conf)
		Listen(cmd, conf.Interval, conf.WarnCount)
		err := cmd.Start()
		execErr(err)
		cmd.Wait()
	}
}
