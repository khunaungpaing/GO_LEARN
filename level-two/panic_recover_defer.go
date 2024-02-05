package leveltwo

import "fmt"

func PanicRecoverDefer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
		fmt.Println("Deferred cleanup in main")
	}()

	fmt.Println("Main function starts")
	performTask()
	fmt.Println("Main function continues")
}

func performTask() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in performTask:", r)
		}
		fmt.Println("Deferred cleanup in performTask")
	}()

	fmt.Println("Performing the task")
	simulateError()
	fmt.Println("Task completed successfully")
}

func simulateError() {
	defer fmt.Println("Deferred cleanup in simulateError")

	fmt.Println("Simulating an error")
	panic("Error occurred!")
}
