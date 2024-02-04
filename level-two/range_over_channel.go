package leveltwo

import (
	"fmt"
	"sync"
)

func RangeOverChannel() {
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		// Close the channel when done sending data
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Using a range loop to receive values until the channel is closed
		for value := range ch {
			fmt.Println("Received:", value)
		}
	}()

	wg.Wait()
}
