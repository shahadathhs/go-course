// if_else.go
// Package basic demonstrates "if-else" statements in Go
package basic

import "fmt"

func IfElse() {
	fmt.Println("--- Simple if Statement ---")
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	fmt.Println("\n--- if-else Statement ---")
	num := 7
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	fmt.Println("\n--- if-else if Ladder ---")
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	fmt.Println("\n--- if with Short Statement ---")
	if length := len("Golang"); length > 5 {
		fmt.Println("Long word")
	}

	fmt.Println("\n--- Nested if Statement ---")
	temp := 30
	if temp > 25 {
		fmt.Println("It's hot")
		if temp > 35 {
			fmt.Println("Stay hydrated!")
		}
	}
}
