package levelone

import "fmt"

// Note : In Go, the iota identifier is used in const declarations
// to create a sequence of related values that increment automatically.
// It starts with 0 and increments by 1 for each subsequent constant in the const block.

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func ConstInGo() {
	// Define constants
	const pi = 3.14159
	const appName = "MyApp"
	const gravity = 9.81

	// Printing values
	fmt.Println("Pi:", pi)
	fmt.Println("Application Name:", appName)
	fmt.Println("Gravity:", gravity)

	fmt.Println("Days of the week:")
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
}
