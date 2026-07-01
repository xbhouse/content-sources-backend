---
type: pull_request
number: 299
title: "Fixes 1470: switch nightly introspect to tasking system"
state: merged
author: rverdile
created: 2023-05-30T20:42:45Z
updated: 2023-06-05T14:12:40Z
closed: 2023-06-05T14:12:36Z
merged: 2023-06-05T14:12:36Z
base_branch: main
head_branch: nightly-tasking
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/299
---

# Pull Request #299: Fixes 1470: switch nightly introspect to tasking system

**Author**: @rverdile
**Created**: May 30, 2023 at 08:42 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `nightly-tasking`

## Description

## Summary
- Renames the `introspect-all` command to `nightly-jobs`
- If new tasking system has been enabled, adds each repository as introspect job to queue, instead of directly invoking introspect for all repositories

## Testing steps
1. Set `NEW_TASKING_SYSTEM` env var or `config.yaml` to true.
2. `go run cmd/external_repos/main.go nightly-jobs` will show `[Enqueued Task]...` for each repository.
3. Main server will need to be running for tasks to process.
4. Setting NEW_TASKING_SYSTEM to false should _not_ use the new tasking system, and instead directly invoke introspection for all repositories.

---

## Discussion

### Comment by @jlsherrill on May 30, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-1470

---

## Reviews

### Review by @jlsherrill - Commented on May 31, 2023 at 08:32 PM UTC

### Review by @jlsherrill - Commented on May 31, 2023 at 08:32 PM UTC

### Review by @rverdile - Commented on May 31, 2023 at 09:32 PM UTC

### Review by @rverdile - Commented on May 31, 2023 at 09:32 PM UTC

### Review by @rverdile - Commented on May 31, 2023 at 09:33 PM UTC

### Review by @jlsherrill - Commented on June 01, 2023 at 06:56 PM UTC

### Review by @jlsherrill - Commented on June 01, 2023 at 06:57 PM UTC

### Review by @jlsherrill - Approved on June 01, 2023 at 07:26 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/299*
