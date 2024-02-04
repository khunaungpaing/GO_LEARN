package leveltwo

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int           { return len(s) }
func (s ByLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByLength) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

func CustomSorting() {
	words := []string{"apple", "banana", "orange", "grape"}

	fmt.Println("Before sorting:", words)

	// Using sort.Sort with custom type ByLength
	sort.Sort(ByLength(words))

	fmt.Println("After sorting by length:", words)
}
