// init.go  
// Package basic demonstrates the use of the `init()` function in Go
package basic

import "fmt"

// NOTES:
// - The `init()` function is a special function in Go that is automatically
//   executed before the `main()` function or any other function in the same package.
// - You can have multiple `init()` functions in the same package (even in different files).
// - Execution order of init functions:
//     1. All package-level variable initializations happen first.
//     2. Then all `init()` functions in the order they are defined (file by file).
// - Common use cases: initializing state, setting up configurations, precomputing data.
// - `init()` cannot take arguments and does not return any values.

func init() {
	fmt.Println("Initializing package1...")
}

func init() {
	fmt.Println("Initializing package2...")
}

func init() {
	fmt.Println("Initializing package3...")
}

// Init demonstrates a normal function call after init functions
func Init() {
	fmt.Println("Inside the main function")
}
