---
type: pull_request
number: 1268
title: "Build: add wait for playwright setup repos"
state: merged
author: dominikvagner
created: 2025-10-29T21:47:23Z
updated: 2025-11-04T11:33:59Z
closed: 2025-11-04T11:33:58Z
merged: 2025-11-04T11:33:58Z
base_branch: main
head_branch: add-valid-setup-repos-wait
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1268
---

# Pull Request #1268: Build: add wait for playwright setup repos

**Author**: @dominikvagner
**Created**: October 29, 2025 at 09:47 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `add-valid-setup-repos-wait`

## Description

## Summary
This adds a step in the Playwright CI that will wait until all the setup repos are valid. This should help alleviate some of the flakiness.



---

## Discussion

### Comment by @dominikvagner on November 03, 2025 at 01:49 PM UTC

> is there a timeout somewhere in the case the repo fails to introspect or snapshot for some reason?

good point, added a 3 minute timeout (last time it took a minute, so I think this should be safe) on that workflow step 🫡 👍🏼 


### Comment by @rverdile on November 03, 2025 at 03:03 PM UTC

/retest

### Comment by @rverdile on November 03, 2025 at 10:28 PM UTC

/retest

---

## Reviews

### Review by @TenSt - Approved on October 29, 2025 at 10:49 PM UTC

### Review by @rverdile - Commented on October 30, 2025 at 01:25 PM UTC

is there a timeout somewhere in the case the repo fails to introspect or snapshot for some reason?

### Review by @rverdile - Approved on November 03, 2025 at 03:03 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1268*
