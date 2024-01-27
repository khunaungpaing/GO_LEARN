package levelone

import "fmt"

// Basic function
func greet(name string) {
	fmt.Println("Hello, " + name + "!")
}

// Function with multiple parameters and a return value
func add(a, b int) int {
	return a + b
}

// Function with named return values
func divide(dividend, divisor float64) (quotient float64, remainder float64) {
	quotient = dividend / divisor
	remainder = dividend - (quotient * divisor)
	return // implicit return of named values
}

// Variadic function (accepts variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func FunctionsInGo() {
	// Basic function call
	greet("John")

	// Function with multiple parameters and return value
	result := add(3, 7)
	fmt.Println("Sum:", result)

	// Function with named return values
	q, r := divide(10.0, 3.0)
	fmt.Printf("Quotient: %f, Remainder: %f\n", q, r)

	// Variadic function call
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers:", total)

	// Anonymous function (closure)
	func() {
		fmt.Println("This is an anonymous function.")
	}()

	// Anonymous function with parameters
	multiply := func(x, y int) int {
		return x * y
	}
	product := multiply(5, 4)
	fmt.Println("Product:", product)
}
