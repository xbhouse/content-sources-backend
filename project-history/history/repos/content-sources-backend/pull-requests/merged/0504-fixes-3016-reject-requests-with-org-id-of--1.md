---
type: pull_request
number: 504
title: "Fixes 3016: Reject requests with org id of -1"
state: merged
author: xbhouse
created: 2023-12-13T21:48:50Z
updated: 2024-02-13T15:16:10Z
closed: 2023-12-18T20:42:53Z
merged: 2023-12-18T20:42:53Z
base_branch: main
head_branch: orgid-check
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/504
---

# Pull Request #504: Fixes 3016: Reject requests with org id of -1

**Author**: @xbhouse
**Created**: December 13, 2023 at 09:48 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `orgid-check`

## Description

## Summary

Adds check to middleware to reject requests if the Org ID is -1

## Testing steps

Make a request to any endpoint with an identity header with Org ID set to -1. Response should return a 403.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on December 13, 2023 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-3016

### Comment by @jlsherrill on December 18, 2023 at 08:42 PM UTC

going ahead and merging as a) 95% of this pr was already qe'd, and b) the 5% left can't easily be tested by qe

---

## Reviews

### Review by @jlsherrill - Approved on December 18, 2023 at 07:24 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/504*
