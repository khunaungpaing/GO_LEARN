package levelone

import "fmt"

// Variadic function to concatenate strings
func concatenateStrings(strings ...string) string {
	result := ""
	for _, str := range strings {
		result += str
	}
	return result
}

func VariadicInGo() {
	result := concatenateStrings("Hello, ", "Go ", "World!")
	fmt.Println("Concatenated String:", result)
}
