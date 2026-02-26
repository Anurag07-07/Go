// Package declaration — required at the top of every Go file
package main

// Import "fmt" for console I/O functions
import "fmt"

// main is the program's entry point
func main() {

	// ── Range Over a Slice (Sum Example) ───────────────────────────────────
	// Declare and initialize 'sum' to 0 — will accumulate the total
	sum := 0

	// Declare a slice of integers inline using a slice literal
	nums := []int{1, 3, 2, 4, 5}

	// 'range nums' iterates over each element in the slice
	// In each iteration:
	//   '_' → the index (we discard it using the blank identifier '_')
	//   'num' → the actual value at that index
	for _, num := range nums {
		sum = sum + num // add each element's value to the running total
	}

	fmt.Println(sum) // Output: 15 (1+3+2+4+5)

	// ── Range Over a Map ───────────────────────────────────────────────────
	// Declare and initialize a map with string keys and string values
	user := map[string]string{"A": "Something", "B": "Something"}

	// 'range user' iterates over each key-value pair in the map
	//   'k' → the key of the current entry
	//   'v' → the value associated with that key
	// NOTE: Map iteration order is NOT guaranteed — output may vary each run
	for k, v := range user {
		fmt.Println(k, v) // Output: A Something  and  B Something (order may vary)
	}

	// ── Range Over a String ─────────────────────────────────────────────────
	// Ranging over a string gives you Unicode code points (runes), not raw bytes
	//   'i' → the byte index of the character within the string
	//   'c' → the Unicode code point (rune) of the character at that position
	for i, c := range "golang" {
		fmt.Println(i, c)      // prints byte index and Unicode code point (int32 value)
		fmt.Println(string(c)) // string(c) converts the rune back to its character representation
	}
	// Example output for 'g':
	//   0 103       ← index 0, Unicode code point of 'g' is 103
	//   g           ← the actual character
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Using range to find max value and iterate over a map of student info
// ─────────────────────────────────────────────────────────────────────────────
// func main() {
// 	// Range over a slice to find the maximum value
// 	scores := []int{88, 45, 92, 67, 77}
// 	max := scores[0] // assume first element is the max
//
// 	for _, score := range scores {
// 		if score > max {
// 			max = score // update max if we find a higher score
// 		}
// 	}
// 	fmt.Println("Max Score:", max) // Output: Max Score: 92
//
// 	// Range over a map — print student name and grade
// 	students := map[string]string{
// 		"Alice": "A",
// 		"Bob":   "B+",
// 		"Carol": "A-",
// 	}
//
// 	for name, grade := range students {
// 		fmt.Printf("Student: %-10s Grade: %s\n", name, grade)
// 	}
// 	// Output (order may vary):
// 	// Student: Alice      Grade: A
// 	// Student: Bob        Grade: B+
// 	// Student: Carol      Grade: A-
//
// 	// Range over a string to count vowels
// 	vowels := 0
// 	for _, ch := range "Hello Gopher" {
// 		switch ch {
// 		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
// 			vowels++
// 		}
// 	}
// 	fmt.Println("Vowels:", vowels) // Output: Vowels: 3
// }
