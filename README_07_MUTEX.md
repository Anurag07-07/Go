# 🔒 README_07 — Mutex & Concurrency Safety

> **File Reference:** `Mutex/mu.go`, `Goroutines/gor.go`

---

## 1. The Problem — Race Conditions

When multiple goroutines access and modify the **same variable** at the same time,
the result is undefined — this is called a **race condition**.

### ❌ Unsafe Code (Race Condition)

```go
package main

import (
    "fmt"
    "sync"
)

var counter int // shared variable

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    counter++ // ← NOT safe! Multiple goroutines here = race condition
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go increment(&wg)
    }
    wg.Wait()
    fmt.Println(counter) // Could print anything: 887, 943, 1000, ...
}
```

**Why?** `counter++` is **not atomic** — it's actually 3 steps:
1. Read `counter`
2. Add 1
3. Write back

Two goroutines can read the same value simultaneously and both write back `old+1`, losing one increment.

### Detect Race Conditions

```bash
go run -race main.go
```
Go's built-in race detector will warn you.

---

## 2. `sync.Mutex` — The Solution

A **Mutex** (Mutual Exclusion) is a lock that ensures only **one goroutine** can access a critical section at a time.

```go
var mu sync.Mutex

mu.Lock()   // Acquire the lock — other goroutines block here
// ... critical section (only one goroutine at a time) ...
mu.Unlock() // Release the lock — next goroutine can proceed
```

### ✅ Safe Code with Mutex

```go
package main

import (
    "fmt"
    "sync"
)

var (
    counter int
    mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    mu.Lock()         // Only one goroutine runs this at a time
    counter++
    mu.Unlock()       // Release so others can proceed
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go increment(&wg)
    }
    wg.Wait()
    fmt.Println(counter) // Always prints: 1000 ✅
}
```

---

## 3. Your Code Explained — `Mutex/mu.go`

```go
package main

import (
    "fmt"
    "sync"
)

// post holds shared state
type post struct {
    views int
    mu    sync.Mutex // Mutex is embedded inside the struct (best practice!)
}

// inc increments views safely
func (p *post) inc(wg *sync.WaitGroup) {
    defer func() {
        p.mu.Unlock() // Always unlock, even if panic
        wg.Done()
    }()
    p.mu.Lock()   // Lock BEFORE accessing views
    p.views += 1  // Safe: only one goroutine here at a time
}

func main() {
    var wg sync.WaitGroup
    myPost := post{views: 0}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go myPost.inc(&wg) // 100 goroutines, all safe!
    }

    wg.Wait()
    fmt.Println(myPost.views) // Always prints: 100 ✅
}
```

**Key design decisions:**
1. `mu sync.Mutex` is **inside the struct** — groups the data it protects
2. Receiver is a **pointer** `*post` — mutex must not be copied
3. `defer` unlocks **even if there's a panic** — prevents deadlock

---

## 4. Always `defer mu.Unlock()`

```go
// ❌ Risky — if something panics, Unlock never runs (deadlock!)
mu.Lock()
doSomething()
mu.Unlock()

// ✅ Safe — defer guarantees Unlock runs no matter what
mu.Lock()
defer mu.Unlock()
doSomething()
```

> **Golden Rule:** Lock → defer Unlock → do work

---

## 5. `sync.RWMutex` — Read/Write Mutex

For data that is **read often but written rarely**, `RWMutex` is more efficient:
- Multiple goroutines can **read simultaneously**
- Only one goroutine can **write** (and blocks all readers)

```go
var rwmu sync.RWMutex

// For READING:
rwmu.RLock()
defer rwmu.RUnlock()
// ... read shared data ...

// For WRITING:
rwmu.Lock()
defer rwmu.Unlock()
// ... write shared data ...
```

### Example — Concurrent Cache

```go
package main

import (
    "fmt"
    "sync"
)

type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func NewCache() *Cache {
    return &Cache{data: make(map[string]string)}
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()         // Exclusive write lock
    defer c.mu.Unlock()
    c.data[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()         // Shared read lock — multiple readers OK
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

func main() {
    cache := NewCache()
    var wg sync.WaitGroup

    // 5 writers
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
        }(i)
    }

    // 20 readers
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            key := fmt.Sprintf("key%d", i%5)
            if val, ok := cache.Get(key); ok {
                fmt.Println("Got:", key, "=", val)
            }
        }(i)
    }

    wg.Wait()
}
```

---

## 6. Mutex vs Channel — When to Use Which

| Scenario | Use |
|----------|-----|
| Protecting shared state (counter, map, struct) | **Mutex** |
| Passing data between goroutines | **Channel** |
| Simple flag / toggle | **Mutex** |
| Pipelining / producer-consumer | **Channel** |
| Caching with many readers | **RWMutex** |

> Go motto: *"Don't communicate by sharing memory; share memory by communicating."*
> But sometimes Mutex is simpler and clearer — use what fits.

---

## 7. `sync.Once` — Run Exactly Once

For initialization that should only happen once (e.g., singleton):

```go
package main

import (
    "fmt"
    "sync"
)

var (
    instance *Config
    once     sync.Once
)

type Config struct {
    DSN string
}

func GetConfig() *Config {
    once.Do(func() {
        // This block runs ONLY ONCE, even with 1000 goroutines
        instance = &Config{DSN: "postgres://localhost/mydb"}
        fmt.Println("Config initialized!")
    })
    return instance
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            cfg := GetConfig()
            fmt.Println("Using DSN:", cfg.DSN)
        }()
    }
    wg.Wait()
    // "Config initialized!" prints ONCE
    // "Using DSN: ..." prints 10 times
}
```

---

## 8. Common Mistakes

### ❌ Mistake 1 — Copying a Mutex

```go
// WRONG — copying a mutex breaks it
myPost := post{views: 0}
postCopy := myPost         // ← copies the mutex too!
go postCopy.inc(&wg)       // ← uses broken copy

// FIX — always use pointer
go myPost.inc(&wg)          // ← receiver is *post, uses original mutex
```

### ❌ Mistake 2 — Locking Twice (Deadlock)

```go
func (p *post) getViews() int {
    p.mu.Lock()
    defer p.mu.Unlock()
    return p.inc_wrong() // ← if inc_wrong also calls Lock() → DEADLOCK!
}
```

**Fix:** Use separate internal methods that assume the lock is already held.

### ❌ Mistake 3 — Not Unlocking on Error

```go
// WRONG
mu.Lock()
data, err := riskyOperation()
if err != nil {
    return err // ← Unlock never called! Everything else will deadlock.
}
mu.Unlock()

// FIX: always defer
mu.Lock()
defer mu.Unlock()           // guaranteed
data, err := riskyOperation()
if err != nil {
    return err              // safe — defer runs
}
```

---

## ✅ Solved Examples

### Example 1 — Safe Bank Account

```go
package main

import (
    "errors"
    "fmt"
    "sync"
)

type BankAccount struct {
    balance float64
    mu      sync.Mutex
}

func (a *BankAccount) Deposit(amount float64) {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.balance += amount
}

func (a *BankAccount) Withdraw(amount float64) error {
    a.mu.Lock()
    defer a.mu.Unlock()
    if amount > a.balance {
        return errors.New("insufficient funds")
    }
    a.balance -= amount
    return nil
}

func (a *BankAccount) Balance() float64 {
    a.mu.RLock()  // only reading — use RLock if RWMutex
    defer a.mu.RUnlock()
    // Note: sync.Mutex doesn't have RLock — use RWMutex for this pattern
    return a.balance
}

func main() {
    account := &BankAccount{balance: 1000}
    var wg sync.WaitGroup

    // 50 deposits of 10 each
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            account.Deposit(10)
        }()
    }

    // 20 withdrawals of 20 each
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            err := account.Withdraw(20)
            if err != nil {
                fmt.Println("Error:", err)
            }
        }()
    }

    wg.Wait()
    fmt.Printf("Final balance: %.2f\n", account.balance)
    // 1000 + 500 - 400 = 1100.00 ✅
}
```

---

### Example 2 — Concurrent Word Counter (from files)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "sync"
)

type WordCounter struct {
    counts map[string]int
    mu     sync.Mutex
}

func (wc *WordCounter) Add(word string) {
    wc.mu.Lock()
    defer wc.mu.Unlock()
    wc.counts[word]++
}

func (wc *WordCounter) CountFile(filename string, wg *sync.WaitGroup) {
    defer wg.Done()

    f, err := os.Open(filename)
    if err != nil {
        return
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        word := strings.ToLower(scanner.Text())
        wc.Add(word)
    }
}

func main() {
    wc := &WordCounter{counts: make(map[string]int)}
    var wg sync.WaitGroup

    files := []string{"file1.txt", "file2.txt", "file3.txt"}
    for _, f := range files {
        wg.Add(1)
        go wc.CountFile(f, &wg)
    }
    wg.Wait()

    for word, count := range wc.counts {
        fmt.Printf("%s: %d\n", word, count)
    }
}
```

---

## 🏋️ Practice Problems

### Problem 1 — Post View Counter (extend your `mu.go`)
Add a `GetViews()` method to the `post` struct that safely returns the view count.
Spawn 100 goroutines to increment and 50 goroutines to read simultaneously.

<details>
<summary>💡 Solution</summary>

```go
package main

import (
    "fmt"
    "sync"
)

type post struct {
    views int
    mu    sync.RWMutex
}

func (p *post) inc(wg *sync.WaitGroup) {
    defer wg.Done()
    p.mu.Lock()
    defer p.mu.Unlock()
    p.views++
}

func (p *post) GetViews() int {
    p.mu.RLock()
    defer p.mu.RUnlock()
    return p.views
}

func main() {
    var wg sync.WaitGroup
    myPost := post{views: 0}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go myPost.inc(&wg)
    }

    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Println("Views so far:", myPost.GetViews())
        }()
    }

    wg.Wait()
    fmt.Println("Final views:", myPost.GetViews()) // 100
}
```
</details>

---

### Problem 2 — Thread-Safe Stack
Implement a concurrent-safe stack with:
- `Push(val int)`
- `Pop() (int, bool)`
- `Peek() (int, bool)`

<details>
<summary>💡 Solution</summary>

```go
package main

import (
    "fmt"
    "sync"
)

type Stack struct {
    data []int
    mu   sync.Mutex
}

func (s *Stack) Push(val int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data = append(s.data, val)
}

func (s *Stack) Pop() (int, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    if len(s.data) == 0 {
        return 0, false
    }
    n := len(s.data)
    val := s.data[n-1]
    s.data = s.data[:n-1]
    return val, true
}

func (s *Stack) Peek() (int, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    if len(s.data) == 0 {
        return 0, false
    }
    return s.data[len(s.data)-1], true
}

func main() {
    s := &Stack{}
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(v int) {
            defer wg.Done()
            s.Push(v)
        }(i)
    }
    wg.Wait()

    for {
        val, ok := s.Pop()
        if !ok {
            break
        }
        fmt.Println("Popped:", val)
    }
}
```
</details>

---

### Problem 3 — Race Condition Fix
The following code has a race condition. Fix it using `sync.Mutex`:

```go
// BROKEN — fix this
var total int

func addToTotal(n int, wg *sync.WaitGroup) {
    defer wg.Done()
    total += n
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 100; i++ {
        wg.Add(1)
        go addToTotal(i, &wg)
    }
    wg.Wait()
    fmt.Println(total) // should be 5050
}
```

<details>
<summary>💡 Solution</summary>

```go
package main

import (
    "fmt"
    "sync"
)

var (
    total int
    mu    sync.Mutex
)

func addToTotal(n int, wg *sync.WaitGroup) {
    defer wg.Done()
    mu.Lock()
    defer mu.Unlock()
    total += n
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 100; i++ {
        wg.Add(1)
        go addToTotal(i, &wg)
    }
    wg.Wait()
    fmt.Println(total) // Always 5050 ✅
}
```
</details>

---

## 🔑 Key Takeaways

| Concept | Remember |
|---------|----------|
| Race condition | Multiple goroutines write shared data simultaneously |
| `sync.Mutex` | Only 1 goroutine in critical section at a time |
| `defer mu.Unlock()` | **Always** defer unlock — prevents deadlocks |
| `sync.RWMutex` | Multiple readers OR one writer — not both |
| Embed mutex in struct | Groups data with its own lock |
| Never copy a Mutex | Always use pointer receivers |
| `sync.Once` | For one-time initialization |
| `-race` flag | `go run -race main.go` detects races |

---

## 📊 sync Package Summary

```
sync package
├── sync.Mutex       → exclusive lock
├── sync.RWMutex     → read-write lock
├── sync.WaitGroup   → wait for goroutines
├── sync.Once        → run code exactly once
├── sync.Map         → concurrent-safe map (no manual locking)
├── sync.Pool        → reuse objects to reduce GC pressure
└── sync.Cond        → condition variables (advanced)
```

### `sync.Map` — Built-in Concurrent Map

```go
var m sync.Map

m.Store("key", "value")               // Set
val, ok := m.Load("key")              // Get
m.Delete("key")                       // Delete
m.Range(func(k, v any) bool {         // Iterate
    fmt.Println(k, v)
    return true // return false to stop
})
```

---

*You've reached the end! 🎉 Review the checklist in [README.md](./README.md) to test your knowledge.*
