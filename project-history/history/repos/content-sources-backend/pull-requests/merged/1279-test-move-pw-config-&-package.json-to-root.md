---
type: pull_request
number: 1279
title: "Test: Move PW config & package.json to root"
state: merged
author: swadeley
created: 2025-11-03T11:08:34Z
updated: 2026-01-13T06:53:50Z
closed: 2025-11-04T12:44:12Z
merged: 2025-11-04T12:44:12Z
base_branch: main
head_branch: swadeley/move_package_json
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1279
---

# Pull Request #1279: Test: Move PW config & package.json to root

**Author**: @swadeley
**Created**: November 03, 2025 at 11:08 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/move_package_json`

## Description

## Summary

Move PW config & package.json to root

  To enable running PW tests from the root of the repo.
  To make it the same for users as the frontend.
   Update PW actions to point to root
   Update gitignore

## Testing steps
builds and tests pass


---

## Discussion

### Comment by @swadeley on November 04, 2025 at 09:51 AM UTC

2025-11-03 14:30:21 [   ERROR] [          MainThread] deploy failed: timed out waiting for ClowdApp-owned resources

### Comment by @swadeley on November 04, 2025 at 09:51 AM UTC

/retest

---

## Reviews

### Review by @TenSt - Changes Requested on November 03, 2025 at 11:58 AM UTC

Please update workflow files - you need to remove/update the `working-directory` parameter in steps were it is used to point to root (by default if it is not set).

### Review by @TenSt - Approved on November 04, 2025 at 10:41 AM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1279*
