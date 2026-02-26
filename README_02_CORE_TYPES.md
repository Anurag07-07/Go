# 🐹 Go (Golang) Mastery Guide — Part 2: Core Types & Functions
> **Series**: `README_01_FOUNDATIONS.md` → **`README_02_CORE_TYPES.md`** → `README_03_OOP_AND_ADVANCED.md` → `README_04_CONCURRENCY_AND_MASTERY.md`

---

## 📚 Table of Contents
1. [Functions — Complete Guide](#1-functions--complete-guide)
2. [Pointers — Pass by Reference](#2-pointers--pass-by-reference)
3. [Maps — Key-Value Powerhouse](#3-maps--key-value-powerhouse)
4. [Range — Iterating Over Collections](#4-range--iterating-over-collections)
5. [Closures — Functions That Remember](#5-closures--functions-that-remember)
6. [Variadic Functions](#6-variadic-functions)
7. [Generics (Go 1.18+)](#7-generics-go-118)
8. [Error Handling](#8-error-handling)
9. [Practice Exercises](#9-practice-exercises)

---

## 1. Functions — Complete Guide

📁 File: `functions/func.go`

### Basic Function:
```go
package main

import "fmt"

// func name(param type) returnType { ... }
func add(a int, b int) int {
    return a + b
}

// Shorthand: same-type parameters can be grouped
func multiply(a, b int) int {
    return a * b
}

func main() {
    result := add(45, 56)        // → 101
    product := multiply(4, 5)   // → 20
    fmt.Println(result, product)
}
```

### Multiple Return Values (Go superpower!):
```go
// Return multiple values — list types in parentheses
func getLanguage() (string, string, string) {
    return "golang", "java", "C++"
}

// Named return values — even cleaner
func minMax(nums []int) (min, max int) {
    min, max = nums[0], nums[0]
    for _, n := range nums[1:] {
        if n < min { min = n }
        if n > max { max = n }
    }
    return  // "naked return" — returns named values
}

func main() {
    lang1, lang2, lang3 := getLanguage()
    fmt.Println(lang1, lang2, lang3) // → golang java C++

    // Use _ to discard values you don't need
    lang1, _, _ := getLanguage()    // only keep first value
    fmt.Println(lang1)              // → golang

    min, max := minMax([]int{3, 1, 8, 2, 7})
    fmt.Printf("Min: %d, Max: %d\n", min, max) // → Min: 1, Max: 8
}
```

### Functions as Values (First-Class Functions):
```go
// A function can be stored in a variable
add := func(a, b int) int {
    return a + b
}
fmt.Println(add(3, 4)) // → 7

// A function can be passed as an argument
func applyOp(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    result := applyOp(6, 7, func(a, b int) int {
        return a * b
    })
    fmt.Println(result) // → 42
}
```

### Function That Returns a Function:
```go
// Higher-order function — returns another function
func processIt() func(a int) int {
    return func(a int) int {
        return a * 2
    }
}

func main() {
    doubler := processIt()        // get the inner function
    fmt.Println(doubler(5))       // → 10
    fmt.Println(doubler(21))      // → 42
}
```

### defer — Run at Function Exit:
```go
func main() {
    defer fmt.Println("World")   // runs LAST (when main returns)
    defer fmt.Println("Middle")  // defers stack in LIFO order
    fmt.Println("Hello")
}
// Output:
// Hello
// Middle
// World

// Common use case: cleanup resources
func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil { return err }
    defer f.Close()    // guaranteed to run even if function panics
    // ... do stuff with f
    return nil
}
```

### 🎯 Solved Example: Mini Calculator
```go
package main

import "fmt"

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }
func divide(a, b int) (int, int) {
    return a / b, a % b  // quotient, remainder
}

// applyOp takes ANY function with signature func(int,int) int
func applyOp(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    fmt.Println("Add:", add(10, 5))           // → 15
    fmt.Println("Subtract:", subtract(10, 5)) // → 5
    fmt.Println("Multiply:", multiply(10, 5)) // → 50
    q, r := divide(17, 5)
    fmt.Printf("Divide: %d quotient, %d remainder\n", q, r) // → 3, 2

    // Pass different operations dynamically
    fmt.Println("Dynamic Add:", applyOp(6, 7, add))      // → 13
    fmt.Println("Dynamic Mul:", applyOp(6, 7, multiply)) // → 42
}
```

---

## 2. Pointers — Pass by Reference

📁 File: `pointers/pointers.go`

> **Go is pass-by-value by default.** Every function gets a COPY of the argument.
> Pointers let you pass the memory ADDRESS so the function can modify the original.

### The Core Concept:
```
Variable 'num':
┌─────────────────────────────────────────┐
│  Address: 0xc0000b4008                  │
│  Value:   42                            │
└─────────────────────────────────────────┘
         │
         │  &num  (address-of operator → creates a pointer)
         ▼
Pointer 'p' of type *int:
┌─────────────────────────────────────────┐
│  Value:   0xc0000b4008  (stores address)│
└─────────────────────────────────────────┘
         │
         │  *p  (dereference operator → read the value at the address)
         ▼
         42  (the actual value)
```

### Pass by Value vs Pass by Pointer:
```go
package main

import "fmt"

// ❌ Pass by VALUE — original is NOT changed
func changeNum(num int) {
    num = 5  // modifies only the LOCAL COPY
    fmt.Println("In changeNum:", num) // → 5
}

// ✅ Pass by POINTER — original IS changed
func changeNumPtr(num *int) {
    *num = 5  // dereference → modify the actual value at that address
    fmt.Println("In changeNumPtr:", *num) // → 5
}

func main() {
    num := 1

    changeNum(num)
    fmt.Println("After changeNum:", num) // → 1 (UNCHANGED!)

    changeNumPtr(&num)   // & = "take the address of num"
    fmt.Println("After changeNumPtr:", num) // → 5 (CHANGED!)
}
```

### Pointer Basics:
```go
func main() {
    x := 42

    p := &x           // p is a *int (pointer to int) — stores address of x
    fmt.Println(p)    // → 0xc0000b4008 (memory address)
    fmt.Println(*p)   // → 42  (dereference: read value at address)

    *p = 100          // change value AT the address
    fmt.Println(x)    // → 100 (x is changed because p points to x!)

    // new() creates a pointer to a new zero-value variable
    q := new(int)     // *int, pointing to 0
    *q = 77
    fmt.Println(*q)   // → 77
}
```

### When to use Pointers:
```go
// ✅ Use pointer receiver when:
// 1. Method needs to MODIFY the struct
// 2. Struct is large (avoid copying overhead)

type Counter struct {
    count int
}

func (c *Counter) Increment() {  // pointer receiver — modifies original
    c.count++
}

func (c Counter) Value() int {   // value receiver — just reads, no mutation
    return c.count
}

func main() {
    c := Counter{}
    c.Increment()
    c.Increment()
    fmt.Println(c.Value()) // → 2
}
```

### 🎯 Solved Example: Swap using Pointers
```go
package main

import "fmt"

func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    x, y := 10, 20
    fmt.Println("Before:", x, y) // → Before: 10 20
    swap(&x, &y)
    fmt.Println("After:", x, y)  // → After: 20 10
}
```

---

## 3. Maps — Key-Value Powerhouse

📁 File: `maps/maps.go`

> Maps are Go's built-in key-value store (like Python's dict or Java's HashMap).

```go
package main

import (
    "fmt"
    "maps" // Go 1.21+
)

func main() {

    // ── Create using make() ──────────────────────────────────────────
    m1 := make(map[string]string)
    m1["name"]    = "golang"
    m1["company"] = "Google"
    fmt.Println(m1)            // → map[company:Google name:golang]
    fmt.Println(m1["name"])    // → golang

    // Accessing missing key → returns ZERO VALUE (not a crash!)
    fmt.Println(m1["missing"]) // → "" (empty string)
    fmt.Println(len(m1))       // → 2

    // ── Delete a key ──────────────────────────────────────────────────
    delete(m1, "name")
    fmt.Println(m1)            // → map[company:Google]

    // ── Clear entire map (Go 1.21+) ───────────────────────────────────
    clear(m1)
    fmt.Println(len(m1))       // → 0

    // ── Map literal (declare + initialize) ───────────────────────────
    m2 := map[string]string{
        "name":  "Anurag",
        "email": "anurag@gmail.com",
    }
    fmt.Println(m2)            // → map[email:anurag@gmail.com name:Anurag]

    // ── Two-value lookup (existence check) ────────────────────────────
    // value, ok := map[key]
    // ok == true  → key exists
    // ok == false → key does NOT exist
    name, ok := m2["name"]
    if ok {
        fmt.Println("Found:", name)       // → Found: Anurag
    } else {
        fmt.Println("Key not found")
    }

    // Check existence only (discard value)
    _, exists := m2["phone"]
    fmt.Println("phone exists?", exists) // → false

    // ── Compare maps ──────────────────────────────────────────────────
    fmt.Println(maps.Equal(m1, m2))      // → false (m1 is empty)
}
```

### Map with complex value types:
```go
// Map of string → slice (common pattern)
courses := map[string][]string{
    "backend": {"Go", "PostgreSQL", "Docker"},
    "frontend": {"React", "TypeScript", "CSS"},
}
fmt.Println(courses["backend"])  // → [Go PostgreSQL Docker]
courses["backend"] = append(courses["backend"], "Redis")

// Map of string → struct
type Student struct {
    Name  string
    Grade int
}

students := map[string]Student{
    "S001": {"Anurag", 90},
    "S002": {"Priya", 85},
}
fmt.Println(students["S001"].Name) // → Anurag
```

### Iterate over a Map:
```go
inventory := map[string]int{
    "apples":  50,
    "bananas": 30,
    "oranges": 20,
}

for product, quantity := range inventory {
    fmt.Printf("%-10s: %d units\n", product, quantity)
}
// Note: map iteration order is RANDOM in Go (by design)
```

### 🎯 Solved Example: Word Frequency Counter
```go
package main

import (
    "fmt"
    "strings"
)

func wordCount(sentence string) map[string]int {
    freq := make(map[string]int)
    words := strings.Fields(sentence) // split by whitespace
    for _, word := range words {
        freq[strings.ToLower(word)]++ // increment; default int is 0, so 0+1=1
    }
    return freq
}

func main() {
    text := "go is great go is fast go is fun"
    freq := wordCount(text)

    for word, count := range freq {
        fmt.Printf("'%s' appears %d times\n", word, count)
    }
    // Output (order may vary):
    // 'go' appears 3 times
    // 'is' appears 3 times
    // 'great' appears 1 times
    // 'fast' appears 1 times
    // 'fun' appears 1 times
}
```

---

## 4. Range — Iterating Over Collections

📁 File: `range/`

```go
package main

import "fmt"

func main() {

    // ── Range over slice ───────────────────────────────────────────────
    nums := []int{10, 20, 30, 40, 50}
    for i, v := range nums {
        fmt.Printf("nums[%d] = %d\n", i, v)
    }

    // ── Range (index only) ─────────────────────────────────────────────
    for i := range nums {
        fmt.Println(i) // → 0 1 2 3 4
    }

    // ── Range (value only, discard index) ─────────────────────────────
    for _, v := range nums {
        fmt.Println(v) // → 10 20 30 40 50
    }

    // ── Range over map ────────────────────────────────────────────────
    scores := map[string]int{"Alice": 90, "Bob": 85}
    for name, score := range scores {
        fmt.Printf("%s scored %d\n", name, score)
    }

    // ── Range over string (iterates runes, not bytes) ──────────────────
    for i, ch := range "Go🚀" {
        fmt.Printf("index=%d, char=%c\n", i, ch)
    }
    // Output:
    // index=0, char=G
    // index=1, char=o
    // index=2, char=🚀

    // ── Range over channel (reads until channel is closed) ────────────
    ch := make(chan int, 3)
    ch <- 1; ch <- 2; ch <- 3
    close(ch)
    for val := range ch {
        fmt.Println(val) // → 1 2 3
    }
}
```

---

## 5. Closures — Functions That Remember

📁 File: `closures/closure.go`

> A **closure** is a function that "closes over" variables from its outer scope.
> The inner function **remembers** those variables even after the outer function returns.

```go
package main

import "fmt"

// counter() returns a function — that inner function is a CLOSURE
// It captures and remembers the 'count' variable
func counter() func() int {
    count := 1   // This variable lives as long as the inner function lives

    return func() int {
        count += 1   // modifies the CAPTURED variable each time
        return count
    }
}

func main() {
    increment := counter()  // increment now holds the inner function
                             // AND its captured 'count' state

    fmt.Println(increment()) // → 2  (count: 1 → 2)
    fmt.Println(increment()) // → 3  (count: 2 → 3)
    fmt.Println(increment()) // → 4  (count: 3 → 4)

    // Create a NEW counter — completely INDEPENDENT state!
    increment2 := counter()
    fmt.Println(increment2()) // → 2  (fresh count, not shared)
}
```

### Why this matters — Independent state:
```go
// Each call to counter() creates a NEW closure with its OWN 'count'
c1 := counter()
c2 := counter()

fmt.Println(c1()) // → 2   (c1's count)
fmt.Println(c1()) // → 3
fmt.Println(c2()) // → 2   (c2's count — INDEPENDENT!)
fmt.Println(c2()) // → 3
fmt.Println(c1()) // → 4   (c1 resumes from 3)
```

### Closure as an adder factory:
```go
// adder returns a closure that adds 'x' to any number passed to it
func adder(x int) func(int) int {
    return func(y int) int {
        return x + y  // 'x' is captured from the outer scope
    }
}

func main() {
    addFive   := adder(5)
    addTen    := adder(10)
    addHundred := adder(100)

    fmt.Println(addFive(3))     // → 8   (5+3)
    fmt.Println(addTen(7))      // → 17  (10+7)
    fmt.Println(addHundred(25)) // → 125 (100+25)
}
```

### 🎯 Real-World Closure: Rate Limiter
```go
package main

import (
    "fmt"
    "time"
)

// makeRateLimiter returns a function that can only be called
// 'maxCalls' times per window, then reports rate-limited
func makeRateLimiter(maxCalls int) func() bool {
    calls := 0
    return func() bool {
        calls++
        if calls > maxCalls {
            return false // rate limited!
        }
        return true
    }
}

func main() {
    limiter := makeRateLimiter(3)

    for i := 0; i < 5; i++ {
        if limiter() {
            fmt.Printf("Request %d: OK\n", i+1)
        } else {
            fmt.Printf("Request %d: RATE LIMITED\n", i+1)
        }
        time.Sleep(100 * time.Millisecond)
    }
}
// Output:
// Request 1: OK
// Request 2: OK
// Request 3: OK
// Request 4: RATE LIMITED
// Request 5: RATE LIMITED
```

---

## 6. Variadic Functions

> Variadic functions accept **zero or more arguments** of the same type using `...`

```go
package main

import "fmt"

// nums ...int means: accept ANY number of int arguments
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {   // nums is treated as []int inside
        total += n
    }
    return total
}

// Mixed — regular param + variadic (variadic must be LAST)
func greet(greeting string, names ...string) {
    for _, name := range names {
        fmt.Printf("%s, %s!\n", greeting, name)
    }
}

func main() {
    fmt.Println(sum(1, 2, 3))           // → 6
    fmt.Println(sum(1, 2, 3, 4, 5))     // → 15
    fmt.Println(sum())                   // → 0

    // Spread a slice into variadic args using ...
    numbers := []int{10, 20, 30}
    fmt.Println(sum(numbers...))         // → 60

    greet("Hello", "Alice", "Bob", "Charlie")
    // → Hello, Alice!
    // → Hello, Bob!
    // → Hello, Charlie!
}
```

---

## 7. Generics (Go 1.18+)

📁 File: `Generics/gene.go`

> Generics allow you to write functions/types that work with ANY type — type-safe reuse.

```go
package main

import "fmt"

// ── Generic Function ─────────────────────────────────────────────────────────
// [T any] means: T can be ANY type
// This function works for []int, []string, []float64 — anything!
func printSlice[T any](items []T) {
    for _, val := range items {
        fmt.Println(val)
    }
}

// [T comparable] means: T must be comparable (supports == and !=)
// This constraint is needed for equality checks
func contains[T comparable](items []T, target T) bool {
    for _, item := range items {
        if item == target {
            return true
        }
    }
    return false
}

// ── Generic Struct ────────────────────────────────────────────────────────────
// stack[T any] works for any element type
type stack[T any] struct {
    elements []T
}

func (s *stack[T]) push(elem T) {
    s.elements = append(s.elements, elem)
}

func (s *stack[T]) pop() (T, bool) {
    var zero T  // zero value of type T
    if len(s.elements) == 0 {
        return zero, false
    }
    last := s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1]
    return last, true
}

func main() {
    // Works with ints
    nums := []int{1, 2, 3, 4, 5}
    printSlice(nums)  // → 1 2 3 4 5

    // Works with strings — same function!
    words := []string{"Go", "is", "awesome"}
    printSlice(words) // → Go is awesome

    // Generic contains check
    fmt.Println(contains(nums, 3))     // → true
    fmt.Println(contains(words, "Java")) // → false

    // Generic stack of strings
    myStack := stack[string]{}
    myStack.push("first")
    myStack.push("second")
    myStack.push("third")

    top, ok := myStack.pop()
    fmt.Println(top, ok) // → third true

    // Generic stack of ints — same type, different T
    numStack := stack[int]{}
    numStack.push(10)
    numStack.push(20)
    val, _ := numStack.pop()
    fmt.Println(val) // → 20
}
```

### Type Constraints:
```go
// Built-in constraints from "golang.org/x/exp/constraints":
// any        → any type (interface{})
// comparable → supports == and !=
// ordered    → supports <, >, <=, >= (ints, floats, strings)

// Custom constraint — type must be int OR float64
type Number interface {
    int | float64
}

func double[T Number](n T) T {
    return n * 2
}

fmt.Println(double(5))     // → 10   (int)
fmt.Println(double(3.14))  // → 6.28 (float64)
```

---

## 8. Error Handling

> Go does NOT have exceptions. Errors are values — returned as the **last return value**.

```go
package main

import (
    "errors"
    "fmt"
)

// ── Returning errors ──────────────────────────────────────────────────────────
// Convention: error is ALWAYS the last return value
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil  // nil means "no error"
}

// ── Custom error type ─────────────────────────────────────────────────────────
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on '%s': %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{Field: "age", Message: "must be non-negative"}
    }
    if age > 150 {
        return &ValidationError{Field: "age", Message: "unrealistic value"}
    }
    return nil
}

func main() {
    // ── Pattern 1: Check error immediately ──────────────────────────
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result) // → Result: 5
    }

    // ── Pattern 2: Handle error and return early ─────────────────────
    _, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)   // → Error: cannot divide by zero
        return
    }

    // ── Pattern 3: Type assertion on errors ─────────────────────────
    err = validateAge(-5)
    var ve *ValidationError
    if errors.As(err, &ve) {
        fmt.Printf("Field: %s, Msg: %s\n", ve.Field, ve.Message)
    }

    // ── fmt.Errorf for wrapping errors ───────────────────────────────
    wrappedErr := fmt.Errorf("processing failed: %w", err)
    fmt.Println(wrappedErr)
}
```

### The Golden Rule of Go Error Handling:
```
func doSomething() (Result, error) {
    1. Try the operation
    2. If error → return zero value + error
    3. If success → return result + nil
}

caller:
    result, err := doSomething()
    if err != nil {
        // handle or return the error — NEVER ignore it
    }
    // use result safely here
```

---

## 9. Practice Exercises

### 🟢 Beginner:
1. Write a function `isEven(n int) bool` and test it for numbers 1–10
2. Write a function that takes a name and returns a greeting string
3. Use a map to store 5 country capitals and print all pairs

### 🟡 Intermediate:
4. Create a closure `multiplier(factor int) func(int) int` that multiplies any number by the factor
5. Write a generic `max[T]` function that returns the larger of two values
6. Implement a `wordFrequency(text string) map[string]int` function

### 🔴 Advanced:
7. Create a function `pipeline` that takes any number of `func(int) int` operations and applies them in sequence: `pipeline(2, double, addFive, square)` → applies double, then addFive, then square
8. Implement a generic `filter[T any](items []T, predicate func(T) bool) []T` function
9. Build a simple in-memory cache using a map + closure that stores results of expensive computations

---

## 🗺️ What's Next?

Continue to **[README_03_OOP_AND_ADVANCED.md](./README_03_OOP_AND_ADVANCED.md)** →
> **Structs | Methods | Interfaces | Embedding | Polymorphism | Dependency Injection**
