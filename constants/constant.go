// Package declaration — every Go source file must belong to a package
package main

// Import "fmt" for formatted output functions like Println
import "fmt"

// main is the program entry point
func main() {

	// ── Single Constant Declaration ────────────────────────────────────
	// 'const' declares a constant — its value CANNOT be changed after declaration
	// Unlike variables, constants must have a value at declaration time
	// Here: name is of type 'string' and holds the value "Golang"
	const name string = "Golang"
	fmt.Println(name) // Output: Golang

	// ── Multiple Constants Block ───────────────────────────────────────
	// You can group multiple constants using a const block with parentheses
	// This is cleaner than writing 'const' for each one individually
	const (
		hello = "Anurag"  // 'hello' is an untyped string constant — Go infers the type
		buddy = "someone" // 'buddy' is also an untyped string constant
	)

	// Printing both constants from the block
	fmt.Println(hello) // Output: Anurag
	fmt.Println(buddy) // Output: someone
}

// ────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Using constants for mathematical and config values
// ────────────────────────────────────────────────────────────────────────────
// func main() {
// 	// Mathematical constant
// 	const Pi float64 = 3.14159
//
// 	// App-level config constants grouped in a block
// 	const (
// 		AppName    = "MyGoApp"     // untyped string constant
// 		MaxRetries = 3             // untyped integer constant
// 		IsDebug    = true          // untyped boolean constant
// 	)
//
// 	radius := 5.0
// 	area := Pi * radius * radius   // constants can be used in expressions
//
// 	fmt.Println("App:", AppName)       // Output: App: MyGoApp
// 	fmt.Println("Max Retries:", MaxRetries) // Output: Max Retries: 3
// 	fmt.Println("Debug Mode:", IsDebug)    // Output: Debug Mode: true
// 	fmt.Println("Circle Area:", area)      // Output: Circle Area: 78.53975
//
// 	// NOTE: constants cannot be reassigned — this would cause a compile error:
// 	// Pi = 3.0  ← ERROR: cannot assign to Pi (declared const)
// }