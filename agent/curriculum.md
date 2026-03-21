# CS-Daily Curriculum

**Philosophy:** Each challenge builds on the previous. No random jumps.
Progress is shaped by real feedback — if something is shaky, we stay until it's solid.

---

## Revised Structure (2026-03-21)

### Weekly Rhythm
- **Mon morning:** 2-3 articles + why this matters
- **Tue-Wed:** Read, experiment, ask questions (commute-friendly)
- **Thu morning:** Small coding exercise — from scratch, no scaffold
- **Fri morning:** Warmup / code reading / fix a bug
- **Weekend:** Longer project — something real and demo-able

### Rules
- No new concept until the previous one is solid
- Every exercise: write from scratch, no big scaffolds
- Weekend projects must be demo-able: "look what I built"
- Code reading challenges: once every ~2 weeks (find bug / explain / refactor)
- Before each challenge: state what you need to know, what to write from scratch, what's OK to google

---

## Track A: Go Language Mastery

### Phase 1 — Foundations
- [x] A-00: Assessment — types, slices, maps, structs, pointer receivers, error handling
- [x] A-01: Maps & Slices — TwoSum, WordFreq, Chunk
- [x] W-01: Rate Limiter — maps, mutexes, HTTP middleware (early detour — exposed design gap)
- [ ] A-02: **Struct design** — how to decompose a problem into types, design from scratch
- [ ] A-03: Error handling — custom errors, %w, errors.Is/As, when to use what (syntax known, judgment needed)
- [ ] A-04: Goroutines properly — scheduler, leaks, WaitGroup (fill conceptual gap under W-01)
- [ ] A-05: Mutex + channels — natural extension, W-01 revisited with real understanding
- [ ] A-06: HTTP — net/http internals, routing, middleware (now with goroutine foundation)
- [ ] A-07: Context — deadlines, cancellation
- [ ] A-08: Testing — mocking with interfaces, benchmarks
- [ ] A-09: Strings & Bytes
- [ ] A-10: Generics

### Phase 2 — Idiomatic Go
- [ ] A-11: Functional patterns — closures, higher-order functions
- [ ] A-12: Concurrency patterns — worker pool, fan-out/fan-in, pipeline
- [ ] A-13: Design patterns — options pattern, repository
- [ ] A-14: Performance — profiling, escape analysis

---

## Track B: CS Fundamentals (in Go)
Starts when A-track Phase 1 is solid (~A-06).

- [ ] B-01: Arrays vs slices internals
- [ ] B-02: Linked lists
- [ ] B-03: Stacks & queues
- [ ] B-04: Hash maps — implement a basic one
- [ ] B-05: Binary trees
- [ ] B-06: Heaps / priority queue
- [ ] B-07: Graphs — BFS, DFS
- [ ] B-08: Sorting algorithms
- [ ] B-09: Binary search
- [ ] B-10: Two pointers, sliding window
- [ ] B-11: Recursion & backtracking
- [ ] B-12: Dynamic programming

---

## Track C: Distributed Systems
Starts after A-06 (HTTP). Weekend projects bridge A and C tracks.

- [ ] C-01: HTTP server — net/http, routing, middleware, JSON API
- [ ] C-02: PostgreSQL — database/sql, connection pooling, migrations
- [ ] C-03: Testing HTTP APIs — integration tests, test containers
- [ ] C-04: Redis — caching, TTL, pub/sub
- [ ] C-05: Redis — distributed locks, rate limiting
- [ ] C-06: Kafka — producers, consumers
- [ ] C-07: Kafka — exactly-once, offset management
- [ ] C-08: gRPC — proto definitions, server/client, streaming (backlogged per Michail interest)
- [ ] C-09: Docker — containerizing a Go app
- [ ] C-10: Full system — API + Postgres + Redis + Kafka

---

## Weekend Projects (planned)
| Week | Concept | Project idea |
|------|---------|-------------|
| W1 (2026-03-22) | Struct design | CLI tool: something small and demo-able |
| W2 | Error handling | CLI: real errors, graceful messages |
| W3 | Goroutines | Concurrent word counter or downloader |
| W4 | Mutex + channels | Something with real concurrency |

---

## Code Reading Challenges (every ~2 weeks)
Real stdlib or open-source Go code. Task: understand it, explain it, find the bug, or refactor it.
Builds the "read someone else's code" muscle alongside writing.

---

## Key Observations (from code review history)
**Actual gap:** Design thinking — can implement when given structure; struggles deciding what structs/fields/decomposition to use from scratch.
**Interfaces:** Already solid — implicit satisfaction, static compliance checks, polymorphism. Don't repeat this.
**Error handling:** Syntax known (%w, Unwrap, custom errors). Needs judgment — when to wrap, when to sentinel, when to return nil.
**Aliasing:** Fixed and absorbed. copy + make pattern internalized.
**Control flow:** Early returns mostly there. Watch for fall-through patterns.

---

## Resources
- https://go.dev/tour/
- https://gobyexample.com/
- https://go.dev/doc/effective_go
- https://100go.co/
- https://gophercises.com/
