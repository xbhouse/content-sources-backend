---
repo: patchman-engine
pr_number: 898
period: production
source: history/repos/patchman-engine/pull-requests/merged/
---

# Deep Dive: patchman-engine PR #898

## Problem / Motivation

By early 2022, Patchman had rich compliance data in PostgreSQL but no standardized way to export it for platform analytics. Red Hat Insights uses Floorist (DataHub) to run scheduled SQL queries and upload results to S3 for downstream reporting. Without Floorist configuration in the ClowdApp deployment, patch compliance metrics were invisible to the broader analytics pipeline.

## Solution Approach

@josef-hak added Floorist (DataHub) job definitions to `deploy/clowdapp.yaml`, defining SQL export queries that run on a schedule and publish patch compliance aggregates. This was a deployment-only change — no application code paths changed — but it established a pattern reused years later for template metrics on both patchman-engine and content-sources-backend.

## Key Design Decisions

- **SQL-in-manifest**: Queries live in the ClowdApp YAML rather than application code, keeping analytics concerns separate from the evaluation engine.
- **Additive deployment**: Zero impact on running services; Floorist jobs are independent CronJobs.
- **Platform alignment**: Uses the same DataHub mechanism as other Insights services.

## Impact

Floorist became the standard analytics export path for Patchman. Later PRs (#1654, content-sources #1538) optimized and extended these queries as template and advisory features grew. This PR marks the moment Patch data joined the platform-wide analytics ecosystem.

## Review Notes

Minimal discussion — checklist-only PR description, merged within one day. @MichaelMraka approved. No alternatives documented; Floorist was the established platform pattern.

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
