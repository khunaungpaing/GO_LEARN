package leveltwo

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		fmt.Println(string(char))
		time.Sleep(100 * time.Millisecond)
	}
}

func GoRoutines() {
	// Creating and starting two goroutines
	go printNumbers()
	go printLetters()

	// Sleeping to allow goroutines to execute
	time.Sleep(1 * time.Second)

	fmt.Println("GoRoutines() function exits.")
}
