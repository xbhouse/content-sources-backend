---
type: pull_request
number: 393
title: "Remove the need for the packages endpoint to visit whole partition"
state: merged
author: semtexzv
created: 2020-10-15T12:47:51Z
updated: 2021-03-16T09:29:48Z
closed: 2020-10-15T19:08:20Z
merged: 2020-10-15T19:08:20Z
base_branch: master
head_branch: packages-index-scan
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/393
---

# Pull Request #393: Remove the need for the packages endpoint to visit whole partition

**Author**: @semtexzv
**Created**: October 15, 2020 at 12:47 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `packages-index-scan`

## Description

We currently require a `system_package_<part>`  table scan for packages endpoint. This adds computed column, which is then included in the index, and is used for the counts. This should remove the need to do full table scan on the partition, and remove the worst offender in the slowness of packages endpoint.

---

## Reviews

### Review by @josef-hak - Approved on October 15, 2020 at 03:13 PM UTC

Looks, good. I hope it will help.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/393*
