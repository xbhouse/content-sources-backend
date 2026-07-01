---
type: pull_request
number: 322
title: "Fixes 1882: only refresh running tasks"
state: merged
author: rverdile
created: 2023-06-28T14:58:34Z
updated: 2023-06-28T18:16:52Z
closed: 2023-06-28T18:16:49Z
merged: 2023-06-28T18:16:49Z
base_branch: main
head_branch: heartbeat-404
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/322
---

# Pull Request #322: Fixes 1882: only refresh running tasks

**Author**: @rverdile
**Created**: June 28, 2023 at 02:58 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `heartbeat-404`

## Description

## Summary
We've been seeing an error message that suggests there is a scenario where a worker will attempt to refresh the heartbeat of the previous (but no longer in-progress) task. This makes it so now only the running task's heartbeat will be refreshed.

## Testing steps
1. Use the tasking system by creating/introspecting/deleting repositories. Everything should be working as normal.
2. Try exiting the server with Ctrl+C at different points while tasks are running (or finished). You should not see any warnings or errors.

---

## Discussion

### Comment by @jlsherrill on June 28, 2023 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-1882

---

## Reviews

### Review by @jlsherrill - Approved on June 28, 2023 at 05:53 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/322*
