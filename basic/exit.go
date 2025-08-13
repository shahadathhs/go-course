// exit.go
// Package basic demonstrates how os.Exit() works in Go.
//
// 📌 Important Notes:
// 1. `os.Exit(code int)` terminates the program immediately with the given exit code.
// 2. Deferred functions are NOT executed when `os.Exit()` is called.
// 3. An exit code of 0 means "success", any non-zero code means "error".
// 4. Code after os.Exit() is unreachable.
//
// 📌 Syntax:
//    os.Exit(statusCode)
//
// Example:
//    os.Exit(0)  // Successful exit
//    os.Exit(1)  // Error exit
//
// 🔹 Difference from panic():
//    - panic() runs deferred statements before exiting
//    - os.Exit() skips all deferred statements

package basic

import (
	"fmt"
	"os"
)

func Exit() {
	defer fmt.Println("Deferred statement") // ❌ This will NOT run if os.Exit() is called

	fmt.Println("Starting the main function")

	// Exit with status code of 1 (indicates an error)
	os.Exit(1)

	// ❌ This will never be executed
	fmt.Println("End of main function")
}
