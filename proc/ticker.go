package proc

import (
	"time"
)

// InitTicker Starts timer for process on stdin
// Ticker will communicate every (interval / warnCount) seconds, the signals sent are as follows:
// 1 will be sent if a warning needs to be sent without shutting down server
// 2 will be sent if the warn count is zero and the server needs to restart
func InitTicker(interval int, warnCount int, Stat, Rem chan<- int) {

	var ticker *time.Ticker = time.NewTicker(time.Duration(interval/warnCount) * time.Second)
	var I int = interval / warnCount

	for {
		select {
		case <-ticker.C:
			warnCount--
			if warnCount == 0 {
				Stat <- 2
				ticker.Stop()
				return
			}

			Stat <- 1
			Rem <- (I * warnCount)
		}
	}
}
