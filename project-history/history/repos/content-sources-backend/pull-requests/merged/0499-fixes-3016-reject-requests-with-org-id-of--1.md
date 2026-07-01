---
type: pull_request
number: 499
title: "Fixes 3016: Reject requests with org id of -1"
state: merged
author: xbhouse
created: 2023-12-07T15:22:08Z
updated: 2023-12-18T21:00:39Z
closed: 2023-12-13T19:01:10Z
merged: 2023-12-13T19:01:10Z
base_branch: main
head_branch: orgid-check
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/499
---

# Pull Request #499: Fixes 3016: Reject requests with org id of -1

**Author**: @xbhouse
**Created**: December 07, 2023 at 03:22 PM UTC
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

### Comment by @jlsherrill on December 07, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-3016

### Comment by @jlsherrill on December 13, 2023 at 02:54 PM UTC

Two small comments, otherwise looks great and works fine!

### Comment by @Andrewgdewar on December 13, 2023 at 03:34 PM UTC

/retest

### Comment by @mayurilahane on December 13, 2023 at 07:00 PM UTC

lgtm

---

## Reviews

### Review by @jlsherrill - Commented on December 11, 2023 at 01:33 PM UTC

### Review by @xbhouse - Commented on December 11, 2023 at 02:16 PM UTC

### Review by @jlsherrill - Commented on December 13, 2023 at 02:53 PM UTC

### Review by @jlsherrill - Commented on December 13, 2023 at 02:54 PM UTC

### Review by @jlsherrill - Approved on December 13, 2023 at 06:09 PM UTC

ACK!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/499*
