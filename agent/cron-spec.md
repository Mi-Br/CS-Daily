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

### Step 4: Generate next challenge (when ready)
Triggered when:
- Current challenge is marked complete in progress.json, OR
- Michail says "ready for next" in Telegram

Generation process:
1. Read `agent/curriculum.md` → find next uncompleted challenge
2. Check `progress.json` → check for weak spots that need reinforcement
3. Search web for 1-2 real articles/blog posts on the concept
4. Write challenge to `challenges/YYYY-MM-DD_<title>/README.md`
5. Push to repo
6. Send challenge to Telegram (Hustle group)

## Challenge file format

```markdown
# Challenge: [Title]
**Track:** A/B/C  **ID:** A-XX  **Date:** YYYY-MM-DD

## 📖 Read First
- [Article title](URL) — why: one sentence on what you'll get from it
- [Article title](URL) — why: ...

## 🎯 The Challenge
[Clear problem statement]

## ✅ Done When
- [ ] Criterion 1
- [ ] Criterion 2
- [ ] Tests written and passing
- [ ] `go vet` passes with no warnings

## 💡 Hints (read only if stuck)
<details>
<summary>Hint 1</summary>
...
</details>

## 📁 Where to put your code
`challenges/YYYY-MM-DD_<title>/`
```

## Telegram messages
- Reviews: sent to "Hustle Michail & July" group (-5271500205)
- Nudges: same group
- New challenges: same group
