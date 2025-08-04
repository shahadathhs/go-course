// data_types.go
// Package basic demonstrates Go’s built-in data types,
// including primitive, composite, derived, and user-defined types.
package basic

import (
	"fmt"
	"time"
)

// Custom types and type aliases
// Type alias
type MyInt = int

// New named type
type Currency float64

// Struct type
type Person struct {
	Name    string
	Age     int
	Married bool
}

// Interface type
type Stringer interface {
	String() string
}

// Implement Stringer for Person
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %q, Age: %d, Married: %t}", p.Name, p.Age, p.Married)
}

// Function type
type Adder func(int, int) int

// DataTypes prints examples of different Go data types
func DataTypes() {
	// ----- Primitive types -----
	var b bool = true
	var s string = "Hello, Go!"
	var i int = -42
	var ui uint = 42
	var f32 float32 = 3.14
	var f64 float64 = 2.718281828459045
	var c64 complex64 = 1 + 2i
	var c128 complex128 = complex(3, 4)

	fmt.Println("Boolean:", b)
	fmt.Println("String:", s)
	fmt.Println("Integers:", i, ui)
	fmt.Println("Floats:", f32, f64)
	fmt.Println("Complex:", c64, c128)

	// Zero values
	var zInt int
	var zStr string
	var zBool bool
	fmt.Println("\nZero values -> int:", zInt, "; string:", zStr, "; bool:", zBool)

	// ----- Composite types -----
	// Array
	arr := [3]int{1, 2, 3}
	// Slice
	slice := []string{"a", "b", "c"}
	// Map
	m := map[string]int{"x": 10, "y": 20}
	// Struct
	person := Person{Name: "Alice", Age: 30, Married: true}
	// Pointer
	pi := 3.14
	ptr := &pi
	// Function value
	add := Adder(func(x, y int) int { return x + y })

	fmt.Println("\nArray:", arr)
	fmt.Println("Slice:", slice)
	fmt.Println("Map:", m)
	fmt.Println("Struct:", person)
	fmt.Println("Pointer:", ptr, "points to", *ptr)
	fmt.Println("Function add(3,4):", add(3, 4))

	// ----- Other types -----
	// Time (struct)
	now := time.Now()
	// Channel (nil channel)
	var ch chan int

	fmt.Println("\nTime:", now.Format(time.RFC3339))
	fmt.Println("Channel (nil):", ch)

	// ----- Type conversions -----
	var x int = 100
	var y float64 = float64(x)
	fmt.Println("\nConvert int to float64:", y)

	// ----- Constants -----
	const Pi = 3.1415926535
	const Truth = true
	fmt.Println("\nConstants Pi:", Pi, "Truth:", Truth)

	// ----- Type assertion and empty interface -----
	var any interface{} = "sample"
	str, ok := any.(string)
	fmt.Println("\nEmpty interface holds ->", str, "; asserted string?", ok)

	// ----- Demonstrate interface implementation -----
	var s1 Stringer = person
	fmt.Println("\nInterface Stringer ->", s1.String())
}
