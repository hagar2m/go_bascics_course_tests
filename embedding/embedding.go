package embedding

import (
	"fmt"
)

// Define a struct called "Person" with methods.
type Person struct {
	FirstName string
	LastName  string
}

// Method for greeting.
func (p *Person) Greet() {
	fmt.Printf("Hello, my name is %s %s.\n", p.FirstName, p.LastName)
}

// Define another struct called "Employee" that embeds "Person".
type Employee struct {
	Person // Embed the "Person" struct
	EmployeeID int
}

// Method for displaying employee information.
func (e *Employee) DisplayInfo() {
	fmt.Printf("Employee ID: %d\n", e.EmployeeID)
}

// func main() {
// 	// Create an "Employee" instance.
// 	employee := Employee{
// 		Person: Person{
// 			FirstName: "John",
// 			LastName:  "Doe",
// 		},
// 		EmployeeID: 12345,
// 	}

// 	// Call methods on the "Employee" instance.
// 	employee.Greet()
// 	employee.DisplayInfo()
// }
