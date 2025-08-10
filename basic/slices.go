// slices.go Package basic demonstrates slice usage in Go
package basic

import (
	"fmt"
	"slices"
)

func Slices() {

	fmt.Println("--- Declare and Initialize Slices ---")
	var emptySlice []int // nil slice
	fmt.Println("Empty slice:", emptySlice, "len:", len(emptySlice), "cap:", cap(emptySlice))

	numbers := []int{1, 2, 3}
	fmt.Println("Literal slice:", numbers, "len:", len(numbers), "cap:", cap(numbers))

	sliceMake := make([]int, 5)       // len=5, cap=5
	sliceMakeCap := make([]int, 5, 8) // len=5, cap=8
	fmt.Println("Make slice:", sliceMake, "len:", len(sliceMake), "cap:", cap(sliceMake))
	fmt.Println("Make slice w/capacity:", sliceMakeCap, "len:", len(sliceMakeCap), "cap:", cap(sliceMakeCap))

	fmt.Println("\n--- Slice from an Array ---")
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4] // index 1 to 3
	fmt.Println("Slice1 from array:", slice1, "len:", len(slice1), "cap:", cap(slice1))
	// Capacity is measured from the starting index to the end of the underlying array

	fmt.Println("\n--- Appending to a Slice ---")
	slice1 = append(slice1, 6)
	fmt.Println("After appending 6:", slice1, "len:", len(slice1), "cap:", cap(slice1))

	sliceGrow := []int{1, 2}
	fmt.Println("Before growth:", sliceGrow, "len:", len(sliceGrow), "cap:", cap(sliceGrow))
	sliceGrow = append(sliceGrow, 3, 4, 5) // exceeds capacity, triggers new array allocation
	fmt.Println("After growth:", sliceGrow, "len:", len(sliceGrow), "cap:", cap(sliceGrow))

	fmt.Println("\n--- Copying Slices ---")
	copySlice := make([]int, len(slice1))
	copy(copySlice, slice1)
	fmt.Println("Copied slice:", copySlice)

	fmt.Println("\n--- Iterating Over Slices ---")
	for i, v := range slice1 {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	fmt.Println("\n--- Accessing Elements ---")
	fmt.Println("Element at index 2 of slice1:", slice1[2])

	fmt.Println("\n--- Comparing Slices ---")
	if slices.Equal(slice1, copySlice) {
		fmt.Println("slice1 is equal to copySlice")
	} else {
		fmt.Println("slice1 is NOT equal to copySlice")
	}

	fmt.Println("\n--- Multidimensional Slices ---")
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	fmt.Println("\n--- Slice Expressions ---")
	fmt.Println("slice1[1:]:", slice1[1:])
	fmt.Println("slice1[:2]:", slice1[:2])
	fmt.Println("slice1[:]:", slice1[:])

	fmt.Println("\n--- Array vs Slice ---")
	arrExample := [3]int{1, 2, 3}
	sliceFromArr := arrExample[:]
	sliceFromArr[0] = 99
	fmt.Println("Array after modifying slice:", arrExample)

	fmt.Println("\n--- Slices in Go ---")
	// Slices are reference types — they point to an underlying array.
	// len = current number of elements accessible.
	// cap = elements from starting index to end of underlying array.
	// Appending beyond capacity causes allocation of a new underlying array.
	// Assigning one slice to another makes them share the same array data.
	// Arrays are fixed-length; slices are dynamic and more flexible.
}
