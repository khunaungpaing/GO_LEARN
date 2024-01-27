package levelone

import "fmt"

// Person represents an individual with a name and age.
type Human struct {
	Name string
	Age  int
}

// Student represents a student, embedding the Person type and adding a Mark field.
type Student struct {
	Human
	Mark int
}

func EmbedStructInGo() {
	student := Student{
		Human: Human{
			Name: "John",
			Age:  20,
		},
		Mark: 90,
	}

	// Accessing fields of the embedded Person struct
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)

	// Accessing fields of the Student struct
	fmt.Println("Mark:", student.Mark)
}
