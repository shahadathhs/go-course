// Package basic demonstrates the usage of `range` in Go
package basic

import "fmt"

func Range() {

	// --- 1. RANGE OVER STRING ---
	message := "Hello World"

	// When ranging over a string, Go iterates over runes (Unicode code points)
	// i  → byte index of the rune in the string
	// v  → rune (Unicode code point) at that index
	for i, v := range message {
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

	// --- 2. RANGE OVER STRING WITH MULTI-BYTE RUNES ---
	// This shows how range handles Unicode characters (multi-byte)
	message2 := "Golang ❤"
	fmt.Println("\nString with Unicode:")
	for i, r := range message2 {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
	}

	// --- 3. RANGE OVER SLICE ---
	nums := []int{10, 20, 30}
	fmt.Println("\nRange over slice:")
	for i, val := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, val)
	}

	// --- 4. RANGE OVER ARRAY ---
	arr := [3]string{"apple", "banana", "cherry"}
	fmt.Println("\nRange over array:")
	for i, val := range arr {
		fmt.Printf("Index: %d, Value: %s\n", i, val)
	}

	// --- 5. RANGE OVER MAP ---
	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("\nRange over map (order is random):")
	for key, val := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, val)
	}

	// --- 6. RANGE OVER CHANNEL ---
	fmt.Println("\nRange over channel:")
	ch := make(chan string)
	go func() {
		ch <- "first"
		ch <- "second"
		close(ch) // Must close channel for range to end
	}()
	for msg := range ch {
		fmt.Println("Received:", msg)
	}

	// --- 7. IGNORING INDEX OR VALUE ---
	fmt.Println("\nIgnoring values in range:")

	// Ignore value, keep index
	for i := range nums {
		fmt.Println("Index only:", i)
	}

	// Ignore index, keep value
	for _, val := range nums {
		fmt.Println("Value only:", val)
	}

	// --- 8. GOTCHA: STRING INDEXES ---
	// i is byte index, not character position if runes are multi-byte
	word := "Go❤"
	fmt.Println("\nString index details:")
	for i, r := range word {
		fmt.Printf("Byte Index: %d, Rune: %c\n", i, r)
	}
}
