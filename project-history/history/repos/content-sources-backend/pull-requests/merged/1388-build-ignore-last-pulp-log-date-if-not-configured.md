---
type: pull_request
number: 1388
title: "Build: ignore last pulp log date if not configured"
state: merged
author: jlsherrill
created: 2026-02-03T21:03:45Z
updated: 2026-02-04T19:20:49Z
closed: 2026-02-04T19:20:43Z
merged: 2026-02-04T19:20:43Z
base_branch: main
head_branch: silence_perf
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1388
---

# Pull Request #1388: Build: ignore last pulp log date if not configured

**Author**: @jlsherrill
**Created**: February 03, 2026 at 09:03 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `silence_perf`

## Description

## Summary

This stops collecting the pulp log metric if not configured, resulting in the error not logged in the perf testing environment.

## Testing steps



---

## Discussion

### Comment by @swadeley on February 04, 2026 at 11:09 AM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on February 03, 2026 at 09:49 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1388*
