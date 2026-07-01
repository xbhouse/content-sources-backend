---
type: pull_request
number: 874
title: "fix(InventoryTable): avoids broken systems page by using different pr\u2026"
state: closed
author: mkholjuraev
created: 2022-09-29T14:23:26Z
updated: 2024-04-03T09:22:07Z
closed: 2022-10-11T09:42:44Z
base_branch: master
head_branch: fix-inventory
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/874
---

# Pull Request #874: fix(InventoryTable): avoids broken systems page by using different pr…

**Author**: @mkholjuraev
**Created**: September 29, 2022 at 02:23 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `fix-inventory`

## Description

This PR fixes the broken systems page by using a different prop for the tags modal. The reloading inventory issue still exists.

To test:

1. run this PR against stage/beta
2. navigate to systems page
3. leave this open idle for 2-3 mins
4. observe that the systems table gets reloaded, but the page does not brake.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/874*
