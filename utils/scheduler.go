package utils

import (
	"time"
)

func Scheduler(timeout time.Duration, action func()) {
	ticker := time.NewTicker(timeout)

	
	go func() {
		defer ticker.Stop()
		
		for range ticker.C {
			action()
		}
	}()
}