package proc

import (
	"os"
	"os/exec"
)

var javpath string

func execErr(err error) {
	if err != nil {
		panic(err)
	}
}

func build(javpath string, jarfile string, workdir string, args []string) (cmd *exec.Cmd) {

	if javpath == "java" {
		javpath, _ = exec.LookPath("java")
	}

	s := []string{javpath + "-jar" + jarfile}

	var jvmargs []string = append(s, args...)

	cmd = &exec.Cmd{
		Path:   javpath,
		Args:   jvmargs,
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
		Stderr: os.Stderr,
		Dir:    workdir,
	}

	return cmd
}

func inject(args ...[]string) {

}

// Hook starts JVM instance and reads stdout
func Hook(javpath string, jarfile string, workdir string, args []string) {
	cmd := build(javpath, jarfile, workdir, args)
	cmd.Run()
}
