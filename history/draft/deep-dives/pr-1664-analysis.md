# backend#1664 — LWLP-626: Lightwell JFrog bridge

Author: @etsien (with substantial follow-up from @jlsherrill)  
Merged: 2026-08-27  
URL: https://github.com/content-services/content-sources-backend/pull/1664  
Period: lightwell  
Follow-up: #1695 (LWLP-946 JFrog bridge updates)

## Problem

Lightwell package versions and builds live only in Pulp. JFrog still needs to *hear about* advisories. Lightwell emits those notifications. The first design consumed the general Insights notifications queue — millions of messages — and would fan out one message per org. That is the wrong bus for an advisory bridge. This is not the Insights partner-repos feature.

## Approach

@etsien added `pkg/jfrog_bridge` (unit tests, lint, `jfrog_bridge.enabled`). @jlsherrill pushed back on the queue: too heavy, and “we will send multiple messages (one for each org), and may not send any at all.” He opened a platform-mq PR for a dedicated topic. The topic name changed from `platform.lightwell.advisory_created` (underscores illegal) to `platform.lightwell.advisory-created`. Justin also made Lightwell notification code publish to that topic once, added an admin test endpoint, and fixed lint.

CI stuck in `PgListener.WaitForNotification` (LISTEN/NOTIFY timeout); tests were rerun after that diagnosis.

## Design decisions

- Dedicated Kafka topic, not the platform notifications firehose. Cost and semantics: one advisory event, not per-org Insights notifications.
- Bridge is optional (`jfrog_bridge.enabled`). Package list/detail keep working with the flag off — they never queried JFrog.
- Admin can inject a test payload. Needed because the JFrog endpoint is not in every ephemeral.

## Impact

First Lightwell integration that is *not* a Pulp content path. Combined with OSV sync (#1609), Lightwell is a Pulp-backed catalog plus an advisory product that happens to share content-sources’ API process. Network API (#1666) continued the advisory/search surface. Pulp stayed the version store.

## Quote

> “consuming from the general notifications queue is too heavy, iirc there are millions of messages being processed through regularly. … A solution would be to use a new topic, and modify the lightwell notification code to also send to this topic just once.” — @jlsherrill
