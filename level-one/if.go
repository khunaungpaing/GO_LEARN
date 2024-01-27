package levelone

import "fmt"

func IfElseInGo(temperature int) {

	// If else
	if temperature > 30 {
		fmt.Println("It's hot outside.")
	} else if temperature >= 20 && temperature <= 30 {
		fmt.Println("It's a pleasant day.")
	} else {
		fmt.Println("It's cold outside.")
	}

	// If else with Short variable declaration
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5.")
	} else {
		fmt.Println("x is not greater than 5.")
	}
}
