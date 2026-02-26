// Package declaration — must be present in every Go file
package main

// Import "fmt" to use Println for displaying output
import "fmt"

// main is the program's entry point
func main() {

	// ── Style 1: Variable with Explicit Type ──────────────────────────────
	// 'var' keyword declares a variable, 'name' is the variable name,
	// 'string' is the type, and "golang" is the value assigned to it
	var name string = "golang"
	fmt.Println(name) // Output: golang

	// ── Style 2: Variable with Type Inference (no explicit type) ──────────
	// Go automatically infers the type from the value — here it's bool (true)
	var isActive = true
	fmt.Println(isActive) // Output: true

	// ── Style 3: Short Variable Declaration (Walrus Operator :=) ─────────
	// := declares AND assigns in one step; Go infers the type automatically
	// Most commonly used inside functions — cannot be used at package level
	first_name := "Anurag"
	fmt.Println(first_name) // Output: Anurag

	// ── Style 4: Declare first, assign later ──────────────────────────────
	// When only declaring without a value, you MUST specify the type
	// The zero value for string is "" (empty string)
	var company string   // declared but not yet assigned
	company = "Google"   // now assigned a value
	fmt.Println(company) // Output: Google

	// ── Style 5: Float variable with explicit type ─────────────────────────
	// float32 is a 32-bit floating-point number; float64 is also available (more precise)
	var price float32 = 45.23
	fmt.Println(price) // Output: 45.23
}

// ──────────────────────────────────────────────────────────────────────
// EXAMPLE: Declaring multiple variables of different types
// ──────────────────────────────────────────────────────────────────────
// func main() {
// 	var city string = "Mumbai"     // string variable with explicit type
// 	population := 20000000         // int variable using short declaration
// 	var isCapital bool = false     // boolean variable
// 	var area float64 = 603.4       // float64 variable (more precision)
//
// 	fmt.Println(city)              // Output: Mumbai
// 	fmt.Println(population)        // Output: 20000000
// 	fmt.Println(isCapital)         // Output: false
// 	fmt.Println(area)              // Output: 603.4
//
// 	// You can also declare multiple variables at once using 'var' block:
// 	var (
// 		x int     = 10
// 		y float32 = 3.14
// 		z string  = "Go"
// 	)
// 	fmt.Println(x, y, z)           // Output: 10 3.14 Go
// }
