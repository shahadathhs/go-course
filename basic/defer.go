// Package basic demonstrates the usage of `defer` in Go, including syntax,
// execution order, and common use cases.
package basic

import "fmt"

func Defer() {
	/*
		📌 SYNTAX:
			defer functionCall(arguments)

		📌 KEY POINTS:
		1. `defer` schedules a function call to be run after the surrounding function returns.
		2. Deferred calls are executed **in LIFO order** (Last In, First Out).
		3. Deferred calls execute even if the function returns early or panics (before program crashes).
		4. Common use case: resource cleanup (closing files, unlocking mutexes, closing DB connections).
		5. The arguments to deferred functions are evaluated **immediately**, not when the deferred function executes.
	*/

	process(10)
}

func process(i int) {
	// These deferred statements are executed in reverse order (LIFO)
	defer fmt.Println("Deferred i value:", i) // Argument evaluated now (will print original i value)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")

	// Normal execution
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:", i)

	/*
		📌 Expected output order:
		1. "Normal execution statement"
		2. "Value of i: 11"
		3. "Third deferred statement executed"
		4. "Second deferred statement executed"
		5. "First deferred statement executed"
		6. "Deferred i value: 10"
	*/
}
