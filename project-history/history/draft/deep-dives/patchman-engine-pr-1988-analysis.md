---
repo: patchman-engine
pr_number: 1988
period: convergence
source: history/repos/patchman-engine/pull-requests/merged/
---

# Deep Dive: patchman-engine PR #1988

## Problem / Motivation

Cyndi, the platform inventory sync pipeline, changed its event structure. The monolithic `system_platform` table mixed host inventory metadata with patch evaluation state, making it increasingly difficult to align with upstream inventory changes (RHINENG-21214).

## Solution Approach

@Dugowitch split storage into `system_inventory` (host metadata from Cyndi) and `system_patch` (patch/advisory state), connected via a backward-compatible `system_platform` **view** with INSTEAD OF triggers so existing application queries continued working during transition.

## Key Design Decisions

- **View compatibility layer**: Application code keeps querying `system_platform` while underlying tables split — reduces blast radius.
- **Trigger-based routing**: INSTEAD OF triggers on the view translate inserts/updates to correct underlying tables.
- **Long review cycle**: Dec 2025 → Feb 2026 merge reflects migration complexity and extensive @MichaelMraka review.

## Impact

Major inventory sync architecture change. Prerequisite for reliable inventory group and workspace features in patchman-ui. Sourcery review guide describes split as "host metadata vs patch state" separation.

## Review Notes

Multiple review rounds from @MichaelMraka. Sourcery noted patch coverage at 6% — migration-heavy PR with limited unit test surface.

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
