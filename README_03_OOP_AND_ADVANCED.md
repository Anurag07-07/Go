# 🐹 Go (Golang) Mastery Guide — Part 3: OOP Patterns in Go
> **Series**: `README_01_FOUNDATIONS.md` → `README_02_CORE_TYPES.md` → **`README_03_OOP_AND_ADVANCED.md`** → `README_04_CONCURRENCY_AND_MASTERY.md`

---

## 📚 Table of Contents
1. [Structs — Go's Custom Type System](#1-structs--gos-custom-type-system)
2. [Methods — Functions Attached to Types](#2-methods--functions-attached-to-types)
3. [Interfaces — The Contract System](#3-interfaces--the-contract-system)
4. [Struct Embedding — Go's Composition Model](#4-struct-embedding--gos-composition-model)
5. [Polymorphism in Go](#5-polymorphism-in-go)
6. [Dependency Injection Pattern](#6-dependency-injection-pattern)
7. [The Empty Interface (any)](#7-the-empty-interface-any)
8. [Type Assertions & Type Switches](#8-type-assertions--type-switches)
9. [Stringer Interface](#9-stringer-interface)
10. [Practice Exercises](#10-practice-exercises)

---

## 1. Structs — Go's Custom Type System

📁 File: `structs/structs.go`

> A **struct** is a custom data type that groups related fields together.
> It's Go's answer to **classes** — but simpler and more explicit.

### Basic Struct:
```go
package main

import "fmt"

// Define a struct type with 'type name struct { ... }'
type Person struct {
    Name    string
    Age     int
    Email   string
    IsAdmin bool
}

func main() {

    // ── Method 1: Named field initialization (RECOMMENDED) ───────────
    p1 := Person{
        Name:    "Anurag",
        Age:     25,
        Email:   "anurag@gmail.com",
        IsAdmin: true,
    }
    fmt.Println(p1)        // → {Anurag 25 anurag@gmail.com true}
    fmt.Println(p1.Name)   // → Anurag (dot notation)
    p1.Age = 26            // update a field
    fmt.Println(p1.Age)    // → 26

    // ── Method 2: Positional initialization (NOT recommended — fragile) ──
    p2 := Person{"Bob", 30, "bob@gmail.com", false}
    fmt.Println(p2)

    // ── Method 3: Zero-value struct (all fields = zero values) ────────
    var p3 Person
    // p3.Name == "" | p3.Age == 0 | p3.IsAdmin == false | p3.Email == ""
    p3.Name = "Charlie"
    fmt.Println(p3)

    // ── Pointer to struct ─────────────────────────────────────────────
    // Go automatically dereferences pointer — use p4.Name, NOT (*p4).Name
    p4 := &Person{Name: "Diana", Age: 28}
    p4.Age = 29            // same as (*p4).Age = 29  (auto-deref)
    fmt.Println(p4.Age)    // → 29
}
```

### Zero Values Reminder:
```go
// If you don't set a field, Go uses the zero value:
type Order struct {
    ID       string    // ""
    Amount   float32   // 0
    Paid     bool      // false
    Quantity int       // 0
    Tags     []string  // nil
}
```

### Anonymous Struct (one-off, no type declaration needed):
```go
// Great for test fixtures or ad-hoc data
product := struct {
    Name  string
    Price float32
}{
    Name:  "Laptop",
    Price: 75000.00,
}
fmt.Printf("Product: %s, Price: %.2f\n", product.Name, product.Price)
```

### Constructor Function Pattern (idiomatic Go):
```go
// Go doesn't have a 'new' keyword for custom types.
// Convention: write a "constructor" function named New<Type> or new<Type>
func NewPerson(name string, age int, email string) *Person {
    // Validate inputs here
    if age < 0 {
        age = 0
    }
    return &Person{
        Name:  name,
        Age:   age,
        Email: email,
    }
}

p := NewPerson("Anurag", 25, "anurag@gmail.com")
```

---

## 2. Methods — Functions Attached to Types

> A **method** is a function with a **receiver** — it's attached to a specific type.
> Syntax: `func (receiver TypeName) MethodName(params) returnType { ... }`

### Value Receiver vs Pointer Receiver:
```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

// ── VALUE RECEIVER — (r Rectangle) ────────────────────────────────────────────
// Gets a COPY of the struct — modifications don't affect the original
// Use for: read-only operations, small structs
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// ── POINTER RECEIVER — (r *Rectangle) ────────────────────────────────────────
// Gets a POINTER to the struct — modifications AFFECT the original
// Use for: when method needs to modify the struct, or struct is large
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor   // modifies original (via pointer)
    r.Height *= factor
}

func (r *Rectangle) String() string {
    return fmt.Sprintf("Rectangle(%.1f × %.1f)", r.Width, r.Height)
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}

    // Call methods — Go auto-takes the address when needed
    fmt.Println(rect.Area())      // → 15 (5 × 3)
    fmt.Println(rect.Perimeter()) // → 16 (2 × (5+3))

    // Scale modifies the original rect
    rect.Scale(2)
    fmt.Println(rect.Width)       // → 10
    fmt.Println(rect.Height)      // → 6

    // Works on pointer too — Go auto-dereferences
    pRect := &Rectangle{Width: 4, Height: 2}
    fmt.Println(pRect.Area())     // → 8
}
```

### Method chaining (fluent API):
```go
type QueryBuilder struct {
    table  string
    where  string
    limit  int
}

func (q *QueryBuilder) From(table string) *QueryBuilder {
    q.table = table
    return q  // return pointer to allow chaining
}

func (q *QueryBuilder) Where(condition string) *QueryBuilder {
    q.where = condition
    return q
}

func (q *QueryBuilder) Limit(n int) *QueryBuilder {
    q.limit = n
    return q
}

func (q *QueryBuilder) Build() string {
    return fmt.Sprintf("SELECT * FROM %s WHERE %s LIMIT %d",
        q.table, q.where, q.limit)
}

func main() {
    query := (&QueryBuilder{}).
        From("users").
        Where("age > 18").
        Limit(10).
        Build()

    fmt.Println(query)
    // → SELECT * FROM users WHERE age > 18 LIMIT 10
}
```

---

## 3. Interfaces — The Contract System

📁 File: `interafaces/inter.go`

> An **interface** defines a set of method signatures — a **contract**.
> Any type that implements ALL the methods in an interface AUTOMATICALLY satisfies it.
> **No `implements` keyword needed** — this is called **implicit implementation**.

### Interface Definition & Implementation:
```go
package main

import (
    "fmt"
    "math"
)

// Define the interface — the CONTRACT
type Shape interface {
    Area() float64
    Perimeter() float64
}

// ── Concrete type 1: Circle ───────────────────────────────────────────────────
type Circle struct {
    Radius float64
}

// Circle implements Shape by having both Area() and Perimeter() methods
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// ── Concrete type 2: Rectangle ───────────────────────────────────────────────
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// ── Function that works with ANY Shape ────────────────────────────────────────
// This is POLYMORPHISM: one function, many concrete types
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// ── Total area of any collection of shapes ────────────────────────────────────
func totalArea(shapes []Shape) float64 {
    total := 0.0
    for _, s := range shapes {
        total += s.Area()
    }
    return total
}

func main() {
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 6}

    printShapeInfo(c) // → Area: 78.54, Perimeter: 31.42
    printShapeInfo(r) // → Area: 24.00, Perimeter: 20.00

    // Slice of Shape interface — holds DIFFERENT concrete types
    shapes := []Shape{c, r, Circle{Radius: 3}}
    fmt.Printf("Total area: %.2f\n", totalArea(shapes)) // → 107.68
}
```

### Your Real Code — Payment Gateway Example:
```go
package main

import "fmt"

// paymenter interface — the payment CONTRACT
// Any type with a pay(float32) method automatically satisfies this
type paymenter interface {
    pay(amount float32)
}

// payment struct — HIGH LEVEL, depends only on the interface
type payment struct {
    gateway paymenter  // any value that satisfies paymenter
}

// makePayment delegates to whatever gateway was injected
func (p payment) makePayment(amount float32) {
    p.gateway.pay(amount)  // dynamic dispatch!
}

// razorpay — concrete implementation 1
type razorpay struct{}
func (r razorpay) pay(amount float32) {
    fmt.Println("making payment using razorpay", amount)
}

// stripe — concrete implementation 2
type stripe struct{}
func (s stripe) pay(amount float32) {
    fmt.Println("making payment using stripe", amount)
}

func main() {
    // Inject Stripe
    stripeGw := stripe{}
    p := payment{gateway: stripeGw}
    p.makePayment(100) // → making payment using stripe 100

    // Swap to Razorpay — ZERO changes to payment struct!
    razorpayGw := razorpay{}
    p2 := payment{gateway: razorpayGw}
    p2.makePayment(250) // → making payment using razorpay 250
}
```

---

## 4. Struct Embedding — Go's Composition Model

📁 File: `structs/structs.go`

> Go has **no inheritance**. Instead it has **embedding** — composition over inheritance.
> When you embed a struct, you get all its fields and methods promoted to the outer type.

```go
package main

import (
    "fmt"
    "time"
)

// ── Base struct: customer ──────────────────────────────────────────────────────
type customer struct {
    name string
}

// ── Embedded struct: order EMBEDS customer ────────────────────────────────────
// This is NOT inheritance — it's composition!
type order struct {
    id        string
    amount    float32
    status    string
    customer            // EMBEDDED — no field name, just the type
    createdAt time.Time
}

func main() {
    // Initialize embedded struct using type name as the field key
    o := order{
        id:     "ORD-001",
        amount: 1999.99,
        status: "pending",
        customer: customer{    // initialize embedded struct
            name: "Anurag",
        },
        createdAt: time.Now(),
    }

    // ── Promoted field access — shorthand (embedding magic!) ─────────
    fmt.Println(o.name)           // → Anurag  (promoted from customer)

    // ── Explicit access ───────────────────────────────────────────────
    fmt.Println(o.customer.name)  // → Anurag  (also valid)

    fmt.Println(o.id)             // → ORD-001
    fmt.Println(o.amount)         // → 1999.99
}
```

### Embedding with methods:
```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return a.Name + " makes a sound"
}

type Dog struct {
    Animal  // embed Animal — Dog "inherits" Speak()
    Breed string
}

// Dog can OVERRIDE Speak() if needed
func (d Dog) Speak() string {
    return d.Name + " says: Woof!"
}

func main() {
    d := Dog{
        Animal: Animal{Name: "Rex"},
        Breed:  "German Shepherd",
    }

    fmt.Println(d.Speak())         // → Rex says: Woof! (Dog's own Speak)
    fmt.Println(d.Animal.Speak())  // → Rex makes a sound (Animal's Speak)
    fmt.Println(d.Name)            // → Rex (promoted field)
}
```

### Real-world: Constructor with embedding:
```go
// Constructor function — idiomatic Go pattern
func newOrder(id string, amount float32, status string) *order {
    return &order{
        id:        id,
        amount:    amount,
        status:    status,
        createdAt: time.Now(),
    }
}

// Method with pointer receiver — modifies the original
func (o *order) changeStatus(newStatus string) {
    o.status = newStatus
}

// Method with pointer receiver — getter
func (o *order) getAmount() float32 {
    return o.amount
}

func main() {
    myOrder := newOrder("ORD-101", 199.99, "pending")
    fmt.Println("Status:", myOrder.status)        // → pending
    fmt.Println("Amount:", myOrder.getAmount())   // → 199.99

    myOrder.changeStatus("shipped")
    fmt.Println("New Status:", myOrder.status)    // → shipped

    myOrder.customer = customer{name: "Priya"}
    fmt.Println("Customer:", myOrder.name)        // → Priya (promoted!)
}
```

---

## 5. Polymorphism in Go

> **Polymorphism** = one interface, many concrete implementations.
> The RIGHT method is called at **runtime** based on the actual type.

```go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Name() string
}

type Circle struct{ Radius float64 }
type Square struct{ Side float64 }
type Triangle struct{ Base, Height float64 }

func (c Circle) Area() float64   { return math.Pi * c.Radius * c.Radius }
func (c Circle) Name() string    { return "Circle" }

func (s Square) Area() float64   { return s.Side * s.Side }
func (s Square) Name() string    { return "Square" }

func (t Triangle) Area() float64 { return 0.5 * t.Base * t.Height }
func (t Triangle) Name() string  { return "Triangle" }

// printArea works with ANY Shape — polymorphic behavior!
func printArea(s Shape) {
    fmt.Printf("%-10s area = %.2f\n", s.Name(), s.Area())
}

func largestShape(shapes []Shape) Shape {
    largest := shapes[0]
    for _, s := range shapes[1:] {
        if s.Area() > largest.Area() {
            largest = s
        }
    }
    return largest
}

func main() {
    shapes := []Shape{
        Circle{Radius: 5},
        Square{Side: 4},
        Triangle{Base: 6, Height: 8},
    }

    for _, s := range shapes {
        printArea(s)
    }
    // → Circle     area = 78.54
    // → Square     area = 16.00
    // → Triangle   area = 24.00

    biggest := largestShape(shapes)
    fmt.Printf("\nLargest: %s (%.2f)\n", biggest.Name(), biggest.Area())
    // → Largest: Circle (78.54)
}
```

---

## 6. Dependency Injection Pattern

> This is one of Go's most important patterns for clean, testable code.

```go
package main

import "fmt"

// ── Define interfaces (contracts/ports) ───────────────────────────────────────
type Logger interface {
    Log(message string)
}

type Database interface {
    Save(data string) error
    FindByID(id string) (string, error)
}

// ── Concrete implementations ──────────────────────────────────────────────────
type ConsoleLogger struct{}
func (l ConsoleLogger) Log(msg string) {
    fmt.Printf("[LOG] %s\n", msg)
}

type InMemoryDB struct {
    store map[string]string
}
func (db *InMemoryDB) Save(data string) error {
    db.store[data] = data
    return nil
}
func (db *InMemoryDB) FindByID(id string) (string, error) {
    val, ok := db.store[id]
    if !ok {
        return "", fmt.Errorf("not found: %s", id)
    }
    return val, nil
}

// ── UserService — depends on INTERFACES, not concrete types ──────────────────
// This makes it easy to test (inject mock logger/db) and swap implementations
type UserService struct {
    db     Database
    logger Logger
}

func NewUserService(db Database, logger Logger) *UserService {
    return &UserService{db: db, logger: logger}
}

func (s *UserService) CreateUser(name string) error {
    s.logger.Log("Creating user: " + name)
    return s.db.Save(name)
}

func main() {
    // Wire up dependencies
    db := &InMemoryDB{store: make(map[string]string)}
    logger := ConsoleLogger{}
    service := NewUserService(db, logger)

    service.CreateUser("Anurag")  // → [LOG] Creating user: Anurag
    service.CreateUser("Priya")   // → [LOG] Creating user: Priya

    // To test: create MockDB and MockLogger — no changes to UserService!
}
```

---

## 7. The Empty Interface (any)

> `any` (alias for `interface{}`) can hold a value of **ANY type**.
> Use it carefully — you lose type safety.

```go
package main

import "fmt"

func printAnything(v any) {
    fmt.Printf("Type: %T, Value: %v\n", v, v)
}

func main() {
    printAnything(42)          // → Type: int, Value: 42
    printAnything("hello")     // → Type: string, Value: hello
    printAnything(true)        // → Type: bool, Value: true
    printAnything([]int{1,2,3}) // → Type: []int, Value: [1 2 3]

    // A slice of any — can hold mixed types
    mixed := []any{1, "two", 3.0, true, nil}
    for _, v := range mixed {
        fmt.Println(v)
    }
}
```

---

## 8. Type Assertions & Type Switches

> When you have an interface value, use type assertions to get the concrete type back.

```go
package main

import "fmt"

func describe(i any) {
    // ── Type Switch — most idiomatic way ─────────────────────────────
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d (doubled: %d)\n", v, v*2)
    case string:
        fmt.Printf("String: %q (length: %d)\n", v, len(v))
    case bool:
        fmt.Printf("Bool: %t\n", v)
    case []int:
        fmt.Printf("Int slice with %d elements\n", len(v))
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

func main() {
    describe(42)
    describe("Golang")
    describe(true)
    describe([]int{1, 2, 3})
    describe(3.14)

    // ── Single type assertion ─────────────────────────────────────────
    var val any = "hello"

    // Safe assertion (recommended) — returns (value, ok)
    s, ok := val.(string)
    if ok {
        fmt.Println("Got string:", s) // → Got string: hello
    }

    // Unsafe assertion — PANICS if wrong type!
    // str := val.(int)  // PANIC: interface conversion
}
```

---

## 9. Stringer Interface

> The `fmt.Stringer` interface lets you control how your type prints.

```go
package main

import "fmt"

type Temperature struct {
    Celsius float64
}

// Implementing fmt.Stringer interface
// If a type has a String() string method, fmt uses it automatically
func (t Temperature) String() string {
    return fmt.Sprintf("%.1f°C (%.1f°F)", t.Celsius, t.Celsius*9/5+32)
}

type Color struct {
    R, G, B uint8
}

func (c Color) String() string {
    return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

func main() {
    temp := Temperature{Celsius: 37.5}
    fmt.Println(temp)          // → 37.5°C (99.5°F)  (uses String() automatically!)
    fmt.Printf("Body temp: %v\n", temp) // same

    color := Color{R: 255, G: 128, B: 0}
    fmt.Println(color)         // → #FF8000
}
```

---

## 10. Practice Exercises

### 🟢 Beginner:
1. Create a `Car` struct with fields (Make, Model, Year, Price) and print all car details
2. Add a `IsAffordable(budget float64) bool` method to Car
3. Create a `Vehicle` interface with `Speed() int` and `FuelType() string` methods

### 🟡 Intermediate:
4. Implement the `Vehicle` interface for `Car`, `Bike`, and `Truck` structs
5. Write a function `printVehicleDetails(vehicles []Vehicle)` that loops over any collection of vehicles
6. Create a base `Address` struct and embed it into `Person` and `Company` structs

### 🔴 Advanced:
7. Build a notification system: create a `Notifier` interface with `Send(to, message string) error`, then implement `EmailNotifier`, `SMSNotifier`, and `SlackNotifier`. Write a `NotificationService` struct that takes `[]Notifier` and broadcasts to all
8. Implement a plugin system where each plugin satisfies `Plugin interface { Name() string; Execute(input string) string }`. Register plugins in a map and execute by name
9. Create a type hierarchy: `Animal` → `Pet` (embedding) → implement `Speak()` interface for `Dog`, `Cat`, `Parrot` — demonstrate polymorphism and method overriding

---

## Interface Design Principles

```
┌─────────────────────────────────────────────────────────────────────────┐
│  Go Interface Best Practices                                            │
├─────────────────────────────────────────────────────────────────────────┤
│  1. Keep interfaces SMALL (1-3 methods is ideal)                        │
│     → io.Reader has 1 method. fmt.Stringer has 1 method.               │
│                                                                         │
│  2. "Accept interfaces, return concrete types."  — Go proverb           │
│     func process(r io.Reader) *Result { ... }  ✅                       │
│     func process(f *os.File) io.Reader { ... } ✅                       │
│                                                                         │
│  3. Name 1-method interfaces with "-er" suffix                          │
│     Reader, Writer, Closer, Stringer, Handler, Logger                  │
│                                                                         │
│  4. Design for the CONSUMER, not the implementer                        │
│     The package that USES the interface should DEFINE it.               │
│                                                                         │
│  5. Implicit implementation enables zero-coupling                       │
│     Third-party types can satisfy your interfaces without modification! │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 🗺️ What's Next?

Continue to **[README_04_CONCURRENCY_AND_MASTERY.md](./README_04_CONCURRENCY_AND_MASTERY.md)** →
> **Goroutines | Channels | WaitGroups | Mutexes | Select | Real-world Patterns | Go Module System**
