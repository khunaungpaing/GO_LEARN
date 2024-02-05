package leveltwo

import (
	"fmt"
	"regexp"
)

func RegularExpressionInGo() {
	// Basic Matching
	re := regexp.MustCompile(`\d+`)
	match := re.MatchString("123abc")
	fmt.Println("1. Basic Matching:", match)

	// Finding Matches
	re = regexp.MustCompile(`\d+`)
	firstMatch := re.FindString("123abc456def")
	allMatches := re.FindAllString("123abc456def", -1)
	fmt.Println("2. Finding Matches:")
	fmt.Println("   First Match:", firstMatch)
	fmt.Println("   All Matches:", allMatches)

	// Submatches
	re = regexp.MustCompile(`(\d+)([a-z]+)`)
	submatches := re.FindStringSubmatch("123abc")
	fmt.Println("3. Submatches:")
	fmt.Println("   Submatches:", submatches)
	if len(submatches) >= 3 {
		fmt.Println("   Number:", submatches[1])
		fmt.Println("   Letters:", submatches[2])
	}

	// Replace
	re = regexp.MustCompile(`\d+`)
	replaced := re.ReplaceAllString("123abc456def", "X")
	fmt.Println("4. Replace:", replaced)

	// Flags
	re = regexp.MustCompile(`(?i)go`)
	flagMatch := re.MatchString("Go is awesome!")
	fmt.Println("5. Flags:", flagMatch)
}
