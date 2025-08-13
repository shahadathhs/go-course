// panic.go
// Package basic demonstrates the usage of panic in Go.
// -----------------------------------------------------
// **Definition**
// `panic()` is a built-in function in Go that stops normal execution of the current goroutine.
// It starts panicking, unwinding the stack, and running all deferred calls before crashing the program.
// Panics are typically used for unrecoverable errors.
//
// **Syntax**
// panic(value interface{})
//
// **Key Points**
// 1. `panic` stops the normal execution immediately after it’s called.
// 2. All deferred functions in the current goroutine are executed in LIFO order.
// 3. After deferred calls finish, the program terminates unless `recover()` is used.
// 4. `panic` is often used for programmer errors, not regular error handling.
// 5. You can panic with any type (`string`, `error`, struct, etc.).
//
// **Common Use Cases**
// - Unexpected unrecoverable situations (e.g., array index out of range).
// - Internal sanity checks failing.
// - Library code detecting improper usage.
//
// **Warning**
// Avoid using panic for normal error handling — prefer `error` returns.

package basic

import "fmt"

// Panic demonstrates how panic works in Go.
func Panic() {
	fmt.Println("=== Panic Demonstration ===")

	// Example of a valid input (no panic)
	processPanic(10)

	fmt.Println("---------------------------")

	// Example of an invalid input (triggers panic)
	// Comment/uncomment to test
	processPanic(-3)

	// This will not run after panic
	fmt.Println("This will NOT be printed if panic occurs above")
}

func processPanic(input int) {
	// Deferred calls still run before program crashes
	defer fmt.Println("[Deferred] Statement 1: Always runs before panic exit")
	defer fmt.Println("[Deferred] Statement 2: Runs before Statement 1")

	if input < 0 {
		fmt.Println("Before panic")
		// Trigger panic with a descriptive message
		panic(fmt.Sprintf("Invalid input: %d (must be non-negative)", input))

		// These lines are unreachable after panic
		// fmt.Println("After panic")
		// defer fmt.Println("Deferred after panic")
	}

	// Normal execution when no panic
	fmt.Println("Processing input:", input)
}
