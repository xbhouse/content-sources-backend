---
type: pull_request
number: 61
title: "Increase kafka timeout"
state: merged
author: semtexzv
created: 2020-01-15T10:02:07Z
updated: 2021-03-16T09:26:55Z
closed: 2020-01-15T16:25:35Z
merged: 2020-01-15T16:25:35Z
base_branch: master
head_branch: listener-timeout
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/61
---

# Pull Request #61: Increase kafka timeout

**Author**: @semtexzv
**Created**: January 15, 2020 at 10:02 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `listener-timeout`

## Description

*No description provided*

---

## Discussion

### Comment by @beav on January 15, 2020 at 01:11 PM UTC

was it timing out before?

### Comment by @semtexzv on January 15, 2020 at 04:18 PM UTC

Yes, we have timeouts when we receive large upload batches, this change is to see whether it is just a matter of kafka being overloaded or other issue.

---

## Reviews

### Review by @beav - Approved on January 15, 2020 at 04:21 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/61*
