package leveltwo

import (
	"fmt"
	"sync"
)

func CloseChannel() {
	// Creating a channel
	ch := make(chan int)

	// Use a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Goroutine to send data to the channel
	wg.Add(1)
	go func() {
		close(ch)
		defer wg.Done()
		for i := 1; i <= 5; i++ {

			//Check if the channel is closed
			_, ok := <-ch
			if ok {
				ch <- i
			} else {
				fmt.Println("Channel is closed")
			}
		}
		// Close the channel when done sending data

	}()

	// Wait for goroutines to finish
	wg.Wait()

}
