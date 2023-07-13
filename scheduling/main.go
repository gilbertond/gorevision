package main

import (
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	// timeAfterBlocking()
	// runInIntervals()
	cronRobfig()
}

// blocks execution for 10 secs
func timeAfterBlocking() {
	timeBefore := time.Now()
	timeAfter := <-time.After(time.Second * 10)
	fmt.Printf("\n%v\n", timeBefore)
	fmt.Printf("%v\n", timeAfter)
}

// run in intervals.
// prob: runs forever and no way to stop since no control of underlying time.Ticker
func runInIntervals() {
	for tick := range time.Tick(time.Second * 2) { // note: 1st print occurs after the delay
		fmt.Printf("\nAfter 2 seconds: %v", tick.Format("2006-0-02 15:04:05"))
	}
}

// https://stephenafamo.com/blog/better-scheduling-in-go/
// timer with a controlled start time and end
func customTimer(ctx context.Context, start time.Time, delay time.Duration) <-chan time.Time {
	// chann := make(chan time.Time, 1)

	// 	validate start time. if start == time.Time{}
	// 	if !start.IsZero() {
	// 		if time.Until(start) < 0 {   // equivalent to t.Sub(time.Now())
	//
	// 		}
	// 	}

	return nil
}

// cron lib

// https://pkg.go.dev/github.com/robfig/cron
func cronRobfig() {
	// Note: funcs are executed in their own routines
	cr := cron.New()
	cr.AddFunc("*/1 * * * *", func() {
		fmt.Printf("every minute")
	})
	cr.AddFunc("TZ=America/Denver 43 12 * * * *", func() {
		fmt.Printf("run daily at 12:40 denver time")
	})
	cr.AddFunc("@every 0h0m1s", func() {
		fmt.Printf("run every second")
	})
	cr.Start()

	time.Sleep(10 * time.Second)

	cr.Stop() // does not stop already running jobs
}
