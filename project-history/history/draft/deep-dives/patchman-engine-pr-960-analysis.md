---
repo: patchman-engine
pr_number: 960
period: production
source: history/repos/patchman-engine/pull-requests/merged/
---

# Deep Dive: patchman-engine PR #960

## Problem / Motivation

Red Hat Insights migrated from `account_number` to `org_id` as the canonical tenant identifier. Patchman-engine stored accounts keyed by account number, but inventory events and identity headers were transitioning to org-scoped IDs. Without migration, data isolation risk and inconsistent multi-tenancy would block platform alignment (SPM-1482).

## Solution Approach

@psegedy added an `org_id` column to accounts, dual-write logic for new accounts, backfill for existing records, and special handling when Kafka messages still carried only `account_number` (inventory had not yet sent org_id in all event types). A follow-up org_id populator CJI (SPM-1550) completed the backfill in production.

## Key Design Decisions

- **Dual identifier period**: Continue creating accounts from Kafka `account_number` while accepting org_id from identity headers — pragmatic bridge during platform transition.
- **Cache safety**: Explicitly avoided `findAccount` in listeners to prevent cross-tenant cache collisions when account_number of one user could numerically match org_id of another.
- **Non-destructive org_id**: Never overwrite org_id when multiple null-account records exist and an empty account_number header arrives.

## Impact

Foundational identity migration enabling all subsequent org-scoped APIs. Blocked data leaks during Insights' platform-wide identity overhaul. Every content-sources endpoint later reused the same org_id patterns from platform-go-middlewares.

## Review Notes

@MichaelMraka requested rebase before merge. Author documented todo items inline and resolved cache leak concerns in PR description.

> **Notable quote** (@MichaelMraka): "Please rebase it on top of current master."

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
