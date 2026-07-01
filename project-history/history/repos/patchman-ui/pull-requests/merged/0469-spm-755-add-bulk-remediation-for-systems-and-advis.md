---
type: pull_request
number: 469
title: "SPM-755 Add bulk remediation for systems and advisories"
state: merged
author: jiridostal
created: 2021-03-18T10:56:38Z
updated: 2022-08-02T08:38:38Z
closed: 2021-03-25T10:23:57Z
merged: 2021-03-25T10:23:57Z
base_branch: master
head_branch: bulkRemedation
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/469
---

# Pull Request #469: SPM-755 Add bulk remediation for systems and advisories

**Author**: @jiridostal
**Created**: March 18, 2021 at 10:56 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `bulkRemedation`

## Description

This PR implement the Patch side of bulk remediation.
**There is a bug that prevents the remediation from creating - the wizard goes to the last step and than fails, platform is working on it**

The purpose is to demonstrate the behavior on Patch side.

---

## Discussion

### Comment by @jiridostal on March 18, 2021 at 01:24 PM UTC

Fixed!

### Comment by @jiridostal on March 25, 2021 at 02:47 PM UTC

:tada: This PR is included in version 1.14.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.14.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.14.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Changes Requested on March 18, 2021 at 01:20 PM UTC

I am afraid the PatchRemediationButton component is missing from the PR commit.

### Review by @mkholjuraev - Commented on March 18, 2021 at 02:46 PM UTC

### Review by @mkholjuraev - Commented on March 18, 2021 at 02:47 PM UTC

It looks that 2 remediation modals are being mounted in the Systems page, when I close one model, the other is displayed. Otherwise working as described.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/469*
