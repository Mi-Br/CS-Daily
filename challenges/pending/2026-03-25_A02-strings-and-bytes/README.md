# A-02: Strings & Bytes

**Assigned:** 2026-03-25  
**Track:** A (Fundamentals)  
**Package:** `a02`

---

## Background

In Go, strings are immutable byte slices — not arrays of characters. This has real consequences: indexing a string gives you a `byte`, not a `rune`. Iterating with `range` gives you `rune`s (Unicode code points). Understanding this distinction is essential before moving to any string-heavy work.

This challenge has three exercises. Do them in order — they build on each other.

---

## Exercise 1: Reverse Words

Write a function that reverses the order of words in a sentence.

```go
func ReverseWords(s string) string
```

- `"hello world"` → `"world hello"`
- `"  go is great  "` → `"great is go"` (trim leading/trailing spaces, normalize internal spaces)
- `""` → `""`

**Do not** use `strings.Reverse` (it doesn't exist) or any external package beyond the standard library.

---

## Exercise 2: Caesar Cipher

Write an encoder and decoder for the Caesar cipher.

```go
func CaesarEncode(s string, shift int) string
func CaesarDecode(s string, shift int) string
```

Rules:
- Only shift **ASCII letters** (a–z, A–Z). Leave everything else unchanged (spaces, digits, punctuation).
- Wrap around: `z` + 1 = `a`, `Z` + 1 = `A`
- `shift` can be negative or larger than 26 — handle it (hint: use modulo)
- `CaesarDecode(CaesarEncode(s, n), n) == s` must always hold

Examples:
- `CaesarEncode("Hello, World!", 3)` → `"Khoor, Zruog!"`
- `CaesarEncode("xyz", 3)` → `"abc"`
- `CaesarDecode("Khoor, Zruog!", 3)` → `"Hello, World!"`

---

## Exercise 3: Run-Length Encoding

Implement a simple run-length encoder and decoder.

```go
func RLEEncode(s string) string
func RLEDecode(s string) string
```

Run-length encoding compresses repeated characters:
- `"aaabbc"` → `"3a2b1c"`
- `"abcd"` → `"1a1b1c1d"`
- `""` → `""`

For decoding:
- `"3a2b1c"` → `"aaabbc"`
- Assume valid encoded input (you don't need to handle malformed strings)

---

## What to Submit

- `strings.go` — your implementations
- `strings_test.go` — your tests (write them yourself!)

**Test requirements:**
- At least 2 test cases per function
- Cover the edge cases mentioned above (empty string, wrap-around for Caesar, etc.)

---

## Hints (read only if stuck)

<details>
<summary>Exercise 1</summary>

`strings.Fields` splits on any whitespace and trims automatically. Then reverse the resulting slice.

</details>

<details>
<summary>Exercise 2</summary>

Cast to `rune` to do arithmetic. To wrap: `((r - base + shift) % 26 + 26) % 26 + base`. The `+26` before the final `%26` handles negative shifts.

</details>

<details>
<summary>Exercise 3</summary>

For encoding: iterate with `range`, track current char and count. When char changes, write `count + char` to a `strings.Builder`. For decoding: iterate two characters at a time — digit(s) then letter.

</details>
