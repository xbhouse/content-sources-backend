---
type: pull_request
number: 789
title: "chore(InventoryDetail) remove feature flag, show patch set only when it is present"
state: merged
author: mkholjuraev
created: 2022-04-26T09:45:07Z
updated: 2022-05-16T15:12:49Z
closed: 2022-04-26T12:29:59Z
merged: 2022-04-26T12:29:59Z
base_branch: master
head_branch: integrate-patch-set
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/789
---

# Pull Request #789: chore(InventoryDetail) remove feature flag, show patch set only when it is present

**Author**: @mkholjuraev
**Created**: April 26, 2022 at 09:45 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `integrate-patch-set`

## Description

I had discussion with UX and they agreed (for now) to remove completely patch set info on System detail header when it is not present.

Also this PR removes patch set flag **Assign to patch set** on system detail actions. 

---

## Discussion

### Comment by @jiridostal on May 16, 2022 at 03:12 PM UTC

:tada: This PR is included in version 1.47.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.47.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.47.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gabipodolnikova - Approved on April 26, 2022 at 12:25 PM UTC

Yes, this is better than having a blank space. Thanks! :)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/789*
