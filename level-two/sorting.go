package leveltwo

import (
	"fmt"
	"sort"
)

// Custom type for sorting
type Person struct {
	Name string
	Age  int
}

func sortAll(slice interface{}, descending bool) {
	switch s := slice.(type) {
	case []int:
		if descending {
			sort.Sort(sort.Reverse(sort.IntSlice(s)))
		} else {
			sort.Ints(s)
		}
		fmt.Printf("Sorted int slice: %v\n", s)

	case []string:
		if descending {
			sort.Sort(sort.Reverse(sort.StringSlice(s)))
		} else {
			sort.Strings(s)
		}
		fmt.Printf("Sorted string slice: %v\n", s)

	case []Person:
		// Custom sorting for the Person struct based on Age
		if descending {
			sort.Slice(s, func(i, j int) bool {
				return s[i].Age > s[j].Age
			})
		} else {
			sort.Slice(s, func(i, j int) bool {
				return s[i].Age < s[j].Age
			})
		}
		fmt.Printf("Sorted Person slice: %v\n", s)

	default:
		fmt.Println("Unsupported type for sorting")
	}
}

func SortingInGo() {

	intSlice := []int{4, 2, 7, 1, 9, 3}
	stringSlice := []string{"apple", "banana", "orange", "grape"}
	personSlice := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	fmt.Println("Before sorting:")
	fmt.Println("intSlice:", intSlice)
	fmt.Println("stringSlice:", stringSlice)
	fmt.Println("personSlice:", personSlice)

	// Sorting in ascending order
	sortAll(intSlice, false)
	sortAll(stringSlice, false)
	sortAll(personSlice, false)

	// Sorting in descending order
	sortAll(intSlice, true)
	sortAll(stringSlice, true)
	sortAll(personSlice, true)
}
