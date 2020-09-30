package utils

import (
	"io/ioutil"
)

func checkFileErr(err error) {

	if err != nil {
		panic(err)
	}
}

// ReadArgsFile reads a file containing the JVM arguments to start server. 
func ReadArgsFile(path string) (jvmArgs []rune) {

	data, err := ioutil.ReadFile(path)
	checkFileErr(err)

	println([]rune(string(data)))
	return jvmArgs
}