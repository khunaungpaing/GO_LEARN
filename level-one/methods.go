package levelone

import "fmt"

// Define a type named 'Counter'
type Counter struct {
	count int
}

// Method associated with the 'Counter' type
func (c *Counter) Increment() {
	c.count++
}

func MethodInGo() {
	// Create an instance of 'Counter'
	counter := Counter{}

	// Call the 'Increment' method on the 'Counter' instance
	counter.Increment()

	// Print the updated count
	fmt.Println("Count:", counter.count)
}
