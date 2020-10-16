package proc

import (
	"fmt"
	"sort"
	"time"
)

// InitTicker Starts timer for process on stdin
func InitTicker(interval int, warnCount int, Stat, Rem chan<- int) {

	var C int
	var S int = int(time.Now().Unix())
	var I int = interval / warnCount
	var warnArr []int = getWarnTimes(S, I, warnCount)
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	go func(Stat, Rem chan<- int) {
		for {
			<-ticker.C
			C = int(time.Now().Unix())
			q := sort.Search(len(warnArr), func(q int) bool { return warnArr[q] >= C })

			if warnCount == 0 {
				Stat <- 2
				Rem <- 0
				break
			} else if q < len(warnArr) && warnArr[q] == C {
				warnCount--
				fmt.Printf("%v\n", warnArr)
				Rem <- (I * warnCount)
				Stat <- 1
			}
		}

		ticker.Reset(1 * time.Second)
	}(Stat, Rem)
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
