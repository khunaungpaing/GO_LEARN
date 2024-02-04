package leveltwo

import (
	"fmt"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func sum(in <-chan int, wg *sync.WaitGroup) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		sum := 0
		for n := range in {
			sum += n
		}
		out <- sum
		wg.Done()
	}()
	return out
}

func ChannelDirectionInGo() {
	fmt.Println("This is Pipeline Pattern")
	fmt.Println("------------------------")
	nums := []int{1, 2, 3, 4, 5}

	// Creating channels for each stage
	genChan := generator(nums...)
	squareChan := square(genChan)
	newOne := squareChan
	for n := range newOne {
		fmt.Println(n)
	}

	// Creating a WaitGroup to wait for the sum stage to finish
	var wg sync.WaitGroup

	// Starting the sum stage
	wg.Add(1)
	sumChan := sum(squareChan, &wg)

	// Waiting for the sum stage to finish
	go func() {
		wg.Wait()
	}()

	// Reading the final result from the sum channel
	result := <-sumChan
	fmt.Println("Final Result:", result)
}
