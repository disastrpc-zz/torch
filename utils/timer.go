package utils

import (
	"fmt"
	"io"
	"time"
)

// getDuration returns a time.Duration type representing the amount of seconds in i
func getDuration(i int64) time.Duration {
	var timeInterval = time.Duration(i)
	return timeInterval
}

// InitTimer Starts timer for process on stdin
func InitTimer(interval int64, stdin io.WriteCloser) *time.Timer {
	timer := time.NewTimer(time.Second * getDuration(interval))

	go func() {
		fmt.Printf("Started %v sec counter\n", interval)
		<-timer.C
		print("timer stopped\n")
		stdin.Write([]byte("stop\n"))
	}()

	return timer
}
