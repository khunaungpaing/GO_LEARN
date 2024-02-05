package leveltwo

import (
	"fmt"
)

func FormatStrings() {
	// 1. Basic String Formatting
	name := "John"
	age := 30
	fmt.Printf("1. Name: %s, Age: %d\n", name, age)

	// 2. Width and Precision
	price := 29.99
	fmt.Printf("2. Price: %.2f\n", price)

	// 3. String Alignment
	text := "Hello"
	fmt.Printf("3a. |%10s|\n", text)  // Right-aligned
	fmt.Printf("3b. |%-10s|\n", text) // Left-aligned

	// 4. Integer Formatting
	number := 42
	fmt.Printf("4. Decimal: %d, Binary: %b, Hexadecimal: %x\n", number, number, number)

	// 5. Sprintf for String Assignment
	formattedString := fmt.Sprintf("5. Name: %s, Age: %d", name, age)
	fmt.Println(formattedString)

	// 6. String Concatenation
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Println("6. Full Name:", fullName)
}
