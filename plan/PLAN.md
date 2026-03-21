# Training Plan

**Last updated:** 2026-03-21  
**Updated by:** Julia

This is our shared document. Michail adds notes, feedback, and flags things he finds difficult.
Julia reads this before creating any new challenge and adjusts accordingly.

---

## Current Week

→ See `weeks/2026-W12.md`

---

## Upcoming Weeks (rough order, subject to change)

| Week | Concept | Why |
|------|---------|-----|
| W12 (this week) | Struct design | W-01 exposed that design from scratch is the real gap |
| W13 | Error handling | Syntax known, judgment not yet |
| W14 | Goroutines properly | Fill the conceptual gap under W-01 |
| W15 | Mutex + channels | W-01 revisited with real understanding |
| W16 | HTTP | Now with goroutine foundation |

---

## What Michail Finds Difficult

> **📝 Michail — add your notes here. Anything you struggled with, found confusing, or want to revisit.**
> Julia reads this before every new challenge.

- Struct composition — deciding what types I need from scratch, what fields they should have
- Mutexes — understand the principle but couldn't write it from scratch
- HTTP middleware — same, understand it but need more foundation
- Reading other people's code is still slow for me
- Writing tests — know the pattern but need more muscle memory

---

## What's Working

- Pointer vs value receivers — clicked
- Slice aliasing and copy pattern — absorbed
- Table-driven tests — adopted
- Static interface compliance checks (`var _ Interface = ...`) — using naturally

---

## Pace & Rhythm

- **Weekdays:** ~1h per morning
- **Commute:** Reading (phone-friendly articles)
- **Weekends:** Longer project — something real and demo-able

---

## Rules (agreed 2026-03-21)

1. No new concept until the previous one is solid
2. Every exercise: write from scratch, no big scaffolds
3. Each challenge states: what you need to know, what to write from scratch, what's OK to google
4. Code reading challenge every ~2 weeks
5. Weekend project must be demo-able: "look what I built"
6. Michail adds notes to this file → Julia incorporates before next challenge
