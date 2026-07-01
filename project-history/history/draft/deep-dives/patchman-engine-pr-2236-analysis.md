---
repo: patchman-engine
pr_number: 2236
period: convergence
source: history/repos/patchman-engine/pull-requests/merged/
---

# Deep Dive: patchman-engine PR #2236

## Problem / Motivation

The manager pod ran database migrations in an init container. During rolling deploys with multiple replicas, several pods could simultaneously attempt migration and compete for `pg_advisory_lock(123)`. ClowdApp cannot set maxSurge on manager (public web service), making concurrent migrators inevitable — and hard to debug when two pods both try to migrate.

## Solution Approach

@TenSt moved migration to a dedicated ClowdApp **Job** (single runner), replaced manager's `db-migration` init with lightweight `check-for-db` (matching listener/evaluator pattern), and added retry logic in `entrypoint.sh`.

## Key Design Decisions

- **Job vs init container**: One migration runner per deploy; app pods block until schema ready.
- **Fail-safe rollout**: New pods fail init if migration fails; old pods keep serving.
- **Retry on migrate failure**: Configurable resilience for transient DB errors.

## Impact

Deployment reliability improvement for all future schema migrations. Paired with runbook documentation (PR #2244). Addresses a class of production incidents during large migrations.

## Review Notes

@dominikvagner: "looks good to me! ✅". Sourcery suggested making `MIGRATION_MAX_RETRIES` configurable via ClowdApp parameter.

> **Notable quote** (@dominikvagner): "looks good to me! ✅"

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
