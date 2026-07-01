---
type: pull_request
number: 1532
title: "Test: run post-test cleanup sequentially"
state: merged
author: swadeley
created: 2026-06-10T09:14:02Z
updated: 2026-06-10T18:16:16Z
closed: 2026-06-10T18:16:16Z
merged: 2026-06-10T18:16:16Z
base_branch: main
head_branch: swadeley/update_cleanup_to_be_sequential
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1532
---

# Pull Request #1532: Test: run post-test cleanup sequentially

**Author**: @swadeley
**Created**: June 10, 2026 at 09:14 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/update_cleanup_to_be_sequential`

## Description



## Summary
Run post-test cleanup sequentially
Parallel Promise.all teardown caused race conditions when template deletes and RHSM container disconnect ran at the same time.

## Testing steps

Integration template tests pass

---

## Reviews

### Review by @TenSt - Approved on June 10, 2026 at 02:20 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1532*
