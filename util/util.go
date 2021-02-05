package util

import (
	"fmt"
	"io"
	"time"
)

func LogFor(w io.Writer, t time.Duration, prefix string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(t)
		done <- true
	}()
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Fprintf(w, "%s: message at ticker %s\n", prefix, t)
		}
	}

}
