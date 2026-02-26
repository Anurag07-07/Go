// Package declaration — required at the top of every Go file
package main

// Import multiple packages inside a parenthesized block
import (
	"fmt"  // "fmt" provides Println, Printf, etc. for console output
	"time" // "time" provides date/time utilities like time.Now(), Weekday constants, etc.
)

// main is the entry point of the Go program
func main() {

	// ── Basic If-Else If-Else ──────────────────────────────────────────────
	// Declare variable 'temp' with short declaration; value is 15
	temp := 15

	// if: checks if temp is less than or equal to 15
	if temp <= 15 {
		fmt.Println("Cold Day")
	} else if temp >= 15 && temp <= 25 {
		// else if: checks if temp is between 15 and 25 (inclusive)
		// '&&' is the logical AND operator — both conditions must be true
		fmt.Println("Moderate Day")
	} else {
		// else: runs when none of the above conditions are true (temp > 25)
		fmt.Println("Hot Day")
	}

	// ── If with Initialization Statement ─────────────────────────────────
	// Go allows you to declare a variable directly inside the if statement
	// 'age := 15' is the init statement, scoped only to this if-else block
	// After the semicolon comes the actual condition: age >= 18
	if age := 15; age >= 18 {
		fmt.Println("Adult")
	} else {
		// age is still accessible inside the else block (same scope)
		fmt.Println("Teenager")
	}

	// NOTE: Go does NOT have a ternary operator (condition ? a : b)
	// Use a full if-else block instead

	// ── Basic Switch Statement ─────────────────────────────────────────────
	// Switch evaluates 'i' and matches it against each case
	// Unlike many languages, Go does NOT need 'break' — it stops at first match
	var i int = 1
	switch i {
	case 1:
		fmt.Println("one") // matches when i == 1
	case 2:
		fmt.Println("two") // matches when i == 2
	case 3:
		fmt.Println("three") // matches when i == 3
	default:
		// default runs when no case matches (like 'else' in if-else)
		fmt.Println("four")
	}

	// ── Switch with Multiple Conditions in One Case ────────────────────────
	// time.Now().Weekday() returns the current day of the week as a constant
	// case time.Saturday, time.Sunday: matches if today is Saturday OR Sunday
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		// Runs for Monday through Friday
		fmt.Println("Work Days")
	}

	// ── Type Switch ───────────────────────────────────────────────────────
	// A type switch checks the DYNAMIC TYPE of an interface{} value
	// Declaring 'whoAmI' as an anonymous function (inline function stored in a variable)
	// It accepts one parameter 'i' of type 'interface{}' — can hold any type
	whoAmI := func(i interface{}) {
		// i.(type) is the special syntax for type switching
		// 't' captures the actual typed value for use inside each case
		switch t := i.(type) {
		case int:
			// Runs when i holds an int value
			fmt.Println("Its an Integer")
		case string:
			// Runs when i holds a string value
			fmt.Println("Its an String")
		case bool:
			// Runs when i holds a bool value
			fmt.Println("Its an Boolean")
		default:
			// %T is the format verb that prints the type of the value
			fmt.Printf("Unknown type: %T\n", t)
		}
	}

	// Call the anonymous function with a string argument — will print "Its an String"
	whoAmI("golang")
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Grading system using if-else and switch
// ─────────────────────────────────────────────────────────────────────────────
// func main() {
// 	score := 85
//
// 	// if-else chain to determine grade
// 	if score >= 90 {
// 		fmt.Println("Grade: A")
// 	} else if score >= 75 {
// 		fmt.Println("Grade: B")   // Output: Grade: B
// 	} else if score >= 60 {
// 		fmt.Println("Grade: C")
// 	} else {
// 		fmt.Println("Grade: F")
// 	}
//
// 	// Switch to print a message about the grade
// 	grade := "B"
// 	switch grade {
// 	case "A":
// 		fmt.Println("Excellent!")
// 	case "B":
// 		fmt.Println("Good Job!")  // Output: Good Job!
// 	case "C":
// 		fmt.Println("Average")
// 	default:
// 		fmt.Println("Needs Improvement")
// 	}
//
// 	// Type switch example
// 	var val interface{} = 42  // could be any type
// 	switch v := val.(type) {
// 	case int:
// 		fmt.Println("Integer:", v)  // Output: Integer: 42
// 	case string:
// 		fmt.Println("String:", v)
// 	}
// }
