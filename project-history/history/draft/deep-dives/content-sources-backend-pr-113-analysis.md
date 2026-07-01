---
repo: content-sources-backend
pr_number: 113
period: expansion
source: history/repos/content-sources-backend/pull-requests/merged/
---

# Deep Dive: content-sources-backend PR #113

## Problem / Motivation

Manual/CLI introspection (PR #55) did not scale. Repositories needed event-driven introspection when created or updated via the platform — HMSCONTENT-151.

## Solution Approach

@avisiedo added `kafka-introspect` command consuming Kafka messages to trigger introspection. Schema-generated Go types (`make build` regenerates `.gen.go`), kafkacat test tooling, validation across three ephemeral Kafka pool types (default, managed-kafka, real-managed-kafka).

## Key Design Decisions

- **Schema codegen**: Kafka message schemas generate Go types — contract-first approach.
- **Ephemeral validation matrix**: Tested against all Kafka pool configurations before merge.
- **Producer on bulk_create**: @jlsherrill requested kafka message production when repos created/updated via API, not just consumed.

## Impact

Event-driven architecture foundation. Enabled platform integration — other services can trigger introspection. 13106 packages discovered in first ephemeral test run.

## Review Notes

@jlsherrill on missing header: "We should probably not make it required" for X-Rh-Insights-Request-Id. @avisiedo converted TODO comments to Jira cards before merge.

> **Notable quote** (@avisiedo): "IntrospectionUrl returned 13106 new packages"

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
