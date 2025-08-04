package basic

import "fmt"

const pi = 3.14
const GRAVITY = 9.81

func Constants() {

	const days int = 7

	const (
		monday        = 1
		tuesday       = 2
		wednesday     = 3
		thursday  int = 4
	)

	// * using iota
	const (
		sunday = iota
		// 🔍 What is iota?
		// iota is a pre-declared identifier in Go used in constant declarations.
		// It simplifies the creation of incrementing numbers for constants.
		// It resets to 0 every time a new const block is started.
		monday2
		tuesday2
		wednesday2
		thursday2
		friday
		saturday
	)

	// apply days = 10 days = 20 days = 30 days = 40

	fmt.Println(days)
	fmt.Println(monday, tuesday, wednesday, thursday)
	fmt.Println(sunday, monday2, tuesday2, wednesday2, thursday2, friday, saturday)

	fmt.Println(pi)
	fmt.Println(GRAVITY)
}
