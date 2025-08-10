// Package basic demonstrates various ways to define and use functions in Go.
package basic

import "fmt"

func Functions() {
	// Syntax: func <name>(parameters) returnType { ... }

	// --- 1. BASIC FUNCTION CALL ---

	// Functions are declared with: func <name>(parameters) returnType { ... }
	sum := add(1, 2)
	fmt.Println("1 + 2 =", sum) // Output: 3

	// Functions can be called directly
	fmt.Println("2 + 3 =", add(2, 3))

	// --- 2. ANONYMOUS FUNCTIONS ---

	// Anonymous functions don't have a name — useful for quick one-time operations
	greet := func() {
		fmt.Println("Hello from an Anonymous Function")
	}
	greet() // Call it like any other function

	// --- 3. FUNCTIONS AS VARIABLES ---

	// You can assign a function to a variable and call it later
	operation := add
	result := operation(3, 5)
	fmt.Println("3 + 5 =", result)

	// --- 4. PASSING FUNCTIONS AS ARGUMENTS ---

	// Here we pass 'add' as an argument to another function
	result = applyOperation(5, 3, add)
	fmt.Println("5 + 3 =", result)

	// We can also pass anonymous functions
	result = applyOperation(5, 3, func(a, b int) int {
		return a * b
	})
	fmt.Println("5 * 3 =", result)

	// --- 5. RETURNING FUNCTIONS (CLOSURES) ---

	// createMultiplier returns a function that multiplies by a given factor
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 =", multiplyBy2(6))

	// createMultiplier(3) returns a function that multiplies by 3
	multiplyBy3 := createMultiplier(3)
	fmt.Println("7 * 3 =", multiplyBy3(7))

	// Each returned function keeps its own copy of the environment variable (factor)
	// This is called a closure
}

// --- 6. FUNCTION DEFINITIONS ---

// add is a simple function with two int parameters and one int return value
func add(a, b int) int {
	return a + b
}

// applyOperation is a higher-order function (accepts a function as an argument)
// operation is a function parameter with signature: func(int, int) int
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// createMultiplier is a function that returns another function
// The returned function multiplies its input by 'factor'
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
