package main

import (
	"fmt"
)

//channels are pipes that connect concurrent goroutines
//you can SEND values into channels from one goroutine and RECEIVE those values into another goroutine
func main() {
	channel := make(chan int, 5)       //Create a channel with buffer size 5
	go someFunc(cap(channel), channel) //cap returns capacity of type. chan, array, etc

	for i := range channel {
		fmt.Println(i)
	}
}

func someFunc(n int, channel chan int) {
	for i := range channel {
		channel <- i //push value to channel
	}

	close(channel) //Only necessary when a receiver needs to know if there are no more values coming
}
