package levelone

import "fmt"

// Define a struct type "Circle" that implements the "Shape" interface
type Circle struct {
	Name   string
	Radius float64
}

// Implement the "Area" method for the "Circle" type
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) GetName() string {
	return c.Name
}

// Define a struct type "Rectangle" that implements the "Shape" interface
type Rectangle struct {
	Name   string
	Width  float64
	Height float64
}

// Implement the "Area" method for the "Rectangle" type
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) GetName() string {
	return r.Name
}

func printArea(s Shape) {
	fmt.Printf("Area of %s: %f\n", s.GetName(), s.Area())
}

func InterfaceImplInGo() {
	// Create instances of Circle and Rectangle
	circle := Circle{Name: "sample circle", Radius: 5.0}
	rectangle := Rectangle{Name: "sample rectangle", Width: 4.0, Height: 6.0}

	// Call the printArea function with Circle and Rectangle
	printArea(circle)
	printArea(rectangle)
}
