---
type: pull_request
number: 1134
title: "HMS-8627: start snapshot for repo once it's enabled"
state: merged
author: xbhouse
created: 2025-06-27T14:37:28Z
updated: 2025-07-02T15:12:49Z
closed: 2025-07-02T15:12:49Z
merged: 2025-07-02T15:12:49Z
base_branch: main
head_branch: 8627
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1134
---

# Pull Request #1134: HMS-8627: start snapshot for repo once it's enabled

**Author**: @xbhouse
**Created**: June 27, 2025 at 02:37 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `8627`

## Description

## Summary

- Fixes issue where a snapshot task was not triggered immediately when updating a repo from introspection to snapshotting
- Requires this [test change](https://github.com/content-services/content-sources-frontend/pull/556) for playwright snapshot tests to pass

## Testing steps

1. Create a repo with snapshot set to false
2. Update the repo to set snapshot to true. Introspect and snapshot tasks for that repo should start immediately

---

## Discussion

### Comment by @jlsherrill on June 27, 2025 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-8627

---

## Reviews

### Review by @rverdile - Approved on July 02, 2025 at 03:05 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1134*
