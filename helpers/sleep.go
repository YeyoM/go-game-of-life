package helpers

import "time"

func Sleep(miliseconds int) {
	duration := time.Duration(miliseconds) * time.Millisecond
	time.Sleep(duration)
}
