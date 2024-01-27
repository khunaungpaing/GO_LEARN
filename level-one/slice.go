package levelone

import "fmt"

func SliceInGo() {
	//In Go, a slice is a flexible, dynamically-sized view into the elements of an array.

	// Syntax: make([]T, length, capacity)
	// Creating a slice
	mySlice := make([]int, 3, 5)
	fmt.Println("Slice:", mySlice)

	// Modifying the slice
	mySlice[0] = 1
	mySlice[1] = 2
	mySlice[2] = 3
	fmt.Println("Modified Slice:", mySlice)

	// Appending elements
	mySlice = append(mySlice, 4, 5, 6)
	fmt.Println("Appended Slice:", mySlice)

	// Slicing the slice
	subSlice := mySlice[1:4]
	fmt.Println("Sub-Slice:", subSlice)

	// Length and Capacity
	fmt.Printf("Length: %d, Capacity: %d\n", len(mySlice), cap(mySlice))

}
