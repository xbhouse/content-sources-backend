---
type: pull_request
number: 395
title: "Fixes 2631: handle bucket hostname with proto"
state: merged
author: jlsherrill
created: 2023-09-20T15:25:26Z
updated: 2023-09-20T17:23:02Z
closed: 2023-09-20T17:22:51Z
merged: 2023-09-20T17:22:51Z
base_branch: main
head_branch: 2631_2
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/395
---

# Pull Request #395: Fixes 2631: handle bucket hostname with proto

**Author**: @jlsherrill
**Created**: September 20, 2023 at 03:25 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2631_2`

## Description

## Summary
Sometimes the objectstore  hostname has a protocol/scheme, sometimes it doesn't.  Handle both cases (with a test!)

## Testing steps
snapshotting works in ephemeral, we'll test stage after merge

---

## Discussion

### Comment by @jlsherrill on September 20, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-2631

### Comment by @jlsherrill on September 20, 2023 at 05:23 PM UTC

merged with blessing from @swadeley 

---

## Reviews

### Review by @Andrewgdewar - Approved on September 20, 2023 at 05:12 PM UTC

Confirmed working in ephemeral, all tests/checks passing.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/395*
