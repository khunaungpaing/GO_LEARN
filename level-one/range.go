package levelone

import "fmt"

func RangeInGo() {
	// Array and Slice
	numbersArray := [4]int{1, 2, 3, 4}
	numbersSlice := []int{5, 6, 7, 8}

	// Range with arrays
	fmt.Println("Range with arrays:")
	for index, value := range numbersArray {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Range with slices
	fmt.Println("\nRange with slices:")
	for index, value := range numbersSlice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// String
	message := "Hello, Go!"

	// Range with strings
	fmt.Println("\nRange with strings:")
	for index, char := range message {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// Map
	studentGrades := map[string]int{
		"Alice":   92,
		"Bob":     78,
		"Charlie": 88,
	}

	// Range with maps
	fmt.Println("\nRange with maps:")
	for name, grade := range studentGrades {
		fmt.Printf("%s's Grade: %d\n", name, grade)
	}

	// Channel
	numbersChannel := make(chan int, 3)
	numbersChannel <- 1
	numbersChannel <- 2
	numbersChannel <- 3
	close(numbersChannel)

	// Range with channels
	fmt.Println("\nRange with channels:")
	for num := range numbersChannel {
		fmt.Println(num)
	}
}
