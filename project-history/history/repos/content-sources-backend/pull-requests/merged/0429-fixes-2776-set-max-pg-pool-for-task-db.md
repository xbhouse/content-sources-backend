---
type: pull_request
number: 429
title: "Fixes 2776: set max pg pool for task db"
state: merged
author: jlsherrill
created: 2023-10-12T20:30:51Z
updated: 2023-10-12T23:36:47Z
closed: 2023-10-12T23:35:42Z
merged: 2023-10-12T23:35:42Z
base_branch: main
head_branch: 2776_fix
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/429
---

# Pull Request #429: Fixes 2776: set max pg pool for task db

**Author**: @jlsherrill
**Created**: October 12, 2023 at 08:30 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2776_fix`

## Description

## Summary

Previously the db pool max limit was defaulting to 8 in ephemeral.  This wasn't enough for 2 tasks to run at the same time, due to needing a connection for:
* heartbeat update
* cancel listener
* updating payloads
* finishing the task

If enough of these things happened at once, it could exceede 8.

## Testing steps

ci passes

---

## Discussion

### Comment by @jlsherrill on October 12, 2023 at 08:31 PM UTC

https://issues.redhat.com/browse/HMS-2776

---

## Reviews

### Review by @rverdile - Approved on October 12, 2023 at 08:32 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/429*
