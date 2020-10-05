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
func Listen(cmd *exec.Cmd, interval int64, msgInterval int64) {
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
				utils.InitTicker(interval, msgInterval, stdin)
				listenChans(stdin)
			}

			fmt.Printf("%v %s\n", utils.Constamp, scanner.Text())

		}
	}(reader)
}

func listenChans(stdin io.WriteCloser) {

	var Time chan int = make(chan int)
	var Flag chan int = make(chan int)

	select {
	case Flag <- 2:
		fmt.Println("HERE IN FLAG2")
		message := fmt.Sprintf("say [Torch] server is restarting in %d minutes\n", <-Time)
		stdin.Write([]byte(message))
	case Flag <- 1:
		fmt.Println("HERE IN FLAG1")
		stdin.Write([]byte("say [Torch] server is restarting in 30 seconds\n"))
		time.Sleep(30 * time.Second)
		stdin.Write([]byte("stop\n"))
	}
}
