// Package declaration — required at the start of every Go file
package main

// Import "fmt" for printing output to the console
import "fmt"

// NOTE: In Go, there is ONLY the 'for' keyword for all types of loops.
// There is no 'while' or 'do-while' — 'for' handles every loop pattern.

// main is the program's entry point
func main() {

	// ── While-style Loop (condition only) ───────────────────────────────────
	// This 'for' loop behaves like a while loop in other languages
	// It runs as long as the condition (i <= 3) is true
	i := 1       // initialize counter variable before the loop
	for i <= 3 { // loop condition: continue while i is less than or equal to 3
		fmt.Println(i) // print the current value of i
		i = i + 1      // manually increment i to avoid an infinite loop
	}
	// Output: 1  2  3

	// ── Infinite Loop ──────────────────────────────────────────────────────
	// A 'for' with no condition runs forever (infinite loop)
	// 'break' immediately exits the loop — without it, this would run forever
	for {
		break // exit the loop immediately on first iteration
	}

	// ── Classic C-style For Loop ───────────────────────────────────────────
	// Syntax: for init; condition; post { }
	//   init:      i := 0   — runs once before the loop starts
	//   condition: i < 5    — checked before each iteration; loop stops when false
	//   post:      i++      — runs after each iteration (increment)
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // 'continue' skips the rest of this iteration and moves to the next
		}
		fmt.Println(i) // prints 0, 1, 3, 4 (2 is skipped due to continue)
	}

	// ── Range-based Loop ────────────────────────────────────────────────────
	// 'range 3' generates values 0, 1, 2 (the upper bound 3 is EXCLUDED)
	// This is a shorthand — 'i' holds the current iteration index
	for i := range 3 { // i goes: 0, 1, 2 (3 is excluded)
		fmt.Println(i)
	}
	// Output: 0  1  2
}

// ──────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Summing numbers and iterating over a slice
// ──────────────────────────────────────────────────────────────────────────────
// func main() {
// 	// While-style: find the first number > 10 starting from 1
// 	n := 1
// 	for n <= 10 {
// 		n++
// 	}
// 	fmt.Println("First number > 10:", n) // Output: 11
//
// 	// Classic for loop: print multiplication table of 3
// 	for i := 1; i <= 5; i++ {
// 		fmt.Printf("3 x %d = %d\n", i, 3*i)
// 	}
// 	// Output:
// 	// 3 x 1 = 3
// 	// 3 x 2 = 6
// 	// ...
//
// 	// Range loop: iterate over a slice
// 	fruits := []string{"apple", "banana", "cherry"}
// 	for i, fruit := range fruits {
// 		fmt.Printf("Index %d: %s\n", i, fruit)
// 	}
// 	// Output:
// 	// Index 0: apple
// 	// Index 1: banana
// 	// Index 2: cherry
// }
