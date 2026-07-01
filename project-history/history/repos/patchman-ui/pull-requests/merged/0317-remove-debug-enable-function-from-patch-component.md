---
type: pull_request
number: 317
title: "Remove debug enable function from patch component mount"
state: merged
author: karelhala
created: 2020-10-23T06:46:56Z
updated: 2020-10-26T10:05:08Z
closed: 2020-10-26T09:28:23Z
merged: 2020-10-26T09:28:23Z
base_branch: master
head_branch: remove-enable-global-filter
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/317
---

# Pull Request #317: Remove debug enable function from patch component mount

**Author**: @karelhala
**Created**: October 23, 2020 at 06:46 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remove-enable-global-filter`

## Description

### No need to use `insights.chrome.enable.globalFilter();`

When navigating to patch and then to any other app it shows the global filter no matter if we hid it or not. This is caused by using `insights.chrome.enable.globalFilter();`, this function should be used only for local debugging and not stored in your code, it'll set localStorage value which enables global filter everywhere.

---

## Discussion

### Comment by @karelhala on October 23, 2020 at 06:47 AM UTC

ping @mkholjuraev @jiridostal

### Comment by @jiridostal on October 26, 2020 at 10:05 AM UTC

:tada: This PR is included in version 0.27.11 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.27.11)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.27.11)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/317*
