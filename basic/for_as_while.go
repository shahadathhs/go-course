// for_as_while.go
// Demonstrates how to use a Go "for" loop as a "while" loop
package basic

import "fmt"

func ForAsWhile() {
	fmt.Println("--- Basic While-like Loop ---")
	counter := 0
	for counter < 5 {
		fmt.Println("Counter:", counter)
		counter++
	}

	fmt.Println("\n--- While-like Loop with Break ---")
	sum := 0
	for {
		sum += 10
		fmt.Println("Sum:", sum)
		if sum >= 50 {
			break
		}
	}

	fmt.Println("\n--- While-like Loop with Continue ---")
	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Number:", num)
		num++
	}

	fmt.Println("\n--- Countdown using While-like Loop ---")
	countdown := 5
	for countdown > 0 {
		fmt.Println("Countdown:", countdown)
		countdown--
	}
	fmt.Println("Liftoff!")
}
