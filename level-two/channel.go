package leveltwo

import (
	"fmt"
	"sync"
	"time"
)

func basicChannel() {
	fmt.Println("---- Basic Channel ----")
	ch := make(chan int)

	// Send and receive operations
	go func() {
		ch <- 42 // Send value 42 to the channel
	}()
	value := <-ch // Receive a value from the channel
	fmt.Println("Received from channel:", value)
}

func bufferedChannel() {
	fmt.Println("---- Buffered Channel ----")
	ch := make(chan int, 3) // Buffered channel with capacity 3

	// Buffered channel allows non-blocking sends until capacity is reached
	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 // Uncommenting this line would cause a deadlock since the buffer is full

	for i := 0; i < 3; i++ {
		value := <-ch
		fmt.Println("Received from buffered channel:", value)
	}
}

func closingChannel() {
	fmt.Println("---- Closing Channel ----")
	ch := make(chan int)

	go func() {
		defer close(ch) // Close the channel when done
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	// Range over the channel until it's closed
	for value := range ch {
		fmt.Println("Received from closed channel:", value)
	}
}

func selectStatement() {
	fmt.Println("---- Select Statement ----")
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 42
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 100
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	default:
		fmt.Println("Non-blocking case")
	}
}

func unidirectionalChannels() {
	fmt.Println("---- Unidirectional Channels ----")
	sendCh := make(chan<- int) // Send-only channel
	recvCh := make(<-chan int) // Receive-only channel

	go func() {
		sendCh <- 42
	}()

	go func() {
		value := <-recvCh
		fmt.Println("Received from receive-only channel:", value)
	}()
}

func fanOutFanIn() {
	fmt.Println("---- Fan-Out, Fan-In ----")

	input := make(chan int)
	output := make(chan int)

	// Fan-out: Multiple workers
	for i := 1; i <= 3; i++ {
		go func(workerID int) {
			defer func() {
				if workerID == 3 {
					close(output)
				}
			}()

			for value := range input {
				result := value * workerID
				output <- result
			}
		}(i)
	}

	// Fan-in: Single worker
	go func() {
		defer close(input)
		for i := 1; i <= 5; i++ {
			input <- i
		}
	}()

	// Wait for all workers to finish
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func(workerID int) {
			defer wg.Done()
			for range input {
				// Consume values from input channel
			}
		}(i)
	}

	// Wait for all consumers to finish
	wg.Wait()

	// Receive results from the output channel
	for result := range output {
		fmt.Println("Received from fan-in channel:", result)
	}
}

func cancellationWithChannels() {
	fmt.Println("---- Cancellation with Channels ----")
	ch := make(chan int)
	cancelCh := make(chan struct{})

	// Worker goroutine
	go func() {
		for {
			select {
			case <-cancelCh:
				fmt.Println("Cancellation requested. Exiting.")
				return
			case value := <-ch:
				fmt.Println("Received from cancellation channel:", value)
			}
		}
	}()

	// Simulate sending values to the channel
	for i := 1; i <= 3; i++ {
		ch <- i
	}

	// Send cancellation signal
	close(cancelCh)
}

func ChannelInGo() {
	basicChannel()
	bufferedChannel()
	closingChannel()
	selectStatement()
	unidirectionalChannels()
	fanOutFanIn()
	cancellationWithChannels()
}
