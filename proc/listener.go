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
func Listen(cmd *exec.Cmd,
	interval int,
	msgInterval int,
	Con chan<- string,
	Stat chan<- int,
	reader io.Reader) {

	Time := make(chan int, msgInterval)
	Flag := make(chan int, msgInterval)

	stdout, err := cmd.StdoutPipe()

	go func(reader io.Reader,
		Flag, Time chan int,
		Con chan<- string) {

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {

			if strings.Contains(scanner.Text(), "For help, type") {
				msg := utils.FormatLog("Started Torch ticker")
				fmt.Print(msg)
				InitTicker(interval, msgInterval, Flag, Time)
			}

			Con <- scanner.Text()
		}
	}(reader, Flag, Time, Con)

	go func(interval, msgInterval int,
		Flag, Time chan int,
		Stat chan<- int) {

		for {
			select {
			case F := <-Flag:
				if F == 1 {
					T := <-Time
					Stat <- 1
				} else if F == 2 {
					Stat <- 2
					close(Flag)
					close(Time)
				}
			default:
				continue
			}
			time.Sleep(1 * time.Second)
		}
	}(interval, msgInterval, Flag, Time, Stat)
}
