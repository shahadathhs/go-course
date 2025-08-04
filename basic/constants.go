package basic

import (
	"fmt"
	"math"
)

// Package-level constants:
const (
	// Numeric (untyped and typed)
	pi               = 3.14159       // untyped floating-point
	Gravity  float64 = 9.80665       // typed floating-point
	EarthAge         = 4.543e9       // untyped scientific notation
	MaxInt32         = math.MaxInt32 // constant from another package

	// Boolean
	Enabled  = true
	Disabled = false

	// String
	Language = "Go"
	Version  = "1.20"
)

// Grouped constants with explicit types and defaults:
const (
	dayCount int = 7 // typed
	// next constants inherit type int
	hourCount   = 24
	minuteCount = 60
	secondCount = 60
)

// Simple weekday enumeration with iota:
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// Skipping values and custom start:
const (
	_     = iota       // skip zero
	CodeA              // 1
	_                  // skip 2
	CodeB              // 3
	Start = iota + 100 // 4 + 100 = 104
)

// Bit-mask flags using iota and shifting:
const (
	FlagNone    = 0
	FlagRead    = 1 << iota // 1 << 1 = 2
	FlagWrite               // 1 << 2 = 4
	FlagExecute             // 1 << 3 = 8
	FlagAll     = FlagRead | FlagWrite | FlagExecute
)

// Constant expressions:
const (
	CircleAreaFactor = pi * 2 // untyped numeric expression
	SquareSide       = 5
	SquareArea       = SquareSide * SquareSide
)

// You can’t change a constant at run-time — attempt to do so is a compile-time error.

func Constants() {
	fmt.Println("---- Numeric Constants ----")
	fmt.Println("pi           =", pi)
	fmt.Println("Gravity      =", Gravity)
	fmt.Printf("EarthAge     = %.2e years\n", EarthAge)
	fmt.Println("MaxInt32     =", MaxInt32)

	fmt.Println("\n---- Boolean Constants ----")
	fmt.Println("Enabled      =", Enabled)
	fmt.Println("Disabled     =", Disabled)

	fmt.Println("\n---- String Constants ----")
	fmt.Println("Language     =", Language)
	fmt.Println("Version      =", Version)

	fmt.Println("\n---- Time Units ----")
	fmt.Println("Days         =", dayCount)
	fmt.Println("Hours in day =", hourCount)
	fmt.Println("Minutes      =", minuteCount)
	fmt.Println("Seconds      =", secondCount)

	fmt.Println("\n---- Weekday Enumeration ----")
	fmt.Println("Sunday       =", Sunday)
	fmt.Println("Monday       =", Monday)
	fmt.Println("... Saturday =", Saturday)

	fmt.Println("\n---- Custom iota Skips ----")
	fmt.Println("CodeA        =", CodeA)
	fmt.Println("CodeB        =", CodeB)
	fmt.Println("Start        =", Start)

	fmt.Println("\n---- Bit-mask Flags ----")
	fmt.Printf("FlagRead     = %b\n", FlagRead)
	fmt.Printf("FlagWrite    = %b\n", FlagWrite)
	fmt.Printf("FlagExecute  = %b\n", FlagExecute)
	fmt.Printf("FlagAll      = %b\n", FlagAll)

	fmt.Println("\n---- Constant Expressions ----")
	fmt.Println("CircleAreaFactor =", CircleAreaFactor)
	fmt.Println("SquareArea       =", SquareArea)
}
