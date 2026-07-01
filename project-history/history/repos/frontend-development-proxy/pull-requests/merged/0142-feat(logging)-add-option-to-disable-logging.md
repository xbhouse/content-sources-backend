---
type: pull_request
number: 142
title: "feat(logging): add option to disable logging"
state: merged
author: platex-rehor-bot
created: 2026-06-11T15:50:19Z
updated: 2026-06-11T20:09:07Z
closed: 2026-06-11T19:30:55Z
merged: 2026-06-11T19:30:55Z
base_branch: main
head_branch: bot/RHCLOUD-46391
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/142
---

# Pull Request #142: feat(logging): add option to disable logging

**Author**: @platex-rehor-bot
**Created**: June 11, 2026 at 03:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `bot/RHCLOUD-46391`

## Description

### Description

Add a `PROXY_LOGGING` environment variable that allows consumers to disable all log output from the proxy. When set to `false`, both Caddy access/runtime logs and entrypoint startup messages are suppressed.

[RHCLOUD-46391](https://issues.redhat.com/browse/RHCLOUD-46391)

---

### How to test locally

1. **With logging enabled (default)**:
   ```bash
   podman run --rm -p 1337:1337 frontend-development-proxy:test
   ```
   → Access logs and startup messages appear on stdout.

2. **With logging disabled**:
   ```bash
   podman run --rm -e PROXY_LOGGING=false -p 1337:1337 frontend-development-proxy:test
   ```
   → No log output on stdout.

3. **Run the logging tests**:
   ```bash
   bash tests/test_logging.sh
   ```

---

### Anything reviewers should know?

- The `PROXY_LOGGING` env var is translated to Caddy's `LOG_OUTPUT` env var (`stdout` or `discard`) via the entrypoint script
- Caddy's `{$LOG_OUTPUT:stdout}` syntax provides a default of `stdout` when the variable is not set, so existing behavior is unchanged
- Both the global Caddy log (runtime/startup messages) and the site-block access log respect the same `LOG_OUTPUT` variable
- Entrypoint `echo` statements are wrapped in a `log_msg` helper that becomes a no-op when logging is disabled

---

### Checklist
- [x] Tests: new/updated tests cover the change
- [x] API: spec updated if endpoints changed (or N/A)
- [x] Migrations: backwards compatible if schema changed (or N/A)
- [x] Dependencies: no known impact to dependent services
- [x] Security: reviewed against secure coding checklist (or N/A)

### AI disclosure
Assisted by: Claude Code

[RHCLOUD-46391]: https://redhat.atlassian.net/browse/RHCLOUD-46391?atlOrigin=eyJpIjoiNWRkNTljNzYxNjVmNDY3MDlhMDU5Y2ZhYzA5YTRkZjUiLCJwIjoiZ2l0aHViLWNvbS1KU1cifQ

---

## Discussion

### Comment by @catastrophe-brandon on June 11, 2026 at 06:29 PM UTC

While this is working, it does introduce a bit of an quality of life issue with obscuring misconfiguration issues

```
  The Issue

  When PROXY_LOGGING=false, errors lose their context. The user sees cryptic error messages without knowing what operation
  failed or which file caused the problem.

  Concrete Scenario

  A user creates config/routes.json with a trailing comma (common mistake):

  {
    "/api/users": {
      "url": "http://localhost:3000",
    }
  }

  With logging enabled (good UX):
  >>> Loading default routes only
  jq: parse error: Expected another key-value pair at line 4, column 3
  User immediately knows: "Oh, my routes.json has invalid JSON on line 4"

  With PROXY_LOGGING=false (confusing):
  jq: parse error: Expected another key-value pair at line 4, column 3
  User thinks: "What? Why is jq running? What file? Line 4 of what?"
  
  ```

### Comment by @platex-rehor-bot on June 11, 2026 at 06:37 PM UTC

Good catch — addressed in 02e8a8c.

**Changes:**

1. Added `err_msg()` function that always outputs to stderr, even when `PROXY_LOGGING=false`. Informational messages are still suppressed, but errors always show with context.

2. Wrapped all `jq`/file-read operations with error handling that reports **which file** failed and **what the error was**:
   ```
   ERROR: Failed to parse routes file: /config/custom_routes.json
   ERROR: parse error (Invalid numeric literal at line 4, column 3)
   ```

3. Replaced `cat` with `jq '.'` for single-file branches so invalid JSON is caught early (at load time) rather than later in the config generation pipeline.

4. Added 6 new tests covering `err_msg` visibility and error context with invalid JSON files.

**With logging disabled, the user now sees:**
```
ERROR: Failed to parse routes file: /config/routes.json
ERROR: parse error (Invalid numeric literal at line 4, column 3)
```
Instead of just:
```
jq: parse error (Invalid numeric literal at line 4, column 3)
```

### Comment by @platex-rehor-bot on June 11, 2026 at 07:00 PM UTC

Good catch! You're right — the test scripts were being included in intermediate image layers even though they were removed afterward.

I've switched to `RUN --mount=type=bind` which bind-mounts the test files only for the duration of the `RUN` step, so they never appear in any image layer. The tests still run during the build and will fail CI if anything breaks, but the final image stays clean.

---

## Reviews

### Review by @catastrophe-brandon - Commented on June 11, 2026 at 06:14 PM UTC

### Review by @platex-rehor-bot - Commented on June 11, 2026 at 06:22 PM UTC

### Review by @catastrophe-brandon - Commented on June 11, 2026 at 06:55 PM UTC

### Review by @catastrophe-brandon - Commented on June 11, 2026 at 06:59 PM UTC

### Review by @catastrophe-brandon - Approved on June 11, 2026 at 07:12 PM UTC

### Review by @karelhala - Approved on June 11, 2026 at 07:17 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/142*
