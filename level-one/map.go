package levelone

import "fmt"

func MapInGo() {
	// Create an empty map
	studentGrades := make(map[string]int)

	// Add key-value pairs to the map
	studentGrades["Alice"] = 92
	studentGrades["Bob"] = 78
	studentGrades["Charlie"] = 88

	// Access values using keys
	aliceGrade := studentGrades["Alice"]
	fmt.Println("Alice's Grade:", aliceGrade)

	// Check if a key exists
	if grade, exists := studentGrades["Bob"]; exists {
		fmt.Println("Bob's Grade:", grade)
	} else {
		fmt.Println("Bob's Grade not found.")
	}

	// Iterating over the map
	fmt.Println("\nAll Grades:")
	for name, grade := range studentGrades {
		fmt.Printf("%s's Grade: %d\n", name, grade)
	}

	// Delete a key-value pair
	delete(studentGrades, "Alice")

	// Check if a key exists after deletion
	if _, exists := studentGrades["Alice"]; exists {
		fmt.Println("Alice's Grade:", studentGrades["Alice"])
	} else {
		fmt.Println("Alice's Grade not found after deletion.")
	}
}
