# 🐹 Go (Golang) — Zero to Master Complete Guide

> A complete, hands-on reference for mastering Go from absolute zero to production-ready.
> Every concept is tied to actual code in this repository with **solved examples** and **practice problems**.

---

## 🗺️ Study Roadmap

Read in order. Each file builds on the previous one.

| # | File | Topics | Level |
|---|------|--------|-------|
| 1 | [📘 README_01_FOUNDATIONS.md](./README_01_FOUNDATIONS.md) | Hello World, Variables, Constants, Enums, Data Types, Conditionals, Loops, Arrays, Slices | 🟢 Beginner |
| 2 | [📗 README_02_CORE_TYPES.md](./README_02_CORE_TYPES.md) | Functions, Pointers, Maps, Range, Closures, Variadic, Generics, Error Handling | 🟡 Intermediate |
| 3 | [📙 README_03_OOP_AND_ADVANCED.md](./README_03_OOP_AND_ADVANCED.md) | Structs, Methods, Interfaces, Embedding, Polymorphism, Dependency Injection | 🟠 Advanced |
| 4 | [📕 README_04_CONCURRENCY_AND_MASTERY.md](./README_04_CONCURRENCY_AND_MASTERY.md) | Goroutines, Channels, WaitGroup, Mutex, Select, Worker Pools, Pipelines | 🔴 Expert |
| 5 | [📦 README_05_PACKAGES.md](./README_05_PACKAGES.md) | Go Modules, Custom Packages, Exported Names, go.mod, go get | 🟡 Intermediate |
| 6 | [📁 README_06_FILES.md](./README_06_FILES.md) | os package, Read/Write/Create/Delete files, bufio, File Stats, Directory ops | 🟠 Advanced |
| 7 | [🔒 README_07_MUTEX.md](./README_07_MUTEX.md) | sync.Mutex, RWMutex, Race Conditions, Lock/Unlock, Defer patterns | 🔴 Expert |

---

## 📁 Repository Structure

```
d:\Golang\
│
├── 📘 README_01_FOUNDATIONS.md
├── 📗 README_02_CORE_TYPES.md
├── 📙 README_03_OOP_AND_ADVANCED.md
├── 📕 README_04_CONCURRENCY_AND_MASTERY.md
├── 📦 README_05_PACKAGES.md       ← NEW
├── 📁 README_06_FILES.md          ← NEW
├── 🔒 README_07_MUTEX.md          ← NEW
│
├── hello_world/main.go          → package main, fmt.Println, go run/build
├── variables/main.go            → 5 variable styles, zero values, type inference
├── constants/constant.go        → const, blocks, iota
├── Enums/enums.go               → iota, typed enums
├── simple_values/main.go        → int, float, string, bool
├── conditional/main.go          → if/else, switch
├── loops/loops.go               → for, while-style, infinite, break, continue
├── range/main.go                → range over slice, map, string, channel
├── Arrays/arrays.go             → fixed arrays, 2D arrays
├── Slices/slices.go             → make, append, copy, slicing
├── functions/func.go            → basic, multi-return, defer
├── pointers/pointers.go         → &, *, pass-by-value vs pointer
├── maps/maps.go                 → make, set/get, delete, two-value lookup
├── closures/closure.go          → closure factory, captured variables
├── vfunc/vfunc.go               → variadic functions
├── Generics/gene.go             → [T any], [T comparable]
├── structs/structs.go           → struct, embedding, constructor, methods
├── interafaces/inter.go         → interface, polymorphism, DI
├── Goroutines/gor.go            → go keyword, WaitGroup
├── Channels/chan.go             → make(chan), send/receive, select
├── Mutex/mu.go                  → sync.Mutex, WaitGroup + Mutex
├── files/file.go                → os, Read/Write/Create/Delete files
├── packages/pac.go              → go mod, custom packages
│   ├── auth/credentials.go      → exported functions
│   └── user/user.go             → exported structs
└── practice/prac.go             → practice exercises
```

---

## ⚡ Quick Concept Lookup

| I want to learn... | Go to... |
|--------------------|----------|
| Declare variables | [README_01 §3](./README_01_FOUNDATIONS.md) |
| Loops | [README_01 §7](./README_01_FOUNDATIONS.md) |
| Slices | [README_01 §9](./README_01_FOUNDATIONS.md) |
| Functions | [README_02 §1](./README_02_CORE_TYPES.md) |
| Pointers | [README_02 §2](./README_02_CORE_TYPES.md) |
| Maps | [README_02 §3](./README_02_CORE_TYPES.md) |
| Closures | [README_02 §5](./README_02_CORE_TYPES.md) |
| Generics | [README_02 §7](./README_02_CORE_TYPES.md) |
| Error Handling | [README_02 §8](./README_02_CORE_TYPES.md) |
| Structs | [README_03 §1](./README_03_OOP_AND_ADVANCED.md) |
| Interfaces | [README_03 §3](./README_03_OOP_AND_ADVANCED.md) |
| Goroutines | [README_04 §1](./README_04_CONCURRENCY_AND_MASTERY.md) |
| Channels | [README_04 §3](./README_04_CONCURRENCY_AND_MASTERY.md) |
| Mutex | [README_07](./README_07_MUTEX.md) |
| Custom Packages | [README_05](./README_05_PACKAGES.md) |
| File I/O | [README_06](./README_06_FILES.md) |

---

## 🎯 The "Learn by Doing" Checklist

```
1. READ the concept in the README
2. LOOK at the actual .go file in this repo
3. RUN it:  go run <path>/main.go
4. MODIFY it: change values, break things, fix it
5. WRITE your own version from scratch
```

### ✅ Foundations Checklist
- [ ] Run `hello_world/main.go` and add your own Printf examples
- [ ] Try all 5 variable declaration styles
- [ ] Write constants with iota for weekdays (Mon=1 through Sun=7)
- [ ] Modify a loop: what does `for i := 10; i > 0; i -= 3` print?
- [ ] Add 5 items to a slice, print only the last 3 using slicing

### ✅ Core Types Checklist
- [ ] Write a function that returns `(result int, err error)`
- [ ] Swap two numbers using pointers
- [ ] Create a map from country names to capitals, check if "India" exists
- [ ] Write a closure that keeps a running sum
- [ ] Write a generic `Min[T]` function

### ✅ OOP Checklist
- [ ] Create a `Shape` interface with `Area()` and implement for `Circle` and `Rectangle`
- [ ] Embed `Animal` in `Dog` and override a method
- [ ] Use interface as function parameter for dependency injection

### ✅ Concurrency Checklist
- [ ] Launch 10 goroutines with WaitGroup and verify all finish
- [ ] Send 5 numbers through a channel, square them, print results
- [ ] Use `select` to receive from 2 channels simultaneously
- [ ] Fix a race condition using Mutex

### ✅ Packages, Files & Mutex Checklist
- [ ] Create your own package with an exported function
- [ ] Read a file line by line using `bufio.Scanner`
- [ ] Write text to a new file and then delete it
- [ ] Use `sync.Mutex` to safely increment a counter from 50 goroutines
- [ ] Use `sync.RWMutex` for a concurrent-safe cache

---

## 🧠 Go Mental Models (Quick Reference)

```
┌─ Variables ──────────────────────────────────────────────────────┐
│  var x int = 5        → explicit type                            │
│  x := 5               → short declaration (inside func only)     │
│  var x int            → zero value (0 for int)                   │
└──────────────────────────────────────────────────────────────────┘

┌─ Functions ──────────────────────────────────────────────────────┐
│  func name(a int, b int) (int, error) { }                        │
│  → multiple returns | named returns | defer | first-class        │
└──────────────────────────────────────────────────────────────────┘

┌─ Structs + Methods ──────────────────────────────────────────────┐
│  type Person struct { Name string; Age int }                     │
│  func (p *Person) Greet() string { return "Hi, " + p.Name }     │
└──────────────────────────────────────────────────────────────────┘

┌─ Interfaces (implicit) ──────────────────────────────────────────┐
│  type Stringer interface { String() string }                      │
│  Any type with String() automatically satisfies Stringer         │
└──────────────────────────────────────────────────────────────────┘

┌─ Goroutines + Channels ──────────────────────────────────────────┐
│  go func() { ch <- result }()    // send concurrently            │
│  val := <-ch                      // blocking receive             │
└──────────────────────────────────────────────────────────────────┘

┌─ Mutex (Safe Shared State) ──────────────────────────────────────┐
│  var mu sync.Mutex                                               │
│  mu.Lock(); defer mu.Unlock()   // always defer Unlock           │
└──────────────────────────────────────────────────────────────────┘

┌─ Packages ───────────────────────────────────────────────────────┐
│  go mod init github.com/yourname/project                         │
│  Exported: Uppercase  │  Unexported: lowercase                   │
└──────────────────────────────────────────────────────────────────┘

┌─ File I/O ───────────────────────────────────────────────────────┐
│  os.ReadFile("f.txt")               → read entire file           │
│  os.WriteFile("f.txt", data, 0644) → write entire file          │
│  os.Create / os.Open / os.OpenFile  → fine-grained control       │
└──────────────────────────────────────────────────────────────────┘
```

---

*Happy Gophering! 🚀 — From Zero to Master, one concept at a time.*
