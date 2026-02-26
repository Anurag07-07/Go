// Closures in Go
// A closure is a function that "remembers" the variables from its outer scope,
// even after the outer function has finished executing.
// The inner function "closes over" (captures) those variables.
package main

// Import "fmt" for console output
import "fmt"

// ── Closure Factory Function ───────────────────────────────────────────────────
// 'counter' is a function that RETURNS another function
// Return type: func() int — a function that takes no arguments and returns an int
func counter() func() int {
	// 'count' is a local variable inside 'counter'
	// Normally it would be destroyed when 'counter' returns,
	// but the inner function CAPTURES it — keeping it alive in memory
	count := 1

	// Return an anonymous (unnamed) inner function
	// This inner function forms a CLOSURE over 'count'
	return func() int {
		count += 1   // access and MODIFY the captured 'count' variable each time this is called
		return count // return the updated count value
	}
}

// main is the program entry point
func main() {
	// Call counter() — this executes the outer function and returns the inner function
	// 'increment' now holds a reference to the inner function (plus its captured 'count' state)
	increment := counter()

	// Call the returned inner function (increment)
	// count starts at 1, then += 1 makes it 2 — returns 2
	fmt.Println(increment()) // Output: 2

	// If you call increment() again:
	// count is still remembered (it's 2), then += 1 makes it 3 — returns 3
	// fmt.Println(increment()) // Output: 3 (try uncommenting this!)
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Using closures to create independent counters and an adder
// ─────────────────────────────────────────────────────────────────────────────
// func counter() func() int {
// 	count := 0                    // each closure gets its OWN 'count'
// 	return func() int {
// 		count++                   // increment and return
// 		return count
// 	}
// }
//
// // adder returns a closure that adds 'x' to whatever is passed to it
// func adder(x int) func(int) int {
// 	return func(y int) int {
// 		return x + y             // 'x' is captured from the outer scope
// 	}
// }
//
// func main() {
// 	// Create two INDEPENDENT counters — each has its own captured 'count'
// 	c1 := counter()
// 	c2 := counter()
//
// 	fmt.Println(c1()) // Output: 1   (c1's count: 0 → 1)
// 	fmt.Println(c1()) // Output: 2   (c1's count: 1 → 2)
// 	fmt.Println(c2()) // Output: 1   (c2's count is independent — starts at 0)
//
// 	// Create an adder that always adds 10
// 	addTen := adder(10)
// 	fmt.Println(addTen(5))  // Output: 15  (10 + 5)
// 	fmt.Println(addTen(20)) // Output: 30  (10 + 20)
//
// 	// Create an adder that always adds 100
// 	addHundred := adder(100)
// 	fmt.Println(addHundred(7)) // Output: 107  (100 + 7)
// }
