package leveltwo

import (
	"fmt"
	"sync"
	"time"
)

func sampleWork(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate some work
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroupInGo() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter before starting a new goroutine
		go sampleWork(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers done")
}
