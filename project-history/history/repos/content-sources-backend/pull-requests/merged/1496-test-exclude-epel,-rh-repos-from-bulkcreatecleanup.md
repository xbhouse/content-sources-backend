---
type: pull_request
number: 1496
title: "Test: exclude EPEL, RH repos from BulkCreateCleanupURL test"
state: merged
author: xbhouse
created: 2026-05-14T16:28:47Z
updated: 2026-05-15T15:23:51Z
closed: 2026-05-15T15:23:51Z
merged: 2026-05-15T15:23:50Z
base_branch: main
head_branch: fix-test
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1496
---

# Pull Request #1496: Test: exclude EPEL, RH repos from BulkCreateCleanupURL test

**Author**: @xbhouse
**Created**: May 14, 2026 at 04:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-test`

## Description

## Summary

Excludes EPEL and RH repos from the query in BulkCreateCleanupURL so the test does not intermittently fail if one of the seeded repos is found first

## Testing steps

TestBulkCreateCleanupURL passes consistently

---

## Reviews

### Review by @rverdile - Commented on May 14, 2026 at 07:10 PM UTC

### Review by @xbhouse - Commented on May 14, 2026 at 07:17 PM UTC

### Review by @xbhouse - Commented on May 14, 2026 at 07:19 PM UTC

### Review by @xbhouse - Commented on May 14, 2026 at 09:51 PM UTC

### Review by @TenSt - Dismissed on May 15, 2026 at 09:08 AM UTC

### Review by @rverdile - Commented on May 15, 2026 at 11:30 AM UTC

### Review by @rverdile - Commented on May 15, 2026 at 11:31 AM UTC

### Review by @xbhouse - Commented on May 15, 2026 at 01:01 PM UTC

### Review by @TenSt - Approved on May 15, 2026 at 02:29 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1496*
