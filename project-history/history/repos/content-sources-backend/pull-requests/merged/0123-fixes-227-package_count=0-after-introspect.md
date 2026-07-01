---
type: pull_request
number: 123
title: "Fixes 227: package_count=0 after introspect"
state: merged
author: avisiedo
created: 2022-10-13T22:37:39Z
updated: 2022-10-18T14:01:10Z
closed: 2022-10-18T14:01:10Z
merged: 2022-10-18T14:01:10Z
base_branch: main
head_branch: hmscontent-227-fix-zero-count
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/123
---

# Pull Request #123: Fixes 227: package_count=0 after introspect

**Author**: @avisiedo
**Created**: October 13, 2022 at 10:37 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `hmscontent-227-fix-zero-count`

## Description

- update Introspect to pass dao.Repository by reference, and update callers.

> Sorry, I didn't realize about this during the pr reviews

TODO:

- [x] Add unit test for this change.

Signed-off-by: Alejandro Visiedo <avisiedo@redhat.com>

---

## Discussion

### Comment by @jlsherrill on October 13, 2022 at 11:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-227

### Comment by @avisiedo on October 18, 2022 at 02:00 PM UTC

:+1: merging

---

## Reviews

### Review by @rverdile - Commented on October 17, 2022 at 06:24 PM UTC

### Review by @rverdile - Commented on October 17, 2022 at 06:37 PM UTC

### Review by @rverdile - Commented on October 17, 2022 at 06:39 PM UTC

### Review by @avisiedo - Commented on October 17, 2022 at 06:42 PM UTC

### Review by @avisiedo - Commented on October 17, 2022 at 06:49 PM UTC

### Review by @rverdile - Commented on October 17, 2022 at 06:53 PM UTC

### Review by @rverdile - Commented on October 17, 2022 at 07:10 PM UTC

### Review by @rverdile - Commented on October 17, 2022 at 07:10 PM UTC

### Review by @rverdile - Commented on October 17, 2022 at 07:11 PM UTC

### Review by @rverdile - Commented on October 17, 2022 at 07:12 PM UTC

### Review by @avisiedo - Commented on October 18, 2022 at 06:54 AM UTC

### Review by @rverdile - Approved on October 18, 2022 at 01:19 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/123*
