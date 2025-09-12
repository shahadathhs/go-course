// closures.go
// Package intermediate demonstrates the concept of closures in Go.
// ------------------------------------------------------------------
// Important Notes:
// 1. A closure is a function value that references variables from outside its body.
// 2. The captured variables are "closed over" and live on even after the outer function has returned.
// 3. Each closure maintains its own independent environment.
// 4. Closures can be used for stateful functions, counters, factories, and encapsulation.
// 5. Unlike global variables, closures provide localized state without exposing it globally.
// ------------------------------------------------------------------

package intermediate

import "fmt"

// Closures demonstrates multiple closure use cases.
func Closures() {
	// Example 1: Basic adder closure
	fmt.Println("=== Example 1: Basic adder closure ===")
	sequence := adder()     // returns a closure
	fmt.Println(sequence()) // 1
	fmt.Println(sequence()) // 2
	fmt.Println(sequence()) // 3
	fmt.Println(sequence()) // 4

	// New independent closure
	sequence2 := adder()
	fmt.Println(sequence2()) // 1 (fresh counter)

	// Example 2: Closure with parameters (subtractor)
	fmt.Println("\n=== Example 2: Subtractor closure with parameter ===")
	subtracter := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()
	fmt.Println(subtracter(1)) // 98
	fmt.Println(subtracter(2)) // 96
	fmt.Println(subtracter(3)) // 93

	// Example 3: Closure as function factory
	fmt.Println("\n=== Example 3: Function factory ===")
	multiplierBy2 := multiplier(2)
	multiplierBy3 := multiplier(3)
	fmt.Println("2 * 5 =", multiplierBy2(5))
	fmt.Println("3 * 5 =", multiplierBy3(5))

	// Example 4: Immediate invocation with closure
	fmt.Println("\n=== Example 4: Immediate closure invocation ===")
	result := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println("Sum from immediate closure:", result)

	// Example 5: Closure capturing loop variable
	fmt.Println("\n=== Example 5: Closure capturing loop variable ===")
	funcs := make([]func() int, 3)
	for i := 0; i < 3; i++ {
		funcs[i] = func(n int) func() int {
			return func() int {
				return n
			}
		}(i) // capture current value of i
	}
	for _, f := range funcs {
		fmt.Println(f()) // prints 0, 1, 2
	}
}

// ------------------------------------------------------------------
// Example 1: Closure maintaining state
func adder() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// ------------------------------------------------------------------
// Example 3: Function factory closure
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}
