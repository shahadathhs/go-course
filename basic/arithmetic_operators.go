// arithmetic_operators.go
// Package basic demonstrates Go’s arithmetic operators,
// compound assignments, increment/decrement, and overflow behavior.
package basic

import (
	"fmt"
	"math"
)

// ArithmeticOperators shows usage of all basic arithmetic operators,
// compound assignments, increment/decrement, and overflow/underflow.
func ArithmeticOperators() {
	// ----------------------------------------
	// 1. Simple arithmetic (integer)
	// ----------------------------------------
	var a, b int = 10, 3
	fmt.Println("a =", a, ", b =", b)

	fmt.Println("Addition:         a + b =", a+b)
	fmt.Println("Subtraction:      a - b =", a-b)
	fmt.Println("Multiplication:   a * b =", a*b)
	fmt.Println("Integer Division: a / b =", a/b)
	fmt.Println("Remainder:        a % b =", a%b)

	// ----------------------------------------
	// 2. Floating-point arithmetic
	// ----------------------------------------
	var x, y float64 = 7.5, 2.0
	fmt.Println("\nx =", x, ", y =", y)
	fmt.Println("Addition:       x + y =", x+y)
	fmt.Println("Subtraction:    x - y =", x-y)
	fmt.Println("Multiplication: x * y =", x*y)
	fmt.Println("Division:       x / y =", x/y)

	// Pi approximation constant
	const piApprox = 22.0 / 7.0
	fmt.Println("\nPi approximation (22/7):", piApprox)

	// ----------------------------------------
	// 3. Compound assignment operators
	// ----------------------------------------
	n := 5
	fmt.Printf("\nn initial: %d\n", n)
	n += 3
	fmt.Println("n += 3  ->", n)
	n *= 2
	fmt.Println("n *= 2  ->", n)
	n -= 4
	fmt.Println("n -= 4  ->", n)
	n /= 2
	fmt.Println("n /= 2  ->", n)
	n %= 3
	fmt.Println("n %= 3  ->", n)

	// ----------------------------------------
	// 4. Increment and Decrement
	// ----------------------------------------
	counter := 0
	fmt.Printf("\ncounter initial: %d\n", counter)
	counter++ // only statement form; no ++n
	fmt.Println("counter++ ->", counter)
	counter-- // only statement form; no --n
	fmt.Println("counter-- ->", counter)

	// ----------------------------------------
	// 5. Overflow and Underflow
	// ----------------------------------------
	// Signed integer overflow wraps around
	var maxInt int64 = math.MaxInt64
	fmt.Printf("\nSigned int64 max: %d\n", maxInt)
	maxInt = maxInt + 1
	fmt.Println("After overflow:   ", maxInt)

	// Unsigned integer overflow wraps around
	var uMax uint64 = math.MaxUint64
	fmt.Printf("\nUnsigned uint64 max: %d\n", uMax)
	uMax = uMax + 1
	fmt.Println("After overflow:      ", uMax)

	// Floating-point underflow example
	var tiny float64 = 1e-323
	fmt.Printf("\nTiny float64: %g\n", tiny)
	tiny = tiny / math.MaxFloat64
	fmt.Println("After division by MaxFloat64:", tiny)
}
