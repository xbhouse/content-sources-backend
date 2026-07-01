---
repo: patchman-ui
pr_number: 1164
period: templates-era
source: history/repos/patchman-ui/pull-requests/merged/
---

# Deep Dive: patchman-ui PR #1164

## Problem / Motivation

Enzyme, the React test library used since 2019, was deprecated and incompatible with React 18 and PatternFly 5 upgrades. RHINENG-7975 blocked frontend modernization until tests migrated to React Testing Library (RTL).

## Solution Approach

@mkholjuraev rewrote Enzyme tests to RTL across the codebase, removed unnecessary Redux selector mocks, and cleaned up test utilities. Inventory table tests were **commented out** temporarily to unblock PF5/React migration — fed-modules testing deferred.

## Key Design Decisions

- **RTL over Enzyme**: Industry-standard replacement; aligns with PatternFly 5 ecosystem.
- **Pragmatic scope cut**: Disable fed-modules inventory tests rather than delay entire migration.
- **Multi-PR effort**: Part of #1132–#1158 enzyme refactor series.

## Impact

Unblocked PatternFly 5 migration (#1372) and later Playwright E2E (#1445). Test infrastructure modernized for 2024+ development. 96% patch coverage on changed files per Codecov.

## Review Notes

Large diff across test files. Released label applied post-merge.

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
