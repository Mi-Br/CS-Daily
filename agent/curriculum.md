# CS-Daily Curriculum

**Philosophy:** Each challenge builds on the previous. No random jumps.
Progress is shaped by real feedback — if something is shaky, we stay until it's solid.

---

## Track A: Go Language Mastery

### Phase 1 — Re-entry & Assessment (where we start)
- [ ] A-00: Assessment — 5 short exercises covering types, slices, maps, structs, pointer receivers, error handling
- [ ] A-01: Slices deep dive — internals, append, capacity, copy
- [ ] A-02: Maps — idiomatic use, nil maps, concurrent safety
- [ ] A-03: Structs & methods — value vs pointer receivers (why it matters)
- [ ] A-04: Interfaces — implicit implementation, small interfaces, polymorphism
- [ ] A-05: Error handling — custom errors, wrapping, `errors.Is` / `errors.As`
- [ ] A-06: Goroutines — what they are, how the scheduler works, leaks
- [ ] A-07: Channels — unbuffered vs buffered, direction, select
- [ ] A-08: sync package — Mutex, WaitGroup, Once
- [ ] A-09: Context — deadlines, cancellation, request-scoped values
- [ ] A-10: Testing — table-driven tests, mocking with interfaces, benchmarks
- [ ] A-11: Generics — when to use, constraints, real examples
- [ ] A-12: Modules & packages — project layout, internal packages

### Phase 2 — Idiomatic Go
- [ ] A-13: Functional patterns — closures, higher-order functions
- [ ] A-14: Concurrency patterns — worker pool, fan-out/fan-in, pipeline
- [ ] A-15: Design patterns — options pattern, functional options, repository
- [ ] A-16: Performance — profiling, escape analysis, avoiding allocations

---

## Track B: CS Fundamentals (in Go)

### Data Structures
- [ ] B-01: Arrays vs slices internals
- [ ] B-02: Linked lists — singly, doubly, circular
- [ ] B-03: Stacks & queues
- [ ] B-04: Hash maps — collision, load factor (implement a basic one)
- [ ] B-05: Binary trees — insert, search, traversal
- [ ] B-06: Heaps — min/max heap, priority queue
- [ ] B-07: Graphs — adjacency list/matrix, BFS, DFS

### Algorithms
- [ ] B-08: Sorting — bubble, merge, quick (implement + understand tradeoffs)
- [ ] B-09: Binary search & variants
- [ ] B-10: Two pointers, sliding window
- [ ] B-11: Recursion & backtracking
- [ ] B-12: Dynamic programming — memoization, tabulation

---

## Track C: Distributed Systems

- [ ] C-01: HTTP server — `net/http`, routing, middleware, JSON API
- [ ] C-02: PostgreSQL — `database/sql`, connection pooling, migrations
- [ ] C-03: Testing HTTP APIs — integration tests, test containers
- [ ] C-04: Redis — caching, TTL, pub/sub (first touch)
- [ ] C-05: Redis — distributed locks, rate limiting
- [ ] C-06: Kafka — producers, consumers, consumer groups (first touch)
- [ ] C-07: Kafka — exactly-once, offset management, error handling
- [ ] C-08: gRPC — proto definitions, server/client, streaming
- [ ] C-09: Docker — containerizing a Go app, multi-stage builds
- [ ] C-10: Putting it together — a small system: API + Postgres + Redis + Kafka

---

## Pacing

- 1 challenge per 1-2 days
- After each review: update `progress.json` with observations
- If a concept appears shaky in review → insert a reinforcement challenge before moving on
- Never skip a weak spot — it will surface later at 10x the cost

---

## Resources (reference)

- https://go.dev/tour/ — official tour
- https://gobyexample.com/ — quick reference
- https://go.dev/doc/effective_go — idiomatic Go bible
- https://100go.co/ — 100 common Go mistakes (great for code review context)
- https://gophercises.com/ — free real-world exercises
