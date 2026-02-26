// Package declaration — every Go file must start with a package name
package main

// Import multiple packages inside a parenthesis block
import (
	"fmt"    // "fmt" is for printing output to the console
	"slices" // "slices" package (Go 1.21+) provides helper functions for slices like Equal, Contains, etc.
)

// Slices are like dynamic arrays — they can grow and shrink in size
// They are the MOST USED data structure in Go
// Unlike arrays, slices do NOT have a fixed size

// main is the program's entry point
func main() {

	// ── Declare a nil Slice ────────────────────────────────────────────────
	// var nums []string declares a slice with no elements — it's nil (zero value for slices)
	// A nil slice has length 0 and capacity 0
	var nums []string
	fmt.Println(nums == nil) // Output: true — the slice is nil (not initialized yet)
	fmt.Println(len(nums))   // Output: 0 — no elements

	// ── Create a Slice Using make() ────────────────────────────────────────
	// make([]int, size, capacity) creates a slice with:
	//   size (5)     — initial number of elements (all set to 0 for int)
	//   capacity (50) — max elements before Go must allocate more memory
	var nums1 = make([]int, 5, 50)
	fmt.Println(nums1) // Output: [0 0 0 0 0] — 5 zero-value integers

	// ── Append Elements to a Slice ─────────────────────────────────────────
	// append() adds elements to the END of a slice
	// It returns a new (or grown) slice — you must reassign it back to nums1
	nums1 = append(nums1, 1, 2, 3, 4, 5) // appends 5 new elements after the existing zeros

	// cap() returns the maximum number of elements the slice can hold before reallocation
	fmt.Println(cap(nums1)) // Output: 50 (we set capacity to 50 in make())

	// Print the full slice — first 5 are zeros, then appended 1-5
	fmt.Println(nums1) // Output: [0 0 0 0 0 1 2 3 4 5]

	// ── Modify an Element by Index ─────────────────────────────────────────
	// Access and update index 0 (the first element) to 45
	nums1[0] = 45
	fmt.Println(nums1) // Output: [45 0 0 0 0 1 2 3 4 5]

	// ── Copy a Slice ───────────────────────────────────────────────────────
	// make([]int, len(nums1)) creates a new slice with the same length as nums1
	var nums2 = make([]int, len(nums1))

	// copy(dst, src) copies elements from nums1 (source) to nums2 (destination)
	// This is a DEEP copy — changes to nums2 won't affect nums1
	copy(nums2, nums1) // copy(duplicate array, original array)

	// ── Slice Operator (Sub-slicing) ───────────────────────────────────────
	// You can extract a portion of a slice using [low:high] notation
	// The result includes elements from index 'low' up to (but not including) 'high'
	var nums3 = []int{4, 5, 6} // declare and initialize a slice inline
	fmt.Println(nums3[0:1])    // [low:high] → includes index 0 only → Output: [4]
	fmt.Println(nums3[:1])     // [:high]    → same as [0:1]         → Output: [4]
	fmt.Println(nums3[1:])     // [low:]     → from index 1 to end   → Output: [5 6]

	// Print both nums1 (modified) and nums2 (copy)
	fmt.Println(nums1, nums2)

	// ── Compare Two Slices ─────────────────────────────────────────────────
	// slices.Equal(a, b) returns true if both slices have the same elements in the same order
	fmt.Println(slices.Equal(nums1, nums2)) // Output: true — they are identical copies

	// ── 2D Slice ───────────────────────────────────────────────────────────
	// A slice of slices — like a dynamic 2D matrix
	// Each inner slice can have a different length (unlike 2D arrays)
	var nums5 = [][]int{{1, 2, 3}, {4, 6, 5}}
	fmt.Println(nums5) // Output: [[1 2 3] [4 6 5]]
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Managing a dynamic shopping cart using slices
// ─────────────────────────────────────────────────────────────────────────────
// func main() {
// 	// Start with an empty cart (nil slice)
// 	cart := []string{}
//
// 	// Add items to the cart using append
// 	cart = append(cart, "Apple")
// 	cart = append(cart, "Banana", "Mango") // append multiple at once
// 	fmt.Println("Cart:", cart)              // Output: Cart: [Apple Banana Mango]
// 	fmt.Println("Items:", len(cart))        // Output: Items: 3
//
// 	// View first two items using slice operator
// 	fmt.Println("First 2:", cart[:2])       // Output: First 2: [Apple Banana]
//
// 	// Copy the cart (so original isn't affected)
// 	backup := make([]string, len(cart))
// 	copy(backup, cart)
//
// 	// Modify original cart — backup stays intact
// 	cart[0] = "Grapes"
// 	fmt.Println("Modified Cart:", cart)     // Output: Modified Cart: [Grapes Banana Mango]
// 	fmt.Println("Backup:", backup)          // Output: Backup: [Apple Banana Mango]
//
// 	// Compare cart and backup — they differ now
// 	fmt.Println(slices.Equal(cart, backup)) // Output: false
// }
