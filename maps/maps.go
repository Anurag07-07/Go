// Package declaration — required at the top of every Go source file
package main

// Import packages inside a parenthesized block
import (
	"fmt"  // "fmt" provides console I/O functions like Println, Printf
	"maps" // "maps" package (Go 1.21+) provides utility functions like maps.Equal
)

// main is the program's entry point — execution starts here
func main() {

	// ── Create a Map Using make() ──────────────────────────────────────────
	// make(map[KeyType]ValueType) creates an empty, initialized map
	// Here: keys are strings ("name", "company") and values are strings
	// Unlike arrays/slices, maps are unordered key-value pairs
	m1 := make(map[string]string)

	// ── Set Values in the Map ─────────────────────────────────────────────
	// Use m1["key"] = "value" to insert or update a key-value pair
	m1["name"] = "golang"    // adds key "name" with value "golang"
	m1["company"] = "Google" // adds key "company" with value "Google"
	fmt.Println(m1)          // Output: map[company:Google name:golang]

	// ── Fetch Values from the Map ─────────────────────────────────────────
	// Access a value by its key using m1["key"]
	fmt.Println(m1["name"]) // Output: golang

	// Accessing a key that DOES NOT exist returns the zero value for the value type:
	//   string → "" (empty string)
	//   int    → 0
	//   bool   → false
	fmt.Println(m1["data"]) // Output: "" (empty string — key "data" doesn't exist)

	// ── Map Length ───────────────────────────────────────────────────────
	// len() returns the number of key-value pairs in the map
	fmt.Println(len(m1)) // Output: 2

	// ── Delete a Key ─────────────────────────────────────────────────────
	// delete(mapName, key) removes the specified key and its value from the map
	delete(m1, "name") // removes the "name" key

	// ── Clear Entire Map ──────────────────────────────────────────────────
	// clear(mapName) removes ALL key-value pairs from the map (Go 1.21+)
	// After this, len(m1) == 0 and m1 is an empty (but not nil) map
	clear(m1)

	// ── Initialize a Map with Values in One Line ──────────────────────────
	// Declare and populate a map using map literal syntax: map[K]V{key: value, ...}
	m2 := map[string]string{"name": "Anurag", "email": "anurag@gmail.com"}
	fmt.Println(m2) // Output: map[email:anurag@gmail.com name:Anurag]

	// ── Check if a Key Exists (Two-Value Lookup) ──────────────────────────
	// When reading from a map, Go optionally returns a second boolean value 'ok'
	//   ok == true  → key was found in the map
	//   ok == false → key was NOT found
	// Use '_' to ignore a return value you don't need:
	//   _, ok := m2["name"]  ← only check existence, discard the value
	name, ok := m2["name"]
	fmt.Println(name) // Output: Anurag

	// Use the 'ok' variable to check if the key actually existed
	if ok {
		fmt.Println("User is There") // Output: User is There
	} else {
		fmt.Println("User is not there")
	}

	// ── Compare Two Maps ──────────────────────────────────────────────────
	// maps.Equal(m1, m2) returns true if both maps have the same keys and values
	// m1 was cleared (empty) and m2 has 2 keys — so they are NOT equal
	fmt.Println(maps.Equal(m1, m2)) // Output: false
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Word frequency counter using a map
// ─────────────────────────────────────────────────────────────────────────────
// func main() {
// 	// Count occurrences of each word in a sentence
// 	words := []string{"go", "is", "great", "go", "is", "fast", "go"}
//
// 	// Create a map to store word → count
// 	frequency := make(map[string]int)
//
// 	// Iterate over words and increment count for each word
// 	for _, word := range words {
// 		frequency[word]++ // if key doesn't exist, default int is 0, then +1
// 	}
//
// 	fmt.Println("Word Frequencies:", frequency)
// 	// Output: map[go:3 great:1 is:2 fast:1]
//
// 	// Check if "go" exists
// 	count, exists := frequency["go"]
// 	if exists {
// 		fmt.Println("'go' appeared", count, "times") // Output: 'go' appeared 3 times
// 	}
//
// 	// Delete a word from the map
// 	delete(frequency, "fast")
// 	fmt.Println("After delete:", frequency) // "fast" is removed
// }
