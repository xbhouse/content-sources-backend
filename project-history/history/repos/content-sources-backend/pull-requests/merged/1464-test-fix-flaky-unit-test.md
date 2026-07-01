---
type: pull_request
number: 1464
title: "Test: fix flaky unit test"
state: merged
author: rverdile
created: 2026-04-15T17:20:31Z
updated: 2026-04-15T18:09:36Z
closed: 2026-04-15T18:09:33Z
merged: 2026-04-15T18:09:33Z
base_branch: main
head_branch: fix-test-2
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1464
---

# Pull Request #1464: Test: fix flaky unit test

**Author**: @rverdile
**Created**: April 15, 2026 at 05:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-test-2`

## Description

## Summary
Needs to calculate now() once otherwise started_at could be before queued_at

## Testing steps
1. If you run this test a few times without these changes you'll probably see it fail



---

## Discussion

### Comment by @rverdile on April 15, 2026 at 05:47 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on April 15, 2026 at 05:37 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1464*
