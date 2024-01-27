package levelone

import "fmt"

func PointerInGo() {
	// Declare a variable
	value := 42

	// Declare a pointer that points to the memory address of the variable
	pointer := &value

	// Print the value and the memory address
	fmt.Println("Value:", value)
	fmt.Println("Memory Address:", pointer)

	// Access the value through the pointer
	fmt.Println("Value through Pointer:", *pointer)

	// Modify the value through the pointer
	*pointer = 99

	// The original variable is also modified
	fmt.Println("Modified Value:", value)

	// Create a pointer to an integer with the new function
	newPointer := new(int)

	// Assign a value through the pointer
	*newPointer = 42

	// Print the value through the new pointer
	fmt.Println("Value through New Pointer:", *newPointer)
}
