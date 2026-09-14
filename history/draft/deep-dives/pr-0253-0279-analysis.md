# backend#253 / #279 — Tasking system and snapshot integration

Authors: @rverdile (#253), @jlsherrill (#279)  
Merged: 2023-05-16 / 2023-06-03  
URLs: https://github.com/content-services/content-sources-backend/pull/253 · https://github.com/content-services/content-sources-backend/pull/279  
Period: pulp-integration

## Problem

#248’s snapshot pipeline could only run as a CLI command. Introspection still went through Kafka. A Pulp sync can take minutes and the worker can die; without a durable queue and stored Pulp task hrefs, snapshots could not be a user-facing feature.

## Approach

#253 (@rverdile) added a Postgres tasking system: `tasks/queue`, `tasks/client`, `tasks/worker`. Handlers enqueue; a worker pool dequeues. Behind `new_tasking_system: true` it replaced Kafka for introspect on create/update/introspect. Justin asked for an architecture README and for queue-latency / processing-failure metrics; Ryan had already instrumented both.

#279 (@jlsherrill) wired snapshotting to that queue. `POST /repositories` with `snapshot=true` enqueues a snapshot task. The payload stores Pulp sync/publication/distribution hrefs so a killed worker can resume. QE could not yet observe snapshot completion from the API — @swadeley updated IQE for the bool, but monitoring was still “look at `snapshots.distribution_path`.” Justin disabled Pulp in stage/prod before merge so the flag could land dark.

## Design decisions

- Postgres, not Kafka, for work that must resume. Kafka inspection was deleted in #390 (Sep 2023). The task table is the source of truth (heartbeat, requeue, cancel — later #375, #386).
- Resume is a payload problem. #248 left Pulp task IDs unsaved; #279 stores them so “kill -9 the worker” still produces a snapshot.
- Ship dark. Snapshotting stayed feature-flagged until nightly sync (#387) and the list/trigger APIs (#308, #458) existed.

## Impact

This pair is the execution model the rest of the project uses: every introspect, snapshot, template update, upload conversion, and Lightwell import is a row in `tasks`. Admin task UI (#316) and cancel-in-progress-snapshot (#386) are features of this queue, not of Pulp.

## Quote

> “When creating a repository, if you pass snapshot=true, a task is spawned to snapshot the repository. This snapshot task should save information within the payload such as sync task href, publication task href, etc..., such that if the worker dies, the task should be able to be resumed.” — @jlsherrill, #279 summary
