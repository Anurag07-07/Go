// Package declaration — required at the start of every Go file
package main

// NOTE: If you don't set any field in a struct, Go uses its zero value:
//   string → ""   |   int/float → 0   |   bool → false   |   pointer → nil

// Import required packages
import (
	"fmt"  // "fmt" for printing output to the console
	"time" // "time" for time-related types like time.Time and time.Now()
)

// ── Struct Embedding ───────────────────────────────────────────────────────────
// 'customer' is a simple struct with just a name field
// Structs in Go are custom data types that group related fields together
type customer struct {
	name string // 'name' field of type string — stores the customer's name
}

// 'order' struct embeds the 'customer' struct (struct embedding = composition)
// Embedding lets 'order' inherit the fields and methods of 'customer'
// This is Go's way of achieving composition (not classical inheritance)
type order struct {
	id        string    // unique identifier for the order
	amount    float32   // order total amount as a 32-bit float
	status    string    // current status of the order (e.g., "pending", "completed")
	customer            // EMBEDDED struct — 'order' now has access to customer.name
	createdAt time.Time // stores the date/time the order was created
}

// ── Constructor Function ───────────────────────────────────────────────────────
// 'newOrder' is a constructor-style function that creates and returns an *order
// Returning a pointer (*order) is efficient — avoids copying the entire struct
func newOrder(id string, amount float32, status string) *order {
	// Create an 'order' value using field names (named initialization)
	order := order{
		id:     id,     // set the id field
		amount: amount, // set the amount field
		status: status, // set the status field
	}
	return &order // return a pointer to the order (& takes the address)
}

// ── Method with Pointer Receiver ──────────────────────────────────────────────
// Methods in Go are attached to a type using a receiver
// (o *order) is a POINTER RECEIVER — changes made inside this method affect the original struct
// This method updates the 'status' field of an order
func (o *order) changeStatus(status string) {
	o.status = status // modify the status of the order via its pointer
}

// ── Method with Pointer Receiver (Getter) ─────────────────────────────────────
// This method returns the 'amount' field of the order
// Using a pointer receiver is consistent and allows future mutation if needed
func (o *order) getAmount() float32 {
	return o.amount // read and return the amount value
}

// main is the program's entry point
func main() {

	// ── Struct Embedding in Action ─────────────────────────────────────────
	// Create an 'order' struct using named field initialization
	// The embedded 'customer' struct is initialized using its type name as the field key
	ccs := order{
		id:     "1",         // set order ID
		amount: 45,          // set order amount
		status: "completed", // set order status
		customer: customer{ // initialize the embedded 'customer' struct
			name: "Anurag", // set the customer's name inside the embedded struct
		},
	}

	// Print the entire struct — shows all fields including embedded customer
	fmt.Println(ccs)
	// Output: {1 45 completed {Anurag} 0001-01-01 00:00:00 +0000 UTC}

	// Access embedded struct fields directly:
	// fmt.Println(ccs.name)       → "Anurag"  (promoted from embedded customer)
	// fmt.Println(ccs.customer.name) → "Anurag" (explicit access also works)
}

// ─────────────────────────────────────────────────────────────────────────────
// EXAMPLE: Creating and using an order with methods and constructor
// ─────────────────────────────────────────────────────────────────────────────
// func main() {
// 	// Create an order using the constructor function
// 	myOrder := newOrder("ORD-101", 199.99, "pending")
//
// 	// Access fields on the returned pointer
// 	fmt.Println("Order ID:", myOrder.id)          // Output: Order ID: ORD-101
// 	fmt.Println("Amount:", myOrder.getAmount())   // Output: Amount: 199.99
// 	fmt.Println("Status:", myOrder.status)        // Output: Status: pending
//
// 	// Set the creation timestamp
// 	myOrder.createdAt = time.Now()
//
// 	// Change the order status using the method (modifies via pointer)
// 	myOrder.changeStatus("shipped")
// 	fmt.Println("Updated Status:", myOrder.status) // Output: Updated Status: shipped
//
// 	// Set the embedded customer's name
// 	myOrder.customer = customer{name: "Priya"}
// 	fmt.Println("Customer:", myOrder.name)          // Output: Customer: Priya
//
// 	// Anonymous struct: one-off struct used directly without a type definition
// 	product := struct {
// 		name  string
// 		price float32
// 	}{
// 		"Laptop", 75000.00,
// 	}
// 	fmt.Println("Product:", product.name, "Price:", product.price)
// 	// Output: Product: Laptop Price: 75000
// }
