package levelone

import (
	"errors"
	"fmt"
)

// Function with multiple return values
func calculateAverageAndSum(numbers []float64) (float64, float64, error) {
	if len(numbers) == 0 {
		// Returning an error if the input list is empty
		return 0, 0, errors.New("empty list")
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	average := sum / float64(len(numbers))

	// Returning the average, sum, and a nil error
	return average, sum, nil
}

func MultipleReturnInGo() {
	// Sample list of numbers
	numberList := []float64{5, 10, 15, 20}

	// Calling the function with multiple return values
	average, sum, err := calculateAverageAndSum(numberList)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Average:", average)
		fmt.Println("Sum:", sum)
	}

	// Handling an error case with an empty list
	emptyList := []float64{}
	average, sum, err = calculateAverageAndSum(emptyList)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Average:", average)
		fmt.Println("Sum:", sum)
	}
}
