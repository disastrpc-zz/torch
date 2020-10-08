// Formula used to calculate restarts
// start time = UNIX time int representing time the timer started
// interval = user supplied value which represents time between restarts in seconds
// warn count = How many restart warnings should be sent out
// warninterval = interval / warn count
// ex: for an interval of 7200 and a count of 3 the warn interval would be 2400 seconds
// time to restart = start time + interval
// time to warn = (warninterval * warn count) + current UNIX time

// WIP !!

package proc

import (
	"sort"
	"time"
)

// InitTicker Starts timer for process on stdin
func InitTicker(interval int, warnCount int, Flag, Time chan int) {

	var C int
	var S int = int(time.Now().Unix())
	var I int = interval / warnCount
	var warnArr []int = getWarnTimes(S, I, warnCount)
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	go func() {
		for {
			<-ticker.C
			C = int(time.Now().Unix())
			q := sort.Search(len(warnArr), func(q int) bool { return warnArr[q] >= C })

			if warnCount == 0 {
				time.Sleep(3)
				Flag <- 2
				Time <- 0
				break
			} else if q < len(warnArr) && warnArr[q] == C {
				warnCount--
				Flag <- 1
				Time <- (I * warnCount)
			}
		}
	}()
}

func getWarnTimes(S, I, C int) []int {

	var arr []int

	for i := 1; i <= C; i++ {
		arr = append(arr, (S + (I * i)))
	}

	// List needs to be sorted for binary search to work.
	sort.Ints(arr)
	return arr
}
