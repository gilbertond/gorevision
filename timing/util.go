package main

import (
	"fmt"
	"time"
)

func timeTrack(start time.Time, operationName string) {
	elapsed:=time.Since(start)
	ms := float64(elapsed) / float64(time.Millisecond)
	fmt.Printf("took %s time %f secs to run\n", operationName, ms)
}
