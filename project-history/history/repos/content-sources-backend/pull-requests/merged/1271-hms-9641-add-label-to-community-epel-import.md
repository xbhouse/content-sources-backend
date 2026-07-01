---
type: pull_request
number: 1271
title: "HMS-9641: add label to community epel import"
state: merged
author: rverdile
created: 2025-10-31T15:45:17Z
updated: 2025-11-03T14:49:29Z
closed: 2025-11-03T14:49:22Z
merged: 2025-11-03T14:49:22Z
base_branch: main
head_branch: epel-label
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1271
---

# Pull Request #1271: HMS-9641: add label to community epel import

**Author**: @rverdile
**Created**: October 31, 2025 at 03:45 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `epel-label`

## Description

## Summary
Adds label to the community epel repo import

## Testing steps
1. Without this PR:
2. Run `make repos-import` note the repo config label for community repos
3. Run it again, check the label again, it will have changed
4. With this PR:
5. The label will not change and it will be something like "EPEL_10_Everything_x86_64" without random characters on the end


---

## Discussion

### Comment by @xbhouse on October 31, 2025 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-9641

### Comment by @rverdile on October 31, 2025 at 08:42 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on October 31, 2025 at 04:24 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1271*
