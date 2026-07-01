---
type: pull_request
number: 163
title: "Fixes 289: Package count accuracy fix"
state: merged
author: Andrewgdewar
created: 2023-01-05T17:51:57Z
updated: 2023-01-11T10:30:40Z
closed: 2023-01-11T10:00:27Z
merged: 2023-01-11T10:00:27Z
base_branch: main
head_branch: CONTENT-289
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/163
---

# Pull Request #163: Fixes 289: Package count accuracy fix

**Author**: @Andrewgdewar
**Created**: January 05, 2023 at 05:51 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `CONTENT-289`

## Description

This fixes an issue seen when two packages were identical but have different names.

In such a case the total count would be incorrect, as the checksum data used for package uniqueness would prevent duplication of the package, but the original package count would still be used after this process was completed. 

This fix makes an actual count of the associated items in the database after the introspection is finished to accurately return the correct number for storage. 

- [x]  Needs tests

---

## Discussion

### Comment by @jlsherrill on January 05, 2023 at 06:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-289

### Comment by @swadeley on January 10, 2023 at 02:16 PM UTC

/retest

### Comment by @swadeley on January 11, 2023 at 10:00 AM UTC

OK, now it shows just one package. Thank you.

---

## Reviews

### Review by @Andrewgdewar - Commented on January 05, 2023 at 05:53 PM UTC

### Review by @rverdile - Commented on January 05, 2023 at 08:31 PM UTC

### Review by @Andrewgdewar - Commented on January 05, 2023 at 10:08 PM UTC

### Review by @rverdile - Approved on January 09, 2023 at 07:54 PM UTC

looks good. tested and working. nice!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/163*
