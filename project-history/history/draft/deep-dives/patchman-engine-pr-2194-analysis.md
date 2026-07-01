---
repo: patchman-engine
pr_number: 2194
period: convergence
source: history/repos/patchman-engine/pull-requests/merged/
---

# Deep Dive: patchman-engine PR #2194

## Problem / Motivation

Downstream services — including content-sources template workflows — needed to react when patch advisories change. Previously, advisory updates were internal to patchman-engine with no event contract for other services (RHINENG-26118).

## Solution Approach

@katarinazaprazna added `patchman.advisory.update` Kafka topic configuration, `AdvisoryUpdateEvent` message contract, and `WriteEvents` batch publisher. Removed `kafkaTopics` from ClowdApp — topics now provisioned externally by Platform team.

## Key Design Decisions

- **External topic provisioning**: Platform team owns Kafka topic lifecycle; app only references config.
- **Typed event contract**: `AdvisoryUpdateEvent` enables content-sources and other consumers to deserialize reliably.
- **Local dev alignment**: docker-compose/scripts updated for local topic testing.

## Impact

Opens advisory lifecycle to event-driven integration. Connects to template-advisory relationship work (PR #2249) and content-sources template update events (PR #1537). Cross-product convergence milestone.

## Review Notes

@MichaelMraka and @TenSt reviewed. Unit tests for serialization and Kafka writing added.

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
