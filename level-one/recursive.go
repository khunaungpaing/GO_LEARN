package levelone

import "fmt"

// Recursive function to calculate the nth Fibonacci number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func RecursiveInGo() {
	result := fibonacci(8)
	fmt.Println("8th Fibonacci number:", result)
}
