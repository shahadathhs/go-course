// for_loop.go
// Package basic demonstrates "for" loops in Go
package basic

import "fmt"

func ForLoop() {
	fmt.Println("--- Basic Counting Loop ---")
	// Traditional for-loop (like C/Java)
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	fmt.Println("\n--- For as While ---")
	// for as a while loop
	j := 0
	for j < 3 {
		fmt.Println("j:", j)
		j++
	}

	fmt.Println("\n--- Infinite Loop ---")
	// Infinite loop with break
	count := 0
	for {
		fmt.Println("count:", count)
		count++
		if count == 3 {
			break
		}
	}

	fmt.Println("\n--- Loop with Continue and Break ---")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println("Odd:", i)
		if i == 7 {
			break // stop when i == 7
		}
	}

	fmt.Println("\n--- Range Over Slice ---")
	slice := []string{"go", "is", "awesome"}
	for idx, val := range slice {
		fmt.Printf("Index: %d, Value: %s\n", idx, val)
	}

	fmt.Println("\n--- Range Over Map ---")
	myMap := map[string]int{"apple": 2, "banana": 3}
	for key, val := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, val)
	}

	fmt.Println("\n--- Nested Loops (Pattern) ---")
	rows := 4
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println("\n--- Go 1.22: Loop Over Range Literal ---")
	// Requires Go 1.22+
	for i := range 5 {
		fmt.Println("i (range 5):", i)
	}

	fmt.Println("\n--- Loop Over String (Unicode) ---")
	text := "Gölang"
	for idx, ch := range text {
		fmt.Printf("Index: %d, Rune: %c\n", idx, ch)
	}
}
