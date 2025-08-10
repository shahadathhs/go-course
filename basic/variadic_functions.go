// Package basic demonstrates variadic functions in Go, including syntax, usage, and examples.
package basic

import "fmt"

func VariadicFunctions() {
	/*
		📌 VARIADIC FUNCTION SYNTAX:
			func functionName(param1 type1, param2 ...type2) returnType {
				// param2 is a slice of type2
			}

		- The `...` before the last parameter type means the function accepts zero or more arguments of that type.
		- Inside the function, `param2` is treated as a slice.
		- You can pass a slice by expanding it with `sliceVar...` when calling the function.
	*/

	// Calling variadic function with individual ints
	sequence, total := sum(1, 20, 30, 40, 50, 60)
	fmt.Println("Sequence:", sequence, "Total:", total)

	sequence2, total2 := sum(2, 40, 36, 40, 50, 60)
	fmt.Println("Sequence:", sequence2, "Total:", total2)

	// Calling variadic function by passing a slice expanded with ...
	numbers := []int{1, 2, 3, 4, 5, 9}
	sequence3, total3 := sum(3, numbers...)
	fmt.Println("Sequence:", sequence3, "Total:", total3)
}

// sum takes an int `sequence` and a variadic number of ints `nums`
// Returns the sequence and the sum of nums
func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
