---
type: pull_request
number: 1059
title: "Test: fix name creation in repo tests"
state: merged
author: xbhouse
created: 2025-04-04T17:50:20Z
updated: 2025-04-07T17:08:12Z
closed: 2025-04-07T17:08:12Z
merged: 2025-04-07T17:08:12Z
base_branch: main
head_branch: fix-name-creation
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1059
---

# Pull Request #1059: Test: fix name creation in repo tests

**Author**: @xbhouse
**Created**: April 04, 2025 at 05:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-name-creation`

## Description

## Summary

We are seeing 403s in the integration repo [test runs](https://github.com/content-services/content-sources-playwright/actions/runs/14269364245/job/39998792719#step:10:119) from these 2 tests when they call the cleanup helper, likely because we were passing a function instead of a string. This fixes the name creation to call the function before passing it to the helper.

## Testing steps


---

## Discussion

### Comment by @xbhouse on April 07, 2025 at 05:08 PM UTC

merging to try to fix integration suite

---

## Reviews

### Review by @swadeley - Approved on April 04, 2025 at 06:17 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1059*
