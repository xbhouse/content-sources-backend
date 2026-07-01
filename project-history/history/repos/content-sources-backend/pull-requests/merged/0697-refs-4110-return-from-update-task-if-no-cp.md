---
type: pull_request
number: 697
title: "Refs 4110: return from update task if no cp"
state: merged
author: jlsherrill
created: 2024-06-10T12:59:39Z
updated: 2024-06-10T15:02:44Z
closed: 2024-06-10T14:03:11Z
merged: 2024-06-10T14:03:11Z
base_branch: main
head_branch: 4110_bz
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/697
---

# Pull Request #697: Refs 4110: return from update task if no cp

**Author**: @jlsherrill
**Created**: June 10, 2024 at 12:59 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4110_bz`

## Description

## Summary
Return from the update task if candlepin is not configured

## Testing steps

In ephemeral try to update a repository and check the status of the task, it should be successful

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 10, 2024 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-4110

### Comment by @jlsherrill on June 10, 2024 at 01:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @swadeley on June 10, 2024 at 02:02 PM UTC

It works, thanks for the fix

---

## Reviews

### Review by @swadeley - Approved on June 10, 2024 at 02:03 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/697*
