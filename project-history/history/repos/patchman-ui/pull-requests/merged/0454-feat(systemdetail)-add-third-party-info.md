---
type: pull_request
number: 454
title: "feat(SystemDetail): add third party info"
state: merged
author: mkholjuraev
created: 2021-03-03T22:37:29Z
updated: 2021-08-09T06:56:31Z
closed: 2021-03-05T09:31:30Z
merged: 2021-03-05T09:31:30Z
base_branch: master
head_branch: third-party
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/454
---

# Pull Request #454: feat(SystemDetail): add third party info

**Author**: @mkholjuraev
**Created**: March 03, 2021 at 10:37 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `third-party`

## Description

resolves: https://issues.redhat.com/browse/SPM-676

Currently, we have to make an extra backend request to get third-party info from the Patch engine even if we have all the rest of the data fetched by Inventory.

If InventoryDetailHead had getEntities feature enabled as InventoryTable component, we would just populate everything with one call. I suggest discussing it with the FEC team.

---

## Discussion

### Comment by @jiridostal on March 05, 2021 at 09:44 AM UTC

:tada: This PR is included in version 1.12.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.12.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.12.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/454*
