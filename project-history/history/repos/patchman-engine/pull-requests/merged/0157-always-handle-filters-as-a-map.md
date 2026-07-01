---
type: pull_request
number: 157
title: "Always handle filters as a map"
state: merged
author: semtexzv
created: 2020-03-02T13:58:07Z
updated: 2021-03-16T09:27:35Z
closed: 2020-03-02T15:09:31Z
merged: 2020-03-02T15:09:31Z
base_branch: master
head_branch: filters-map
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/157
---

# Pull Request #157: Always handle filters as a map

**Author**: @semtexzv
**Created**: March 02, 2020 at 01:58 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `filters-map`

## Description

Current code which handles filters as an array has an error, which results in default filters always being applied.

This PR changes this behavior, and always handles filters as a map keyed by field name to which a filter is applied.

---

## Reviews

### Review by @josef-hak - Approved on March 02, 2020 at 03:09 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/157*
