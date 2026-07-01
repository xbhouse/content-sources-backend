---
type: pull_request
number: 1332
title: "HMS-9846: add retain_days_limit to cleanup deployment"
state: merged
author: rverdile
created: 2025-12-11T19:00:51Z
updated: 2025-12-11T19:21:22Z
closed: 2025-12-11T19:21:19Z
merged: 2025-12-11T19:21:19Z
base_branch: main
head_branch: add-retain-to-deploy
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1332
---

# Pull Request #1332: HMS-9846: add retain_days_limit to cleanup deployment

**Author**: @rverdile
**Created**: December 11, 2025 at 07:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `add-retain-to-deploy`

## Description

## Summary
The cleanup deployment is missing the OPTIONS_SNAPSHOT_RETAIN_DAYS_LIMIT variable, so it is not using the expected value and snapshots are not being cleaned up as expected

## Testing steps



---

## Discussion

### Comment by @rverdile on December 11, 2025 at 07:12 PM UTC

to unblock: https://github.com/content-services/content-sources-frontend/pull/777

---

## Reviews

### Review by @xbhouse - Approved on December 11, 2025 at 07:18 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1332*
