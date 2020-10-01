package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Config represents the configuration settings for Torch.
type Config struct {
	JavPath  string
	JarFile  string
	JvmArgs  []string
	WorkDir  string
	Interval int64
}

//ParseConf reads config file and returns reference to Config structure. Parameters are used to overwrite file values with command line args.
func ParseConf(jp, jf *string, ja *[]string, i *int) *Config {

	f, err := os.Open("torch.conf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open config file. %v", err)
	}
	defer f.Close()

	conf := new(Config)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := strings.Trim(scanner.Text(), "\x0a\x0b\x09\x20\x0d")
		switch {
		case strings.HasPrefix(t, "\x2f\x2f"):
			// Ignores comment lines
		case strings.HasPrefix(t, "java_path"):
			if *jp != "java" {
				conf.JavPath = *jp
			} else if a := splitPref(t)[1]; a == "java" {
				conf.JavPath = "java"
			} else {
				conf.JavPath = splitPref(t)[1]
			}
		case strings.HasPrefix(t, "server_jar"):
			if jf != nil {
				conf.JarFile = *jf
			} else {
				conf.JarFile = splitPref(t)[1]
			}
		case strings.HasPrefix(t, "java_args"):
			if ja != nil {
				conf.JvmArgs = *ja
			} else {
				conf.JvmArgs = splitPref(t)
			}
		case strings.HasPrefix(t, "reboot_interval"):
			if *i != 0 {
				conf.Interval = int64(*i)
			} else {
				conf.Interval, err = strconv.ParseInt(splitPref(t)[1], 10, 32)
			}
		}
	}

	fmt.Printf("%v", conf.JarFile)
	conf.WorkDir, err = ParseWorkingDir(conf.JarFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	return conf
}

func splitPref(s string) (slice []string) {
	return strings.Split(s, " = ")
}

// ParseWorkingDir returns the working directory of the server jar file.
func ParseWorkingDir(path string) (string, error) {

	var err error
	var p string = filepath.Dir(path)
	if p == "." {
		err = errors.New("error: unable to parse working directory")
	} else {
		err = nil
	}

	return p, err
}

// ReadArgsFile reads a file containing the JVM arguments to start server. Returns byte array containing JVM command.
func ReadArgsFile(path string) (jvmArgs []byte) {

	jvmArgs, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return jvmArgs
}
