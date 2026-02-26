# 🐹 Go (Golang) Mastery Guide — Part 4: Concurrency & Mastery
> **Series**: `README_01_FOUNDATIONS.md` → `README_02_CORE_TYPES.md` → `README_03_OOP_AND_ADVANCED.md` → **`README_04_CONCURRENCY_AND_MASTERY.md`**

---

## 📚 Table of Contents
1. [Goroutines — Lightweight Concurrency](#1-goroutines--lightweight-concurrency)
2. [sync.WaitGroup — Wait for Goroutines](#2-syncwaitgroup--wait-for-goroutines)
3. [Channels — Goroutine Communication](#3-channels--goroutine-communication)
4. [Buffered Channels](#4-buffered-channels)
5. [Channel Direction (Read-only / Write-only)](#5-channel-direction-read-only--write-only)
6. [select — Multi-Channel Operations](#6-select--multi-channel-operations)
7. [sync.Mutex — Prevent Race Conditions](#7-syncmutex--prevent-race-conditions)
8. [Real-world Patterns](#8-real-world-patterns)
9. [Go Modules & Project Structure](#9-go-modules--project-structure)
10. [Quick Reference Cheat Sheet](#10-quick-reference-cheat-sheet)
11. [Master-Level Practice Exercises](#11-master-level-practice-exercises)

---

## 1. Goroutines — Lightweight Concurrency

📁 File: `Goroutines/gor.go`

> A **goroutine** is a lightweight thread managed by the Go runtime.
> Starting thousands of goroutines is perfectly normal in Go (each starts with ~2KB stack).
> Launch with `go functionName(args)`.

### Basic Goroutine:
```go
package main

import (
    "fmt"
    "time"
)

func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

func main() {
    // Without 'go' — runs synchronously, blocks main goroutine
    greet("Alice")   // → Hello, Alice!

    // With 'go' — runs concurrently in a new goroutine
    go greet("Bob")  // starts BUT main might exit before it runs!

    // Without proper synchronization, main exits before goroutine executes
    // BAD solution: time.Sleep (never rely on this in production!)
    time.Sleep(time.Millisecond * 100)
    fmt.Println("Main done")
}
```

### Goroutine Lifecycle:
```
Main goroutine starts:
│
├─→ go greet("Bob") ──────────────────────────────→ Goroutine: Hello, Bob!
│                                                         ↑
│                                                    (scheduled by runtime)
│
├── fmt.Println("Main done")
│
└── main() returns → ALL goroutines are killed!
```

### Anonymous goroutine:
```go
func main() {
    // Immediately invoked goroutine (IIFE)
    go func() {
        fmt.Println("I'm an anonymous goroutine!")
    }()

    // Pass values to avoid closure capture issues
    for i := 0; i < 5; i++ {
        go func(id int) {  // pass i as id (NOT capturing i from outer scope)
            fmt.Println("Goroutine:", id)
        }(i)
    }

    time.Sleep(time.Second)
}
```

### ⚠️ Common Goroutine Gotcha — Closure Capture:
```go
// ❌ WRONG — all goroutines capture the SAME 'i' variable!
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i)  // prints 3, 3, 3 instead of 0, 1, 2!
    }()
}

// ✅ CORRECT — pass 'i' as parameter
for i := 0; i < 3; i++ {
    go func(id int) {   // 'id' is a fresh copy for each goroutine
        fmt.Println(id) // prints 0, 1, 2 (in any order)
    }(i)
}
```

---

## 2. sync.WaitGroup — Wait for Goroutines

📁 File: `Goroutines/gor.go`

> `sync.WaitGroup` lets you wait for a collection of goroutines to finish.
> Three methods: `Add(n)`, `Done()`, `Wait()`

```go
package main

import (
    "fmt"
    "sync"
)

// task simulates some work done by a goroutine
func task(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // ALWAYS call Done() when this goroutine finishes
                     // defer ensures it runs even if the function panics

    fmt.Printf("Task %d: started\n", id)
    // ... simulate work ...
    fmt.Printf("Task %d: done\n", id)
}

func main() {
    var wg sync.WaitGroup  // zero value is ready to use

    for i := 0; i < 10; i++ {
        wg.Add(1)          // tell WaitGroup: "one more goroutine to wait for"
        go task(i, &wg)    // start goroutine — pass POINTER to wg (must share state)
    }

    wg.Wait()              // BLOCK until all goroutines call Done()
    fmt.Println("All tasks completed!")
}
// Output (order varies — it's concurrent!):
// Task 3: started
// Task 0: started
// Task 7: started
// ... (non-deterministic order)
// All tasks completed!
```

### WaitGroup + Return values via channel:
```go
package main

import (
    "fmt"
    "sync"
)

func processItem(id int, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    result := fmt.Sprintf("Processed item %d", id)
    results <- result
}

func main() {
    var wg sync.WaitGroup
    results := make(chan string, 10)  // buffered channel

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go processItem(i, results, &wg)
    }

    // Close channel when all goroutines finish
    go func() {
        wg.Wait()
        close(results)
    }()

    // Collect all results
    for result := range results {
        fmt.Println(result)
    }
}
```

---

## 3. Channels — Goroutine Communication

📁 File: `Channels/chan.go`

> **"Don't communicate by sharing memory; share memory by communicating."** — Go Proverb
> Channels are typed conduits through which goroutines send and receive values.

```
Goroutine A                   Goroutine B
    │                              │
    │  ch <- "hello"               │
    │──────────→ [ channel ] ──────│→  msg := <-ch
    │            (pipe)            │
```

### Channel Basics:
```go
package main

import "fmt"

func main() {
    // Create an unbuffered channel of type string
    // make(chan Type)
    message := make(chan string)

    // Send MUST be in a goroutine because unbuffered channels BLOCK
    // until the receiver is ready
    go func() {
        message <- "ping"  // send "ping" to the channel
    }()

    msg := <-message       // receive — BLOCKS until data arrives
    fmt.Println(msg)       // → ping
}
```

### Channel for synchronization (done channel pattern):
```go
package main

import "fmt"

func task(done chan bool) {
    defer func() {
        done <- true  // signal: "I'm finished"
    }()
    fmt.Println("Processing...")
    // ... long work ...
}

func main() {
    done := make(chan bool)
    go task(done)

    <-done  // BLOCK here until task sends 'true'
    fmt.Println("Task finished!")
}
```

### Channel for results (sum example):
```go
package main

import "fmt"

func sum(result chan int, a int, b int) {
    result <- a + b  // send result through channel
}

func main() {
    result := make(chan int)

    go sum(result, 4, 5)    // runs concurrently

    res := <-result         // wait and receive the result
    fmt.Println("Sum:", res) // → Sum: 9
}
```

### Iterating over a channel with range:
```go
package main

import "fmt"

func processChan(numChan chan int) {
    // 'range' reads from channel until it's CLOSED
    for num := range numChan {
        fmt.Println("Processing number:", num)
    }
    fmt.Println("Channel closed, done processing")
}

func main() {
    numChan := make(chan int, 5)

    // Send values
    go func() {
        for i := 1; i <= 5; i++ {
            numChan <- i
        }
        close(numChan)  // MUST close so range loop terminates!
    }()

    processChan(numChan)
}
// Output:
// Processing number: 1
// Processing number: 2
// Processing number: 3
// Processing number: 4
// Processing number: 5
// Channel closed, done processing
```

---

## 4. Buffered Channels

> An unbuffered channel blocks until BOTH sender and receiver are ready.
> A **buffered channel** has internal storage — send doesn't block until buffer is full.

```go
package main

import "fmt"

func main() {
    // Unbuffered: blocks on send until receiver is ready
    ch1 := make(chan int)    // capacity = 0

    // Buffered: can hold up to N items before blocking
    ch2 := make(chan int, 3) // capacity = 3

    // Buffered send is NON-BLOCKING while buffer has space
    ch2 <- 10    // OK — buffer: [10]
    ch2 <- 20    // OK — buffer: [10, 20]
    ch2 <- 30    // OK — buffer: [10, 20, 30]
    // ch2 <- 40 // DEADLOCK — buffer is full, no receiver!

    fmt.Println(<-ch2) // → 10 (FIFO order)
    fmt.Println(<-ch2) // → 20
    fmt.Println(<-ch2) // → 30

    // Real-world: Email queue
    emailChan := make(chan string, 100)
    done := make(chan bool)

    // Consumer goroutine
    go func() {
        defer func() { done <- true }()
        for email := range emailChan {
            fmt.Println("Sending email to:", email)
        }
    }()

    // Producer
    for i := 0; i < 5; i++ {
        emailChan <- fmt.Sprintf("user%d@gmail.com", i)
    }
    close(emailChan)

    <-done  // wait for all emails to be "sent"
}
```

### Buffered Channel Capacity:
```go
ch := make(chan int, 5)
fmt.Println(len(ch), cap(ch)) // → 0 5  (0 items, capacity 5)
ch <- 1; ch <- 2
fmt.Println(len(ch), cap(ch)) // → 2 5  (2 items, capacity 5)
```

---

## 5. Channel Direction (Read-only / Write-only)

> You can restrict a channel to only sending or only receiving.
> This makes function contracts clear and prevents accidental misuse.

```go
package main

import "fmt"

// chan<- string  → write-only (can only SEND into this channel)
func emailSender(emailChan chan<- string, done chan<- bool) {
    defer func() { done <- true }()

    emails := []string{"a@mail.com", "b@mail.com", "c@mail.com"}
    for _, email := range emails {
        emailChan <- email  // ✅ can send
        // msg := <-emailChan // ❌ COMPILE ERROR: cannot receive from send-only channel
    }
    close(emailChan)
}

// <-chan string   → read-only (can only RECEIVE from this channel)
func emailReceiver(emailChan <-chan string, done chan<- bool) {
    defer func() { done <- true }()

    for email := range emailChan {
        fmt.Println("Received email:", email)
        // emailChan <- "reply" // ❌ COMPILE ERROR: cannot send to receive-only channel
    }
}

func main() {
    emailChan := make(chan string, 10)
    done1 := make(chan bool)
    done2 := make(chan bool)

    go emailSender(emailChan, done1)    // passes emailChan as chan<-string
    go emailReceiver(emailChan, done2)  // passes emailChan as <-chan string

    <-done1
    <-done2
}
```

---

## 6. select — Multi-Channel Operations

📁 File: `Channels/chan.go`

> `select` waits on multiple channel operations, like a switch for channels.
> It picks whichever channel is ready first — if multiple are ready, picks one at random.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    chan1 := make(chan int)
    chan2 := make(chan string)

    // Send on both channels concurrently
    go func() { chan1 <- 42 }()
    go func() { chan2 <- "pong" }()

    // select reads from whichever channel has data first
    for i := 0; i < 2; i++ {
        select {
        case val := <-chan1:
            fmt.Println("Received from chan1:", val)   // → 42
        case val := <-chan2:
            fmt.Println("Received from chan2:", val)   // → pong
        }
    }
}
```

### select with default (non-blocking):
```go
func main() {
    ch := make(chan int, 1)

    // Non-blocking select with 'default'
    select {
    case val := <-ch:
        fmt.Println("Received:", val)
    default:
        fmt.Println("No data available — proceeding without blocking")
    }
}
```

### select with timeout:
```go
func main() {
    result := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)    // simulates slow operation
        result <- "done"
    }()

    select {
    case res := <-result:
        fmt.Println("Got result:", res)
    case <-time.After(1 * time.Second):  // timeout after 1 second
        fmt.Println("Timeout! Operation took too long.")
    }
}
```

### Fan-out / Fan-in pattern:
```go
package main

import (
    "fmt"
    "sync"
)

// Fan-out: distribute work to multiple goroutines
func fanOut(input <-chan int, workers int) []<-chan int {
    channels := make([]<-chan int, workers)
    for i := 0; i < workers; i++ {
        out := make(chan int)
        channels[i] = out
        go func(ch chan int) {
            for v := range input {
                ch <- v * 2  // double each value
            }
            close(ch)
        }(out)
    }
    return channels
}

// Fan-in: merge multiple channels into one
func merge(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    merged := make(chan int, 10)

    output := func(c <-chan int) {
        defer wg.Done()
        for v := range c {
            merged <- v
        }
    }

    wg.Add(len(channels))
    for _, c := range channels {
        go output(c)
    }

    go func() {
        wg.Wait()
        close(merged)
    }()

    return merged
}
```

---

## 7. sync.Mutex — Prevent Race Conditions

> When multiple goroutines access shared memory, use `sync.Mutex` to prevent race conditions.
> `Lock()` → exclusive access. `Unlock()` → release.

```go
package main

import (
    "fmt"
    "sync"
)

// ── WITHOUT Mutex (RACE CONDITION) ────────────────────────────────────────────
// counter++ is NOT atomic — goroutines can read/write simultaneously!
var count int

// ── WITH Mutex (SAFE) ─────────────────────────────────────────────────────────
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()         // acquire the lock — only ONE goroutine at a time
    defer c.mu.Unlock() // always release when done (defer is safe here!)
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

func main() {
    counter := &SafeCounter{}
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }

    wg.Wait()
    fmt.Println("Final count:", counter.Value()) // → 1000 (always correct with mutex)
}
```

### sync.RWMutex for read-heavy workloads:
```go
// RWMutex: multiple goroutines can READ simultaneously
// But only ONE goroutine can WRITE at a time
type Cache struct {
    mu    sync.RWMutex
    store map[string]string
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()          // exclusive write lock
    defer c.mu.Unlock()
    c.store[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()         // shared read lock (multiple readers OK!)
    defer c.mu.RUnlock()
    val, ok := c.store[key]
    return val, ok
}
```

---

## 8. Real-world Patterns

### Worker Pool Pattern:
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Job represents work to be done
type Job struct {
    ID   int
    Data string
}

// Worker processes jobs from the queue
func worker(id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        // Simulate work
        time.Sleep(10 * time.Millisecond)
        result := fmt.Sprintf("Worker %d processed job %d: %s", id, job.ID, job.Data)
        results <- result
    }
}

func main() {
    const numWorkers = 5
    const numJobs = 20

    jobs := make(chan Job, numJobs)
    results := make(chan string, numJobs)
    var wg sync.WaitGroup

    // Start workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- Job{ID: j, Data: fmt.Sprintf("task-%d", j)}
    }
    close(jobs)

    // Close results when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()

    // Collect results
    for result := range results {
        fmt.Println(result)
    }
}
```

### Pipeline Pattern:
```go
// Stage 1: generate numbers
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Stage 2: square them
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    // Build pipeline: generate → square
    nums := generate(2, 3, 4, 5)
    results := square(nums)

    for result := range results {
        fmt.Println(result) // → 4 9 16 25
    }
}
```

---

## 9. Go Modules & Project Structure

### Initialize a new module:
```bash
# In your project root:
go mod init github.com/yourname/projectname

# This creates go.mod — tracks dependencies
```

### go.mod file:
```
module github.com/anurag/myapp

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    gorm.io/gorm v1.25.0
)
```

### Common Go commands:
```bash
go run main.go          # compile + run
go build ./...          # build all packages
go test ./...           # run all tests
go mod tidy             # clean up dependencies
go get package@version  # add a dependency
go vet ./...            # check for common mistakes
go fmt ./...            # format all code
```

### Recommended project structure:
```
myproject/
├── cmd/
│   └── main.go           ← entry point
├── internal/             ← private packages (not importable externally)
│   ├── handler/
│   ├── service/
│   └── repository/
├── pkg/                  ← public packages (importable externally)
├── api/                  ← API definitions, proto files
├── configs/              ← configuration files
├── go.mod
├── go.sum
└── README.md
```

### Writing Tests:
```go
// filename: calculator_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    got := add(2, 3)
    want := 5
    if got != want {
        t.Errorf("add(2,3) = %d; want %d", got, want)
    }
}

func TestAddNegative(t *testing.T) {
    got := add(-1, -2)
    if got != -3 {
        t.Errorf("Expected -3, got %d", got)
    }
}

// Table-driven tests (idiomatic Go testing)
func TestAddTable(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"zeros", 0, 0, 0},
        {"mixed", 10, -5, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

---

## 10. Quick Reference Cheat Sheet

### Goroutine & Channel Summary:
```
┌─────────────────────────────────────────────────────────────────────────┐
│  Goroutines & Channels — Quick Reference                                │
├──────────────────────────┬──────────────────────────────────────────────┤
│  go func()               │  Launch goroutine                            │
│  make(chan T)             │  Unbuffered channel (synchronous)            │
│  make(chan T, N)          │  Buffered channel (async up to N items)      │
│  ch <- value             │  Send to channel (blocks if full)            │
│  value := <-ch           │  Receive from channel (blocks if empty)      │
│  close(ch)               │  Close channel (no more sends)               │
│  for v := range ch       │  Receive until channel closed                │
│  select { case: }        │  Wait on multiple channels                   │
│  time.After(d)           │  Channel that fires after duration           │
├──────────────────────────┴──────────────────────────────────────────────┤
│  sync.WaitGroup          │  Wait for N goroutines to finish             │
│  wg.Add(1)               │  Increment counter                           │
│  wg.Done()               │  Decrement counter (defer wg.Done())        │
│  wg.Wait()               │  Block until counter = 0                    │
│  sync.Mutex              │  Exclusive lock for shared memory            │
│  mu.Lock() / mu.Unlock() │  Acquire/release the lock                   │
│  sync.RWMutex            │  Multiple readers OR one writer              │
└─────────────────────────────────────────────────────────────────────────┘
```

### Go Operators:
```go
// Arithmetic
+  -  *  /  %    // add, subtract, multiply, divide, modulo
++  --           // increment, decrement (statement, not expression in Go)

// Comparison
==  !=  <  >  <=  >=

// Logical
&&  ||  !        // AND, OR, NOT

// Bitwise
&  |  ^  <<  >>  // AND, OR, XOR, left-shift, right-shift

// Assignment
=  :=            // assign, short declare+assign
+=  -=  *=  /=  %=   // compound assignment
```

### Built-in Functions Summary:
```go
len(x)           // length of string/slice/map/array/channel
cap(x)           // capacity of slice/channel
make(T, args)    // create slice/map/channel
new(T)           // allocate zero-value T, returns *T
append(s, ...v)  // append to slice
copy(dst, src)   // copy slice elements
delete(m, key)   // delete from map
close(ch)        // close a channel
panic(v)         // stop execution, unwind stack
recover()        // recover from panic (use with defer)
print/println    // low-level (use fmt.Print instead)
```

### Panic and Recover:
```go
func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    return a / b, nil  // will panic if b == 0
}

func main() {
    result, err := safeDiv(10, 0)
    if err != nil {
        fmt.Println("Error:", err) // → Error: recovered from panic: runtime error: integer divide by zero
    } else {
        fmt.Println(result)
    }
}
```

---

## 11. Master-Level Practice Exercises

### 🟢 Goroutine Basics:
1. Launch 5 goroutines that each print "Goroutine N working" and wait for all to finish using WaitGroup
2. Create a goroutine that counts to 10 slowly; make main wait using a done channel
3. Write a concurrent ping-pong: two goroutines passing "ping" and "pong" back and forth 5 times

### 🟡 Intermediate Concurrency:
4. Build a **concurrent number squarer**: given `[]int{1..10}`, square each number in a separate goroutine and collect results in order
5. Implement a **timeout wrapper**: a function that calls any `func() string` but returns an error if it takes more than N milliseconds
6. Create a **pub-sub system**: a publisher sends messages to a channel, multiple subscriber goroutines each read and process them

### 🔴 Advanced Patterns:
7. Build a full **worker pool**: 3 workers process 20 jobs from a queue concurrently; collect and print all results
8. Implement a **concurrent web scraper**: fetch 10 URLs concurrently (using `http.Get`), collect results, handle errors from each
9. Build a **rate-limited API client**: can only make 5 requests per second using a ticker and channel

---

## 🏆 Go Mastery Roadmap

```
Level 1 — Foundations ✅ (README_01)
  ✓ Variables, Constants, Types
  ✓ Conditionals, Loops
  ✓ Arrays, Slices

Level 2 — Core Types ✅ (README_02)
  ✓ Functions (multi-return, higher-order)
  ✓ Pointers
  ✓ Maps
  ✓ Closures
  ✓ Generics
  ✓ Error Handling

Level 3 — OOP Patterns ✅ (README_03)
  ✓ Structs + Methods
  ✓ Interfaces + Polymorphism
  ✓ Embedding (Composition)
  ✓ Dependency Injection

Level 4 — Concurrency ✅ (README_04)
  ✓ Goroutines
  ✓ Channels (unbuffered, buffered, directional)
  ✓ WaitGroup
  ✓ Mutex / RWMutex
  ✓ Select
  ✓ Worker Pool, Pipeline, Fan-out/Fan-in

Level 5 — Production (Your Next Steps) 🎯
  → REST APIs with net/http or Gin
  → Database: GORM + PostgreSQL
  → gRPC services
  → Docker + Kubernetes deployment
  → Profiling & Benchmarking
  → Advanced testing (mocks, integration tests)
```

---

## 📚 Recommended Resources

| Resource | Link | Level |
|----------|------|-------|
| **Go Tour** | [go.dev/tour](https://go.dev/tour) | Beginner |
| **Effective Go** | [go.dev/doc/effective_go](https://go.dev/doc/effective_go) | Intermediate |
| **Go by Example** | [gobyexample.com](https://gobyexample.com) | Beginner–Intermediate |
| **Go Proverbs** | [go-proverbs.github.io](https://go-proverbs.github.io) | All |
| **Go Playground** | [play.golang.org](https://play.golang.org) | Practice online |
| **Standard Library Docs** | [pkg.go.dev/std](https://pkg.go.dev/std) | Reference |
| **Concurrency Patterns** | [talks.golang.org/2012/concurrency.slide](https://talks.golang.org/2012/concurrency.slide) | Advanced |

---

> 🎓 **You've covered everything!** Practice every concept by running the code in your own directory.
> The fastest way to master Go is: **read** → **type** → **modify** → **break** → **fix** → **repeat**.
