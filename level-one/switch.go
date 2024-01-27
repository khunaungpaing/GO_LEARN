package levelone

import "fmt"

func SwitchInGo() {
	day := "Wednesday"

	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Tuesday", "Wednesday":
		fmt.Println("It's a weekday.")
	case "Thursday":
		fmt.Println("Almost there to the weekend.")
	case "Friday":
		fmt.Println("TGIF!")
	default:
		fmt.Println("It's the weekend.")
	}
}
