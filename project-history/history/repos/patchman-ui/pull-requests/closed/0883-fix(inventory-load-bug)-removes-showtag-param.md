---
type: pull_request
number: 883
title: "fix(inventory-load-bug): removes showTag param"
state: closed
author: adonispuente
created: 2022-10-07T19:19:53Z
updated: 2022-10-10T16:52:01Z
closed: 2022-10-10T16:52:01Z
base_branch: master
head_branch: inventory2
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/883
---

# Pull Request #883: fix(inventory-load-bug): removes showTag param

**Author**: @adonispuente
**Created**: October 07, 2022 at 07:19 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `inventory2`

## Description

Just for testing atm
Please test to see if you encounter the Inventory reload bug.

To test, head over to systems tab (anywhere theres an inventory component) 
- wait at least 3 minutes
- paginate to the next page


notes: if youre going to retest, paginate back to the original 20 results and refresh.
Desired result:
Inventory loads no problem

Attempting to fix:
Inventory component breaking on certain loads

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/883*
