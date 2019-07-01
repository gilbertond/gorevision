package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.increment("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.getValue("somekey"))
}

//******************************************************************************
//sync.Mutex for synchronous
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

//increment the current value of the given key
func (c *SafeCounter) increment(key string) {
	c.mux.Lock()   //lock so only one goroutine at a time accesses
	c.v[key]++     //modify map as you like
	c.mux.Unlock() //unlock after
}

//return current value of the counter for the given key
func (c *SafeCounter) getValue(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock() //defer makes sure the mutex is unlocked so that one goroutine can access map value
	return c.v[key]
}
