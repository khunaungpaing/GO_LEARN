package leveltwo

import (
	"fmt"
	"time"
)

func SelectInGo() {
	// Creating two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1: Sends a message to ch1 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from Goroutine 1"
	}()

	// Goroutine 2: Sends a message to ch2 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from Goroutine 2"
	}()

	// Using select to wait for either of the channels to receive a message
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	}

	// Additional work can be done after receiving a message
	fmt.Println("Select function exits.")
}
