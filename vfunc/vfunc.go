// Variadic Functions in Go
// A variadic function can accept a variable number of arguments of the same type
// The '...' (three dots) syntax is used to declare and call variadic functions
package main

// Import "fmt" for console output
import "fmt"

// ── Variadic Function ──────────────────────────────────────────────────────────
// 'sum' accepts any number of int arguments using '...int'
// Inside the function, 'a' is treated as a regular slice of ints ([]int)
// You can pass 0, 1, or many arguments when calling this function
func sum(a ...int) int {
	total := 0 // initialize total to 0 before accumulation

	// Range over 'a' (which behaves as []int inside the function)
	// '_' discards the index; 'v' holds the current element's value
	for _, v := range a {
		total += v // add each element to the running total
	}

	return total // return the final accumulated total
}

// main is the program entry point
func main() {
	// ── Pass Individual Arguments ──────────────────────────────────────────
	// All arguments are passed directly — Go collects them into a slice internally
	fmt.Println(sum(1, 2, 3, 4, 5)) // Output: 15 (1+2+3+4+5)

	// ── Pass a Slice Using Spread Operator ─────────────────────────────────
	// 'nums...' UNPACKS the slice — spreads its elements as individual arguments
	// Without '...', you'd get a type mismatch error ([]int vs ...int)
	nums := []int{4, 5, 6}
	fmt.Println(sum(nums...)) // Output: 15 (4+5+6)
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Variadic join function and mixed usage
// ─────────────────────────────────────────────────────────────────────────────
// import (
// 	"fmt"
// 	"strings"
// )
//
// // joinWords joins any number of strings with a separator
// func joinWords(sep string, words ...string) string {
// 	return strings.Join(words, sep)  // 'words' is a []string inside the function
// }
//
// // maxVal returns the largest of any number of int arguments
// func maxVal(nums ...int) int {
// 	max := nums[0]    // assume first is the largest
// 	for _, n := range nums {
// 		if n > max {
// 			max = n
// 		}
// 	}
// 	return max
// }
//
// func main() {
// 	// Call with individual string arguments
// 	result := joinWords(", ", "Go", "Python", "Java", "Rust")
// 	fmt.Println(result) // Output: Go, Python, Java, Rust
//
// 	// Call with a pre-existing slice using spread operator
// 	languages := []string{"HTML", "CSS", "JavaScript"}
// 	fmt.Println(joinWords(" | ", languages...)) // Output: HTML | CSS | JavaScript
//
// 	// Find max value
// 	fmt.Println(maxVal(3, 9, 1, 7, 5)) // Output: 9
//
// 	// Spread a slice into maxVal
// 	scores := []int{88, 42, 95, 67}
// 	fmt.Println(maxVal(scores...)) // Output: 95
// }
