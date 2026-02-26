// Every Go file must belong to a package. 'main' is the entry-point package.
package main

// Import the "fmt" package which provides formatted I/O functions (like Println, Printf, etc.)
import "fmt"

// main() is the entry-point function — Go starts executing your program from here
func main() {
	// fmt.Println prints the given string to the console followed by a newline
	fmt.Println("Hello World")
}

// ─────────────────────────────────────────────
// EXAMPLE: Printing different types of values
// ─────────────────────────────────────────────
// func main() {
// 	fmt.Println("Hello, Gopher!")   // prints a string
// 	fmt.Println(42)                  // prints an integer
// 	fmt.Println(3.14)                // prints a float
// 	fmt.Println(true)                // prints a boolean
// 	fmt.Printf("Name: %s, Age: %d\n", "Anurag", 25) // formatted print
// }

// ─────────────────────────────────────────────
// HOW TO RUN / BUILD
// ─────────────────────────────────────────────
// To compile the Go app into an executable binary:
//   go build filename.go

// To compile AND immediately run the Go app:
//   go run filename.go