---
type: pull_request
number: 1298
title: "Test: Log but not fail if template not found"
state: merged
author: swadeley
created: 2025-11-13T11:40:34Z
updated: 2026-01-13T06:54:03Z
closed: 2025-11-13T14:13:13Z
merged: 2025-11-13T14:13:12Z
base_branch: main
head_branch: swadeley/log_not_fail_on_template_cleanup
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1298
---

# Pull Request #1298: Test: Log but not fail if template not found

**Author**: @swadeley
**Created**: November 13, 2025 at 11:40 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/log_not_fail_on_template_cleanup`

## Description

  ## Summary

Fail in cleanup should not pass errors up to the test.

## Testing steps
tests pass



---

## Discussion

### Comment by @TenSt on November 13, 2025 at 02:13 PM UTC

Merging this as I believe I have a flaky test that is failing because of it.

---

## Reviews

### Review by @TenSt - Approved on November 13, 2025 at 01:38 PM UTC

lgtm 👍 

### Review by @rverdile - Commented on November 13, 2025 at 03:04 PM UTC

### Review by @swadeley - Commented on November 13, 2025 at 05:35 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1298*
