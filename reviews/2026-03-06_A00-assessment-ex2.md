# Code Review: A-00 Exercise 2 — Struct + Pointer Receivers
**Date:** 2026-03-06  
**File reviewed:** `challenges/2026-03-04_A00-assessment/bank_account.go` + `bank_account_test.go`  
**Verdict:** ✅ Pass (with important notes) — continue to remaining exercises

---

## Summary

Good first pass. The logic is correct, pointer receiver reasoning is solid, and tests cover the core happy/sad paths. But there are a few issues that matter — one of them is a meaningful miss against the spec. Let's go through them.

---

## 🔴 Critical

### 1. You implemented `Print()` instead of `String()`

The challenge explicitly asked for `String() string`. This isn't pedantry — it's important because `String() string` satisfies the `fmt.Stringer` interface:

```go
type Stringer interface {
    String() string
}
```

When you implement `Stringer`, Go's `fmt` package automatically calls it when you do `fmt.Println(account)`. It's one of Go's fundamental interfaces. Renaming it to `Print()` means your type *doesn't* implement the interface, and you lose all of that for free.

**Fix:** Rename `Print()` to `String()`. No other changes needed.

```go
// Before
func (acc Account) Print() string { ... }

// After
func (acc Account) String() string { ... }
```

Then test it like this:
```go
fmt.Println(test_account) // will call String() automatically
```

---

## 🟡 Significant

### 2. Snake_case field names — Go isn't Python

```go
type Account struct {
    balance        float64
    account_number string  // ← wrong
    account_owner  Customer
}
```

Go uses **camelCase** for all identifiers. `account_number` should be `accountNumber`. `account_owner` should be `accountOwner`. This isn't style preference — it's the language convention, and tools like `golint` will flag it.

### 3. Withdraw doesn't guard against negative amounts

```go
func (acc *Account) Withdraw(amount float64) error {
    if (acc.balance - amount) >= 0 {
        acc.balance -= amount
        return nil
    }
    return errors.New("Error: Withrdawing more than balance")
}
```

What happens if someone calls `acc.Withdraw(-100)`? The condition `-100 >= 0` is false — but instead of subtracting, you'd... wait, actually `balance - (-100)` would be positive, so it *passes* and *adds* 100 to the balance. You'd get a "withdrawal" that increases the balance.

The current condition hides this. A clearer, explicit approach:

```go
func (acc *Account) Withdraw(amount float64) error {
    if amount < 0 {
        return errors.New("withdrawal amount cannot be negative")
    }
    if amount > acc.balance {
        return errors.New("insufficient funds")
    }
    acc.balance -= amount
    return nil
}
```

### 4. Error message formatting

Go error messages by convention:
- **Lowercase** (no capital first letter)
- **No punctuation** at the end
- **No "Error:" prefix** — callers add context when wrapping

So `errors.New("Error: Withrdawing more than balance")` should be `errors.New("insufficient funds")`.

Also — typo: "Withrdawing" → "Withdrawing".

---

## 🟢 What You Got Right

**Pointer receiver reasoning is exactly right.** The comments:
```go
// we modify balance so pointer receiver makes sense
func (acc *Account) Withdraw(amount float64) error { ... }

// read only so value pointer is fine
func (acc Account) Balance() float64 { ... }
```

This is the correct mental model. Pointer receiver when you're mutating state, value receiver when you're just reading. You articulated it correctly — that's the key thing the exercise was testing.

**The `Deposit` negative guard is correct.** Checking `amount < 0` before doing anything is the right pattern.

---

## 🧪 Test Review

**Good:** Table-structure avoided, but the cases cover what matters (deposit, negative deposit, overdraft). That's the minimum viable test surface.

**Issues:**

**a) Shared mutable state across subtests.** You initialize `test_account` once and run three subtests against it. The deposit test increases the balance by 100 before the overdraft test runs — so the overdraft test is checking against 600, not 500. It still passes, but the test behavior is now order-dependent. If you reorder the tests, results may change.

Better pattern: initialize a fresh account inside each `t.Run`:
```go
t.Run("Test deposit", func(t *testing.T) {
    acc := Account{balance: 500, ...}
    // test against acc
})
```

**b) Use `t.Fatal` where continuation doesn't make sense.** In the deposit test:
```go
if err != nil {
    t.Error("Got error depositing amount, expect none")
} else {
    if test_account.Balance() != cb+dep {
        t.Error("Balance after deposit does not match")
    }
}
```

If `Deposit` returns an error, the `else` block won't run anyway because of the `if/else`. But using `t.Fatal` (or `t.Fatalf`) is cleaner — it stops the test immediately when a prerequisite fails. Use `t.Error` when you want to collect multiple failures; `t.Fatal` when further assertions don't make sense.

**c) No test for negative withdrawal.** Given the bug described above, a test like this would have caught it:
```go
t.Run("Negative withdrawal amount", func(t *testing.T) {
    err := test_account.Withdraw(-100)
    if err == nil {
        t.Error("Expected error for negative withdrawal amount")
    }
})
```

---

## Pointer Receiver — One More Thing

You said "read only so value pointer is fine" for `Balance()`. Correct. But be consistent: if `Account` has *any* pointer receiver methods, Go recommends that *all* methods use pointer receivers (or none). Mixing pointer and value receivers on the same type can lead to subtle bugs with interfaces.

In practice: if your struct is mutated by some methods, make everything a pointer receiver. It's simpler and avoids the confusion.

---

## What's Left in A-00

You've done Exercise 1 (TopN) and Exercise 2 (BankAccount). Still outstanding:
- Exercise 3: Interface + multiple types (Shape, Circle, Rectangle, Triangle)
- Exercise 4: Error handling (ParseConfig + custom error type)
- Exercise 5: Goroutines + channels (FetchAll)

These get progressively harder. Exercise 5 is where most developers hit friction — the ordering constraint is where it gets interesting. Take your time on it.

---

## TL;DR

| Issue | Severity |
|---|---|
| `Print()` instead of `String()` — misses `fmt.Stringer` | 🔴 Fix it |
| `account_number` snake_case — use `accountNumber` | 🟡 Fix it |
| Negative withdrawal not guarded | 🟡 Fix it |
| Error message formatting (uppercase, typo) | 🟡 Fix it |
| Shared mutable state in tests | 🟡 Fix it |
| No negative withdrawal test | 🟡 Add it |
| Pointer receiver reasoning | ✅ Solid |
| Deposit negative guard | ✅ Correct |
