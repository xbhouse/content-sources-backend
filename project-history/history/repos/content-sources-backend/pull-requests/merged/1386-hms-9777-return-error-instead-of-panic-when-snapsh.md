---
type: pull_request
number: 1386
title: "HMS-9777: return error instead of panic when snapshot not found"
state: merged
author: rverdile
created: 2026-02-03T18:35:39Z
updated: 2026-02-04T15:12:00Z
closed: 2026-02-04T15:11:55Z
merged: 2026-02-04T15:11:55Z
base_branch: main
head_branch: panic
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1386
---

# Pull Request #1386: HMS-9777: return error instead of panic when snapshot not found

**Author**: @rverdile
**Created**: February 03, 2026 at 06:35 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `panic`

## Description

## Summary
There is a very infrequent error coming through where it seems there may be a template using a repository with no snapshots. This will make that task to error instead of panic.

## Testing steps
You can't create a template with missing snapshots, so this is hard to reproduce. Tests should pass.


---

## Discussion

### Comment by @xbhouse on February 03, 2026 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-9777

---

## Reviews

### Review by @xbhouse - Approved on February 03, 2026 at 09:50 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1386*
