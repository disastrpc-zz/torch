// Formula used to calculate restarts
// start time = UNIX time int representing time the timer started
// interval = user supplied value which represents time between restarts in seconds
// warn count = How many restart warnings should be sent out
// warninterval = interval / warn count
// ex: for an interval of 7200 and a count of 3 the warn interval would be 2400 seconds
// time to restart = start time + interval
// time to warn = (warninterval * warn count) + current timepackage utils

// WIP !!

package utils

import (
	"fmt"
	"io"
	"time"
)

// InitTicker Starts timer for process on stdin
func InitTicker(interval int64, warnCount int64, stdin io.WriteCloser) {

	//var stop int64 = time.Now().Unix() + interval
	var ticker *time.Ticker = time.NewTicker(5 * time.Second)
	var Flag chan int = make(chan int)
	var Time chan int = make(chan int)
	var p int = 1

	//start := time.Now().Unix()

	go func() {
		fmt.Println("STARTED TICKER")
		for warnCount != 0 {
			<-ticker.C
			var warnInterval int64 = interval / warnCount
			switch {
			case int64(p) == warnCount:
				fmt.Println("HERE1")
				ticker.Stop()
				Flag <- 1
				close(Flag)
			case (warnInterval*int64(p))+time.Now().Unix() >= time.Now().Unix():
				fmt.Println("HERE2")
				message := fmt.Sprintf("say [Torch] server is restarting in %d minutes\n", (warnInterval*int64(p))/60)
				stdin.Write([]byte(message))
				Flag <- 2
				close(Flag)
				Time <- (int(warnInterval) * p) / 60
				close(Time)
				p++
			}

			fmt.Println(time.Now().Unix())
		}
	}()
}
