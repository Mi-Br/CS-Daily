# Code Review: A-00 Ex4 (errors.go) — Final Revision
**Commit:** b98e32a — "errors challenge completed, inc code review feedback"  
**Date:** 2026-03-10  
**Reviewer:** Julia 🌸

---

## Test Results

```
go test ./...
--- FAIL: TestShapeEdgeCases/Negative_Shape (interfaces_test.go:64)
FAIL
```

**errors.go tests: ✅ ALL PASS**  
**interfaces.go test: ❌ Negative_Shape still fails**

---

## errors.go — What You Fixed ✅

This is a big step forward. The critical blockers from v2 are gone:

- ✅ **Compiles** — no more return-inside-loop killing compilation
- ✅ **Loop completes** — `return` moved outside, iterates all lines correctly
- ✅ **Key normalisation** — `strings.ToLower(k)` now correctly catches `Host`, `PORT`, `DEBUG` etc.
- ✅ **Port range validation** — `< 1 || > 65535` is correct
- ✅ **`%w` wrapping** — `fmt.Errorf("... %w ...", err)` present for port parse errors
- ✅ **`Unwrap()` method** — ConfigError now participates in `errors.Is/As` chains
- ✅ **errors_test.go added** — coverage for port range, missing host, bad bool, unknown keys

---

## Remaining Issues

### 1. Missing host — only catches `host=`, not absent `host`

Your test covers this case:
```go
{name: "missing host", input: `host=\n port=65534\n debug=true`}
```
That's `host=` with an empty value — and your code catches it:
```go
case "host":
    if value == "" {
        return cfg, &ConfigError{Field: key, Err: errors.New("missing value")}
    }
```

But what if `host` is never present in the input at all? Your function would happily return `Config{Host: "", Port: 65534, Debug: true}` with `nil` error. The spec says "Missing host → error" — which should mean both cases.

**Fix:** Add a post-parse check after the loop:
```go
if cfg.Host == "" {
    return cfg, &ConfigError{Field: "host", Err: errors.New("missing or empty")}
}
return cfg, nil
```

### 2. Go error message conventions — capitalisation

Go convention: error strings must be lowercase, no trailing punctuation (they get embedded in larger messages like `"failed to parse config: invalid integer ..."`).

```go
// ❌ Current
errors.New("Invalid integer %w, unable to convert")
errors.New("Invalid port range expect 1 ... 65535, got: %d")
errors.New("Error reading file")

// ✅ Fix
fmt.Errorf("invalid integer: %w", err)
fmt.Errorf("port out of range (1–65535): %d", int_val)
errors.New("error reading scanner")
```

Also: `fmt.Errorf("Invalid integer %w, unable to convert", err)` — `%w` mid-string with trailing text is unusual. Put the wrapped error at the end for clarity.

### 3. interfaces.go — Negative_Shape test still failing

```
--- FAIL: TestShapeEdgeCases/Negative_Shape
    interfaces_test.go:64: Physics broken, negative Area or Perimeter
```

Your `Rectangle.Area()` returns `r.a * r.b` — if either dimension is negative, the product can be negative. You need to decide: either reject negative inputs in the constructor/method, or take `math.Abs`. The test expects non-negative area even for negative inputs.

### 4. interfaces.go — Spurious `copy(s, s_copy)` still present

```go
s_copy := slices.Clone(s)
copy(s, s_copy)  // ← this line does nothing useful
slices.SortFunc(s_copy, ...)
```

`slices.Clone` already creates a new independent slice. `copy(s, s_copy)` then overwrites the original `s` with identical data — mutates the caller's slice for no gain. Delete that line.

---

## Verdict

**errors.go: Conditional Pass ✅**  
Core logic is correct, tests pass, compiles cleanly. Two small things to clean up (post-parse host check + error casing) — these are polish items, not blockers. If you fix them in a follow-up commit that's fine.

**interfaces.go: Not Yet Passing ❌**  
One failing test + the spurious copy. Both are small fixes.

**Suggested next commit:** Fix the two interfaces.go issues (negative area + remove the copy line). Then A-00 is fully done and we move to A-01.

---

## Overall Progress on A-00

| Exercise | Status |
|----------|--------|
| Ex1 TopN | ✅ Pass |
| Ex2 BankAccount | ✅ Pass |
| Ex3 Interfaces | 🔧 Failing test (negative area) |
| Ex4 Errors | ✅ Conditional pass (minor polish) |

Fix Ex3 and you're done with A-00. 🎯
