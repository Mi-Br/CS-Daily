# Daily Cron Job Spec

## Schedule
Runs every day at **09:00 Amsterdam time** (configurable).

## What it does

### Step 1: Check for commits
- Fetch latest commits from `git@github.com:Mi-Br/CS-Daily.git`
- Look at diffs in `challenges/` directory since last review
- If no new commits → send a gentle nudge to Telegram (Hustle group)
- If commits found → proceed to review

### Step 2: Code Review
For each changed `.go` file:
- Read the code
- Review as a senior staff engineer:
  - Correctness (does it solve the problem?)
  - Idiomatic Go (naming, error handling, style)
  - Edge cases (what's not handled?)
  - Performance considerations
  - Test coverage
- Write review to `reviews/YYYY-MM-DD_<challenge>.md`
- Push review file to repo
- Send summary to Telegram (Hustle group)

### Step 3: Update progress
- Update `agent/progress.json`:
  - Add observations about what was strong/weak
  - Flag weak spots if a concept appears shaky
  - If challenge is complete + review passed → mark completed

### Step 4: Generate reading (Monday, ahead of the week)
Triggered on Monday morning, or when a new concept is coming up.

Generation process:
1. Read `agent/curriculum.md` → find the concept for the current/next week
2. Search web for 2-4 real articles/blog posts on that concept
3. Write reading file to `reading/YYYY-MM-DD_topic.md` (see format below)
4. Update `reading/README.md` index
5. Push to repo
6. Send short note to CS Daily group: "Reading for this week is up → reading/YYYY-MM-DD_topic.md"

### Step 5: Generate warmup (Thu/Fri, after reading)
Triggered on Thu or Fri morning, only after the week's reading file exists.

Generation process:
1. Check that `reading/YYYY-MM-DD_topic.md` for this concept exists
2. Write warmup to `warmups/YYYY-MM-DD_topic/README.md` + `warmup.go` + `warmup_test.go` + `go.mod`
3. Update `warmups/README.md` index
4. Push to repo
5. Send short note to CS Daily group: "Warmup is up → warmups/YYYY-MM-DD_topic/"

### Step 6: Generate next weekend challenge (when ready)
Triggered when:
- Current challenge is marked complete in progress.json, OR
- Michail says "ready for next"

Generation process:
1. Read `agent/curriculum.md` → find next uncompleted challenge
2. Check `progress.json` → check for weak spots that need reinforcement
3. Write challenge to `challenges/pending/YYYY-MM-DD_<title>/README.md`
4. Push to repo
5. Send short note to CS Daily group: "Weekend challenge is up → challenges/pending/..."

## File formats

### Reading file: `reading/YYYY-MM-DD_topic.md`
```markdown
# [Topic Name]
**Date:** YYYY-MM-DD | **Concept:** ... | **Est. read time:** ~X min

## Why this matters
[1-2 sentences — why this concept is relevant to what Michail is building]

## Read these (in order)
1. **[Title](URL)** — why: one sentence
2. **[Title](URL)** — why: one sentence

## Key concepts to come away with
[Bullet points of the core ideas, with short code snippets if helpful]

## Questions to answer after reading
[3-5 questions Michail should be able to answer before doing the warmup]

## Your Notes
> (Michail adds notes here as he reads)
```

### Warmup folder: `warmups/YYYY-MM-DD_topic/`
```
README.md       ← exercise description (prereading link, task, done-when checklist)
warmup.go       ← stub with types + empty function signature
warmup_test.go  ← comment-only scaffold ("write two tests: ...")
go.mod
```

### Challenge folder: `challenges/pending/YYYY-MM-DD_title/`
```markdown
# Challenge: [Title]
**Track:** A/B/C  **ID:** A-XX  **Date:** YYYY-MM-DD

## The Challenge
[Clear problem statement — no scaffolded solution]

## Done When
- [ ] Criterion 1
- [ ] Tests written and passing
- [ ] `go test -race ./...` passes
- [ ] `go vet ./...` clean

## Hints (read only if stuck)
<details><summary>Hint 1</summary>...</details>
```

## Messaging
- All messages go to **CS Daily Michail & July** group (chat ID: -5145388620)
- Never send to Hustle group for CS Daily content
- Keep messages short: "Reading for this week is up → reading/2026-03-23_mutexes.md"
