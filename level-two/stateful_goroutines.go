package leveltwo

import (
	"fmt"
	"sync"
)

// In Go programming language, goroutines are lightweight threads of execution.
// They are managed by the Go runtime and provide a way to execute concurrent code.
// By default, goroutines are stateless, meaning they don't retain any state between function calls.
// However, you can achieve stateful behavior in goroutines using various techniques.

func usingClousers() {
	var wg sync.WaitGroup

	// Create a goroutine with state using closure
	statefulGoroutine := func(initialState int) func() {
		state := initialState

		return func() {
			defer wg.Done()
			fmt.Println(state)
			state++
		}
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go statefulGoroutine(i)()
	}

	wg.Wait()
}

type StatefulCounter struct {
	state int
	mu    sync.Mutex
}

func (s *StatefulCounter) Increment() {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println(s.state)
	s.state++
}

func usingStruct() {
	var wg sync.WaitGroup
	counter := StatefulCounter{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
}

func statefulCh(stateCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	state := <-stateCh
	fmt.Println(state)
	stateCh <- state + 1
}

func usingChannel() {
	var wg sync.WaitGroup
	stateCh := make(chan int, 3)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go statefulCh(stateCh, &wg)
		stateCh <- i
	}

	wg.Wait()
	close(stateCh)
}

func StatefulGoroutine() {
	fmt.Println("Using Clouser")
	usingClousers()
	fmt.Println("\nUsing Struct")
	usingStruct()
	fmt.Println("\nUsing Channel")
	usingChannel()
}
