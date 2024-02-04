package leveltwo

import "fmt"

func ChannelBuffer() {
	// Creating a buffered channel with a capacity of 3
	bufferedChannel := make(chan int, 3)

	// Sending values to the buffered channel
	bufferedChannel <- 1
	bufferedChannel <- 2
	bufferedChannel <- 3

	// Uncommenting the line below would result in a runtime error
	// bufferedChannel <- 4 // channel is full

	// Receiving values from the buffered channel
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	// Uncommenting the line below would result in a runtime error
	// fmt.Println(<-bufferedChannel) // channel is empty
}
