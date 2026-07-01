---
type: pull_request
number: 567
title: "Fixes 3353: Snapshot  inconsistent for API/UI"
state: merged
author: Andrewgdewar
created: 2024-02-06T21:36:14Z
updated: 2024-02-13T18:00:40Z
closed: 2024-02-13T17:37:59Z
merged: 2024-02-13T17:37:59Z
base_branch: main
head_branch: HMS-3353
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/567
---

# Pull Request #567: Fixes 3353: Snapshot  inconsistent for API/UI

**Author**: @Andrewgdewar
**Created**: February 06, 2024 at 09:36 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-3353`

## Description

## Summary

Snapshot create-on-edit is inconsistent between API and UI 
This was caused by the URL not being checked correctly before enqueueing snapshot/introspect tasks. 
This card corrects that issue.

## Testing steps

- From the api or UI, update a repository (patch) with the original URL in the request.  
Without this change, introspection/snapshot tasks will always be enqueued. 
With it, the tasks will correctly, be ignored, as those two tasks are contingent on a URL change.




---

## Discussion

### Comment by @jlsherrill on February 06, 2024 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-3353

### Comment by @swadeley on February 07, 2024 at 02:40 PM UTC

lgtm

### Comment by @swadeley on February 07, 2024 at 04:00 PM UTC

/retest

### Comment by @swadeley on February 09, 2024 at 12:04 PM UTC

/retest

### Comment by @swadeley on February 12, 2024 at 01:45 PM UTC

/retest

---

## Reviews

### Review by @Andrewgdewar - Commented on February 06, 2024 at 09:41 PM UTC

### Review by @rverdile - Commented on February 07, 2024 at 03:21 PM UTC

### Review by @Andrewgdewar - Commented on February 08, 2024 at 10:15 PM UTC

### Review by @Andrewgdewar - Commented on February 09, 2024 at 09:02 PM UTC

### Review by @rverdile - Commented on February 12, 2024 at 02:10 PM UTC

### Review by @Andrewgdewar - Commented on February 12, 2024 at 09:49 PM UTC

### Review by @rverdile - Approved on February 13, 2024 at 03:14 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/567*
