---
repo: content-sources-backend
pr_number: 142
period: expansion
source: history/repos/content-sources-backend/pull-requests/merged/
---

# Deep Dive: content-sources-backend PR #142

## Problem / Motivation

Introspection parsing logic lived in `dao/external_resources.go` — hard to test independently, reuse, or version. HMSCONTENT-200 extracted it into a standalone library.

## Solution Approach

@rverdile moved parsing to `github.com/content-services/yummy` (companion PR yummy#7). Backend refactored to import yummy as Go module. Architectural feedback welcomed in review.

## Key Design Decisions

- **Shared library pattern**: First of several content-services libs (yummy, zest, tang, caliri).
- **Module replace for dev**: Local path replace in go.mod for parallel development.
- **Companion PRs**: Backend + yummy changes coordinated.

## Impact

@jlsherrill: "Overall, a big +1, i don't have any major comments." Established micro-library architecture for content-services. @xbhouse later became primary yummy maintainer (7 PRs).

## Review Notes

Merged Dec 2022 after month-long review. Foundation for comps/groups parsing work in 2023.

> **Notable quote** (@jlsherrill): "Overall, a big +1, i don't have any major comments"

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
