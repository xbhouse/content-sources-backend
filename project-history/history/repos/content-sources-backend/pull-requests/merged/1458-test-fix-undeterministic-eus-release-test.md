---
type: pull_request
number: 1458
title: "Test: Fix undeterministic eus release test"
state: merged
author: swadeley
created: 2026-04-13T09:39:06Z
updated: 2026-04-13T10:18:42Z
closed: 2026-04-13T10:18:42Z
merged: 2026-04-13T10:18:41Z
base_branch: main
head_branch: swadeley/fix_flaky_template_test
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1458
---

# Pull Request #1458: Test: Fix undeterministic eus release test

**Author**: @swadeley
**Created**: April 13, 2026 at 09:39 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/fix_flaky_template_test`

## Description


## Summary
Fix undeterministic eus release test

The subtest filters with found[0].ExtendedRelease
 and when found[0] happens to be the standard template
 ExtendedRelease: "" applies no filter, and three templates are returned not one.
## Testing steps



---

## Reviews

### Review by @TenSt - Approved on April 13, 2026 at 10:15 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1458*
