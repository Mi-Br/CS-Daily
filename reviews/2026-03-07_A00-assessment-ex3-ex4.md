# Code Review — A-00 Assessment: Exercises 3 & 4
**Date:** 2026-03-07  
**Reviewer:** Julia (CS-Daily Agent)  
**Files reviewed:** `interfaces.go`, `errors.go`  
**Commit:** `72fb0bc`

---

## Exercise 3: Interfaces (`interfaces.go`)

### What works well ✅

- Interface definition is clean and idiomatic. `Area()` and `Perimeter()` with `float64` — exactly right.
- All three shapes implement `Shape` correctly. Go's implicit interface satisfaction is used properly.
- Heron's formula is correct. The `s` variable is well-named (semi-perimeter).
- `PrintShapeInfo` does the job. `%.2f` formatting is a thoughtful touch.
- The commented-out `LargestAreaShape` implementation is great self-documentation — you explored the manual loop, recognized a cleaner alternative, kept both for comparison. That thinking process is the right instinct.

### Issues to address 🔴

**1. `LargestAreaShape` returns the wrong element (critical bug)**

```go
slices.SortFunc(slist, func(a, b Shape) int {
    return cmp.Compare(a.Area(), b.Area())
})
return slist[0]
```

`SortFunc` sorts in **ascending order** by default (smallest first). So `slist[0]` is the *smallest* area, not the largest. You need `slist[len(slist)-1]`, or reverse the comparator:

```go
return cmp.Compare(b.Area(), a.Area()) // descending
```

Then `slist[0]` is the largest.

**2. Function name mismatch**

The challenge asks for `LargestArea(shapes []Shape) Shape`. You named it `LargestAreaShape`. Not a correctness issue in isolation, but if a caller or test uses the specified signature, it won't compile. In real codebases, spec compliance matters — name it exactly as stated unless you have a good reason not to.

**3. Mutating the input slice**

`slices.SortFunc(slist, ...)` **sorts the caller's slice in place**. The caller doesn't expect their `[]Shape` to be reordered after calling `LargestArea`. This is the same aliasing issue from Exercise 1 — Go slices are reference types. If you need to sort, work on a copy:

```go
func LargestArea(shapes []Shape) Shape {
    sorted := make([]Shape, len(shapes))
    copy(sorted, shapes)
    slices.SortFunc(sorted, func(a, b Shape) int {
        return cmp.Compare(b.Area(), a.Area())
    })
    return sorted[0]
}
```

Or better for this case: a single linear scan is O(n) and avoids the copy entirely:

```go
func LargestArea(shapes []Shape) Shape {
    if len(shapes) == 0 {
        return nil
    }
    largest := shapes[0]
    for _, s := range shapes[1:] {
        if s.Area() > largest.Area() {
            largest = s
        }
    }
    return largest
}
```

This is idiomatic Go for "find max in a collection." Sorting a collection just to find a max is O(n log n) when O(n) is available. Your commented-out manual loop was actually the *better* algorithm — you abandoned the right solution for a fancier one that's slower and breaks the input.

**4. Missing nil/empty guard**

No guard for `len(slist) == 0`. Accessing `slist[0]` on an empty slice panics. The original commented-out version had this handled. Don't lose it when you refactor.

---

### Minor notes 🟡

- `PrintShapeInfo` uses `fmt.Printf` without a newline at the end. Usually you'd want `\n`. Minor, but noticeable.
- Unused import risk: make sure `cmp` and `slices` are both actually used after final cleanup.
- No tests for `interfaces.go`. The challenge requires at least basic tests. Where's `interfaces_test.go`?

---

## Exercise 4: Error Handling (`errors.go`)

### What works well ✅

- `ConfigError` struct exists with `field` and `err` fields — right shape.
- `strings.Cut` for key/value splitting — exactly the idiomatic approach.
- `bufio.Scanner` for line-by-line reading — solid choice.
- `strconv.ParseBool` for Debug — correct.

### Issues to address 🔴

**1. `ConfigError` doesn't implement the `error` interface**

```go
func (cr ConfigError) Error() error {
    return errors.New(fmt.Sprintf("%s is %s", cr.field, cr.err))
}
```

The `error` interface requires `Error() string`, not `Error() error`. This code **doesn't compile**. Fix:

```go
func (e ConfigError) Error() string {
    return fmt.Sprintf("field %q: %v", e.field, e.err)
}
```

**2. `ParseConfig` signature doesn't match the spec**

The challenge asks for `ParseConfig(data string) (Config, error)`. You return `(Config, ConfigError)`. Two problems:

- `ConfigError` isn't an `error` (see above), so this doesn't satisfy the interface.
- Returning a concrete error type instead of the `error` interface is almost always wrong in Go. The caller can't use `errors.Is` / `errors.As` without jumping through hoops. Return `error` and let the concrete type live underneath.

**3. Empty `ConfigError{}` return pattern**

```go
return cfg, ConfigError{}
```

In Go, a "no error" return is `nil`, not a zero-value struct. Since `ConfigError` is a value type, a zero-value `ConfigError{}` is *not nil* — any caller doing `if err != nil` will incorrectly see an error. This is a classic Go footgun. Fix: change return type to `error` and return `nil` for success.

**4. Port error uses wrong variable in error message**

```go
p_int, err := strconv.ParseInt(v, 10, 64)
if err != nil {
    return cfg, ConfigError{field: string(p_int), err: fmt.Errorf("Not valid integer")}
}
```

`p_int` is `0` when `ParseInt` fails. `string(0)` is `"\x00"`, not `"0"`. The field should be the key name (`"Port"` / `"port"`), not the parsed value:

```go
return cfg, ConfigError{field: "port", err: fmt.Errorf("not a valid integer: %w", err)}
```

**5. Port range validation missing**

The spec requires: "If `port` < 1 or > 65535 → error". This check is absent entirely.

**6. Missing `host` validation isn't reliable**

You check `len(v) == 0` for an empty host value, but there's no check at the end of parsing that verifies `host` was present at all. If the config string has no `host=` line, no error is returned. You need a post-parse check:

```go
if cfg.Host == "" {
    return cfg, &ConfigError{field: "host", err: fmt.Errorf("missing")}
}
```

**7. `return cfg, ConfigError{}` inside the loop — only first line parsed**

```go
for scanner.Scan() {
    // ...
    if found {
        // handle key
        return cfg, ConfigError{}  // ← exits after the first key!
    }
}
```

The function returns after processing the very first key-value pair. A config with 3 lines (`host`, `port`, `debug`) only processes `host` and returns. Remove this inner return and return after the loop finishes.

**8. Key casing inconsistency**

The spec shows lowercase keys (`host`, `port`, `debug`). Your code checks `"Port"`, `"Host"`, `"Debug"` (capitalized). The test input `"port=8080"` would not match `"Port"`. Pick one and be consistent; lowercase is what the spec shows.

**9. Error message style**

`"Not valid integer"` — Go error messages should be lowercase and not end with punctuation: `"not a valid integer"`. See: [https://github.com/golang/go/wiki/CodeReviewComments#error-strings](https://github.com/golang/go/wiki/CodeReviewComments#error-strings).

**10. `fmt.Errorf` without `%w`**

The spec says: *"Use `fmt.Errorf` with `%w` for wrapping where appropriate."* None of your `fmt.Errorf` calls use `%w`. When wrapping an existing error (like from `strconv.ParseInt`), use `%w` so callers can do `errors.Is(err, strconv.ErrSyntax)` etc.

---

### What a correct `ConfigError` looks like

```go
type ConfigError struct {
    Field string
    Err   error
}

func (e *ConfigError) Error() string {
    return fmt.Sprintf("config field %q: %v", e.Field, e.Err)
}

func (e *ConfigError) Unwrap() error {
    return e.Err
}
```

Exporting `Field` and `Err` lets callers inspect them. Implementing `Unwrap` makes `errors.As` / `errors.Is` work through the chain.

---

## Summary

| | Ex 3 | Ex 4 |
|---|---|---|
| Compiles | ✅ (probably) | ❌ (Error() return type wrong) |
| Correct output | ❌ (returns min, not max) | ❌ (loop exits early, no port range check) |
| Spec compliance | 🟡 (name mismatch) | ❌ (signature, key casing) |
| Idiomatic Go | 🟡 (mutates input) | ❌ (zero-value error pattern) |
| Tests | ❌ (missing) | ❌ (missing) |

---

## What to fix (priority order)

1. **`errors.go`:** Fix `Error() string`, fix the early return in the loop, fix key casing, add port range check, add post-parse host check, return `nil` for success.
2. **`interfaces.go`:** Fix `LargestAreaShape` — either reverse comparator + use last element, or use the linear scan approach (which was already correct in your commented code!).
3. **Add `interfaces_test.go`** — at minimum: test each shape's Area/Perimeter, test LargestArea with mixed shapes.
4. **Stop mutating the input** in `LargestAreaShape`.

---

## Trend note

This is the third review where aliasing / mutation of input appears. It's clearly a pattern. Before you write any function that takes a slice: ask yourself, *"am I allowed to modify this?"* Default answer: no, unless the caller explicitly passes ownership. This will come up constantly in real Go codebases.

The `ConfigError.Error() error` mistake is interesting — it shows that the *shape* of the interface wasn't fully internalized yet. The `error` interface is one of the most important in Go. Worth spending 10 minutes re-reading [https://go.dev/blog/error-handling-and-go](https://go.dev/blog/error-handling-and-go) before Ex 5.

Good effort pushing on a Saturday. Now fix these and push again — Ex 5 (goroutines) is the most fun one. 🚀
