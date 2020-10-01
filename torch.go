package main

import (
	"fmt"
	"os"
	"torch/utils"

	"github.com/akamensky/argparse"
)

func main() {

	parser := argparse.NewParser("Torch", "Wraps around JVM invocation to provide useful administrative functions")
	var javpath *string = parser.String("j", "java-path",
		&argparse.Options{
			Required: false,
			Help:     "Java arguments to execute",
			Default:  "java",
		})
	var jarfile *string = parser.String("J", "server-jar",
		&argparse.Options{
			Required: false,
			Help:     "Provide path to server jar file",
			Default:  nil,
		})
	var jvmargs *[]string = parser.List("a", "jvm-args",
		&argparse.Options{
			Required: false,
			Help:     "Provide JVM arguments that server-jar will execute",
			Default:  nil,
		})
	var interval *int = parser.Int("i", "interval",
		&argparse.Options{
			Required: false,
			Help:     "Interval server will restart. In seconds",
			Default:  0,
		})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	conf := utils.ParseConf(javpath, jarfile, jvmargs, interval)

	fmt.Printf("%v", conf.Interval)
	// proc.Hook(conf.JavPath, conf.JarFile, conf.WorkDir, conf.JvmArgs)
}