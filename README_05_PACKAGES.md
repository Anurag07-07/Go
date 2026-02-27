# 📦 README_05 — Packages & Go Modules

> **File Reference:** `packages/pac.go`, `packages/auth/credentials.go`, `packages/user/user.go`

---

## 1. What is a Package?

A **package** is a directory of `.go` files that share the same `package` declaration.
Every Go file must belong to a package. Go programs are built from packages.

```
package main   ← executable entry point (special)
package auth   ← reusable library package
package user   ← reusable library package
```

### Two Types of Packages

| Type | `package` name | Purpose |
|------|---------------|---------|
| **main** | `package main` | Produces an executable binary |
| **library** | anything else | Reusable code imported by other packages |

---

## 2. Go Modules (`go.mod`)

A **module** is a collection of packages. It is the unit of versioning and distribution.

### Initialize a Module
```bash
go mod init github.com/yourname/projectname
```

This creates `go.mod`:
```
module github.com/golang

go 1.21
```

### Module Commands Cheatsheet

| Command | What it does |
|---------|-------------|
| `go mod init <name>` | Create a new module |
| `go mod tidy` | Add missing / remove unused dependencies |
| `go get <pkg>` | Download and add a dependency |
| `go mod download` | Download all dependencies |
| `go list -m all` | List all module dependencies |

---

## 3. Exported vs Unexported Names

Go uses **capitalization** to control visibility — there are no `public`/`private` keywords.

```go
// In package auth:
func LoginWithCredentials() { }  // ✅ Exported — usable outside this package
func validateToken() { }         // ❌ Unexported — only usable inside package auth

type User struct {
    Name  string  // ✅ Exported field
    token string  // ❌ Unexported field
}
```

> **Rule:** If the name starts with an **Uppercase** letter → exported (public).
> If it starts with a **lowercase** letter → unexported (package-private).

---

## 4. Your Code Explained — `packages/`

### Project Structure
```
packages/
├── go.mod                  ← module: github.com/golang
├── pac.go                  ← package main (entry point)
├── auth/
│   └── credentials.go      ← package auth
└── user/
    └── user.go             ← package user
```

### `auth/credentials.go`
```go
package auth

import "fmt"

// LoginWithCredentials is Exported (Uppercase L)
// It can be called from any package that imports "github.com/golang/auth"
func LoginWithCredentials(username string, password string) {
    fmt.Println("login user using", username, password)
}
```

### `user/user.go`
```go
package user

// User struct is Exported — accessible from other packages
type User struct {
    Email string   // Exported field
    Name  string   // Exported field
}
```

### `pac.go` — Main entry point
```go
package main

import (
    "github.com/golang/auth"   // import the auth package
    "github.com/golang/user"   // import the user package
)

func main() {
    // Call exported function from auth package
    auth.LoginWithCredentials("Anurag", "772002")

    // Use exported struct from user package
    u := user.User{
        Email: "Anurag@gmail.com",
        Name:  "Anurag",
    }
    println(u.Email)
}
```

**Output:**
```
login user using Anurag 772002
Anurag@gmail.com
```

---

## 5. How Import Paths Work

```go
import "github.com/golang/auth"
//      └── module name ──┘└─ subdirectory ─┘
```

- The **module name** is set in `go.mod`
- The **subdirectory** is the folder relative to the module root
- You access functions via `packagename.FunctionName`

### Aliasing Imports
```go
import (
    auth "github.com/golang/auth"   // alias: same name, explicit
    u    "github.com/golang/user"   // alias: shorter name
    _    "github.com/some/pkg"      // blank import: run init() only
    .    "github.com/some/pkg"      // dot import: no prefix needed (avoid!)
)
```

---

## 6. The `init()` Function

Every package can have an `init()` function. It runs **automatically** before `main()`.

```go
package auth

import "fmt"

func init() {
    fmt.Println("auth package initialized")
}

func LoginWithCredentials(username, password string) {
    fmt.Println("login user using", username, password)
}
```

**Execution order:**
1. Package-level variables initialized
2. `init()` runs (in import order)
3. `main()` runs

---

## 7. Installing Third-Party Packages

```bash
# Install a package
go get github.com/gin-gonic/gin

# Install a specific version
go get github.com/gin-gonic/gin@v1.9.1

# Remove unused packages
go mod tidy
```

After `go get`, the package is added to `go.mod` and `go.sum`.

---

## ✅ Solved Examples

### Example 1 — Create a `math` utility package

**File: `mathutil/math.go`**
```go
package mathutil

// Add returns the sum of two integers
func Add(a, b int) int {
    return a + b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
    return a * b
}

// max is unexported — only usable inside mathutil
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

**File: `main.go`**
```go
package main

import (
    "fmt"
    "github.com/yourname/project/mathutil"
)

func main() {
    sum := mathutil.Add(3, 7)
    product := mathutil.Multiply(4, 5)
    fmt.Println("Sum:", sum)         // Sum: 10
    fmt.Println("Product:", product) // Product: 20
    // mathutil.max(1, 2) ← ERROR: unexported
}
```

---

### Example 2 — Package with a struct and methods

**File: `greet/greet.go`**
```go
package greet

import "fmt"

type Greeter struct {
    Language string
}

func (g Greeter) Hello(name string) {
    switch g.Language {
    case "Hindi":
        fmt.Println("Namaste,", name)
    case "Spanish":
        fmt.Println("Hola,", name)
    default:
        fmt.Println("Hello,", name)
    }
}
```

**File: `main.go`**
```go
package main

import "github.com/yourname/project/greet"

func main() {
    eng := greet.Greeter{Language: "English"}
    eng.Hello("Anurag") // Hello, Anurag

    hin := greet.Greeter{Language: "Hindi"}
    hin.Hello("Anurag") // Namaste, Anurag
}
```

---

### Example 3 — Multiple packages working together (like your `packages/` folder)

```go
// config/config.go
package config

type AppConfig struct {
    DBHost string
    DBPort int
    AppEnv string
}

func Default() AppConfig {
    return AppConfig{
        DBHost: "localhost",
        DBPort: 5432,
        AppEnv: "development",
    }
}
```

```go
// logger/logger.go
package logger

import "fmt"

func Info(msg string) {
    fmt.Println("[INFO]", msg)
}

func Error(msg string) {
    fmt.Println("[ERROR]", msg)
}
```

```go
// main.go
package main

import (
    "github.com/yourname/app/config"
    "github.com/yourname/app/logger"
)

func main() {
    cfg := config.Default()
    logger.Info("Starting app on " + cfg.DBHost)
}
// Output: [INFO] Starting app on localhost
```

---

## 🏋️ Practice Problems

### Problem 1 — Calculator Package
Create a `calculator` package with these exported functions:
- `Add(a, b float64) float64`
- `Subtract(a, b float64) float64`
- `Divide(a, b float64) (float64, error)` — return error if b == 0

Import it in `main.go` and test all three functions.

<details>
<summary>💡 Solution</summary>

```go
// calculator/calc.go
package calculator

import "errors"

func Add(a, b float64) float64      { return a + b }
func Subtract(a, b float64) float64 { return a - b }

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

```go
// main.go
package main

import (
    "fmt"
    "github.com/you/proj/calculator"
)

func main() {
    fmt.Println(calculator.Add(10, 5))       // 15
    fmt.Println(calculator.Subtract(10, 5))  // 5
    result, err := calculator.Divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err) // Error: division by zero
    } else {
        fmt.Println(result)
    }
}
```
</details>

---

### Problem 2 — Student Package
Create a `student` package with:
- A `Student` struct with `Name string`, `Grade int`
- A `NewStudent(name string, grade int) Student` constructor
- A method `(s Student) IsPassing() bool` — passing if Grade >= 50
- Import and use it in `main.go`

<details>
<summary>💡 Solution</summary>

```go
// student/student.go
package student

type Student struct {
    Name  string
    Grade int
}

func NewStudent(name string, grade int) Student {
    return Student{Name: name, Grade: grade}
}

func (s Student) IsPassing() bool {
    return s.Grade >= 50
}
```

```go
// main.go
package main

import (
    "fmt"
    "github.com/you/proj/student"
)

func main() {
    s1 := student.NewStudent("Anurag", 85)
    s2 := student.NewStudent("Jay", 40)

    fmt.Println(s1.Name, "passing:", s1.IsPassing()) // Anurag passing: true
    fmt.Println(s2.Name, "passing:", s2.IsPassing()) // Jay passing: false
}
```
</details>

---

### Problem 3 — `init()` Ordering
Predict the output of this program:

```go
// package a
func init() { fmt.Println("init A") }

// package b (imports a)
func init() { fmt.Println("init B") }

// main (imports b)
func init() { fmt.Println("init Main") }
func main() { fmt.Println("main()") }
```

<details>
<summary>💡 Answer</summary>

```
init A
init B
init Main
main()
```
Go initializes packages in **dependency order** — deepest first.
</details>

---

## 🔑 Key Takeaways

| Concept | Remember |
|---------|----------|
| Package = folder | One `package` name per directory |
| Uppercase = exported | `LoginWithCredentials` ✅, `validateToken` ❌ |
| `go mod init` | Creates the module, run once per project |
| `go mod tidy` | Cleans up dependencies, run often |
| Import path | `module_name/subdirectory` |
| `init()` | Runs automatically before `main()`, in import order |

---

*Next → [📁 README_06_FILES.md](./README_06_FILES.md) — File I/O in Go*
