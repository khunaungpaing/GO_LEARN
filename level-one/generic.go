package levelone

import "fmt"

// Constraint: Numeric
type Numeric interface {
	int | float64
}

// Generic function that calculates the sum of two values of any type satisfying the Numeric constraint
func Sum[T Numeric](a, b T) T {
	return a + b
}

func GenericInGo() {

	a, b := 3, 7
	x, y := 3.5, 2.0

	// Using the generic Sum function with different types
	sumInts := Sum(a, b)
	fmt.Println("Sum of Ints:", sumInts)

	sumFloats := Sum(x, y)
	fmt.Println("Sum of Floats:", sumFloats)

}
