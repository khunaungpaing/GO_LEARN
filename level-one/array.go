package levelone

import "fmt"

func ArrayInGo() {
	// Declaring and initializing an array
	numbers := [5]int{1, 2, 3, 4, 5}

	// Accessing and printing elements
	fmt.Println("First Element:", numbers[0])
	fmt.Println("Second Element:", numbers[1])

	// Modifying an element
	numbers[2] = 10

	// Printing the modified array
	fmt.Println("Modified Array:", numbers)

	// Length of the array
	length := len(numbers)
	fmt.Println("Length of the Array:", length)
}
