---
type: pull_request
number: 1362
title: "HMS-9776: add logs for failed snapshot increment count"
state: merged
author: TenSt
created: 2026-01-15T12:18:25Z
updated: 2026-01-16T15:25:43Z
closed: 2026-01-16T15:25:43Z
merged: 2026-01-16T15:25:43Z
base_branch: main
head_branch: stepan/HMS-9776-use-uudify-string-func-for-snap-increment-and-reset
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1362
---

# Pull Request #1362: HMS-9776: add logs for failed snapshot increment count

**Author**: @TenSt
**Created**: January 15, 2026 at 12:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `stepan/HMS-9776-use-uudify-string-func-for-snap-increment-and-reset`

## Description

## Summary
This PR adds logs for failed `InternalOnly_IncrementFailedSnapshotCount` runs to investigate why we have empty UUID in the query.

## Testing steps



---

## Discussion

### Comment by @xbhouse on January 15, 2026 at 12:30 PM UTC

https://issues.redhat.com/browse/HMS-9776

### Comment by @TenSt on January 16, 2026 at 09:53 AM UTC

> Did you get a chance to look into why this function was being called with an empty uuid?

Hey, Ryan! Thanks for mentioning this, I've missed this part. The investigation is in progress, I'll provide details once I'll find something.

### Comment by @TenSt on January 16, 2026 at 12:26 PM UTC

@rverdile I've reverted the fix for this PR and added logs to see the exact error from the `Run` function to investigate more why we have an empty UUID. The fix with using `UuidifyString` we can apply later.

---

## Reviews

### Review by @rverdile - Commented on January 15, 2026 at 06:28 PM UTC

Did you get a chance to look into why this function was being called with an empty uuid? 

### Review by @rverdile - Approved on January 16, 2026 at 03:09 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1362*
