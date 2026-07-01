---
type: pull_request
number: 671
title: "Fixes 4160: fix incorrect red hat repo label"
state: merged
author: rverdile
created: 2024-05-16T14:01:27Z
updated: 2024-05-16T20:31:20Z
closed: 2024-05-16T14:10:21Z
merged: 2024-05-16T14:10:21Z
base_branch: main
head_branch: rh-dup-label
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/671
---

# Pull Request #671: Fixes 4160: fix incorrect red hat repo label

**Author**: @rverdile
**Created**: May 16, 2024 at 02:01 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `rh-dup-label`

## Description

## Summary
Duplicate/incorrect label was causing `make repos-import` to fail

## Testing steps
1. Run `make repos-import`
2. Should not fail

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 16, 2024 at 08:31 PM UTC

https://issues.redhat.com/browse/HMS-4160

---

## Reviews

### Review by @swadeley - Approved on May 16, 2024 at 02:02 PM UTC

ACK

### Review by @jlsherrill - Approved on May 16, 2024 at 02:02 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/671*
