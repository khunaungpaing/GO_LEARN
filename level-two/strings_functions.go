package leveltwo

import (
	"fmt"
	"strconv"
	"strings"
)

func ManipulateStrings() {
	// Length of a String
	str := "Hello, World!"
	length := len(str)
	fmt.Println("1. Length of the string:", length)

	// Concatenation
	str1 := "Hello"
	str2 := "World"
	result := str1 + ", " + str2
	fmt.Println("2. Concatenated string:", result)

	// Substring
	substr := str[7:12]
	fmt.Println("3. Substring:", substr)

	// String Comparison
	str1 = "Hello"
	str2 = "World"
	isEqual := (str1 == str2)
	fmt.Println("4. Strings are equal:", isEqual)

	// String Conversion
	num := 42
	strOut := strconv.Itoa(num)
	fmt.Println("5. Converted to string:", strOut)

	// String Splitting
	str = "apple,orange,banana"
	substrings := strings.Split(str, ",")
	fmt.Println("6. Split strings:", substrings)

	// String Joining
	substrings = []string{"apple", "orange", "banana"}
	joinedString := strings.Join(substrings, ",")
	fmt.Println("7. Joined string:", joinedString)

	// String Searching
	str = "Hello, World!"
	contains := strings.Contains(str, "World")
	index := strings.Index(str, "World")
	fmt.Println("8. Contains 'World':", contains)
	fmt.Println("9. Index of 'World':", index)
}
