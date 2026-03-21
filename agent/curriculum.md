# CS-Daily Curriculum

**Philosophy:** Each challenge builds on the previous. No random jumps.
Progress is shaped by real feedback — if something is shaky, we stay until it's solid.

---

## Revised Structure (2026-03-21)

### Weekly Rhythm
- **Mon morning:** 2-3 articles + why this matters
- **Tue-Wed:** Read, experiment, ask questions
- **Thu morning:** Small coding exercise — from scratch, no scaffold
- **Fri morning:** Warmup / code reading / fix a bug
- **Weekend:** Longer project — something real and completable

### Rules
- No new concept until the previous one is solid
- Every exercise: write from scratch, no big scaffolds
- Weekend projects must be demo-able: "look what I built"
- Code reading challenges: once every ~2 weeks (find bug / explain / refactor)
- Before each challenge: state clearly what you need to know, what to write from scratch, what's OK to google

---

## Track A: Go Language Mastery

### Phase 1 — Foundations (in progress)
- [x] A-00: Assessment — types, slices, maps, structs, pointer receivers, error handling
- [x] A-01: Maps & Slices — TwoSum, WordFreq, Chunk
- [ ] A-02: Interfaces — implicit satisfaction, small interfaces, why they enable testing
- [ ] A-03: Error handling — custom errors, %w, errors.Is / errors.As
- [ ] A-04: Goroutines properly — scheduler, leaks, WaitGroup
- [ ] A-05: Mutex + channels — natural extension of goroutines
- [ ] A-06: HTTP — net/http internals, routing, middleware (AFTER goroutines)
- [ ] A-07: Strings & Bytes
- [ ] A-08: Context — deadlines, cancellation
- [ ] A-09: Testing — table-driven, mocking with interfaces, benchmarks
- [ ] A-10: Generics

### Phase 2 — Idiomatic Go
- [ ] A-11: Functional patterns — closures, higher-order functions
- [ ] A-12: Concurrency patterns — worker pool, fan-out/fan-in, pipeline
- [ ] A-13: Design patterns — options pattern, repository
- [ ] A-14: Performance — profiling, escape analysis

---

## Track B: CS Fundamentals (in Go)
Starts when A-track Phase 1 is solid.

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
- [ ] C-08: gRPC — proto definitions, server/client, streaming
- [ ] C-09: Docker — containerizing a Go app
- [ ] C-10: Full system — API + Postgres + Redis + Kafka

---

## Weekend Projects (planned)
| Week | Concept | Project |
|------|---------|---------|
| W1 (2026-03-22) | Interfaces | CLI tool: quiz app or unit converter |
| W2 | Error handling | CLI: file parser / config reader |
| W3 | Goroutines | Concurrent word counter or downloader |
| W4 | Mutex + channels | Something with real concurrency |

---

## Resources
- https://go.dev/tour/
- https://gobyexample.com/
- https://go.dev/doc/effective_go
- https://100go.co/
- https://gophercises.com/
