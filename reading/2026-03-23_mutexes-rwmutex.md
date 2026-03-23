# Mutexes & RWMutex
**Date:** 2026-03-23 | **Concept:** sync.Mutex, sync.RWMutex | **Est. read time:** ~30 min

---

## Why this matters

You already used `sync.Mutex` in W-01 without fully understanding it — it worked, but
you couldn't have written it from scratch. This fills that gap. Mutexes are everywhere
in concurrent Go code. Once this clicks, goroutines and channels will make much more sense too.

---

## Read these (in order)

1. **[Go by Example: Mutexes](https://gobyexample.com/mutexes)**  
   Why: Short, concrete. Shows the problem (data race) and the solution side by side. Start here.

2. **[Go by Example: RWMutex](https://gobyexample.com/rw-mutexes)**  
   Why: Extends the above. Explains why you'd want read-only locks and when to use each.

3. **[The Go Blog: Share Memory by Communicating](https://go.dev/blog/codelab-share)**  
   Why: Official Go perspective on concurrency. Short. Sets up the mental model before you hit channels later.

4. **[100 Go Mistakes #58–60: Mutexes](https://100go.co/#not-understanding-race-problems-58)**  
   Why: Real mistakes people make. Directly relevant to what tripped you up in W-01.

---

## Key concepts to come away with

**sync.Mutex — exclusive lock**
- Only one goroutine can hold it at a time
- Use when you're **reading AND writing** shared state
- `Lock()` → do work → `Unlock()` (always `defer mu.Unlock()` right after `Lock()`)

**sync.RWMutex — reader/writer lock**
- Multiple goroutines can **read simultaneously** — `RLock()` / `RUnlock()`
- Only one goroutine can **write** — `Lock()` / `Unlock()` (blocks all readers too)
- Use when reads are frequent and writes are rare

**The rule:**
```
Writing anything?  → Lock() / Unlock()
Read-only?         → RLock() / RUnlock()
```

**Why does it matter which you use?**
- Using `Lock()` for reads works, but it's wasteful — you're blocking other readers for no reason
- Using `RLock()` for writes is a data race — multiple goroutines can corrupt shared state simultaneously
- `go test -race` will catch the second mistake; the first just hurts performance

**defer is your friend:**
```go
func (r *RateLimiter) Allow(id string) bool {
    r.mu.Lock()
    defer r.mu.Unlock()  // guaranteed to run even if you return early
    // ...
}
```
Without `defer`, every early `return` needs a manual `Unlock()`. Miss one = deadlock.

---

## Questions to answer after reading

Before doing the warmup, make sure you can answer these without looking:

1. What's the difference between `Lock()` and `RLock()`?
2. Why does `Allow()` need `Lock()` but `CountRequests()` only needs `RLock()`?
3. What happens if you use `RLock()` inside a method that also writes to the map?
4. Why do we almost always use `defer mu.Unlock()` instead of calling `Unlock()` directly?

---

## Your Notes

> Add anything here as you read — things that clicked, things that confused you, questions.
> I'll read this before the warmup.

