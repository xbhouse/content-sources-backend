---
type: pull_request
number: 468
title: "fix(SonarCube): issues are resolved"
state: merged
author: mkholjuraev
created: 2021-03-17T13:52:20Z
updated: 2021-08-09T06:56:25Z
closed: 2021-03-19T10:07:54Z
merged: 2021-03-19T10:07:54Z
base_branch: master
head_branch: fix-sonarcube-issues
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/468
---

# Pull Request #468: fix(SonarCube): issues are resolved

**Author**: @mkholjuraev
**Created**: March 17, 2021 at 01:52 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-sonarcube-issues`

## Description

resolves: https://issues.redhat.com/browse/SPM-792

SonarCube reliability bugs are fixed. Maintainability bugs will be fixed in the next PR

---

## Discussion

### Comment by @jiridostal on March 19, 2021 at 10:16 AM UTC

:tada: This PR is included in version 1.13.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.13.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.13.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @jiridostal - Changes Requested on March 18, 2021 at 09:34 AM UTC

isSet function is not needed at all. Some resources:

- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Logical_OR
- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Logical_AND
- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Optional_chaining

### Review by @mkholjuraev - Commented on March 18, 2021 at 09:46 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/468*
