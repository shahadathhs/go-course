// arrays.go Package basic demonstrates array usage in Go
package basic

import "fmt"

func Arrays() {
	fmt.Println("--- Declare and Initialize Arrays ---")
	var numbers [5]int
	fmt.Println("Empty array:", numbers)

	numbers[4] = 20
	fmt.Println("After setting index 4:", numbers)

	numbers[0] = 9
	fmt.Println("After setting index 0:", numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("\n--- Access Array Elements ---")
	fmt.Println("Third element:", fruits[2])

	fmt.Println("\n--- Copying Arrays (Value Copy) ---")
	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray // value copy

	copiedArray[0] = 100
	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", copiedArray)

	fmt.Println("\n--- Iterating Over Arrays (Index & Value) ---")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index %d: %d\n", i, numbers[i])
	}

	fmt.Println("\n--- Iterating Over Arrays (Using range) ---")
	for _, v := range numbers { // underscore ignores index
		fmt.Printf("Value: %d\n", v)
	}

	fmt.Println("\n--- Array Length ---")
	fmt.Println("The length of numbers array is", len(numbers))

	fmt.Println("\n--- Comparing Arrays ---")
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{10, 2, 3}
	fmt.Println("Array1 is equal to Array2:", array1 == array2)

	fmt.Println("\n--- Multidimensional Arrays ---")
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	fmt.Println("\n--- Array Pointers (Reference to Array) ---")
	originalArray2 := [3]int{1, 2, 3}
	var pointerArray *[3]int = &originalArray2
	pointerArray[0] = 100
	fmt.Println("Original array (modified via pointer):", originalArray2)
	fmt.Println("Pointer array:", *pointerArray)

	fmt.Println("\n--- Arrays in Go ---")
	// Arrays in Go have fixed length — cannot grow or shrink.
	// Assigning an array to another copies all elements.
	// Use a pointer (*[N]T) if you want to share the same array without copying.
	// For dynamic collections, use slices instead.
}
