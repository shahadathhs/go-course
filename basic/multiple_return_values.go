// Package basic demonstrates multiple return values in Go.
package basic

import (
	"errors"
	"fmt"
)

func MultipleReturnValues() {
	/*
		📌 SYNTAX for multiple return values:
			func functionName(param1 type1, param2 type2, ...) (returnType1, returnType2, ...) {
				// code block
				return value1, value2
			}

		📌 BEST PRACTICES:
		1. Place `error` as the last return value when returning multiple values.
		2. Use **named return values** when it improves clarity.
		3. Use `_` (blank identifier) to ignore values you don’t need.
	*/

	// --- 1. BASIC MULTIPLE RETURN VALUES ---

	// divide returns two integers: quotient and remainder
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %v, Remainder: %v\n", q, r)

	// You can ignore a return value using `_` (blank identifier)
	qOnly, _ := divide(9, 4)
	fmt.Printf("Quotient only: %v\n", qOnly)

	// --- 2. MULTIPLE RETURNS WITH ERROR HANDLING ---

	// compare returns a string and an error
	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// Example where a > b
	result, err = compare(5, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

// --- 3. FUNCTION DEFINITIONS ---

// divide returns quotient and remainder of integer division
// Named return values (quotient, remainder) are declared in the signature
// This allows us to return without explicitly listing them: `return`
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return // Implicitly returns quotient and remainder
}

// compare returns a string describing the comparison result, or an error if equal
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("unable to compare: values are equal")
	}
}
