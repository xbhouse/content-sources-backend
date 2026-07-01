---
type: pull_request
number: 1111
title: "Refresh System detail tables on template change"
state: merged
author: leSamo
created: 2023-08-24T14:24:44Z
updated: 2024-01-08T13:49:57Z
closed: 2023-08-28T13:06:24Z
merged: 2023-08-28T13:06:24Z
base_branch: master
head_branch: refresh-template-change
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1111
---

# Pull Request #1111: Refresh System detail tables on template change

**Author**: @leSamo
**Created**: August 24, 2023 at 02:24 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `refresh-template-change`

## Description

https://issues.redhat.com/browse/SPM-2062

How to test:

1. Go to system detail page
2. Try to apply/remove template multiple times from the "Actions" toolbar in top-right corner of the page
3. Make sure Advisories and Packages tables are refreshed after applying/removing template and show the correct data

---

## Discussion

### Comment by @leSamo on August 24, 2023 at 04:27 PM UTC

/retest

### Comment by @mkholjuraev on August 30, 2023 at 08:07 PM UTC

:tada: This PR is included in version 1.64.0 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.64.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on August 28, 2023 at 12:34 PM UTC

LGTM! I keep getting 403 errors from the API while assigning systems. But codewise, I am pretty sure that it works

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1111*
