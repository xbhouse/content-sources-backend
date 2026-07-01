---
type: pull_request
number: 475
title: "Refs 3053: Move to new bucket name"
state: merged
author: jlsherrill
created: 2023-11-17T19:05:20Z
updated: 2023-11-17T19:30:50Z
closed: 2023-11-17T19:30:44Z
merged: 2023-11-17T19:30:44Z
base_branch: main
head_branch: 3050_2
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/475
---

# Pull Request #475: Refs 3053: Move to new bucket name

**Author**: @jlsherrill
**Created**: November 17, 2023 at 07:05 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3050_2`

## Description

## Summary
Change to a new bucket name for our clowder provided resource.  The previous name was reused across buckets resulting in both buckets having the same 'name' internally.

## Testing steps
IQE tests Pass

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 17, 2023 at 07:06 PM UTC

https://issues.redhat.com/browse/HMS-3053

### Comment by @jlsherrill on November 17, 2023 at 07:06 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on November 17, 2023 at 07:30 PM UTC

task doesn't really require QE testing, merging! 

---

## Reviews

### Review by @xbhouse - Approved on November 17, 2023 at 07:09 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/475*
