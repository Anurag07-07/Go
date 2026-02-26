// Package declaration — every Go file must start with a package name
package main

// Import "fmt" package to use Println for printing output to the console
import "fmt"

// main is the entry point of the program — execution begins here
func main() {
	// ── Integer ──────────────────────────────────────────
	// Prints the result of arithmetic expression 1+2 (= 3)
	fmt.Println(1 + 2)

	// ── String ───────────────────────────────────────────
	// Prints a string literal "Hello Golang" to the console
	fmt.Println("Hello Golang")

	// ── Boolean ──────────────────────────────────────────
	// Prints the boolean literal 'true' to the console
	fmt.Println(true)

	// ── Float ────────────────────────────────────────────
	// Prints a floating-point number 10.5 directly
	fmt.Println(10.5)

	// Prints the result of dividing two floats: 10.5 / 5.5 ≈ 1.909090...
	fmt.Println(10.5 / 5.5)
}

// ─────────────────────────────────────────────────────────
// EXAMPLE: Simple values with variables for clarity
// ─────────────────────────────────────────────────────────
// func main() {
// 	age := 25               // integer value
// 	name := "Anurag"        // string value
// 	isStudent := false      // boolean value
// 	gpa := 8.9              // float value
//
// 	fmt.Println(age)        // Output: 25
// 	fmt.Println(name)       // Output: Anurag
// 	fmt.Println(isStudent)  // Output: false
// 	fmt.Println(gpa)        // Output: 8.9
// 	fmt.Println(age + 5)    // Output: 30  (integer arithmetic)
// 	fmt.Println(gpa * 2)    // Output: 17.8 (float arithmetic)
// }
