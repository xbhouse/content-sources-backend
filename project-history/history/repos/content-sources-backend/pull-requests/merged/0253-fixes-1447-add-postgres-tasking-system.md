---
type: pull_request
number: 253
title: "Fixes 1447: add postgres tasking system"
state: merged
author: rverdile
created: 2023-04-19T16:42:12Z
updated: 2023-05-16T17:14:21Z
closed: 2023-05-16T17:14:21Z
merged: 2023-05-16T17:14:21Z
base_branch: main
head_branch: tasking-system
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/253
---

# Pull Request #253: Fixes 1447: add postgres tasking system

**Author**: @rverdile
**Created**: April 19, 2023 at 04:42 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `tasking-system`

## Description

## Summary
- Adds a postgres tasking system to for handling introspect events.
- Adds the tasks, tasks/queue, tasks/client, and tasks/worker packages.

tasks/queue
- A queue used by client/worker to enqueue and handle the tasks

tasks/client
- `TaskClient` is an interface to enqueue tasks. Right now it's really just a wrapper around `queue.Enqueue(task)`. This is used in the handler layer to enqueue introspect tasks for the `create`, `update`, and `introspect` endpoints.

tasks/worker
- `worker_pool.go` has an interface `TaskWorkerPool` which is consumed by `cmd/content-sources/main.go` to configure and start the workers.
- `worker.go` is the individual worker started/stopped by WorkerPool. This is only meant to be used by WorkerPool, not directly. I separated them to potentially allow more control over the individual workers.


## Testing steps
1. Still working on a few things
2. If you add `new_tasking_system: True` to your config.yaml, it will use this for introspect instead of kafka.
3. Should work right now for all endpoints that trigger introspect (`create`, `update`, `introspect`).
4. For now, if you want to test the tasking system more thoroughly I suggest modifying the introspect handler to do other things, then trigger it via the introspect endpoint.


Todo:
- [x] Finish cleaning up repository handler
- [x] Add links to the code in the tasking system architecture doc (I'll try this after merge)
- [x] Add more tests
- [x] Refactor pgqueue db connection such that a transaction can be used to write tests.

---

## Discussion

### Comment by @jlsherrill on April 19, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-1447

### Comment by @jlsherrill on April 20, 2023 at 06:24 PM UTC

could you add a README in the docs directory detailing the arch of the tasking system and explaining its features? 

### Comment by @rverdile on May 01, 2023 at 07:07 PM UTC

I was trying to add integration/queue tests, but it's actually a lot to think about because the database connection works differently for `pgqueue.go`, and so decent change probably has to be made there to pass a transaction through to it.

I've pushed up everything else, which should be good to go.

### Comment by @jlsherrill on May 10, 2023 at 07:16 PM UTC

Overall looks good, thoughts on incorporating these metrics into the tasking system?  https://github.com/content-services/content-sources-backend/blob/main/pkg/instrumentation/metrics.go#L105-L117

to capture latency (time from queued and then picked up) and tasking failures (failures with processing, not with external factors like failed repo syncs)

### Comment by @rverdile on May 10, 2023 at 07:20 PM UTC

Those should be there already. Are they working?

Latency: https://github.com/content-services/content-sources-backend/blob/68e92b7975ae5518c5dc237b9b2097e75ce7c599/pkg/tasks/worker/worker.go#L81

Result: https://github.com/content-services/content-sources-backend/blob/68e92b7975ae5518c5dc237b9b2097e75ce7c599/pkg/tasks/worker/worker.go#L107-L110

### Comment by @jlsherrill on May 10, 2023 at 07:20 PM UTC

ohhh!  nice work, i overlooked it, let me check!

### Comment by @jlsherrill on May 10, 2023 at 07:21 PM UTC

yes, working great!!

### Comment by @jlsherrill on May 11, 2023 at 02:03 PM UTC

/retest

### Comment by @jlsherrill on May 11, 2023 at 02:04 PM UTC

[test]

### Comment by @swadeley on May 16, 2023 at 04:07 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Commented on April 20, 2023 at 05:22 PM UTC

### Review by @jlsherrill - Commented on April 20, 2023 at 05:25 PM UTC

### Review by @jlsherrill - Commented on April 20, 2023 at 05:53 PM UTC

### Review by @rverdile - Commented on April 20, 2023 at 06:09 PM UTC

### Review by @rverdile - Commented on April 21, 2023 at 09:09 PM UTC

Leaving some comments to give context to some details and draw attention to cases we might prefer to handle differently.

### Review by @jlsherrill - Commented on April 25, 2023 at 08:11 PM UTC

### Review by @jlsherrill - Commented on May 03, 2023 at 09:27 PM UTC

### Review by @jlsherrill - Commented on May 03, 2023 at 09:28 PM UTC

### Review by @jlsherrill - Commented on May 03, 2023 at 09:39 PM UTC

### Review by @jlsherrill - Commented on May 03, 2023 at 09:48 PM UTC

### Review by @rverdile - Commented on May 09, 2023 at 08:39 PM UTC

### Review by @rverdile - Commented on May 09, 2023 at 08:42 PM UTC

### Review by @rverdile - Commented on May 09, 2023 at 08:45 PM UTC

### Review by @jlsherrill - Commented on May 10, 2023 at 04:10 PM UTC

### Review by @jlsherrill - Commented on May 10, 2023 at 04:25 PM UTC

### Review by @jlsherrill - Approved on May 10, 2023 at 07:24 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/253*
