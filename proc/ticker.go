package proc

import (
	"time"
)

// InitTicker Starts timer for process on stdin
func InitTicker(interval int, warnCount int, Stat, Rem chan<- int) {

	var ticker *time.Ticker = time.NewTicker(time.Duration(interval/warnCount) * time.Second)
	var I int = interval / warnCount
	var C int = warnCount

	for {
		select {
		case <-ticker.C:
			C--
			if C == 0 {
				Stat <- 2
				ticker.Stop()
				return
			}

			Stat <- 1
			Rem <- (I * C)
		}
	}

}
