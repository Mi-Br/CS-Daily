# W12 Thu Exercise — Struct Design from Scratch

**Assigned:** 2026-03-25 (Thursday)  
**Track:** W12 — Struct Design  
**Estimated time:** ~45 min

---

## The Problem

You're building a small library system. Given the requirements below, design the structs and implement the logic — **start from a blank file**.

### Requirements

- A **library** holds a collection of books
- Each **book** has: a title, an author, a year published, and whether it's currently checked out
- You can **add** a book to the library
- You can **check out** a book by title — fails if already checked out
- You can **return** a book by title — fails if not checked out
- You can **list** all available books (not checked out)

### What to implement

```go
// You decide what the structs look like.
// Implement these four functions:

func (l *Library) Add(...)
func (l *Library) Checkout(title string) error
func (l *Library) Return(title string) error
func (l *Library) Available() []Book
```

### Done when

- [ ] All 4 methods work correctly
- [ ] `Checkout` and `Return` return meaningful errors (not just `nil`/`errors.New("error")`)
- [ ] `go vet` passes clean
- [ ] At least 3 tests written by you

---

## What to Figure Out Yourself

- What fields does `Book` need?
- What does `Library` hold internally — slice or map? Why?
- How do you look up a book by title efficiently?
- What's the right error message when checkout fails because it's already out?

## OK to Google

- `errors.New` / `fmt.Errorf`
- Slice vs map tradeoffs

## NOT OK to Google

- "Go library struct example" — design it yourself first
