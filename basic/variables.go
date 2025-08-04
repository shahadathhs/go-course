package basic

import "fmt"

var globalVar int= 12

func Variables() {
	var init int
	var a string = "Hello"
	b := "World"
	c := "!"
	init = 10
	fmt.Println(a+" "+b+c, init)

	// Default Values of each type
	// Numeric = 0
	var i int
	var f float64
	// Boolean = false
	var b1 bool
	// String = ""
	var s string
	// Pointer, Function, Interface, Slice, Map, Struct = nil
	fmt.Printf("%v %v %v %q\n", i, f, b1, s)

	// Scope
	{
		var x int = 10
		fmt.Println(x)
	}
	x := 11
	fmt.Println(x)
	
	fmt.Println("globalVar: ", globalVar)
	
	globalVar = 13
	fmt.Println("globalVar: ", globalVar)
	scopeExample()
}

func scopeExample() {
	var x int = 10
	fmt.Println(x)
	fmt.Println("globalVar: ", globalVar)
}
