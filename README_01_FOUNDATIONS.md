# 🐹 Go (Golang) Mastery Guide — Part 1: Foundations
> **Series**: `README_01_FOUNDATIONS.md` → `README_02_CORE_TYPES.md` → `README_03_OOP_AND_ADVANCED.md` → `README_04_CONCURRENCY_AND_MASTERY.md`

---

## 📚 Table of Contents
1. [How Go Works — Big Picture](#1-how-go-works--big-picture)
2. [Hello World](#2-hello-world)
3. [Variables — All 5 Styles](#3-variables--all-5-styles)
4. [Constants & Enums (iota)](#4-constants--enums-iota)
5. [Data Types Deep Dive](#5-data-types-deep-dive)
6. [Conditionals (if / else / switch)](#6-conditionals-if--else--switch)
7. [Loops — One Ring to Rule Them All](#7-loops--one-ring-to-rule-them-all)
8. [Arrays — Fixed Size Collections](#8-arrays--fixed-size-collections)
9. [Slices — Dynamic Power Arrays](#9-slices--dynamic-power-arrays)
10. [Practice Exercises](#10-practice-exercises)

---

## 1. How Go Works — Big Picture

```
Your Code (.go)
      │
      ▼
  go build         ← compiles into a native binary (no JVM, no interpreter)
      │
      ▼
 ./myapp.exe       ← runs directly on the OS — super fast!

OR: go run main.go ← compiles + runs in one step (dev mode)
```

### Key facts to memorize:
| Fact | Detail |
|------|--------|
| **Statically typed** | Types are checked at compile time, not runtime |
| **Compiled** | Go code compiles to native machine code |
| **Garbage collected** | No manual memory management like C |
| **Concurrency built-in** | Goroutines & channels are first-class citizens |
| **No classes** | Uses structs + methods + interfaces instead |
| **No inheritance** | Uses composition (embedding) instead |
| **Implicit interfaces** | No `implements` keyword needed |

---

## 2. Hello World

📁 File: `hello_world/main.go`

```go
// Every Go file must belong to a package. 'main' is the entry-point package.
package main

// Import the "fmt" package which provides formatted I/O functions
import "fmt"

// main() is the entry-point function — Go starts executing your program from here
func main() {
    // fmt.Println prints the given string to the console followed by a newline
    fmt.Println("Hello World")
}
```

### 🔑 Key Rules:
- Every `.go` file **must** start with `package <name>`
- The `main` package + `func main()` is the program's entry point
- Imports are **mandatory** — unused imports = **compile error**
- `fmt.Println` adds a newline; `fmt.Print` does NOT

### 📦 Common fmt functions:
```go
fmt.Println("Hello")            // → Hello\n
fmt.Print("Hello ")             // → Hello  (no newline)
fmt.Printf("Name: %s Age: %d\n", "Anu", 25) // → Name: Anu Age: 25
fmt.Sprintf("Result: %d", 42)  // → returns string (doesn't print)
```

### 🎯 Format Verbs Cheat Sheet:
| Verb | Meaning | Example |
|------|---------|---------|
| `%s` | string | `"hello"` |
| `%d` | integer | `42` |
| `%f` | float | `3.14` |
| `%.2f` | float (2 decimal) | `3.14` |
| `%v` | any value (default) | `{Anu 25}` |
| `%+v` | struct with field names | `{name:Anu age:25}` |
| `%T` | type of value | `int`, `string` |
| `%t` | boolean | `true` |
| `%p` | pointer address | `0xc000b4008` |

### ▶️ How to Run:
```bash
go run hello_world/main.go          # compile + run
go build hello_world/main.go        # compile into binary only
```

---

## 3. Variables — All 5 Styles

📁 File: `variables/main.go`

```go
package main

import "fmt"

func main() {

    // ── Style 1: Explicit Type ────────────────────────────────────────
    var name string = "golang"
    fmt.Println(name) // Output: golang

    // ── Style 2: Type Inference ───────────────────────────────────────
    // Go sees 'true' and infers type bool automatically
    var isActive = true
    fmt.Println(isActive) // Output: true

    // ── Style 3: Short Declaration (:=) ───────────────────────────────
    // Most commonly used! Declares AND assigns in one step.
    // ONLY works inside functions (not at package level)
    firstName := "Anurag"
    fmt.Println(firstName) // Output: Anurag

    // ── Style 4: Declare first, assign later ──────────────────────────
    var company string   // zero value → "" (empty string)
    company = "Google"
    fmt.Println(company) // Output: Google

    // ── Style 5: Float ────────────────────────────────────────────────
    var price float32 = 45.23
    fmt.Println(price) // Output: 45.23
}
```

### 🔑 Zero Values (default when declared but not assigned):
| Type | Zero Value |
|------|-----------|
| `int`, `float32`, `float64` | `0` |
| `string` | `""` (empty string) |
| `bool` | `false` |
| pointer, slice, map, channel | `nil` |

### Multiple variables in one block:
```go
var (
    x int     = 10
    y float32 = 3.14
    z string  = "Go"
)
fmt.Println(x, y, z) // → 10 3.14 Go
```

### ⚠️ Rules to Remember:
```go
// ❌ := cannot be used at package level
package main
x := 10         // COMPILE ERROR

// ✅ Must use var at package level
var x = 10      // OK

// ❌ Declared but never used = compile error
func main() {
    x := 5      // COMPILE ERROR: x declared but not used
}

// ✅ Use _ (blank identifier) to intentionally discard
_, err := someFunc()    // ignore the first return value
```

### 🎯 Solved Example: Student Profile
```go
package main

import "fmt"

func main() {
    // Create a student profile using different variable styles
    var studentName string = "Anurag Sharma"    // Style 1
    age := 21                                    // Style 3 (most common)
    var gpa float64 = 8.75                      // Style 1 with float
    var isEnrolled = true                       // Style 2

    fmt.Printf("Name    : %s\n", studentName)
    fmt.Printf("Age     : %d\n", age)
    fmt.Printf("GPA     : %.2f\n", gpa)
    fmt.Printf("Enrolled: %t\n", isEnrolled)
}
// Output:
// Name    : Anurag Sharma
// Age     : 21
// GPA     : 8.75
// Enrolled: true
```

---

## 4. Constants & Enums (iota)

📁 File: `constants/` and `Enums/enums.go`

### Constants:
```go
package main

import "fmt"

// const declares a value that CANNOT be changed after declaration
const Pi = 3.14159
const AppName = "GoMaster"
const MaxRetries = 3

// Multiple constants in a block
const (
    StatusOK    = 200
    StatusNotFound = 404
    StatusError = 500
)

func main() {
    fmt.Println(Pi)          // 3.14159
    fmt.Println(AppName)     // GoMaster
    fmt.Println(StatusOK)    // 200
}
```

### Enums using `iota`:
```go
package main

import "fmt"

// iota is a special Go keyword used inside const blocks
// It automatically increments: 0, 1, 2, 3...
type Weekday int

const (
    Sunday    Weekday = iota  // 0
    Monday                    // 1 (iota auto-increments)
    Tuesday                   // 2
    Wednesday                 // 3
    Thursday                  // 4
    Friday                    // 5
    Saturday                  // 6
)

// String-based enum (your actual code style)
type OrderStatus string

const (
    Received  OrderStatus = "received"
    Confirmed OrderStatus = "confirmed"
    Prepared  OrderStatus = "prepared"
    Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
    fmt.Println("Changing order status to", status)
}

func main() {
    fmt.Println(Sunday, Monday, Saturday) // → 0 1 6
    changeOrderStatus(Received)           // → Changing order status to received
    changeOrderStatus(Delivered)          // → Changing order status to delivered
}
```

### 🎯 iota Tricks:
```go
type ByteSize float64

const (
    _           = iota             // ignore first value (0)
    KB ByteSize = 1 << (10 * iota) // 1 << 10 = 1024
    MB                             // 1 << 20 = 1,048,576
    GB                             // 1 << 30 = 1,073,741,824
)

fmt.Println(KB) // 1024
fmt.Println(MB) // 1.048576e+06
fmt.Println(GB) // 1.073741824e+09
```

---

## 5. Data Types Deep Dive

### Integer types:
```go
var a int    = 42       // platform-dependent (32 or 64 bit)
var b int8   = 127      // -128 to 127
var c int16  = 32767    // -32768 to 32767
var d int32  = 2147483647
var e int64  = 9223372036854775807
var f uint   = 42       // unsigned: 0 to max (no negatives)
```

### Float types:
```go
var f32 float32 = 3.14   // 32-bit float (~7 decimal digits precision)
var f64 float64 = 3.14159265358979  // 64-bit float (~15 decimal digits) ← prefer this
```

### String type:
```go
var s string = "Hello, Gopher!"

// String is a sequence of bytes (UTF-8 encoded)
fmt.Println(len(s))        // length in bytes: 14
fmt.Println(s[0])          // byte at index 0: 72 (ASCII for 'H')
fmt.Println(string(s[0]))  // → "H"

// String concatenation
greeting := "Hello" + ", " + "World!"  // → "Hello, World!"

// Multi-line strings using backticks
multiLine := `Line 1
Line 2
Line 3`
```

### Boolean:
```go
var isGo bool = true
fmt.Println(!isGo)         // → false (NOT operator)
fmt.Println(true && false) // → false (AND)
fmt.Println(true || false) // → true  (OR)
```

### Type Conversion (explicit — Go doesn't auto-convert):
```go
var i int = 42
var f float64 = float64(i)    // int → float64
var u uint = uint(f)          // float64 → uint

// String ↔ int conversion (use strconv package)
import "strconv"
s := strconv.Itoa(42)          // int → string: "42"
n, _ := strconv.Atoi("42")     // string → int: 42
```

---

## 6. Conditionals (if / else / switch)

📁 File: `conditional/`

### Basic if/else:
```go
package main

import "fmt"

func main() {
    age := 20

    if age >= 18 {
        fmt.Println("Adult")
    } else if age >= 13 {
        fmt.Println("Teenager")
    } else {
        fmt.Println("Child")
    }
}
```

### if with initialization statement (Go-specific pattern):
```go
// You can declare a variable IN the if statement itself!
// The variable is scoped to the if/else block only.
if score := 85; score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B")   // ← this runs
} else {
    fmt.Println("Grade: C")
}
// 'score' is NOT accessible here
```

### switch statement:
```go
package main

import "fmt"

func main() {
    day := "Monday"

    switch day {
    case "Saturday", "Sunday":   // multiple values in one case
        fmt.Println("Weekend!")
    case "Monday":
        fmt.Println("Start of work week")
    case "Friday":
        fmt.Println("Almost weekend!")
    default:
        fmt.Println("Regular weekday")
    }
}
```

### switch without expression (acts like if/else chain):
```go
score := 75

switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 70:
    fmt.Println("C")  // ← this runs
default:
    fmt.Println("F")
}
```

### 🎯 Solved Example: FizzBuzz
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 20; i++ {
        switch {
        case i%15 == 0:
            fmt.Println("FizzBuzz")
        case i%3 == 0:
            fmt.Println("Fizz")
        case i%5 == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}
```

---

## 7. Loops — One Ring to Rule Them All

📁 File: `loops/`

> In Go, there is **only ONE loop keyword: `for`**. It replaces while, do-while, and foreach.

### Style 1: Classic C-style for loop:
```go
for i := 0; i < 5; i++ {
    fmt.Println(i) // 0, 1, 2, 3, 4
}
```

### Style 2: While loop (just condition):
```go
count := 0
for count < 5 {    // works just like while(count < 5)
    fmt.Println(count)
    count++
}
```

### Style 3: Infinite loop:
```go
for {              // runs forever until break
    fmt.Println("running...")
    break          // use break to exit
}
```

### Style 4: for-range loop (iterate over collections):
```go
// Over a slice
fruits := []string{"Apple", "Banana", "Mango"}
for index, value := range fruits {
    fmt.Printf("fruits[%d] = %s\n", index, value)
}
// Output:
// fruits[0] = Apple
// fruits[1] = Banana
// fruits[2] = Mango

// Over a map
person := map[string]string{"name": "Anurag", "city": "Mumbai"}
for key, value := range person {
    fmt.Printf("%s: %s\n", key, value)
}

// Over a string (iterates over runes/characters)
for i, ch := range "Go!" {
    fmt.Printf("index=%d char=%c\n", i, ch)
}

// Ignore index with _
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```

### break and continue:
```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue    // skip 3, go to next iteration
    }
    if i == 7 {
        break       // stop the loop entirely
    }
    fmt.Println(i)
}
// Output: 0 1 2 4 5 6
```

### 🎯 Solved Example: Sum of first N numbers
```go
package main

import "fmt"

func main() {
    n := 10
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    fmt.Printf("Sum of 1 to %d = %d\n", n, sum) // → Sum of 1 to 10 = 55
}
```

### 🎯 Solved Example: Multiplication Table
```go
package main

import "fmt"

func main() {
    num := 7
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d × %d = %d\n", num, i, num*i)
    }
}
// Output:
// 7 × 1 = 7
// 7 × 2 = 14
// ...
// 7 × 10 = 70
```

---

## 8. Arrays — Fixed Size Collections

📁 File: `Arrays/arrays.go`

> Arrays in Go have a **fixed size** — once declared, the size CANNOT change.

```go
package main

import "fmt"

func main() {

    // Declare array of 4 ints — all initialized to 0 (zero value)
    var nums [4]int
    fmt.Println(len(nums)) // → 4
    nums[0] = 1
    fmt.Println(nums[0])   // → 1

    // Bool array — zero value is false
    var vals [5]bool
    fmt.Println(vals)      // → [false false false false false]

    // String array — zero value is "" (empty string)
    var cars [5]string
    cars[0] = "McLaren"
    fmt.Println(cars)      // → [McLaren    ]

    // Declare and initialize in one line
    numss := [3]int{1, 2, 3}
    fmt.Println(numss)     // → [1 2 3]

    // 2D array — matrix (3 rows × 3 cols)
    grid := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println(grid)      // → [[1 2 3] [4 5 6] [7 8 9]]
    fmt.Println(grid[1][2]) // → 6  (row 1, col 2)
}
```

### Key differences: Array vs Slice:
| Feature | Array | Slice |
|---------|-------|-------|
| Size | Fixed at compile time | Dynamic (can grow) |
| Type | `[3]int` | `[]int` |
| Passing to function | Copies the entire array | Passes a reference |
| Common usage | Rarely used directly | Used everywhere |

---

## 9. Slices — Dynamic Power Arrays

📁 File: `Slices/slices.go`

> Slices are the **most used data structure in Go**. Think of them as dynamic arrays.

```go
package main

import (
    "fmt"
    "slices" // Go 1.21+ — helper functions
)

func main() {

    // ── 1. Nil slice (zero value) ─────────────────────────────────────
    var names []string
    fmt.Println(names == nil) // → true
    fmt.Println(len(names))   // → 0

    // ── 2. make([]T, length, capacity) ───────────────────────────────
    nums := make([]int, 5, 50)  // 5 elements (all 0), capacity 50
    fmt.Println(nums)           // → [0 0 0 0 0]
    fmt.Println(cap(nums))      // → 50

    // ── 3. append() ──────────────────────────────────────────────────
    nums = append(nums, 10, 20, 30)
    fmt.Println(nums)           // → [0 0 0 0 0 10 20 30]

    // ── 4. Slice literal ─────────────────────────────────────────────
    fruits := []string{"Apple", "Banana", "Mango", "Orange"}

    // ── 5. Slicing (sub-slicing) ──────────────────────────────────────
    fmt.Println(fruits[1:3])  // [low:high] → ["Banana", "Mango"]
    fmt.Println(fruits[:2])   // [:high]    → ["Apple", "Banana"]
    fmt.Println(fruits[2:])   // [low:]     → ["Mango", "Orange"]
    fmt.Println(fruits[:])    // [:]        → ALL elements

    // ── 6. copy() — deep copy ─────────────────────────────────────────
    backup := make([]string, len(fruits))
    copy(backup, fruits)
    fruits[0] = "Grapes"         // modify original
    fmt.Println(backup[0])       // → "Apple" (backup unchanged!)

    // ── 7. Compare slices ─────────────────────────────────────────────
    a := []int{1, 2, 3}
    b := []int{1, 2, 3}
    fmt.Println(slices.Equal(a, b)) // → true

    // ── 8. 2D slice ─────────────────────────────────────────────────
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    fmt.Println(matrix[0][1]) // → 2
}
```

### How slices work internally:
```
Slice header:
┌─────────────────┐
│  pointer  ──────┼─→  [1][2][3][4][5]  (underlying array)
│  length = 3     │     ↑
│  capacity = 5   │     starts here
└─────────────────┘
```

### 🎯 Solved Example: Shopping Cart
```go
package main

import "fmt"

func main() {
    cart := []string{}  // empty cart

    // Add items
    cart = append(cart, "Laptop")
    cart = append(cart, "Mouse", "Keyboard")
    fmt.Println("Cart:", cart)           // → [Laptop Mouse Keyboard]
    fmt.Println("Total items:", len(cart)) // → 3

    // Remove item at index 1 (Mouse)
    // Trick: append elements before + after the index
    cart = append(cart[:1], cart[2:]...)
    fmt.Println("After remove:", cart)   // → [Laptop Keyboard]

    // Iterate over cart
    for i, item := range cart {
        fmt.Printf("%d. %s\n", i+1, item)
    }
}
```

---

## 10. Practice Exercises

### 🟢 Beginner:
1. Print numbers 1 to 100; for multiples of 3 print "Fizz", 5 print "Buzz", both print "FizzBuzz"
2. Declare variables for your personal profile (name, age, GPA, isStudent) and print them formatted
3. Create an array of your 5 favourite movies and print them with their index

### 🟡 Intermediate:
4. Create a slice, add 10 numbers using a loop, then find the sum and average
5. Write a program that uses a switch to convert a number (1-7) to the day of the week
6. Create a const block with OrderStatus enum (Pending/Processing/Shipped/Delivered) and print each

### 🔴 Advanced:
7. Write a function that removes duplicates from a `[]int` slice
8. Create a 2D slice representing a 3×3 tic-tac-toe board and print it nicely
9. Use `iota` to create a `Permission` type with Read=1, Write=2, Execute=4 (bit flags)

---

## 🗺️ What's Next?

Continue to **[README_02_CORE_TYPES.md](./README_02_CORE_TYPES.md)** →
> **Functions | Pointers | Maps | Structs | Methods | Closures | Generics**
