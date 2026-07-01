---
type: pull_request
number: 104
title: "Fixes 152: Smarter introspection job"
state: merged
author: avisiedo
created: 2022-09-12T09:13:21Z
updated: 2022-10-04T15:30:29Z
closed: 2022-10-04T15:20:04Z
merged: 2022-10-04T15:20:04Z
base_branch: main
head_branch: hmscontent-152
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/104
---

# Pull Request #104: Fixes 152: Smarter introspection job

**Author**: @avisiedo
**Created**: September 12, 2022 at 09:13 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `hmscontent-152`

## Description

- It runs every 8 hours, starting at 0:00 hours.
- It executes for every repository with Status != 'Valid'.
- It does not execute for a repository with Status == 'Valid' hat was introspected into the last 24 hours.
- Unit tests for the above.
- Removed: Add livenessProbe to the pod job specification: it was stepped back.


---

## Discussion

### Comment by @jlsherrill on September 12, 2022 at 09:30 AM UTC

https://issues.redhat.com/browse/HMSCONTENT-152

### Comment by @jlsherrill on September 14, 2022 at 03:50 PM UTC

This is good to rebase now that ryan's pr has gone in

### Comment by @rverdile on September 20, 2022 at 07:49 PM UTC

I left a couple more comments about the output messaging, but otherwise seems to function well

### Comment by @avisiedo on September 26, 2022 at 10:14 AM UTC

sorry for the delay! I am on this pr! thanks for the comments! looking into them! :+1: 

### Comment by @avisiedo on September 26, 2022 at 04:16 PM UTC

/retest

### Comment by @mayurilahane on October 04, 2022 at 03:19 PM UTC

note: Verified repo with (status != valid )

---

## Reviews

### Review by @rverdile - Commented on September 19, 2022 at 05:23 PM UTC

### Review by @rverdile - Commented on September 19, 2022 at 05:49 PM UTC

### Review by @avisiedo - Commented on September 20, 2022 at 03:25 PM UTC

### Review by @avisiedo - Commented on September 20, 2022 at 03:31 PM UTC

### Review by @jlsherrill - Commented on September 20, 2022 at 03:32 PM UTC

### Review by @jlsherrill - Commented on September 20, 2022 at 03:36 PM UTC

### Review by @rverdile - Commented on September 20, 2022 at 07:15 PM UTC

### Review by @rverdile - Commented on September 20, 2022 at 07:44 PM UTC

### Review by @rverdile - Commented on September 20, 2022 at 07:46 PM UTC

### Review by @rverdile - Commented on September 20, 2022 at 07:46 PM UTC

### Review by @rverdile - Commented on September 20, 2022 at 07:47 PM UTC

### Review by @rverdile - Commented on September 20, 2022 at 07:47 PM UTC

### Review by @avisiedo - Commented on September 26, 2022 at 10:22 AM UTC

### Review by @rverdile - Commented on September 26, 2022 at 06:52 PM UTC

### Review by @rverdile - Commented on September 26, 2022 at 06:55 PM UTC

### Review by @jlsherrill - Commented on September 27, 2022 at 09:01 PM UTC

### Review by @avisiedo - Commented on September 28, 2022 at 06:27 AM UTC

### Review by @avisiedo - Commented on September 28, 2022 at 06:58 AM UTC

### Review by @rverdile - Approved on September 28, 2022 at 02:45 PM UTC

lgtm!

### Review by @swadeley - Commented on September 29, 2022 at 10:22 AM UTC

### Review by @swadeley - Commented on September 29, 2022 at 10:22 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/104*
