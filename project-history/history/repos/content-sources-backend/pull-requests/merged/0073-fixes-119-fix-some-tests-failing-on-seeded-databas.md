---
type: pull_request
number: 73
title: "Fixes 119: fix some tests failing on seeded database"
state: merged
author: avisiedo
created: 2022-08-05T17:30:48Z
updated: 2022-08-26T15:26:59Z
closed: 2022-08-26T15:26:59Z
merged: 2022-08-26T15:26:59Z
base_branch: main
head_branch: hmscontent-119
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/73
---

# Pull Request #73: Fixes 119: fix some tests failing on seeded database

**Author**: @avisiedo
**Created**: August 05, 2022 at 05:30 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `hmscontent-119`

## Description

- Update TestUpdateEmpty to make it more predictable.
- Fix TestSavePublicUrls test.
- Update TestUpdateBlank to make it more predictable.
- Update TestBulkCreateOneFails to make it more predictable.

Signed-off-by: Alejandro Visiedo <avisiedo@redhat.com>

---

## Discussion

### Comment by @jlsherrill on August 05, 2022 at 06:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-119

### Comment by @avisiedo on August 24, 2022 at 09:35 AM UTC

I have realized more places to add the org_id filter, now it should be better. I have added at some points `require` instead of `assert` specially when the data (error or data models) are used later.

I have rebased the pr so now it should be in better shape

### Comment by @avisiedo on August 26, 2022 at 04:42 AM UTC

/retest

### Comment by @jlsherrill on August 26, 2022 at 12:24 PM UTC

@avisiedo feel free to go ahead and merge, the failing tests will be fixed by https://github.com/content-services/content-sources-backend/pull/86

---

## Reviews

### Review by @avisiedo - Commented on August 05, 2022 at 05:33 PM UTC

### Review by @avisiedo - Commented on August 05, 2022 at 05:34 PM UTC

### Review by @avisiedo - Commented on August 05, 2022 at 05:35 PM UTC

### Review by @avisiedo - Commented on August 05, 2022 at 05:37 PM UTC

### Review by @avisiedo - Commented on August 05, 2022 at 05:38 PM UTC

### Review by @avisiedo - Commented on August 05, 2022 at 05:43 PM UTC

### Review by @rverdile - Commented on August 09, 2022 at 08:16 PM UTC

### Review by @rverdile - Commented on August 09, 2022 at 08:18 PM UTC

### Review by @rverdile - Commented on August 23, 2022 at 07:35 PM UTC

looking good, but there's some merge conflicts now

### Review by @rverdile - Approved on August 25, 2022 at 03:38 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/73*
