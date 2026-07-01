---
type: pull_request
number: 435
title: "Fix store already initialized error."
state: merged
author: Hyperkid123
created: 2021-02-12T10:39:38Z
updated: 2021-02-19T10:25:30Z
closed: 2021-02-12T13:11:34Z
merged: 2021-02-12T13:11:34Z
base_branch: master
head_branch: registry-error
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/435
---

# Pull Request #435: Fix store already initialized error.

**Author**: @Hyperkid123
**Created**: February 12, 2021 at 10:39 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `registry-error`

## Description

Fixes chrome 2 redux compatibility issue.

Note that the systems table might still have an infinite loop. It will require some adjustments in the inventory component. We will fix it later.

cc @jiridostal @mkholjuraev 

---

## Discussion

### Comment by @jiridostal on February 19, 2021 at 10:25 AM UTC

:tada: This PR is included in version 1.11.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.11.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.11.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/435*
