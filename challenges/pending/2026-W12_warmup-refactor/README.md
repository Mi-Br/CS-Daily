# W12 Warmup — Refactor: Rate Limiter Cleanup

**Assigned:** 2026-03-25  
**Track:** W12 — Struct Design  
**Estimated time:** ~20 min  
**Package:** use the W-01 rate limiter code as base

---

## The Task

Your W-01 solution works but has a few rough edges. Refactor it without changing external behaviour — all existing tests must still pass.

### Things to fix

**1. `RemoteAddr` includes the port**

`r.RemoteAddr` returns `"1.2.3.4:5678"` — a new ephemeral port per connection. Two requests from the same client end up as different keys. Fix the middleware to key on IP only:

```go
import "net"

clientID, _, _ := net.SplitHostPort(r.RemoteAddr)
```

**2. Clean up the leftover `// TODO: implement` comment**

It's in `server.go` above the middleware implementation. The code is there — the comment isn't. Remove it.

**3. Consider: should `clientState` be a value or pointer in the map?**

Currently: `map[string]*clientState`. Think through: what breaks if you change it to `map[string]clientState`? Write a one-line comment in the code explaining your choice.

---

## Done When

- [ ] Middleware keys on IP, not IP+port
- [ ] TODO comment removed
- [ ] `map[string]*clientState` vs `map[string]clientState` — choice documented in a comment
- [ ] All existing tests still pass
- [ ] `go vet` clean
