---
repo: content-sources-backend
pr_number: 956
period: convergence
source: history/repos/content-sources-backend/pull-requests/merged/
---

# Deep Dive: content-sources-backend PR #956

## Problem / Motivation

Backend integration tests were manual or absent. HMS-5287 introduced Playwright E2E testing with CI — following frontend adoption and preceding patchman-ui (#1445).

## Solution Approach

@Andrewgdewar added `_playwright-tests` directory and GitHub Actions workflow. **No production code changes** — pure test infrastructure PR.

## Key Design Decisions

- **`_playwright-tests` naming**: Consistency with frontend convention (@Andrewgdewar chose over `./tests/playwright/` to keep node_modules one layer deep).
- **CI gate on PRs**: Automated E2E from introduction.
- **Cross-repo consistency**: Same test stack as content-sources-frontend.

## Impact

Start of 944+ playwright-themed commits in maturity/convergence periods. Enabled HMS-5430 introspect test and template filter tests later.

## Review Notes

@jlsherrill questioned directory naming. @dominikvagner and @swadeley reviewed. Merged Feb 2025.

> **Notable quote** (@jlsherrill): "why use _playwright-tests instead of playwright-tests or ./tests/playwright/ ?"

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
