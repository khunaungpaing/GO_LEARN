package leveltwo

import (
	"fmt"
	"os"
	"text/template"
)

func executeTextTemplate() {
	// Template string
	templateString := "Name: {{.Name}}, Age: {{.Age}}\n"

	// Create a new template
	tmpl, err := template.New("personTemplate").Parse(templateString)
	if err != nil {
		panic(err)
	}

	// Create an instance of the data struct
	person := Person{"John Doe", 30}

	// Execute the template with the data
	err = tmpl.Execute(os.Stdout, person)
	if err != nil {
		panic(err)
	}
}

func executeFileTemplate() {
	// Read template from a file
	tmpl, err := template.ParseFiles("person_template.tmpl")
	if err != nil {
		panic(err)
	}

	// Create a new file for output
	outputFile, err := os.Create("./bin/output.html")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Create an instance of the data struct
	person := Person{"Jane Doe", 25}

	// Execute the template with the data, writing to the new file
	err = tmpl.Execute(outputFile, person)
	if err != nil {
		panic(err)
	}

	fmt.Println("Template output written to ./bin/output.html")
}

func TextTemplate() {
	// Execute text template example
	executeTextTemplate()

	// Execute file template example
	executeFileTemplate()
}
