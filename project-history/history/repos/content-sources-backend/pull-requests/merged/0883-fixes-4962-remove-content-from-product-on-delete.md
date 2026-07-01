---
type: pull_request
number: 883
title: "Fixes 4962: remove content from product on delete"
state: merged
author: dominikvagner
created: 2024-11-11T12:39:47Z
updated: 2024-11-12T14:30:18Z
closed: 2024-11-12T14:30:18Z
merged: 2024-11-12T14:30:18Z
base_branch: main
head_branch: 4962
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/883
---

# Pull Request #883: Fixes 4962: remove content from product on delete

**Author**: @dominikvagner
**Created**: November 11, 2024 at 12:39 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4962`

## Description

## Summary
This PR fixes an issues caused by a recent candlepin update which removed the convenience feature of auto-unlink content on deletion from products, now a separate call to remove content from the product has to be done before deleting the content.

## Testing steps
Integration tests pass and pipeline "Test" step also passes, content can be deleted from candlepin.


---

## Discussion

### Comment by @jlsherrill on November 11, 2024 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-4962

### Comment by @dominikvagner on November 11, 2024 at 01:12 PM UTC

/retest

### Comment by @rverdile on November 11, 2024 at 02:27 PM UTC

/retest

### Comment by @swadeley on November 11, 2024 at 03:50 PM UTC

/retest

### Comment by @dominikvagner on November 12, 2024 at 06:51 AM UTC

/retest

### Comment by @swadeley on November 12, 2024 at 09:48 AM UTC

/retest

### Comment by @swadeley on November 12, 2024 at 11:34 AM UTC

/retest

### Comment by @swadeley on November 12, 2024 at 12:53 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on November 11, 2024 at 02:34 PM UTC

looks good!

### Review by @jlsherrill - Commented on November 12, 2024 at 01:48 PM UTC

### Review by @jlsherrill - Approved on November 12, 2024 at 02:17 PM UTC

ack from me too

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/883*
