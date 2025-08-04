// variables.go
// Package basic demonstrates Go variable declaration, initialization, scope, shadowing, and zero values.
package basic

import "fmt"

// Package-level (global) variables
var (
	// Typed variable with initial value
	GlobalCount int = 12
	// Untyped variable (type inferred)
	globalFlag = true
)

// Variables demonstrates various ways to declare and use variables in Go.
func Variables() {
	// 1. Declaration with var and explicit type
	var x int
	x = 10
	fmt.Println("x:", x)

	// 2. Declaration with var, type, and initialization
	var greeting string = "Hello"

	// 3. Short declaration (type inferred)
	name := "World"

	// 4. Multiple variables
	a, b, c := 1, 2, 3

	// 5. Blank identifier to ignore a value
	_, useful := "ignored", "used"

	// 6. Zero values
	var i int     // defaults to 0
	var f float64 // defaults to 0.0
	var s string  // defaults to ""
	var flag bool // defaults to false
	var ptr *int  // defaults to nil

	// Print values
	fmt.Println(greeting, name)
	fmt.Println("\na, b, c:", a, b, c)
	fmt.Println("useful:", useful)
	fmt.Printf("Zero values -> int: %d, float64: %f, string: %q, bool: %t, ptr: %v\n", i, f, s, flag, ptr)

	// 7. Variable shadowing in inner scope
	value := 100
	fmt.Println("\nouter value:", value)
	{
		// shadows outer 'value'
		value := 200
		fmt.Println("\ninner shadowed value:", value)
	}
	fmt.Println("\nouter value after inner block:", value)

	// 8. Using package-level variables
	fmt.Println("\nGlobalCount:", GlobalCount)
	GlobalCount++
	fmt.Println("\nGlobalCount after increment:", GlobalCount)
	fmt.Println("globalFlag:", globalFlag)

	// Call another function to show scope
	scopeExample()
}

// scopeExample shows that local variables from Variables() are not visible here
func scopeExample() {
	// 'x' from Variables() is not accessible here; this is a new local
	var x int = 50
	fmt.Println("\nscopeExample x:", x)
	fmt.Println("accessing GlobalCount in scopeExample:", GlobalCount)
}
