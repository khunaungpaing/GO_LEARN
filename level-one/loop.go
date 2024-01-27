package levelone

import "fmt"

func LoopInGo() {
	// 1. Basic for loop
	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 2. for loop with a condition
	fmt.Println("\nFor loop with a condition:")
	counter := 0
	for counter < 5 {
		fmt.Println(counter)
		counter++
	}

	// 3. Infinite for loop with break
	fmt.Println("\nInfinite for loop with break:")
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 5 {
			break // exit the loop
		}
	}

	// 4. for loop with range
	fmt.Println("\nFor loop with range:")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 5. Nested for loop
	fmt.Println("\nNested for loop:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}
}
