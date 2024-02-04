package leveltwo

import (
	"fmt"
	"time"
)

func TimeoutInGo() {
	// Creating a channel for timeout
	timeout := make(chan bool, 1)

	// Creating a channel for communication
	result := make(chan string)

	// Goroutine 1: Simulates a long-running task
	go func() {
		time.Sleep(3 * time.Second)
		result <- "Task completed"
	}()

	// Goroutine 2: Waits for the task result or timeout
	go func() {
		select {
		case <-result:
			fmt.Println("Received result:", <-result)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout occurred")
			timeout <- true
		}
	}()

	fmt.Println("Timeout : ", <-timeout)
}
