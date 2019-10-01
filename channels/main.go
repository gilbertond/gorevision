package channels    //change to main if u wanna execute

import (
	"fmt"
)

//channels are pipes that connect concurrent goroutines
//you can SEND values into channels from one goroutine and RECEIVE those values into another goroutine
func main() {
	channel := make(chan int, 5)       //Create a channel with buffer size 5
	go SomeFunc(cap(channel), channel) //cap returns capacity of type. chan, array, etc

	for i := range channel {
		fmt.Println(i)
	}
}
