---
type: pull_request
number: 690
title: "Fixes 3974: add priority to tasks"
state: merged
author: rverdile
created: 2024-06-05T15:38:28Z
updated: 2024-06-10T15:32:19Z
closed: 2024-06-10T15:32:16Z
merged: 2024-06-10T15:32:16Z
base_branch: main
head_branch: add-priority-tasks
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/690
---

# Pull Request #690: Fixes 3974: add priority to tasks

**Author**: @rverdile
**Created**: June 05, 2024 at 03:38 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `add-priority-tasks`

## Description

## Summary
Adds the priority column to tasks, an integer representing the priority of the task. Tasks with a higher integer will be dequeued first. Ready tasks view is sorted by priority then queued_at time i.e. highest priority, then oldest, goes first.

Make the priority of update_template_content tasks 1, so it is always dequeued first.

## Testing steps
1. Run the db migration
2. Change worker_count in your config.yaml to 1
3. Create two EPEL repos, one after the other.
4. While the first one is snapshotting, create or update a template.
5. The order of the task completion should be: snapshot (for the first EPEL repo), update_template_content, snapshot (for the second EPEL repo).

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 05, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-3974

### Comment by @rverdile on June 07, 2024 at 02:59 PM UTC

/retest


### Comment by @rverdile on June 10, 2024 at 01:49 PM UTC

/retest

### Comment by @swadeley on June 10, 2024 at 02:08 PM UTC

Hi

Please rebase to get fix in https://github.com/content-services/content-sources-backend/pull/697

---

## Reviews

### Review by @xbhouse - Approved on June 06, 2024 at 07:10 PM UTC

looks good to me! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/690*
