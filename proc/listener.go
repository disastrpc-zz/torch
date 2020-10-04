package proc

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"torch/utils"
)

// Listen takes pointer to process struct and listens on standard pipes
func Listen(cmd *exec.Cmd, interval int64) {
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		panic(err)
	}

	stdin, err := cmd.StdinPipe()

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(stdout)

	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {

			if strings.Contains(scanner.Text(), "For help, type") {
				stdin.Write([]byte("say stdin test\n"))
				utils.InitTimer(interval, stdin)
			}
			fmt.Printf("\x5b\x54\x6f\x72\x63\x68\x2f\x57\x72\x61\x70\x70\x65\x72\x5d\x20%s\n", scanner.Text())
		}
	}(reader)
}
