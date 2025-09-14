// pointers.go
// Package intermediate demonstrates pointers in Go.
// ------------------------------------------------------------------
// Important Notes:
// 1. A pointer holds the memory address of a value.
// 2. `&` (address-of) operator → gets the memory address of a variable.
// 3. `*` (dereference) operator → accesses the value stored at that address.
// 4. The zero value of a pointer is `nil`.
// 5. Pointers allow functions to modify values outside their local scope.
// 6. Go does not support pointer arithmetic (unlike C/C++).
// 7. Use pointers when:
//    - You want to modify a variable inside a function.
//    - You want to avoid copying large structs/arrays.
//    - You need to represent optional values (`nil` as "no value").
// ------------------------------------------------------------------

package intermediate

import "fmt"

// Pointers demonstrates different pointer use cases.
func Pointers() {
	// Example 1: Basic pointer usage
	fmt.Println("=== Example 1: Basic pointer ===")
	var a int = 10
	var ptr *int = &a // referencing
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", ptr)
	fmt.Println("Value via pointer (dereference):", *ptr)

	// Example 2: Nil pointer
	fmt.Println("\n=== Example 2: Nil pointer ===")
	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("nilPtr is nil")
	}

	// Example 3: Passing pointer to function
	fmt.Println("\n=== Example 3: Pass pointer to function ===")
	fmt.Println("Before modification, a =", a)
	modifyValue(ptr)
	fmt.Println("After modification, a =", a)

	// Example 4: Function returning a pointer
	fmt.Println("\n=== Example 4: Function returning pointer ===")
	newPtr := newInt(42)
	fmt.Println("Pointer returned:", newPtr)
	fmt.Println("Value at pointer:", *newPtr)

	// Example 5: Pointers with structs
	fmt.Println("\n=== Example 5: Pointers with structs ===")
	person := Person{name: "Alice", age: 30}
	fmt.Println("Before update:", person)
	updateAge(&person, 35)
	fmt.Println("After update:", person)

	// Example 6: Pointer to pointer (double pointer)
	fmt.Println("\n=== Example 6: Pointer to pointer ===")
	b := 5
	ptr1 := &b
	ptr2 := &ptr1
	fmt.Println("b:", b)
	fmt.Println("ptr1 (address of b):", ptr1)
	fmt.Println("ptr2 (address of ptr1):", ptr2)
	fmt.Println("Dereferencing ptr2 →", **ptr2)

	// Example 7: Array vs Slice with pointers
	fmt.Println("\n=== Example 7: Pointers with arrays and slices ===")
	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	fmt.Println("Array via pointer:", *arrPtr)

	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println("Modified slice (no explicit pointer needed):", slice)
}

// ------------------------------------------------------------------
// Example 3: Modify value using pointer
func modifyValue(ptr *int) {
	*ptr++ // dereference and increment value
}

// ------------------------------------------------------------------
// Example 4: Return pointer from function
func newInt(val int) *int {
	return &val
}

// ------------------------------------------------------------------
// Example 5: Struct pointer example
type Person struct {
	name string
	age  int
}

func updateAge(p *Person, newAge int) {
	p.age = newAge
}

// ------------------------------------------------------------------
// Example 7: Modify slice
// Note: slices already behave like pointers (reference type).
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 99
	}
}
