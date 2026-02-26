// Package main is the entry point for a standalone Go program.
// This file demonstrates one of Go's most powerful features: INTERFACES.
//
// KEY CONCEPTS COVERED:
//  1. Interface definition  – what methods a type MUST implement
//  2. Structs               – custom data types that group related fields
//  3. Method receivers      – functions attached to a struct
//  4. Dependency Injection  – passing a behaviour (interface) into a struct
//  5. Polymorphism          – one call, different implementations at runtime
package main

import "fmt" // fmt provides formatted I/O functions like Println

// ─────────────────────────────────────────────────────────────────────────────
// INTERFACE DEFINITION
// ─────────────────────────────────────────────────────────────────────────────

// paymenter is an INTERFACE.
// An interface in Go defines a CONTRACT — a list of method signatures.
// Any type that implements ALL methods in the interface satisfies it automatically.
// (No explicit "implements" keyword needed — this is called implicit implementation.)
//
// Here, any type that has a  pay(amount float32)  method is a valid paymenter.
type paymenter interface {
	pay(amount float32) // any concrete payment gateway MUST implement this method
}

// ─────────────────────────────────────────────────────────────────────────────
// STRUCT: payment (the high-level payment processor)
// ─────────────────────────────────────────────────────────────────────────────

// payment is a high-level struct that does NOT care WHICH payment gateway is used.
// It only knows that its gateway field must satisfy the paymenter interface.
//
// This is DEPENDENCY INJECTION:
//   → The actual gateway (razorpay / stripe / etc.) is "injected" from outside.
//   → The payment struct stays decoupled from concrete implementations.
//   → Swapping gateways requires ZERO changes to this struct.
type payment struct {
	gateway paymenter // holds any value that satisfies the paymenter interface
}

// makePayment is a METHOD on the payment struct (value receiver).
// It delegates the actual payment work to whatever gateway was injected.
//
// POLYMORPHISM in action:
//   → If gateway is a razorpay{}, it runs razorpay's pay().
//   → If gateway is a stripe{},   it runs stripe's pay().
//   → The call p.gateway.pay(amount) looks identical in both cases!
func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount) // dynamic dispatch — Go picks the right implementation at runtime
}

// ─────────────────────────────────────────────────────────────────────────────
// CONCRETE TYPES: razorpay & stripe (the actual payment gateways)
// ─────────────────────────────────────────────────────────────────────────────

// razorpay is a concrete (real) struct representing the Razorpay payment gateway.
// It has no fields — it's used only to attach methods.
type razorpay struct{}

// stripe is a concrete struct representing the Stripe payment gateway.
type stripe struct{}

// pay implements the paymenter interface for razorpay.
// The method receiver (r razorpay) means: "this function belongs to the razorpay type".
// Because razorpay now has a pay() method, it AUTOMATICALLY satisfies paymenter.
func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount) // simulates a Razorpay API call
}

// pay implements the paymenter interface for stripe.
// Similarly, stripe also satisfies paymenter without any explicit declaration.
func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount) // simulates a Stripe API call
}

// ─────────────────────────────────────────────────────────────────────────────
// MAIN — program entry point
// ─────────────────────────────────────────────────────────────────────────────

func main() {
	// STEP 1: Choose a concrete gateway.
	// We pick stripe here, but we could easily swap to razorpay{} with one change.
	stripePaymentGw := stripe{} // creates a stripe value (satisfies paymenter)

	// STEP 2: Inject the gateway into the high-level payment struct.
	// The payment struct only sees the paymenter interface — it doesn't know or
	// care that it's actually a stripe under the hood.
	newPayment := payment{
		gateway: stripePaymentGw, // dependency injection via struct literal
	}

	// STEP 3: Trigger the payment.
	// Internally calls → stripe.pay(100) → prints "making payment using stripe 100"
	newPayment.makePayment(100)

	// ── Try swapping to Razorpay (just change stripePaymentGw to razorpay{}) ──
	// razorpayPaymentGw := razorpay{}
	// newPayment2 := payment{gateway: razorpayPaymentGw}
	// newPayment2.makePayment(250)  // prints "making payment using razorpay 250"
}

// ─────────────────────────────────────────────────────────────────────────────
// SUMMARY — Why interfaces matter
// ─────────────────────────────────────────────────────────────────────────────
//
//  ┌──────────────┐        implements        ┌─────────────────────┐
//  │  paymenter   │ ◄────────────────────── │  razorpay / stripe  │
//  │  (interface) │                          │  (concrete structs) │
//  └──────┬───────┘                          └─────────────────────┘
//         │ used by
//  ┌──────▼───────┐
//  │   payment    │  ← only depends on the interface, not the concrete type
//  │   (struct)   │    → easy to test, easy to extend, easy to swap gateways
//  └──────────────┘
//
//  Rule of thumb: "Accept interfaces, return concrete types." — Go proverb
