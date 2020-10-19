package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

// Config represents the configuration settings for Torch.
type Config struct {
	JavPath    string
	JvmArgs    []string
	JarFile    string
	Interval   int
	WarnCount  int
	WarnMsg    string
	RebootMsg  string
	ShowBanner bool
}

//ParseConfig reads config file and returns reference to Config structure. Parameters are used to overwrite file values with command line args.
func ParseConfig(javpath, jarfile *string, jvmargs *[]string, interval *int) *Config {

	var conf Config

	data, err := ioutil.ReadFile("torch.conf")

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open config file. %v", err)
	}

	// Ensure input data doesn't contain \ escapes before getting unmarshaled
	if bytes.Contains(data, []byte("\x5c")) {
		data = bytes.Replace(data, []byte("\x5c"), []byte("\x5c\x5c"), -1)
	}

	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}

	if *javpath != "" {
		conf.JavPath = *javpath
	}

	if *jarfile != "" {
		conf.JarFile = *jarfile
	}

	if len(*jvmargs) != 0 {
		conf.JvmArgs = *jvmargs
	}

	if *interval != 0 {
		conf.Interval = *interval
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

		switch {
		case v.Type().Field(i).Name == "JvmArgs":
			for a := 0; a < len(c.JvmArgs); a++ {
				args = append(args, c.JvmArgs[a])
			}
		case v.Type().Field(i).Name == "JarFile":
			args = append(args, ("\x2d\x6a\x61\x72"))
			args = append(args, s)
			args = append(args, "\x6e\x6f\x67\x75\x69")
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
