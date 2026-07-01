---
repo: content-sources-backend
pr_number: 1537
period: convergence
source: history/repos/content-sources-backend/pull-requests/merged/
---

# Deep Dive: content-sources-backend PR #1537

## Problem / Motivation

When template content changes (new advisories in repos), downstream consumers — especially patchman — need asynchronous notification. HMS-9952 requires a job publishing template update events to Kafka.

## Solution Approach

@swadeley added `SendTemplateUpdateEvents` background job publishing to `platform.content-sources.template` topic. Coordinated with @xbhouse's advisory ID endpoints (#1536) and message structure changes (#1533).

## Key Design Decisions

- **Job vs inline publish**: Batch job for initial/backfill event population.
- **Coordination with #1536**: @rverdile asked whether enqueueing update-template-content jobs would suffice — team clarified backfill vs ongoing event semantics.
- **Kafka topic contract**: New platform.content-sources.template topic.

## Impact

Event-driven template lifecycle — patchman-engine #2194 advisory events and #2249 template-advisory relationships connect here. Cross-product convergence capstone.

## Review Notes

@rverdile: "I'm a little confused about what the ticket is asking for" — team clarified backfill vs ongoing updates in thread.

> **Notable quote** (@rverdile): "I'm a little confused about what the ticket is asking for. Is this understanding correct?"

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
