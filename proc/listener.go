package proc

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
	"torch/utils"
)

func execErr(err error) {
	if err != nil {
		panic(err)
	}
}

// recStout receives all output from server on stdout and pipes it through the Sout chan, this is received by the Listen function
// which handles the output
func (hook *procHook) recSout(Sout chan<- string,
	Stat chan int,
	Rem chan int) {

	scanner := bufio.NewScanner(hook.reader)
	for scanner.Scan() {

		// TODO - Move detect strings to own func
		if strings.Contains(scanner.Text(), "Preparing spawn area: 100%") {
			msg := utils.FormatLog("Started Torch ticker")
			fmt.Fprintf(hook.view.Tv, "%v", msg)
			go InitTicker(hook.conf.Interval, hook.conf.WarnCount, Stat, Rem)
		} else if strings.HasPrefix(scanner.Text(), "at") {
			continue
		}
		Sout <- scanner.Text()
	}
}

// Listen takes pointer to process struct and listens on standard pipes, and on channels comming from ticker
func (hook *procHook) Listen(Sout chan string,
	Stat chan int,
	Rem chan int) {

	go hook.recSout(Sout, Stat, Rem)

	go func(Stat, Rem chan int, stdin io.WriteCloser) {

		view := hook.view.Tv

		for {
			select {
			case S := <-Stat:
				if S == 1 {
					R := <-Rem
					msg := utils.FormatWarn(R)
					hook.stdin.Write(msg)
				} else if S == 2 {
					hook.stdin.Write(utils.FormatWarn(0))
					time.Sleep(5 * time.Second)
					hook.stdin.Write([]byte("\x73\x74\x6f\x70\n"))
					break
				}
			case out := <-Sout:
				fmt.Fprintf(view, "%v\n", out)
			}

		}
	}(Stat, Rem, hook.stdin)
}
