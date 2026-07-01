---
type: pull_request
number: 914
title: "fix(SPM-1802): pass correct remediation issue ID"
state: merged
author: mkholjuraev
created: 2022-11-11T09:43:56Z
updated: 2024-04-03T09:21:56Z
closed: 2022-11-14T11:24:45Z
merged: 2022-11-14T11:24:45Z
base_branch: master
head_branch: fix-package-systems-rem
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/914
---

# Pull Request #914: fix(SPM-1802): pass correct remediation issue ID

**Author**: @mkholjuraev
**Created**: November 11, 2022 at 09:43 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-package-systems-rem`

## Description

Fixes: [SPM-1802](https://issues.redhat.com/browse/SPM-1802) by sending the expected remediation issue ID according to the new validation.

To test:

1. Visit the packages page
2. click on any package name and open the details 
3. select some systems which have updatable package evra version
4. observe that the remediation modal is loaded & works as expected.  

---

## Discussion

### Comment by @mkholjuraev on November 14, 2022 at 11:24 AM UTC

@bastilian thank you for the review. This lets me release the Patch UI with the fix now.

### Comment by @jiridostal on November 14, 2022 at 11:42 AM UTC

:tada: This PR is included in version 1.55.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.55.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.55.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on November 14, 2022 at 11:10 AM UTC

Approved. Thank you @mkholjuraev!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/914*
