---
type: pull_request
number: 1227
title: "Build: add development helper make commands"
state: merged
author: dominikvagner
created: 2025-10-06T13:14:38Z
updated: 2025-10-06T13:51:15Z
closed: 2025-10-06T13:51:15Z
merged: 2025-10-06T13:51:15Z
base_branch: main
head_branch: qol-improvements
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1227
---

# Pull Request #1227: Build: add development helper make commands

**Author**: @dominikvagner
**Created**: October 06, 2025 at 01:14 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `qol-improvements`

## Description

## Summary
This adds 2 new helper make commands.
- `make compose-run`, that starts the container dependencies without initializing any
- `make repos-mvp`, which will import 2 repos needed for a minimal local setup and force them to snapshot

## Testing steps
Verify those commands do what they should.


---

## Reviews

### Review by @TenSt - Approved on October 06, 2025 at 01:32 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1227*
