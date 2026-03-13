# Challenge: Maps & Slices — The Fundamentals
**Track:** A + B  **ID:** A-01  **Date:** 2026-03-13

---

## Why this challenge

Looking at your A-00 code, three patterns kept surfacing:
- **Slice mutation** — you modified input slices when you shouldn't (Ex2, Ex3 flagged in review)
- **Control flow** — missing early returns, fall-through logic (just fixed in Ex5 today)
- **Map usage** — you haven't written much map-heavy code yet

A-01 fixes all three. These are the most common sources of subtle bugs in Go.

---

## 📖 Read First (pick one, ~15 min)

- [Go Slices: usage and internals](https://go.dev/blog/slices-intro) — **read this one**. Explains how slices share underlying arrays and why mutation is dangerous. This is the root cause of your A-00 habit.
- [Go Maps in Action](https://go.dev/blog/maps) — bonus if you have time

---

## 🎯 The Challenges

### Exercise 1: Two Sum

```go
// TwoSum returns the indices of the two numbers in nums that add up to target.
// Assume exactly one solution exists. Each element may only be used once.
// 
// Example: TwoSum([]int{2, 7, 11, 15}, 9) → (0, 1)  (because nums[0] + nums[1] = 9)
// Example: TwoSum([]int{3, 2, 4}, 6) → (1, 2)
func TwoSum(nums []int, target int) (int, int)
```

**Rules:**
- First solve it in O(n²) — two nested loops, get it working
- Then solve it in O(n) using a map
- Both versions should be in the file, the O(n) one is what gets tested
- Do NOT modify the input slice

**Think about:** what does the map store — keys? values? both?

---

### Exercise 2: Word Frequency

```go
// WordFreq returns a map of each word to how many times it appears in the slice.
// Case-sensitive: "Go" and "go" are different words.
//
// Example: WordFreq([]string{"go", "is", "go", "fast"}) → map[go:2 is:1 fast:1]
func WordFreq(words []string) map[string]int
```

**Rules:**
- Do NOT modify the input slice
- Return an empty map (not nil) for empty input
- Think about the zero value of `int` in a map — it makes this elegant

---

### Exercise 3: Chunk

```go
// Chunk splits a slice into sub-slices of the given size.
// The last chunk may be smaller if len(s) is not divisible by n.
//
// Example: Chunk([]int{1,2,3,4,5}, 2) → [[1 2] [3 4] [5]]
// Example: Chunk([]int{1,2,3}, 3)     → [[1 2 3]]
// Example: Chunk([]int{}, 2)          → []
func Chunk(s []int, n int) [][]int
```

**Rules:**
- Each chunk must be an independent copy — modifying a chunk must NOT affect the original slice
- Handle edge cases: empty slice, n larger than len(s), n = 1
- No aliasing. This is the point of this exercise.

**Why this matters:** If you return slices that share memory with the input, the caller can corrupt your data (or you can corrupt theirs). This is one of the most common Go bugs.

---

## ✅ Done When

- [ ] All three functions implemented in `solution.go`
- [ ] Tests written in `solution_test.go` covering at least:
  - TwoSum: basic case, negative numbers, zero as element
  - WordFreq: empty input, repeated words, single word
  - Chunk: even split, uneven split, empty input, n=1, n > len(s)
- [ ] `go test ./...` passes
- [ ] `go vet ./...` no warnings
- [ ] No input slice mutation anywhere

---

## 💡 Hints (read only if stuck)

<details>
<summary>Hint — Two Sum map approach</summary>

You want to answer: "have I seen the complement of this number before?"

For each number `x`, the complement is `target - x`.
Store `num → index` in the map as you go. Before storing, check if the complement already exists.

```go
seen := make(map[int]int)  // value → index
for i, x := range nums {
    complement := target - x
    if j, ok := seen[complement]; ok {
        return j, i
    }
    seen[x] = i
}
```
</details>

<details>
<summary>Hint — Chunk, no aliasing</summary>

Slices in Go share the underlying array. If you do:
```go
chunk := s[start:end]  // ← this ALIASES s
```
...modifying `chunk[0]` will also change `s[start]`.

To make an independent copy:
```go
chunk := make([]int, end-start)
copy(chunk, s[start:end])  // ← independent copy
```
</details>

---

## 📁 Where to put your code

```
challenges/2026-03-13_A01-maps-and-slices/
  solution.go
  solution_test.go
```

---

## 🏖️ Weekend Mini-Project (optional, ~2-3h)

You already built an HTTP API server in A-00. Extend it with a **rate limiter**.

**Goal:** No endpoint should allow more than N requests per second per client (use IP as client ID).

**Spec:**
- Implement a `RateLimiter` struct with `Allow(clientID string) bool`
- Use a fixed window: track request count per client per second
- If `Allow()` returns false → respond with HTTP 429 Too Many Requests
- Wire it as middleware on your existing server

**Why:** This touches goroutines + maps + mutexes (concurrency-safe shared state) — exactly what you just learned. It's also something every real backend has.

**Stretch:** Make the limit configurable per route, not just global.

---

*Reviewed by Julia — your senior staff engineer on call.*
