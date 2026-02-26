// Package declaration — required at the top of every Go file
package main

// Import "fmt" for formatted output functions like Println
import "fmt"

// ── Function WITHOUT Pointer (Pass by Value) ──────────────────────────────────
// This function receives a COPY of 'num' — the original variable in main() is NOT affected
// Go passes function arguments by VALUE by default (a copy is made)
func changeNum(num int) {
	num = 5                          // modifies only the local copy — original is unchanged
	fmt.Println("In ChangeNum", num) // Output: In ChangeNum 5
}

// ── Function WITH Pointer (Pass by Reference) ──────────────────────────────────
// This function receives a POINTER to int (*int) — it can modify the original variable
// '*int' means "pointer to an integer" — it holds the memory address of an int
func changeNum1(num *int) {
	*num = 5                          // '*num' DEREFERENCES the pointer — modifies the actual value at that address
	fmt.Println("In ChangeNum", *num) // '*num' reads the value AT the pointer address → Output: In ChangeNum 5
}

// main is the entry point of the program
func main() {

	// Declare and initialize variable 'num' with value 1
	// 'num' is stored at some memory address (e.g., 0xc0000b4008)
	num := 1

	// ── Call WITHOUT Pointer ───────────────────────────────────────────────
	// A COPY of 'num' is passed; changeNum modifies only its own local copy
	// The original 'num' in main() remains unchanged
	changeNum(num)
	fmt.Println("After change in main", num) // Output: After change in main 1 (unchanged!)

	// ── Call WITH Pointer ──────────────────────────────────────────────────
	// '&num' takes the memory ADDRESS of 'num' — this is a pointer to num
	// changeNum1 receives this address and can directly modify num's value
	changeNum1(&num)
	fmt.Println("After change in main", num) // Output: After change in main 5 (changed!)
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Swapping two numbers using pointers
// ─────────────────────────────────────────────────────────────────────────────
// // swap takes two int pointers and swaps the values they point to
// func swap(a *int, b *int) {
// 	temp := *a  // dereference 'a' to read its value
// 	*a = *b     // set value at 'a's address to value at 'b's address
// 	*b = temp   // set value at 'b's address to the saved temp
// }
//
// func main() {
// 	x := 10
// 	y := 20
//
// 	fmt.Println("Before swap → x:", x, "y:", y) // Output: Before swap → x: 10 y: 20
//
// 	swap(&x, &y)  // pass addresses of x and y
//
// 	fmt.Println("After swap  → x:", x, "y:", y) // Output: After swap  → x: 20 y: 10
//
// 	// ── Pointer basics ──────────────────────────────────────────────────
// 	p := &x              // p is a pointer to x (holds x's memory address)
// 	fmt.Println(*p)      // dereference p → prints the value of x → Output: 20
// 	*p = 100             // change the value at the address p points to
// 	fmt.Println(x)       // x is now 100, because p pointed to x → Output: 100
// }
