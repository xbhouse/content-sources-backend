---
type: pull_request
number: 414
title: "Fix systemdetail"
state: merged
author: mkholjuraev
created: 2021-01-28T15:09:03Z
updated: 2021-08-09T06:54:58Z
closed: 2021-02-01T21:56:20Z
merged: 2021-02-01T21:56:20Z
base_branch: master
head_branch: fix-systemdetail
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/414
---

# Pull Request #414: Fix systemdetail

**Author**: @mkholjuraev
**Created**: January 28, 2021 at 03:09 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-systemdetail`

## Description

resolves: https://issues.redhat.com/browse/SPM-705

Following lines are [deleted](https://github.com/RedHatInsights/patchman-ui/blob/master/src/store/Reducers/SystemDetailStore.js#L27) as this bug is fixed in FEC:
      //we have to send empty object if there is not entity so that inventory does not throw error
      entity: state.entity && state.entity || {},

---

## Discussion

### Comment by @jiridostal on February 01, 2021 at 10:04 PM UTC

:tada: This PR is included in version 1.8.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.8.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.8.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/414*
