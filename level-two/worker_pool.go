package leveltwo

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID       int
	duration time.Duration
}

type Worker struct {
	ID       int
	jobQueue chan Job
}

func NewWorker(id int, jobQueue chan Job) *Worker {
	return &Worker{
		ID:       id,
		jobQueue: jobQueue,
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range w.jobQueue {
		time.Sleep(job.duration)
		fmt.Printf("Worker %d processing job %d\n", w.ID, job.ID)
	}
}

func WorkerPool() {
	now := time.Now()
	numWorkers := 5

	jobQueue := make(chan Job, 10)

	var wg sync.WaitGroup

	// Create and start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(i, jobQueue)
		wg.Add(1)
		go worker.Start(&wg)
	}

	// Create some jobs and send them to the job channel
	for i := 1; i <= 10; i++ {
		job := Job{ID: i, duration: time.Second * time.Duration(i/2)}
		jobQueue <- job
	}

	// Close the job channel to signal that no more jobs will be sent
	close(jobQueue)

	wg.Wait()
	fmt.Println("Duration to Done: ", time.Since(now))
}
