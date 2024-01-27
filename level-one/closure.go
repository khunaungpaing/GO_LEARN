package levelone

import "fmt"

func ClosureInGo() {
	// Outer function
	outerVariable := 10

	// Closure
	closureFunction := func() {
		fmt.Println("Outer variable inside closure:", outerVariable)
	}

	// Call the closure
	closureFunction()
}
