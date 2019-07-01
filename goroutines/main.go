package main

import "fmt"

//goroutine is a lightweight thread managed by go
func main() {
	someFunction("non goroutine is synchronized")

	go someFunction("goroutine which is asynchronized")

	go func(data string) {
		fmt.Println(data)
	}("goroutine asyn anonymous function")

	fmt.Scanln() //press a key before program exits
	fmt.Println("done")
}

func someFunction(data string) {
	for i := 0; i < 3; i++ {
		fmt.Println(data, ":", i)
	}
}
