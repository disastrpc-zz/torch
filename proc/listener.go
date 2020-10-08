package proc

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
	"torch/utils"
)

// Listen takes pointer to process struct and listens on standard pipes, and on channels comming from ticker
func Listen(cmd *exec.Cmd, interval int, msgInterval int) {

	stdout, err := cmd.StdoutPipe()

	Time := make(chan int, msgInterval)
	Flag := make(chan int, msgInterval)

	if err != nil {
		panic(err)
	}

	stdin, err := cmd.StdinPipe()

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(stdout)

	go func(reader io.Reader, Flag, Time chan int) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {

			if strings.Contains(scanner.Text(), "For help, type") {
				msg := utils.FormatLog("Started Torch ticker")
				fmt.Print(msg)
				InitTicker(interval, msgInterval, Flag, Time)
			}

			fmt.Println(scanner.Text())
		}
	}(reader, Flag, Time)

	go func(interval, msgInterval int, Flag, Time chan int, stdin io.WriteCloser) {
		for {
			select {
			case F := <-Flag:
				if F == 1 {
					T := <-Time
					msg := utils.FormatWarn(T)
					stdin.Write(msg)
				} else if F == 2 {
					stdin.Write([]byte("\x73\x74\x6f\x70\n"))
					close(Flag)
					close(Time)
				}
			default:
				continue
			}
			time.Sleep(1 * time.Second)
		}
	}(interval, msgInterval, Flag, Time, stdin)
}
