# Code Review: A-01 Maps & Slices
**Commit:** "A01 done"  
**Date:** 2026-03-14  
**Reviewer:** Julia 🌸

---

## Test Results

```
go test ./... -v  → ALL PASS ✅
go vet ./...      → CLEAN ✅
```

Tests pass. Vet is silent. That's the baseline — now let's look deeper.

---

## Exercise 1: TwoSum — `two_sums.go`

### What's working ✅
- O(n²) brute force is commented out as reference — exactly as asked
- O(n) map solution is correct and idiomatic: `seen := make(map[int]int)`, look up complement before storing — clean
- No input mutation

### Issues

#### 🟡 `ascOrder` is dead code

You wrote a helper to ensure the returned indices are in ascending order. The intent is correct, but the helper is never actually needed:

```go
seen := make(map[int]int)  // stores value → index, built left-to-right

for i, v := range nums {
    need := target - v
    index, exists := seen[need]
    if exists {
        return ascOrder(i, index)  // ← index is ALWAYS < i (it was stored earlier)
    }
    seen[v] = i
}
```

Because you iterate left to right and only check `seen` before inserting, `index` is always the earlier (smaller) index and `i` is always the later (larger) one. `ascOrder` always receives `(bigger, smaller)` and swaps — but `(smaller, bigger)` would be the same result.

You can just write:
```go
return index, i
```

The helper isn't wrong, it just reflects unnecessary uncertainty about your own algorithm. Once you trust the logic, it disappears.

#### 🟡 `else` after `return` — same pattern from A-00

```go
if exists {
    return ascOrder(i, index)
} else {          // ← this else is unreachable after return
    seen[v] = i
}
```

We covered this in A-00. After a `return`, the `else` adds visual noise without meaning. The clean version:

```go
if exists {
    return index, i
}
seen[v] = i
```

This came up in your A-00 reviews. Good to see you building the habit — just not there yet.

#### ❌ Missing test cases (required by README)

The README explicitly asked for:
- ✅ Basic case — covered
- ❌ Negative numbers — not tested
- ❌ Zero as element — not tested

Two Sum with negatives: `TwoSum([]int{-3, 4, 3, 90}, 0)` → `(0, 2)`. Your O(n) map handles this correctly, but it's untested. Test cases you specify are a contract — if you don't write them, you don't know.

---

## Exercise 2: WordFreq — `word_freq.go`

### What's working ✅
- Logic is correct
- Input slice not mutated
- Returns map (not nil) — good

### Issues

#### 🟡 You missed the elegance the README hinted at

The README said: *"Think about the zero value of `int` in a map — it makes this elegant."*

Your code:
```go
key, exists := seen[word]
if exists {
    seen[word] = key + 1
} else {
    seen[word] = 1
}
```

The hint was pointing at this: accessing a missing key in a Go map returns the zero value for the type. For `int`, that's `0`. So you can write:

```go
for _, word := range inp {
    seen[word]++
}
```

When `word` is new, `seen[word]` is `0`, and `0++` becomes `1`. When it's already there, it increments. The entire `if/else` disappears. This is one of the most useful Go patterns — worth internalising now.

#### 🟡 Function name: `WordFrequency` vs `WordFreq`

The README specifies `WordFreq`. You named it `WordFrequency`. Minor inconsistency — in a real codebase, API names matter. Follow the spec.

#### ❌ Missing test cases (required by README)

- ❌ Empty input — not tested
- ❌ Single word — not tested
- ✅ Repeated words + case-sensitive — covered (nicely, with "go"/"Go")

Empty input is particularly important — it's where `nil` vs empty map bugs hide.

---

## Exercise 3: Chunk — `chunk.go`

### What's working ✅
- `make` + `copy` pattern — this is exactly right. You understood the aliasing lesson.
- Loop logic and `end` clamping are correct
- Named subtests in the test file — good practice

### Issues

#### 🐛 `size > len(inp)` guard is a bug

```go
if size == 0 || size > len(inp) || len(inp) == 0 {
    return outp
}
```

When `size > len(inp)`, you return empty. But look at what your loop would do without that guard:

```
inp = [1, 2, 3], size = 10
i = 0, end = min(10, 3) = 3
segment = inp[0:3] = [1, 2, 3]
→ returns [[1, 2, 3]]
```

The spec says *"The last chunk may be smaller if len(s) is not divisible by n"* — a size larger than the slice just means the whole slice is the one (partial) chunk. Your guard incorrectly returns `[]` instead.

The fix is simply removing `size > len(inp)` from the guard condition:
```go
if size == 0 || len(inp) == 0 {
    return outp
}
```

#### ❌ Missing test cases (required by README)

- ❌ `n > len(s)` — not tested (and would have caught the bug above)
- ❌ `n = 1` — not tested
- ✅ Even split, uneven split, empty input — covered

`n = 1` matters because it's the maximum-number-of-chunks case and stresses the copy loop.

---

## Tests — Overall

You wrote tests. Good. But the required cases from the README weren't all covered. In A-01, missing test cases aren't just gaps — they let real bugs through (see Chunk's `size > len(inp)` bug which your tests missed).

**Pattern to fix:** When the README gives you a test list, treat it as a minimum spec. Write those cases first, then add more.

---

## Verdict

| Exercise | Logic | Tests | Issues |
|----------|-------|-------|--------|
| TwoSum   | ✅ Correct | 🟡 Missing cases | `else` after return, dead helper |
| WordFreq | ✅ Correct | 🟡 Missing cases | Missed zero-value elegance, name mismatch |
| Chunk    | 🐛 One bug | 🟡 Missing cases | `size > len(inp)` returns wrong result |

**Overall: Conditional Pass 🟡**

The core skills are there — you got the map pattern, you got the `copy` anti-aliasing lesson, you left the brute force as reference. That's solid. But two patterns keep following you:

1. **`else` after `return`** — third review it appears. Time to kill it permanently.
2. **Missing edge case tests** — you're writing the happy path and stopping. Go further.

---

## Required Fixes (before moving to A-02)

1. Fix `Chunk`: remove `size > len(inp)` from the early-return guard
2. Add missing test cases: TwoSum (negatives, zero), WordFreq (empty), Chunk (n=1, n > len)
3. Remove the `else` in TwoSum after the `return`
4. Rename `WordFrequency` → `WordFreq` to match the spec

Optional but encouraged:
- Simplify WordFreq to `seen[word]++` — one line, no if/else

When those are in, A-01 is done. 🎯

---

*Reviewed by Julia — your senior staff engineer on call.*
