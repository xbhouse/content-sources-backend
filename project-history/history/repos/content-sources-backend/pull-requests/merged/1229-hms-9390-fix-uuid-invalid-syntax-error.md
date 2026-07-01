---
type: pull_request
number: 1229
title: "HMS-9390: fix UUID invalid syntax error"
state: merged
author: TenSt
created: 2025-10-08T10:48:51Z
updated: 2025-10-09T12:27:21Z
closed: 2025-10-09T12:27:21Z
merged: 2025-10-09T12:27:21Z
base_branch: main
head_branch: stepan/HMS-9390-fix-uuid-invalid-syntax-error
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1229
---

# Pull Request #1229: HMS-9390: fix UUID invalid syntax error

**Author**: @TenSt
**Created**: October 08, 2025 at 10:48 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `stepan/HMS-9390-fix-uuid-invalid-syntax-error`

## Description

## Summary

Fix bug: https://issues.redhat.com/browse/HMS-9390

In this PR:
- added usage of `UuidifyString` in `checkForValidRepoUuidsUrls` and `checkForValidSnapshotUuids`
- added unit tests for `UuidifyString`, `checkForValidRepoUuidsUrls` and `checkForValidSnapshotUuids`
- refactored `createRepository` and `createSnapshot` functions for tests setup



---

## Reviews

### Review by @dominikvagner - Commented on October 08, 2025 at 11:02 AM UTC

looks good overall 👍🏼🚀 nice job! 
hopefully this lowers the number of glitchtip reports 😄

have two comments about git stuff and one about the `t.Run` stuff 😅 

### Review by @TenSt - Commented on October 08, 2025 at 01:34 PM UTC

### Review by @TenSt - Commented on October 08, 2025 at 01:36 PM UTC

### Review by @dominikvagner - Commented on October 08, 2025 at 02:31 PM UTC

### Review by @dominikvagner - Commented on October 08, 2025 at 02:44 PM UTC

### Review by @xbhouse - Commented on October 08, 2025 at 04:38 PM UTC

### Review by @TenSt - Commented on October 09, 2025 at 08:00 AM UTC

### Review by @TenSt - Commented on October 09, 2025 at 09:43 AM UTC

### Review by @xbhouse - Approved on October 09, 2025 at 12:09 PM UTC

thank you for all these fixes! lgtm, nice job :) 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1229*
