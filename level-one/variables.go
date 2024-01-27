package levelone

import "fmt"

func GoVariables() {
	// Using var Keyword
	var name string
	//assign value to the variable
	name = "Mg Mg"
	// Printing variable
	fmt.Println(name)

	// Using var Keyword and initializing value with datatype
	var age int = 5
	// Printing variable
	fmt.Println(age)

	// Using var Keyword and initializing value without datatype
	var height = 5
	// Printing variable
	fmt.Println(height)

	// Using Short Variable Declaration (:=)
	isStudent := true
	// Printing variable
	fmt.Println(isStudent)

	// Short variable declaration to declare and initialize multiple variables
	num, str := 10, "hello"

	// Printing variables
	fmt.Println("num:", num)
	fmt.Println("str:", str)

}
