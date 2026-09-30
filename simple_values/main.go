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