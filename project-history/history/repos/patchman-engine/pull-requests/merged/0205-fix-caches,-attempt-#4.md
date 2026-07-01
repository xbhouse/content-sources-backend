---
type: pull_request
number: 205
title: "Fix caches, attempt #4"
state: merged
author: semtexzv
created: 2020-04-01T11:59:05Z
updated: 2020-04-01T13:56:39Z
closed: 2020-04-01T13:56:39Z
merged: 2020-04-01T13:56:39Z
base_branch: master
head_branch: fix-caches-calc
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/205
---

# Pull Request #205: Fix caches, attempt #4

**Author**: @semtexzv
**Created**: April 01, 2020 at 11:59 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-caches-calc`

## Description

- Revert back to incremental calculation
- Sequence of actions is now based on tables, first locking `system_platform`, then updating `system_advisories`, and then updating `advisory_account_data`, instead of interleaving updating 
`system_advisories` and `advisory_account_data` for unpatched and patched advisories
- Fixes to system_culling in previous PR should fix some causes of cache inconsistencies


---

## Reviews

### Review by @josef-hak - Approved on April 01, 2020 at 01:56 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/205*
