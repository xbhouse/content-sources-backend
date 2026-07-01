---
type: pull_request
number: 1551
title: "Build: disable sast coverity check"
state: merged
author: TenSt
created: 2026-06-25T16:55:08Z
updated: 2026-06-25T17:16:52Z
closed: 2026-06-25T17:16:51Z
merged: 2026-06-25T17:16:51Z
base_branch: main
head_branch: stepan/disable-sast-coverity-check
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1551
---

# Pull Request #1551: Build: disable sast coverity check

**Author**: @TenSt
**Created**: June 25, 2026 at 04:55 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `stepan/disable-sast-coverity-check`

## Description

This PR disables sast coverity check as it takes about 28m to run.
We will need to investigate how to configure it later to run faster and not block the pipeline.

---

## Reviews

### Review by @rverdile - Approved on June 25, 2026 at 05:16 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1551*
