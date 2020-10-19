//                 		       GNU GENERAL PUBLIC LICENSE
// 					Version 3, 29 June 2007
// 			Torch - A simple server restart utility written in Go
// 				by disastrpc @ github.com/disastrpc

package main

import (
	"fmt"
	"os"
	"torch/proc"
	"torch/utils"

	"github.com/akamensky/argparse"
)

func main() {

	parser := argparse.NewParser("Torch", "Wraps around JVM invocation to provide useful administrative functions")
	var javpath *string = parser.String("j", "java-path",
		&argparse.Options{
			Required: false,
			Help:     "Java arguments to execute",
			Default:  "",
		})
	var jarfile *string = parser.String("J", "server-jar",
		&argparse.Options{
			Required: false,
			Help:     "Provide path to server jar file",
			Default:  "",
		})
	var jvmargs *[]string = parser.List("a", "jvm-args",
		&argparse.Options{
			Required: false,
			Help:     "Provide JVM arguments that server-jar will execute",
			Default:  []string{},
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

	conf := utils.ParseConfig(javpath,
		jarfile,
		jvmargs,
		interval)

	proc.Hook(conf)

}
