# backend#387 / #407 — Nightly sync and Red Hat repos

Author: @jlsherrill  
Merged: 2023-09-27 / 2023-10-10  
URLs: https://github.com/content-services/content-sources-backend/pull/387 · https://github.com/content-services/content-sources-backend/pull/407  
Period: pulp-integration (#407 is the first week of expansion)

## Problem

#279 snapshotted a repo once, at create. Content drifts. Without a recurring sync, “snapshot” meant “whatever we mirrored on day one.” Separately, Insights users needed Red Hat CDN repos in the same snapshot pipeline as custom URLs — not only BYO content.

## Approach

#387 added snapshotting to nightly jobs: if the snapshot feature is on, enqueue a snapshot for every `snapshot=true` repo not snapshotted in the last ~24 hours (including never-snapshotted). `OPTIONS_ALWAYS_RUN_CRON_TASKS` ignores the time window for debugging. @rverdile asked whether a failed snapshot should retry in 8 hours; the first cut queued only snapshots older than ~20 hours.

#407 loaded an initial Red Hat repo catalog from JSON, one shared GPG key, `org_id = -1`, `origin = red_hat`. `make repos-import` (or ephemeral seed) inserts them. List default hides them unless `origin=red_hat` or the `NewRepositoryFiltering` feature is on. Sync uses the CDN cert. A SQL panic on `repository_configurations` was fixed before merge. #421 briefly reverted RH nightly snapshotting, then it returned.

## Design decisions

- Nightly is “not snapshotted in 24h,” not “sync every repo every night.” Failed snapshots did not get a faster retry in v1 (later: HMS-5225).
- RH repos are shared rows (`org_id=-1`), not copied per customer. Origin filtering and later feature guards (#943) control visibility; Pulp still snapshots them once.
- Feature flags gate list behavior. Default list ≠ “all content the org can use.” That surprised @rverdile during review.

## Impact

Nightly sync is why templates can say “latest” and mean something. RH repos in the same table as custom repos is why `origin` exists and why Lightwell later needed its own import path instead of pretending to be `origin=red_hat`.

## Quote

> “nothing should be returned for a 'new org'. do you have the NewRepositoryFiltering feature turned on? because that will affect the behavior.” — @jlsherrill, #407
