// recover.go
// Package basic demonstrates panic recovery in Go
// ------------------------------------------------------------------
// Important Notes:
// 1. recover() is used inside a deferred function to catch a panic.
// 2. If recover() is called outside of a deferred function → it returns nil.
// 3. recover() stops the panic and resumes normal execution.
// 4. If no panic has happened, recover() returns nil.
// 5. panic → unwinds stack → executes deferred functions → recover can stop it.
// 6. Typical use case: gracefully handle unexpected errors (e.g., invalid input, runtime failure)
// ------------------------------------------------------------------

package basic

import "fmt"

func Recover() {
	// Example 1: Recovery from a panic
	fmt.Println("=== Example 1: Basic recovery ===")
	processRecover()
	fmt.Println("Returned from processRecover")

	// Example 2: No panic → recover() returns nil
	fmt.Println("\n=== Example 2: recover() without panic ===")
	processWithoutPanic()

	// Example 3: Multiple defer & recover
	fmt.Println("\n=== Example 3: Multiple defers with recover ===")
	multipleDeferRecover()
}

// ------------------------------------------------------------------
// Example 1: Recover from a panic
func processRecover() {
	defer func() {
		r := recover() // Captures panic value
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong!") // Triggers panic
	// fmt.Println("End Process")     // Not executed
}

// ------------------------------------------------------------------
// Example 2: recover() without panic → returns nil
func processWithoutPanic() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		} else {
			fmt.Println("No panic occurred")
		}
	}()
	fmt.Println("Running normally, no panic here.")
}

// ------------------------------------------------------------------
// Example 3: Multiple defer & recover
func multipleDeferRecover() {
	defer fmt.Println("Deferred 1 (executes last in LIFO order)")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in middle defer:", r)
		}
	}()
	defer fmt.Println("Deferred 2 (executes first in LIFO order)")

	fmt.Println("Triggering panic now...")
	panic("Panic in multipleDeferRecover")
}
