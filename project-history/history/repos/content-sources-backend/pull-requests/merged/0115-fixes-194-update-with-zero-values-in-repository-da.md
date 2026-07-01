---
type: pull_request
number: 115
title: "Fixes 194: Update with zero values in repository dao"
state: merged
author: rverdile
created: 2022-09-27T15:16:47Z
updated: 2022-10-11T14:06:27Z
closed: 2022-10-11T14:06:27Z
merged: 2022-10-11T14:06:27Z
base_branch: main
head_branch: update-bug
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/115
---

# Pull Request #115: Fixes 194: Update with zero values in repository dao

**Author**: @rverdile
**Created**: September 27, 2022 at 03:16 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `update-bug`

## Description

Previously, zero values were being used to indicate that a field should not be updated on `dao.Repository`. For example, if `Repository.Revision == ""`, don't update. Theoretically, if a Revision number was actually `""`, this value could not be updated. 

This PR introduces a `RepositoryUpdate` structure that uses pointers for all values. Using this, `nil` indicates "do not update", so that `dao.Repository` can now be updated with zero values.

One way to test this would be to host a local yum repo where you can change the revision number. Try changing it then removing it and you should see it update in the database accordingly, after introspection.

---

## Discussion

### Comment by @jlsherrill on September 27, 2022 at 09:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-194

### Comment by @rverdile on September 29, 2022 at 02:50 PM UTC

rebased so now this includes package count

### Comment by @avisiedo on October 11, 2022 at 01:49 PM UTC

:+1: lgtm

---

## Reviews

### Review by @avisiedo - Changes Requested on September 30, 2022 at 11:57 AM UTC

One change change, and one suggestion which if agreed would need changes.

### Review by @rverdile - Commented on October 03, 2022 at 05:54 PM UTC

### Review by @rverdile - Commented on October 03, 2022 at 05:58 PM UTC

### Review by @jlsherrill - Commented on October 04, 2022 at 02:23 PM UTC

### Review by @avisiedo - Commented on October 11, 2022 at 01:47 PM UTC

### Review by @avisiedo - Approved on October 11, 2022 at 01:48 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/115*
