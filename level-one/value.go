package levelone

import "fmt"

func BasicValueInGo() {
	fmt.Println("\nBasic Value in go\n+++++++++++++++")
	// Integer value
	age := 25

	// String value
	name := "John Doe"

	// Boolean value
	isStudent := true
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is student:", isStudent)
	println()
}

func CompositeValuesInGo() {
	fmt.Println("Composite Value in go\n+++++++++++++++")

	// Array of integers
	numbers := [3]int{1, 2, 3}

	// Slice of integers
	moreNumbers := []int{4, 5, 6}

	fmt.Println("Array:", numbers)
	fmt.Println("Slice:", moreNumbers)
	println()
}

func MapValuesInGo() {
	fmt.Println("Map Value in go\n+++++++++++++++")

	// Map with string keys and int values
	scores := map[string]int{
		"Math":    90,
		"English": 85,
		"Science": 92,
	}

	fmt.Println("Scores:", scores)
	println()
}

// Define a struct type
type Person struct {
	Name  string
	Age   int
	Email string
}

func StructValuesInGo() {
	fmt.Println("Struct Value in go\n+++++++++++++++")

	// Create a struct instance
	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	fmt.Println("Person:", person)
	println()
}
