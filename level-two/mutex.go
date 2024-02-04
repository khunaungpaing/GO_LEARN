package leveltwo

import (
	"fmt"
	"sync"
)

func MutexInGo() {
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Use the mutex to protect the critical section
			mu.Lock()
			defer mu.Unlock()

			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final counter value
	fmt.Println("Final Counter Value:", counter)
}
