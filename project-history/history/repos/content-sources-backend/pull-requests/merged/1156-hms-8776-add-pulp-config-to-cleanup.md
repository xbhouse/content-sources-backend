---
type: pull_request
number: 1156
title: "HMS-8776: Add pulp config to cleanup"
state: merged
author: jlsherrill
created: 2025-07-25T17:50:26Z
updated: 2025-07-25T19:22:18Z
closed: 2025-07-25T19:22:17Z
merged: 2025-07-25T19:22:17Z
base_branch: main
head_branch: 8776
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1156
---

# Pull Request #1156: HMS-8776: Add pulp config to cleanup

**Author**: @jlsherrill
**Created**: July 25, 2025 at 05:50 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `8776`

## Description

## Summary

Adds the pulp connection env variables to the cleanup command, these are needed to properly delete an upload request in pulp.

## Testing steps

These only really affect stage/prod, so i don't think we need much testing.  Will monitor stage deployment to check for success 

---

## Discussion

### Comment by @jlsherrill on July 25, 2025 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-8776

### Comment by @jlsherrill on July 25, 2025 at 06:10 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on July 25, 2025 at 06:31 PM UTC

### Review by @xbhouse - Approved on July 25, 2025 at 06:32 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1156*
