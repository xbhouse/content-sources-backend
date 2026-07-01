---
type: pull_request
number: 790
title: "chore(ReviewSystems): fix filters breaking the page"
state: merged
author: mkholjuraev
created: 2022-04-26T12:39:55Z
updated: 2022-05-17T08:50:49Z
closed: 2022-04-27T07:56:48Z
merged: 2022-04-27T07:56:48Z
base_branch: master
head_branch: integrate-patch-set
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/790
---

# Pull Request #790: chore(ReviewSystems): fix filters breaking the page

**Author**: @mkholjuraev
**Created**: April 26, 2022 at 12:39 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `integrate-patch-set`

## Description

This resolves: https://issues.redhat.com/browse/SPM-1461.

To repoduce the issue: 

1. Open the patch set wizard
2. Proceed to Review Systems step
3. Active multiple filter with a light speed
4. Hit on the 'Reser filters' button. At this step, the wizard should break and close.


---

## Discussion

### Comment by @mkholjuraev on April 27, 2022 at 07:49 AM UTC

@RedHatInsights/team-interact I need to make a new release before code freeze. Thus I am merging this PR (as it is only a small fix) if you do not mind. 

### Comment by @jiridostal on May 16, 2022 at 03:12 PM UTC

:tada: This PR is included in version 1.47.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.47.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.47.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/790*
