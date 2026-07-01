---
repo: content-sources-backend
pr_number: 1538
period: convergence
source: history/repos/content-sources-backend/pull-requests/merged/
---

# Deep Dive: content-sources-backend PR #1538

## Problem / Motivation

Content templates needed analytics export matching patchman's Floorist pattern (engine PR #898). HMS-10345 adds template metrics to platform reporting.

## Solution Approach

@jlsherrill added Floorist SQL query with extended release information to clowdapp deployment config — same pattern as patchman-engine analytics jobs.

## Key Design Decisions

- **Mirror patchman pattern**: Proven Floorist approach reused for new domain.
- **Extended release info**: Query captures template release metadata for reporting.

## Impact

Parity between Patch and Content Sources in platform analytics. @TenSt approved. @xbhouse linked HMS-10345.

## Review Notes

Small deployment config PR. Fast merge (Jun 2026).

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
