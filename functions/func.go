// Package declaration — every Go file must start with a package name
package main

// Import "fmt" for formatted I/O (Println, Printf, etc.)
import "fmt"

// ── Regular Function ──────────────────────────────────────────────────────────
// 'func' keyword defines a function
// 'add' is the function name
// Parameters: 'a int' and 'b int' — two integer inputs
// Return type: 'int' — this function returns a single integer
func add(a int, b int) int {
	return a + b // add a and b and return the result
}

// ── Multiple Return Values ─────────────────────────────────────────────────────
// Go natively supports returning multiple values from a single function
// The return types are listed in parentheses: (string, string, string)
func getLanguage() (string, string, string) {
	return "golang", "java", "C++" // returns three string values at once
}

// ── Function That Returns a Function ──────────────────────────────────────────
// 'processIt' is a higher-order function — it returns another function
// Return type: func(a int) int — a function that takes an int and returns an int
func processIt() func(a int) int {
	// The return value is an anonymous (unnamed) function (also called a lambda/closure)
	return func(a int) int {
		return 4 // this inner function always returns 4 (simplified example)
	}
}

// main is the program's entry point
func main() {

	// ── Calling add() ──────────────────────────────────────────────────────
	// Call add with arguments 45 and 56; store the returned int in 'res'
	res := add(45, 56)
	fmt.Println(res) // Output: 101

	// ── Capturing Multiple Return Values ──────────────────────────────────
	// Assign all three returned strings to three separate variables
	lang1, lang2, lang3 := getLanguage()
	// TIP: Use '_' (blank identifier) to discard a return value you don't need:
	// lang1, lang2, _ := getLanguage()  ← ignores the third value (C++)
	fmt.Println(lang1) // Output: golang
	fmt.Println(lang2) // Output: java
	fmt.Println(lang3) // Output: C++

	// ── Calling the Higher-Order Function ─────────────────────────────────
	// processIt() returns a function — here we call it but don't capture the result
	// You could also do: fn := processIt(); result := fn(10)
	processIt()
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Using functions in a real-world mini calculator
// ─────────────────────────────────────────────────────────────────────────────
// // multiply takes two int32 values and returns their product
// func multiply(a, b int) int {
// 	return a * b
// }
//
// // divide returns both the quotient and remainder (multiple return values)
// func divide(a, b int) (int, int) {
// 	return a / b, a % b
// }
//
// // applyOp takes a function as a parameter (function as first-class value)
// func applyOp(a, b int, op func(int, int) int) int {
// 	return op(a, b)
// }
//
// func main() {
// 	sum := add(10, 20)                        // Output: 30
// 	product := multiply(4, 5)                 // Output: 20
// 	quotient, remainder := divide(17, 5)      // Output: 3, 2
//
// 	fmt.Println("Sum:", sum)
// 	fmt.Println("Product:", product)
// 	fmt.Println("Quotient:", quotient, "Remainder:", remainder)
//
// 	// Pass 'add' as a function argument to applyOp
// 	result := applyOp(6, 7, add)
// 	fmt.Println("Applied Op Result:", result) // Output: 13
// }
