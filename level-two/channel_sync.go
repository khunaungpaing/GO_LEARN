package leveltwo

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate some work
		fmt.Printf("Worker %d finished job %d\n", id, job)

		results <- job * 2
	}
}

func ChannelSynchronizationInGo() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go worker(i, jobs, results, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()

	close(results)

	// Collect and print results from the results channel
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
