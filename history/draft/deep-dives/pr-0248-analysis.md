# backend#248 — Fixes 1468: add repo snapshot task

Author: @jlsherrill  
Merged: 2023-05-03  
URL: https://github.com/content-services/content-sources-backend/pull/248  
Period: pulp-integration

## Problem

Until May 2023, content-sources only *introspected* repos: parse `repomd`, store package counts and status. Customers could not pin a system to a frozen copy of a repo. That required mirroring content in Pulp (remote → repository → sync → publication → distribution). There was no first-class snapshot object and no way to run that pipeline as a recoverable job.

## Approach

@jlsherrill added a snapshot task and a snapshot DAO. The first version was a manual command (`go run cmd/external-repos/main.go pulp-create <url> <org>`), not yet wired to the API. It could resume after interruption after remote or repo creation, but not mid-sync, mid-publication, or mid-distribution. Those gaps were explicitly deferred to tasking integration (#279).

The PR also introduced a DAO registry and moved DAO mocks into the `dao` package so tests and tasks could share the same construction path.

## Design decisions

- Pulp is the content store; Postgres stores snapshot metadata (distribution path, later task hrefs). Users never talk to Pulp directly.
- Incomplete resume is acceptable for v1. @rverdile interrupted `pulp-create` mid-flight, reran it, and got “successfully created” with no content at `/pulp/content`. @jlsherrill’s reply: save Pulp task IDs on the content-sources task and, on retry, wait for or resume the in-flight Pulp sync. That landed in #279.
- Feature-flagged. Snapshots were not on in stage/prod when this merged (`no-qe-needed`).

## Impact

This is the architectural turn. Templates, RH repo snapshotting, upload repos, nightly sync, and Lightwell content counts all assume this pipeline exists. The snapshot row is what later PRs associate to templates (#836) and expire (#901).

## Quote

> “What we're gonna need to do is save the task ids on the task itself somehow (to be done during integration) and then as part of the task check that the sync task/publish task is completed, and if not retry it.” — @jlsherrill, 2023-04-26
