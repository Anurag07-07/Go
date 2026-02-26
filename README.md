# 🐹 Go (Golang) — Complete Mastery Notes

> A complete, hands-on reference for mastering Go from zero to production-ready.
> Every concept is tied to the actual code in this repository.

---

## 🗺️ How to Use This Guide

Read in order. Each README builds on the previous one.

| # | File | Topics | Level |
|---|------|--------|-------|
| 1 | [📘 README_01_FOUNDATIONS.md](./README_01_FOUNDATIONS.md) | Hello World, Variables, Constants, Enums, Data Types, Conditionals, Loops, Arrays, Slices | 🟢 Beginner |
| 2 | [📗 README_02_CORE_TYPES.md](./README_02_CORE_TYPES.md) | Functions, Pointers, Maps, Range, Closures, Variadic, Generics, Error Handling | 🟡 Intermediate |
| 3 | [📙 README_03_OOP_AND_ADVANCED.md](./README_03_OOP_AND_ADVANCED.md) | Structs, Methods, Interfaces, Embedding, Polymorphism, Dependency Injection | 🟠 Advanced |
| 4 | [📕 README_04_CONCURRENCY_AND_MASTERY.md](./README_04_CONCURRENCY_AND_MASTERY.md) | Goroutines, Channels, WaitGroup, Mutex, Select, Worker Pools, Pipelines | 🔴 Expert |

---

## 📁 Repository Structure

```
d:\Golang\
│
├── 📘 README_01_FOUNDATIONS.md        ← Start here!
├── 📗 README_02_CORE_TYPES.md
├── 📙 README_03_OOP_AND_ADVANCED.md
├── 📕 README_04_CONCURRENCY_AND_MASTERY.md
│
├── hello_world/
│   └── main.go           → package main, fmt.Println, go run/build
│
├── variables/
│   └── main.go           → 5 variable styles, zero values, type inference
│
├── constants/
│   └── main.go           → const, blocks, package-level constants
│
├── Enums/
│   └── enums.go          → iota, typed enums, string enums
│
├── simple_values/
│   └── main.go           → primitive types: int, float, string, bool
│
├── conditional/
│   └── main.go           → if/else, if with init, switch
│
├── loops/
│   └── main.go           → for, while-style, infinite, break, continue
│
├── range/
│   └── main.go           → range over slice, map, string, channel
│
├── Arrays/
│   └── arrays.go         → fixed-size arrays, 2D arrays, zero values
│
├── Slices/
│   └── slices.go         → make, append, copy, slicing [:], 2D slices
│
├── functions/
│   └── func.go           → basic, multi-return, higher-order, defer
│
├── pointers/
│   └── pointers.go       → &, *, pass-by-value vs pointer, new()
│
├── maps/
│   └── maps.go           → make, set/get, delete, clear, two-value lookup
│
├── closures/
│   └── closure.go        → closure factory, captured variables, independent state
│
├── vfunc/
│   └── vfunc.go          → variadic functions, spread operator
│
├── Generics/
│   └── gene.go           → [T any], [T comparable], generic structs
│
├── structs/
│   └── structs.go        → struct definition, embedding, constructor, methods
│
├── interafaces/
│   └── inter.go          → interface, implicit implementation, polymorphism, DI
│
├── Goroutines/
│   └── gor.go            → go keyword, sync.WaitGroup, anonymous goroutines
│
└── Channels/
    └── chan.go            → make(chan), send/receive, buffered, range, select
```

---

## ⚡ Quick Concept Lookup

| I want to learn... | Go to... |
|--------------------|----------|
| How to declare variables | [README_01 §3](./README_01_FOUNDATIONS.md#3-variables--all-5-styles) |
| How loops work | [README_01 §7](./README_01_FOUNDATIONS.md#7-loops--one-ring-to-rule-them-all) |
| What slices are | [README_01 §9](./README_01_FOUNDATIONS.md#9-slices--dynamic-power-arrays) |
| How functions work | [README_02 §1](./README_02_CORE_TYPES.md#1-functions--complete-guide) |
| What pointers are | [README_02 §2](./README_02_CORE_TYPES.md#2-pointers--pass-by-reference) |
| How maps work | [README_02 §3](./README_02_CORE_TYPES.md#3-maps--key-value-powerhouse) |
| What closures are | [README_02 §5](./README_02_CORE_TYPES.md#5-closures--functions-that-remember) |
| How generics work | [README_02 §7](./README_02_CORE_TYPES.md#7-generics-go-118) |
| How to handle errors | [README_02 §8](./README_02_CORE_TYPES.md#8-error-handling) |
| How structs work | [README_03 §1](./README_03_OOP_AND_ADVANCED.md#1-structs--gos-custom-type-system) |
| What interfaces are | [README_03 §3](./README_03_OOP_AND_ADVANCED.md#3-interfaces--the-contract-system) |
| How embedding works | [README_03 §4](./README_03_OOP_AND_ADVANCED.md#4-struct-embedding--gos-composition-model) |
| What goroutines are | [README_04 §1](./README_04_CONCURRENCY_AND_MASTERY.md#1-goroutines--lightweight-concurrency) |
| How channels work | [README_04 §3](./README_04_CONCURRENCY_AND_MASTERY.md#3-channels--goroutine-communication) |
| How to sync goroutines | [README_04 §2](./README_04_CONCURRENCY_AND_MASTERY.md#2-syncwaitgroup--wait-for-goroutines) |
| How to prevent race conditions | [README_04 §7](./README_04_CONCURRENCY_AND_MASTERY.md#7-syncmutex--prevent-race-conditions) |

---

## 🎯 The "Learn by Doing" Checklist

For each topic, follow this cycle:

```
1. READ the concept in the README
2. LOOK at the actual .go file in this repo
3. RUN it: `go run <path>/main.go`
4. MODIFY it: change values, break things, fix it
5. WRITE your own version from scratch
```

### Checklist:
- [ ] Run `hello_world/main.go` and add your own Printf examples
- [ ] Try all 5 variable declaration styles in `variables/main.go`
- [ ] Modify the loop in your head: what does `for i := 10; i > 0; i -= 3` print?
- [ ] Add 5 items to a slice and print only the last 3 using slicing
- [ ] Write a function that returns (result int, err error)
- [ ] Swap two numbers using pointers
- [ ] Create a map from country names to capitals, check if "India" exists
- [ ] Write a closure that keeps a running sum
- [ ] Create a Shape interface and implement it for Triangle and Circle
- [ ] Launch 10 goroutines with WaitGroup and verify all finish
- [ ] Send 5 numbers through a channel, square them in a goroutine, print results

---

## 🧠 Go Mental Models

```
┌─ Variables ──────────────────────────────────────────────────────────┐
│  var x int = 5  → explicit                                           │
│  x := 5         → inferred (only inside functions)                   │
└──────────────────────────────────────────────────────────────────────┘

┌─ Functions ──────────────────────────────────────────────────────────┐
│  func name(a int, b int) (result int, err error) { }                 │
│  → multiple returns | named returns | defer | first-class values      │
└──────────────────────────────────────────────────────────────────────┘

┌─ Structs (like classes, but explicit) ───────────────────────────────┐
│  type Person struct { Name string; Age int }                         │
│  func (p *Person) Greet() string { return "Hi, " + p.Name }         │
└──────────────────────────────────────────────────────────────────────┘

┌─ Interfaces (implicit, not explicit) ────────────────────────────────┐
│  type Stringer interface { String() string }                          │
│  Any type with String() automatically satisfies Stringer             │
└──────────────────────────────────────────────────────────────────────┘

┌─ Goroutines + Channels ──────────────────────────────────────────────┐
│  go func() { ch <- result }()    // send concurrently                │
│  val := <-ch                      // blocking receive                 │
│  → Goroutines are cheap, channels are safe communication             │
└──────────────────────────────────────────────────────────────────────┘
```

---

*Happy Gophering! 🚀*
