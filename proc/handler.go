package proc

import (
	"os"
	"os/exec"
	"torch/utils"
)

var javpath string

// Config Redifine config

func execErr(err error) {
	if err != nil {
		panic(err)
	}
}

func build(conf *utils.Config) (cmd *exec.Cmd) {

	if javpath == "java" {
		javpath, _ = exec.LookPath("java")
	}

	var args []string = utils.Unpack(conf)

	cmd = &exec.Cmd{
		Path:   conf.JavPath,
		Args:   args,
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
		Stderr: os.Stderr,
		Dir:    conf.WorkDir,
	}

	return cmd
}

func inject(args ...[]string) {

}

// Hook starts JVM instance and reads stdout
func Hook(conf *utils.Config) {

	cmd := build(conf)
	cmd.Run()
}
