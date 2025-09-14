package intermediate

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func StringsAndRunes() {
	// ==============================
	// Strings and Runes in Go
	// ==============================

	// --- What is a String? ---
	// 1. A string in Go is a sequence of bytes (immutable).
	// 2. It can store ASCII or UTF-8 encoded Unicode characters.
	// 3. Strings are immutable → modifying them creates a new string.
	// 4. Strings can be created using double quotes "" (interpreted) or backticks `` (raw, multiline).

	// --- What is a Rune? ---
	// 1. A rune is an alias for int32 and represents a single Unicode code point.
	// 2. Runes are useful for working with Unicode characters, which may require multiple bytes.
	// 3. Rune literals are enclosed in single quotes, e.g., 'a', '日', '😊'.

	// === 1. String Literals ===
	// - Double quotes ("") → interpreted string (escapes like \n, \t work)
	// - Backticks (``) → raw string (no escaping, multiline allowed)
	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, Go!"
	message3 := "Hello, \rGo!"
	rawMessage := `Hello\nGo` // raw string (keeps \n literally)

	fmt.Println("== String Literals ==")
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)
	fmt.Println(message3)

	// === 2. Length of Strings ===
	// - len() → counts BYTES, not runes (Unicode chars may use multiple bytes)
	fmt.Println("\n== Lengths ==")
	fmt.Println("len(rawMessage):", len(rawMessage))
	fmt.Println("Rune count (utf8):", utf8.RuneCountInString(rawMessage))

	// === 3. Indexing ===
	// Strings are byte slices → indexing gives a byte (uint8), not rune
	fmt.Println("\n== Indexing ==")
	fmt.Println("First byte in message:", message[0]) // prints ASCII code
	fmt.Printf("As character: %c\n", message[0])

	// === 4. Concatenation ===
	// Using + operator creates a NEW string (strings are immutable)
	greeting := "Hello "
	name := "Alice"
	fmt.Println("\n== Concatenation ==")
	fmt.Println("greeting + name:", greeting+name)

	// === 5. Comparison ===
	// Strings are compared lexicographically (byte by byte, Unicode aware)
	str1 := "Apple"  // 'A' = 65
	str := "apple"   // 'a' = 97
	str2 := "banana" // 'b' = 98
	str3 := "app"
	fmt.Println("\n== Comparison ==")
	fmt.Println("str1 < str2:", str1 < str2) // true
	fmt.Println("str3 < str1:", str3 < str1) // true
	fmt.Println("str > str1:", str > str1)   // true
	fmt.Println("str > str3:", str > str3)   // true

	// === 6. Iteration ===
	// - Range loop → iterates over RUNES (Unicode aware)
	// - Classic loop with index → iterates BYTES
	fmt.Println("\n== Iteration ==")
	fmt.Println("By bytes:")
	for i := 0; i < len(message); i++ {
		fmt.Printf("Index %d: %v (%c)\n", i, message[i], message[i])
	}

	fmt.Println("By runes:")
	for i, char := range message {
		fmt.Printf("Index %d: %U '%c'\n", i, char, char)
	}

	// === 7. Rune Count ===
	fmt.Println("\n== Rune Count ==")
	fmt.Println("Rune count in greeting:", utf8.RuneCountInString(greeting))

	// === 8. Rune Literals ===
	// - Rune = alias for int32 (represents a Unicode code point)
	var ch rune = 'a'
	jch := '日'
	fmt.Println("\n== Rune Literals ==")
	fmt.Println("Rune 'a':", ch)
	fmt.Println("Rune '日':", jch)
	fmt.Printf("As char: %c, %c\n", ch, jch)

	// === 9. Rune → String Conversion ===
	cstr := string(ch) // convert rune to string
	fmt.Println("\n== Rune to String ==")
	fmt.Println("Rune to string:", cstr)
	fmt.Printf("Type of cstr: %T\n", cstr)

	// === 10. Unicode Strings ===
	const NIHONGO = "日本語" // Japanese text
	fmt.Println("\n== Unicode Strings ==")
	fmt.Println("NIHONGO:", NIHONGO)

	jhello := "こんにちは" // Japanese "Hello"
	for _, runeValue := range jhello {
		fmt.Printf("%c ", runeValue)
	}
	fmt.Println()

	// === 11. Emojis (multi-byte runes) ===
	r := '😊'
	fmt.Println("\n== Emojis ==")
	fmt.Printf("Rune value: %v\n", r)   // int32 code point
	fmt.Printf("Rune as char: %c\n", r) // 😊
	fmt.Printf("Rune as quoted: %q\n", r)

	// === 12. Strings package functions ===
	// strings package provides utility functions
	s := "Go is awesome!"
	fmt.Println("\n== strings package ==")
	fmt.Println("Contains 'Go':", strings.Contains(s, "Go"))
	fmt.Println("HasPrefix 'Go':", strings.HasPrefix(s, "Go"))
	fmt.Println("HasSuffix 'awesome!':", strings.HasSuffix(s, "awesome!"))
	fmt.Println("Index of 'is':", strings.Index(s, "is"))
	fmt.Println("Replace:", strings.ReplaceAll(s, "awesome", "great"))
	fmt.Println("ToUpper:", strings.ToUpper(s))
	fmt.Println("ToLower:", strings.ToLower(s))
	fmt.Println("Split:", strings.Split(s, " "))
	fmt.Println("Join:", strings.Join([]string{"Go", "is", "fun"}, "-"))

	// === 13. Immutability ===
	// Strings cannot be changed → create a new one instead
	original := "golang"
	// original[0] = 'G' // ❌ compile error
	changed := "G" + original[1:] // ✅ new string
	fmt.Println("\n== Immutability ==")
	fmt.Println("Original:", original)
	fmt.Println("Changed:", changed)

	// === 14. String <-> []byte Conversion ===
	// Useful for file I/O and performance
	fmt.Println("\n== Conversion between string and []byte ==")
	b := []byte("Hello") // string → []byte
	fmt.Println("Bytes:", b)

	backToStr := string(b) // []byte → string
	fmt.Println("Back to string:", backToStr)
}
