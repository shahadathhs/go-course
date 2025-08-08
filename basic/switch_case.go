// switch_case.go
// Package basic demonstrates "switch-case" usage in Go
package basic

import "fmt"

func SwitchCase() {
	fmt.Println("--- Basic Switch Statement ---")
	fruit := "pineapple"
	switch fruit {
	case "apple":
		fmt.Println("It's an apple.")
	case "banana":
		fmt.Println("It's a banana.")
	default:
		fmt.Println("Unknown fruit!")
	}

	fmt.Println("\n--- Switch with Multiple Conditions ---")
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend.")
	default:
		fmt.Println("Invalid day.")
	}

	fmt.Println("\n--- Switch Without Expression (Conditional Cases) ---")
	number := 15
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	fmt.Println("\n--- Switch with fallthrough ---")
	num := 2
	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not Two")
	}

	fmt.Println("\n--- Type Switch ---")
	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int, int32:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown type")
	}
}
