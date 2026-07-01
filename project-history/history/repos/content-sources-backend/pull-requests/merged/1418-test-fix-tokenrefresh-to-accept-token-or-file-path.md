---
type: pull_request
number: 1418
title: "Test: Fix tokenRefresh to accept token or file path"
state: merged
author: swadeley
created: 2026-03-20T14:04:37Z
updated: 2026-03-23T09:06:06Z
closed: 2026-03-23T09:06:06Z
merged: 2026-03-23T09:06:06Z
base_branch: main
head_branch: swadeley/fix_ensureValidToken
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1418
---

# Pull Request #1418: Test: Fix tokenRefresh to accept token or file path

**Author**: @swadeley
**Created**: March 20, 2026 at 02:04 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/fix_ensureValidToken`

## Description

## Summary

Fix tokenRefresh to accept token or file path
Tests using storageState must pass path to the json file.

## Testing steps

_playwright-tests/Integration/CanUpdateSystemWithTemplate.spec.ts passes

---

## Reviews

### Review by @TenSt - Approved on March 23, 2026 at 08:34 AM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1418*
