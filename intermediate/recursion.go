// recursion.go
// Package intermediate demonstrates recursion in Go.
// ------------------------------------------------------------------
// Important Notes:
// 1. Recursion = a function calling itself until a base case is met.
// 2. Every recursive function must have a **base case** to avoid infinite calls.
// 3. Common use cases: factorials, Fibonacci, tree traversal, divide-and-conquer algorithms.
// 4. Recursion can be direct (function calls itself) or indirect (two+ functions call each other).
// 5. Go does not optimize tail recursion → deep recursion may cause stack overflow.
// ------------------------------------------------------------------

package intermediate

import "fmt"

// Recursion demonstrates different recursion examples.
func Recursion() {
	// Example 1: Factorial (classic recursion)
	fmt.Println("=== Example 1: Factorial ===")
	fmt.Println("5! =", factorial(5))
	fmt.Println("10! =", factorial(10))

	// Example 2: Sum of digits
	fmt.Println("\n=== Example 2: Sum of digits ===")
	fmt.Println("Sum of digits (9):", sumOfDigits(9))         // 9
	fmt.Println("Sum of digits (12):", sumOfDigits(12))       // 3
	fmt.Println("Sum of digits (12345):", sumOfDigits(12345)) // 15

	// Example 3: Fibonacci sequence
	fmt.Println("\n=== Example 3: Fibonacci ===")
	for i := 0; i < 7; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fibonacci(i))
	}

	// Example 4: Indirect recursion (two functions calling each other)
	fmt.Println("\n=== Example 4: Indirect recursion ===")
	indirectRecursion(5)

	// Example 5: Recursion with slices
	fmt.Println("\n=== Example 5: Recursion with slices ===")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of slice:", sumSlice(nums))

	// Example 6: Infinite recursion (demonstration only)
	// ⚠️ WARNING: Uncommenting this will cause stack overflow!
	// fmt.Println(infiniteRecursion())
}

// ------------------------------------------------------------------
// Example 1: Factorial using recursion
func factorial(n int) int {
	if n == 0 { // Base case
		return 1
	}
	return n * factorial(n-1) // Recursive case
}

// ------------------------------------------------------------------
// Example 2: Sum of digits of a number
func sumOfDigits(n int) int {
	if n < 10 { // Base case
		return n
	}
	return n%10 + sumOfDigits(n/10) // Recursive case
}

// ------------------------------------------------------------------
// Example 3: Fibonacci sequence
func fibonacci(n int) int {
	if n <= 1 { // Base case
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2) // Recursive case
}

// ------------------------------------------------------------------
// Example 4: Indirect recursion
func indirectRecursion(n int) {
	if n <= 0 {
		fmt.Println("Done")
		return
	}
	fmt.Println("indirectRecursion:", n)
	helperIndirect(n - 1)
}

func helperIndirect(n int) {
	if n <= 0 {
		fmt.Println("Helper done")
		return
	}
	fmt.Println("helperIndirect:", n)
	indirectRecursion(n - 1)
}

// ------------------------------------------------------------------
// Example 5: Recursion with slices (sum all elements)
func sumSlice(nums []int) int {
	if len(nums) == 0 { // Base case
		return 0
	}
	return nums[0] + sumSlice(nums[1:]) // Recursive case
}

// ------------------------------------------------------------------
// Example 6: Infinite recursion (DO NOT RUN)
// This will never reach a base case, causing stack overflow.
func infiniteRecursion() int {
	fmt.Println("infiniteRecursion called")
	return infiniteRecursion() // No base case!
}
