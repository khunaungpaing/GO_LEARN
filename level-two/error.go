package leveltwo

import "fmt"

// CustomError is a custom error type that satisfies the error interface
type CustomError struct {
	Code    int
	Message string
}

// Error implements the error interface for CustomError
func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func doSomething() error {
	// Simulate an error condition
	return CustomError{Code: 500, Message: "Something went wrong"}
}

func ErrorInGo() {

	err := doSomething()

	if err != nil {
		// Type assertion to access custom error fields
		if customErr, ok := err.(CustomError); ok {
			fmt.Printf("Custom error code: %d, message: %s\n", customErr.Code, customErr.Message)
		} else {
			// Handle non-custom errors
			fmt.Println("Non-custom error:", err)
		}
	} else {
		fmt.Println("No error occurred.")
	}
}
