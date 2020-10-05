package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// Config represents the configuration settings for Torch.
type Config struct {
	JavPath   string
	JarFile   string
	JvmArgs   []string
	Interval  int64
	WarnCount int64
	WarnMsg   string
	RebootMsg string
}

//ParseArgs reads config file and returns reference to Config structure. Parameters are used to overwrite file values with command line args.
func ParseConfig(jp, jf *string, ja *[]string, i *int) *Config {

	data, err := ioutil.ReadFile("torch.conf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open config file. %v", err)
	}

	var conf Config

	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	return &conf
}

// Unpack takes Config struct and uses reflection to return a
// slice representing the complete JVM argument list.
func Unpack(c *Config) (args []string) {
	v := reflect.ValueOf(*c)

	args = append(args, c.JavPath)

	for i := 0; i < v.NumField(); i++ {
		s := v.Field(i).String()

		if i == 1 {
			args = append(args, "\x2d\x6a\x61\x72")
			args = append(args, s)

		} else {
			switch {
			case v.Type().Field(i).Name == "JarFile":
				args = append(args, s)
			case v.Type().Field(i).Name == "JvmArgs":
				for a := 0; a < len(c.JvmArgs); a++ {
					args = append(args, c.JvmArgs[a])
				}
			}
		}
	}
	return args
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

func splitPref(s string) (slice []string) {
	return strings.Split(s, " = ")
}
