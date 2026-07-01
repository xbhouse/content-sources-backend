---
type: pull_request
number: 1375
title: "Test: Add race-safe random repo generator"
state: merged
author: marusak
created: 2026-01-22T08:56:12Z
updated: 2026-01-22T12:25:06Z
closed: 2026-01-22T12:25:06Z
merged: 2026-01-22T12:25:06Z
base_branch: main
head_branch: race_safe_repo_generator
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1375
---

# Pull Request #1375: Test: Add race-safe random repo generator

**Author**: @marusak
**Created**: January 22, 2026 at 08:56 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `race_safe_repo_generator`

## Description

The current version had issue that if two tests requested repo at the same time, both were free when checked with API. Then they both try to use it and it fails. Thus I am splitting the range of repositories into subset based on number of workers. And also picking up random number from the range rather then subsequently from 1 to 100.

## Summary

## Testing steps



---

## Reviews

### Review by @TenSt - Approved on January 22, 2026 at 11:29 AM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1375*
