// naming_conventions.go
// Package basic demonstrates Go naming conventions for packages, files,
// variables, constants, types, functions, and acronyms.
package basic

import "fmt"

// 1. Package names: lower_case, short, no underscores
//    e.g., "basic", "http", "math"

// 2. File names: lowercase, words separated by underscore
//    e.g., naming_conventions.go, data_types.go

// 3. Exported names: PascalCase (UpperCamelCase)
//   - Types, functions, variables, constants
//   - Acronyms remain uppercase: HTTPServer, JSONParser
//   - Examples:
type Employee struct { // exported struct
	FirstName string // exported field
	LastName  string // exported field
	Age       int    // exported field
}

// HTTPServer demonstrates an exported type with acronym
type HTTPServer struct{}

// NewEmployee is a constructor function for Employee
func NewEmployee(firstName, lastName string, age int) *Employee {
	return &Employee{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

// CalculateSalary is an exported function
func CalculateSalary(emp *Employee, basePay float64) float64 {
	return basePay * 1.2
}

// 4. Unexported names: camelCase (lowerCamelCase)
//   - helper functions, internal variables
//   - acronyms still uppercase internally: parseJSON
func parseJSON(data []byte) (interface{}, error) {
	// implementation...
	return nil, nil
}

// maxRetries is an unexported variable
var maxRetries = 5

// 5. Constants: PascalCase or ALL_CAPS for global constants?
//   - Go style: PascalCase for constants, unless truly symbolic do ALL_CAPS
const (
	DefaultTimeout = 30 // seconds
	MaxDBRetries   = 3
)

// 6. Acronyms in names: always uppercase
//    - e.g., URLParser, XMLHTTPRequest

// Demonstration of naming in action
func NamingConventions() {
	// Using exported constructor
	emp := NewEmployee("Alice", "Smith", 28)
	fmt.Println("Employee: ", emp)

	// Using exported function
	salary := CalculateSalary(emp, 1000)
	fmt.Println("Salary:", salary)

	// Using unexported helper
	_, _ = parseJSON([]byte(`{"key":"value"}`))

	// Using constants and variables
	fmt.Println("maxRetries:", maxRetries)
	fmt.Println("DefaultTimeout:", DefaultTimeout)
	fmt.Println("MaxDBRetries:", MaxDBRetries)
}
