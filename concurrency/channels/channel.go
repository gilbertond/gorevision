package channels

func SomeFunc(n int, channel chan int) {
	for i := 1; i <= n; i++ {
		channel <- i // push value to channel
	}

	close(channel) // Only necessary when a receiver needs to know if there are no more values coming
}
