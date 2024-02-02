package leveltwo

import "fmt"

func ChannelInGo() {
	// Normal channel (both send and receive)
	normalCh := make(chan int)

	// Readonly channel (symbol : <-chan)
	// readOnlyCh := make(<-chan int)

	// Write-only channle (symbol : chan<-)
	// writeOnlyCh := make(chan<- int)

	// Sending a value into the channel
	go func() {
		normalCh <- 42
	}()

	// Receiving a value from the channel
	value := <-normalCh

	fmt.Println("Value received from channel: ", value)
}
