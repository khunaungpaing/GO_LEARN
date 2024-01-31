package leveltwo

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SuperHero struct {
	Name string
}

const iteration = 100000

func AtomicInGo() {
	// Normal primitive type
	normalCounter := normalCounterWithConcurrency()

	// Atomic operations with primitive type
	atomicCounter := atomicCounterWithConcurrency()

	// custom type using atomic value
	customTypeValue := customTypeUsingAtomicValue()

	// Print results
	fmt.Println("Normal Counter:", normalCounter, " not get fully count without mutex")
	fmt.Println("Atomic Counter:", atomic.LoadInt64(&atomicCounter), " get full count without mutex")
	fmt.Println("custom type using Atomic Value: ", customTypeValue)

}

func normalCounterWithConcurrency() int64 {
	var wg sync.WaitGroup
	wg.Add(iteration)
	var sum int64

	for i := 0; i < iteration; i++ {
		go func() {
			defer wg.Done()
			sum++
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
	return sum
}

func atomicCounterWithConcurrency() int64 {
	var wg sync.WaitGroup
	wg.Add(iteration)
	var sum int64

	for i := 0; i < iteration; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&sum, 1)
		}()
	}

	wg.Wait()
	return sum
}

func customTypeUsingAtomicValue() string {
	var wg sync.WaitGroup
	wg.Add(1)

	// Using Atomic Value for custom type
	var av atomic.Value
	av.Store(SuperHero{Name: "AntMan"})

	go func() {
		defer wg.Done()
		sp := av.Load().(SuperHero)
		sp.Name = "Spiderman"
		av.Store(sp)        // output will be Spiderman not antman
		sp.Name = "Ironman" // this will not be override because av.Store() store only value not address ;)
	}()

	wg.Wait()

	return av.Load().(SuperHero).Name
}

