// Package basic demonstrates a complete reference for working with maps in Go.
package basic

import (
	"fmt"
	"maps" // New in Go 1.21 — provides helper functions for maps
)

func Maps() {

	// --- 1. DECLARATION & INITIALIZATION ---

	// Declare a map without initializing — this creates a nil map (cannot be written to yet)
	var nilMap map[string]int
	fmt.Println("nilMap:", nilMap) // nil
	// nilMap["key"] = 1 // ❌ This will panic: assignment to entry in nil map

	// Initialize an empty map using make (ready to use)
	myMap := make(map[string]int) // map[keyType]valueType
	fmt.Println("Empty initialized map:", myMap)

	// Map literal initialization with values
	myMapLiteral := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println("Map literal:", myMapLiteral)

	// --- 2. ADDING & UPDATING ENTRIES ---

	// Adding new keys
	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println("After adding entries:", myMap)

	// Updating an existing key overwrites the value
	myMap["code"] = 21
	fmt.Println("After updating 'code':", myMap)

	// --- 3. RETRIEVING VALUES ---

	// Get value by key
	val := myMap["key1"]
	fmt.Println("Value of 'key1':", val)

	// If key doesn't exist, returns zero-value of value type
	fmt.Println("Value of non-existent 'key':", myMap["key"]) // 0

	// --- 4. CHECKING KEY EXISTENCE ---

	// Use "comma ok" idiom to check if a key exists
	_, ok := myMap["key1"]
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value exists with key1")
	}
	fmt.Println("Does key1 exist?", ok)

	// --- 5. DELETING ENTRIES ---

	delete(myMap, "key1") // If key doesn't exist, delete does nothing
	fmt.Println("After deleting 'key1':", myMap)

	// --- 6. ITERATING OVER A MAP ---

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11

	fmt.Println("Iterating over myMap:")
	for key, value := range myMap {
		fmt.Printf("%s => %d\n", key, value)
	}

	fmt.Println("Iterating over only values:")
	for _, value := range myMap {
		fmt.Println(value)
	}

	// --- 7. CLEARING A MAP ---

	// Go 1.21+ has clear()
	// clear(myMap)
	// fmt.Println("After clear:", myMap)

	// --- 8. COMPARING MAPS ---

	myMap2 := map[string]int{"a": 1, "b": 2}
	myMap3 := map[string]int{"a": 1, "b": 2}

	// Direct comparison using == is not allowed for maps (except nil checks)
	// ✅ Use maps.Equal from "maps" package (Go 1.21+)
	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap3 and myMap2 are equal")
	}

	// --- 9. NIL MAPS VS EMPTY MAPS ---

	var myMap4 map[string]string // nil map
	if myMap4 == nil {
		fmt.Println("myMap4 is nil")
	}
	valStr := myMap4["key"] // Allowed — returns zero value ("")
	fmt.Println("Value from nil map:", valStr)

	// Initialize and add values
	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println("myMap4 after initialization:", myMap4)

	// --- 10. MAP LENGTH ---

	fmt.Println("myMap length is", len(myMap))

	// --- 11. NESTED MAPS ---

	// A map with map values
	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println("Nested map:", myMap5)
}
