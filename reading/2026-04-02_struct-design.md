# Struct Design in Go
**Date:** 2026-04-02 | **Concept:** Struct design, type decomposition, value vs pointer semantics | **Est. read time:** ~35 min

## Why this matters
You've been told to "design the structs yourself" — but without a mental model for how to decompose a problem into types, every blank file is intimidating. This reading gives you that model.

## Read these (in order)

1. **[Effective Go — Composite Literals](https://go.dev/doc/effective_go#composite_literals)** — why: shows how Go thinks about initializing structs; sets up the syntax and idioms you'll use
2. **[Design Philosophy on Data and Semantics — Bill Kennedy](https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html)** — why: value vs pointer semantics is the core Go design decision; this is the definitive guide
3. **[JSON in Go — Go by Example](https://gobyexample.com/json)** — why: you'll need JSON in the todo CLI; type it out, don't just read
4. **[Struct Embedding — Go by Example](https://gobyexample.com/struct-embedding)** — why: quick, practical, directly applicable to the todo CLI if you want to extend it

## Key concepts to come away with

**Value semantics vs pointer semantics**
- Value: `func (b Book) Title()` — copy; safe for small, read-only types
- Pointer: `func (l *Library) Add(b Book)` — mutates; use when state must change
- Rule of thumb: if it changes state → pointer receiver

**Struct as domain model**
- One struct per real-world concept (Book, Library, Task, TodoList)
- Fields = properties that define that concept
- Methods = behaviors that belong to that concept

**Slice vs map for lookup**
- Slice: ordered, iterate, simple — O(n) lookup by field
- Map: O(1) lookup by key — use when "find by X" is a core operation
- For a Library with `FindByTitle()`: map wins

## Questions to answer after reading

1. When should a method use a value receiver vs a pointer receiver?
2. Your Library has a `Checkout(title)` method — should it use a map or a slice? Why?
3. In the todo CLI, what fields does a Task need? What does the TodoList hold internally?
4. Why does `json.Marshal` need exported fields (capitalized)?
5. When would you use struct embedding vs a regular field?

## Your Notes
> (Michail adds notes here as he reads)
