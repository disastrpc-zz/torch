package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
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

//ParseArgs reads config file and returns reference to Config structure. Parameters are used to overwrite file values with command line args.
func ParseArgs(jp, jf *string, ja *[]string, i *int) *Config {

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
			if len(*jf) > 0 {
				conf.JarFile = *jf
			} else {
				conf.JarFile = splitPref(t)[1]
			}

		case strings.HasPrefix(t, "java_args"):
			if len(*ja) > 0 {
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

	conf.WorkDir, err = parseWorkingDir(conf.JarFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	return conf
}

// Unpack takes Config struct and uses reflection to return a
// slice representing the complete JVM argument list.
func Unpack(c *Config) (args []string) {
	v := reflect.ValueOf(*c)
	for i := 0; i < v.NumField(); i++ {
		if i == 1 {
			args = append(args, "\x2d\x6a\x61\x72")
		} else {
			f := v.Field(i)
			s := f.String()
			args = append(args, s)
		}
	}
	return args
}

// ParseWorkingDir returns the working directory of the server jar file.
func parseWorkingDir(path string) (string, error) {

	var err error
	var p string = filepath.Dir(path)
	if p == "." {
		err = errors.New("error: unable to parse working directory")
	} else {
		err = nil
	}

	return p, err
}

func splitPref(s string) (slice []string) {
	return strings.Split(s, " = ")
}
