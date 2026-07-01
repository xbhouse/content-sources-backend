---
type: pull_request
number: 95
title: "Fixes 109: Store and return introspection status metadata"
state: merged
author: rverdile
created: 2022-08-31T19:06:54Z
updated: 2022-09-14T13:53:30Z
closed: 2022-09-14T13:53:29Z
merged: 2022-09-14T13:53:29Z
base_branch: main
head_branch: introspect-status
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/95
---

# Pull Request #95: Fixes 109: Store and return introspection status metadata

**Author**: @rverdile
**Created**: August 31, 2022 at 07:06 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `introspect-status`

## Description

Adds introspection status timestamps, status, and last error to repository. Updated on introspection and returned in list and fetch responses.

To test, create some repositories and list/fetch them. See the new fields. Then, introspect the repositories and list/fetch them again to see how the timestamps/error/status change.

---

## Discussion

### Comment by @jlsherrill on August 31, 2022 at 07:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-109

### Comment by @rverdile on September 09, 2022 at 07:22 PM UTC

Something I noticed, but I don't know if it's a something to fix here or in another issue. 

Let's say there are two repositories with distinct URLs, but the same packages. Then, on introspection, the second repository to be introspected would see no changes to the `rpms` table. In which case, there is considered to be 0 packages inserted.

In the case of this issue, the "LastIntrospectionUpdateTime" would not get updated.

### Comment by @jlsherrill on September 12, 2022 at 12:34 PM UTC

@rverdile oh yes, that 'count' was really just an easy way to see that something happened on first introspection.  I wouldn't use it to determine if something changed in a repository.  In fact we can remove it completely if its causing confusion.

### Comment by @avisiedo on September 14, 2022 at 12:23 PM UTC

No furthers thoughts! LGTM  :+1: 

---

## Reviews

### Review by @swadeley - Commented on September 06, 2022 at 12:40 PM UTC

### Review by @avisiedo - Commented on September 07, 2022 at 02:58 PM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 09:04 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 09:04 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 09:08 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 09:08 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 09:51 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 10:55 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 11:21 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 11:24 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 11:31 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 11:32 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 11:33 AM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 11:37 AM UTC

### Review by @avisiedo - Changes Requested on September 08, 2022 at 11:39 AM UTC

only a few changes

thanks for this pr

### Review by @avisiedo - Commented on September 09, 2022 at 10:28 AM UTC

### Review by @rverdile - Commented on September 09, 2022 at 06:52 PM UTC

### Review by @rverdile - Commented on September 09, 2022 at 06:52 PM UTC

### Review by @rverdile - Commented on September 09, 2022 at 07:00 PM UTC

### Review by @rverdile - Commented on September 09, 2022 at 07:01 PM UTC

### Review by @swadeley - Commented on September 12, 2022 at 03:46 PM UTC

### Review by @rverdile - Commented on September 12, 2022 at 04:42 PM UTC

### Review by @rverdile - Commented on September 12, 2022 at 04:44 PM UTC

### Review by @rverdile - Commented on September 12, 2022 at 04:44 PM UTC

### Review by @jlsherrill - Commented on September 12, 2022 at 08:32 PM UTC

### Review by @jlsherrill - Commented on September 12, 2022 at 08:54 PM UTC

### Review by @rverdile - Commented on September 12, 2022 at 09:44 PM UTC

### Review by @jlsherrill - Approved on September 14, 2022 at 12:57 AM UTC

ACK from me, any further thoughts @avisiedo ?

### Review by @avisiedo - Approved on September 14, 2022 at 12:28 PM UTC

### Review by @swadeley - Changes Requested on September 14, 2022 at 12:38 PM UTC

### Review by @swadeley - Approved on September 14, 2022 at 01:52 PM UTC

LGTM

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/95*
