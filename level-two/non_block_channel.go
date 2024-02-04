package leveltwo

import (
	"fmt"
)

func NonBlockOpera() {
	ch := make(chan int, 1)

	// Non-blocking send
	select {
	case ch <- 42:
		fmt.Println("Sent value to channel")
	default:
		fmt.Println("Channel operation would block, so default case is executed")
	}

	// Non-blocking receive
	select {
	case value := <-ch:
		fmt.Println("Received value from channel:", value)
	default:
		fmt.Println("Channel operation would block, so default case is executed")
	}
}
