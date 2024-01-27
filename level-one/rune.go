package levelone

import "fmt"

func RuneInGo() {
	// Iterating over a string using range and rune
	str := "Hello, 世界"

	for _, char := range str {
		fmt.Printf("Character: %c\n", char)
		fmt.Printf("Unicode Code Point: %U\n", char)
	}
}
