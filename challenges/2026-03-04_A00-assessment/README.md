# Challenge A-00: The Assessment
**Track:** A — Go Language  
**ID:** A-00  
**Date:** 2026-03-04  
**Estimated time:** 2–3 hours across 1–2 days

---

## 📖 Read First

Before coding, read these. They're short and high-signal:

1. **[Go by Example](https://gobyexample.com/)** — Skim the sidebar. You'll recognize most of it. Notice what feels rusty.
2. **[Effective Go — Names & Declarations](https://go.dev/doc/effective_go#names)** — 10 min read. Focus on naming conventions and how Go prefers short names in narrow scopes.
3. **[The Go Blog: Errors are values](https://go.dev/blog/errors-are-values)** — Rob Pike on why Go's error handling is a design choice, not a flaw.

---

## 🎯 The Challenge

Five exercises. Each tests a different layer. Do them in order — don't skip.

---

### Exercise 1: Slice Wrangling
Write a function `TopN(nums []int, n int) []int` that returns the top N largest numbers from a slice, **without sorting the original slice**, in descending order.

```
TopN([]int{3, 1, 4, 1, 5, 9, 2, 6}, 3) → [9, 6, 5]
```

Rules:
- Don't modify the input slice
- Handle edge cases: n > len(nums), empty slice, n <= 0

---

### Exercise 2: Struct + Pointer Receivers
Model a simple **bank account**:

```go
type Account struct { ... }
```

Implement these methods:
- `Deposit(amount float64) error`
- `Withdraw(amount float64) error`
- `Balance() float64`
- `String() string` — human-readable summary

Rules:
- Negative deposits = error
- Withdrawing more than balance = error
- Use pointer receivers where it makes sense — and be intentional about it (you'll be asked why in the review)

---

### Exercise 3: Interface + Multiple Types
Define a `Shape` interface with `Area() float64` and `Perimeter() float64`.

Implement it for:
- `Circle`
- `Rectangle`
- `Triangle` (three sides, use Heron's formula for area)

Write a function `PrintShapeInfo(s Shape)` that prints both values.

Then write a function `LargestArea(shapes []Shape) Shape` that returns the shape with the largest area.

---

### Exercise 4: Error Handling
Write a function `ParseConfig(data string) (Config, error)` where `Config` is:

```go
type Config struct {
    Host    string
    Port    int
    Debug   bool
}
```

The input `data` is a simple format:
```
host=localhost
port=8080
debug=true
```

Rules:
- Return a **custom error type** `ConfigError` that includes the field name and the reason
- If `port` is not a valid integer → error
- If `port` < 1 or > 65535 → error
- Missing `host` → error
- Unknown keys → ignore (don't error)
- Use `fmt.Errorf` with `%w` for wrapping where appropriate

---

### Exercise 5: Goroutines + Channels
Write a function `FetchAll(urls []string) []Result` where:

```go
type Result struct {
    URL      string
    Status   int
    Err      error
}
```

Rules:
- Fetch all URLs **concurrently** (use goroutines)
- Use a real HTTP GET (doesn't matter if URLs are fake — handle the error gracefully)
- Return results **in the same order as the input** (this is the tricky part)
- Don't leak goroutines
---

## ✅ Done When

- [ ] All 5 exercises compile and run
- [ ] Each exercise has at least basic tests (`_test.go` file)
- [ ] `go vet ./...` passes cleanly
- [ ] Code is in `challenges/2026-03-04_A00-assessment/`
- [ ] Committed and pushed to `main`

---

## 💡 Hints (read only if stuck)

<details>
<summary>Exercise 1 — hint</summary>
Think about using a min-heap of size N, or just sort a copy. Both are valid — but they have different tradeoffs. Pick one and be ready to explain why.
</details>

<details>
<summary>Exercise 4 — hint</summary>
`strings.Split(data, "\n")` to get lines, then `strings.Cut(line, "=")` to split key/value. Your `ConfigError` should implement the `error` interface.
</details>

<details>
<summary>Exercise 5 — hint</summary>
Make a slice of Results pre-allocated to len(urls). Launch each goroutine with its index. Use sync.WaitGroup. That way results land in the right slot without needing to sort.
</details>

---

## 📁 Where to put your code

All files go inside `challenges/2026-03-04_A00-assessment/`:

```
challenges/2026-03-04_A00-assessment/
  main.go         (or split into multiple files — your call)
  main_test.go
  go.mod
```

Run `go mod init cs-daily/assessment` inside that folder.

---

*When you push, I'll review. Don't wait until it's perfect — commit what you have and I'll give feedback on it.*
