---
type: pull_request
number: 556
title: "Fixes 3517: job to retry existing failed tasks"
state: merged
author: rverdile
created: 2024-02-01T18:43:45Z
updated: 2024-02-05T14:32:50Z
closed: 2024-02-05T14:32:47Z
merged: 2024-02-05T14:32:47Z
base_branch: main
head_branch: run-retry-job
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/556
---

# Pull Request #556: Fixes 3517: job to retry existing failed tasks

**Author**: @rverdile
**Created**: February 01, 2024 at 06:43 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `run-retry-job`

## Description

## Summary
Adds a job to update the current failed "delete-repository-snapshot" tasks so that they get requeued by the new retry system.

## Testing steps
1. In your config, set `tasking.heartbeat: 5s` and `tasking.retry_wait_upper_bound: 10s`.
2. Create a repository with a snapshot.
3. Run `podman stop cs_pulp_api_1` so you can no longer make requests to pulp without error.
4. Try to delete the repository.
5. It will fail and retry 3 times and then stop. 
6. Run `go run cmd/retry_failed_tasks/main.go --force`
7. You should see the failed task updated and retry 3 more times.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 01, 2024 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-3517

---

## Reviews

### Review by @jlsherrill - Approved on February 02, 2024 at 04:11 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/556*
