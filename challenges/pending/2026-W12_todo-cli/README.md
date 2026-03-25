# W12 Weekend Project — `todo` CLI

**Assigned:** 2026-03-22  
**Track:** W12 — Struct Design  
**Week:** 2026-W12 (due by end of week)

---

## What to Build

A CLI task manager. From scratch — no frameworks, no libraries beyond the standard library.

```
todo add "Buy groceries"   → adds a task
todo list                  → shows all tasks
todo done 1                → marks task 1 complete
todo delete 1              → removes a task
```

Tasks **persist between runs** — saved to a JSON file on disk.

---

## Done When

- [ ] All 4 commands work (`add`, `list`, `done`, `delete`)
- [ ] Data survives closing and reopening the terminal
- [ ] Graceful errors (e.g. `todo done 99` when task 99 doesn't exist)
- [ ] `go vet` passes clean

---

## What to Figure Out Yourself (the point of the exercise)

- What structs do you need? What fields?
- How do you store tasks in memory — slice or map?
- Where does save/load logic live? On what type?
- How do you parse `os.Args`?

---

## OK to Google

- `os.Args` — how to read CLI arguments
- `os.ReadFile` / `os.WriteFile` — file I/O
- `json.Marshal` / `json.Unmarshal`
- `fmt.Fprintf(os.Stderr, ...)` — printing errors

## NOT OK to Google

- "Go todo CLI tutorial" — you'll copy someone's struct design. That's the whole point of this week.

---

## Reading (if you haven't done it yet)

1. [Effective Go — Composite Literals](https://go.dev/doc/effective_go#composite_literals) — 10 min
2. [Data semantics in Go — Bill Kennedy](https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html) — 20 min
3. [JSON in Go — Go by Example](https://gobyexample.com/json) — type it out, don't just read
4. [Struct embedding](https://gobyexample.com/struct-embedding) — 5 min
