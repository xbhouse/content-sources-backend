---
type: pull_request
number: 439
title: "Build: use ubi-micro"
state: merged
author: jlsherrill
created: 2023-10-23T19:40:24Z
updated: 2023-10-25T12:54:51Z
closed: 2023-10-25T12:54:47Z
merged: 2023-10-25T12:54:47Z
base_branch: main
head_branch: micro
labels: []
url: https://github.com/content-services/content-sources-backend/pull/439
---

# Pull Request #439: Build: use ubi-micro

**Author**: @jlsherrill
**Created**: October 23, 2023 at 07:40 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `micro`

## Description

## Summary

Switches to using the ubi-micro container, which appears to have everything we need other than a ca-trust, which we can copy from ubi-minimal

This reduces our security footprint and potential attack vectors

## Testing steps

tests pass


---

## Reviews

### Review by @rverdile - Approved on October 24, 2023 at 08:21 PM UTC

tests pass!

### Review by @mshriver - Approved on October 25, 2023 at 11:13 AM UTC

@jlsherrill leaving this for you to merge

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/439*
