---
type: pull_request
number: 484
title: "Wrap inventory detail with wrapper"
state: merged
author: karelhala
created: 2021-03-31T13:32:39Z
updated: 2021-04-01T11:56:43Z
closed: 2021-03-31T17:49:43Z
merged: 2021-03-31T17:49:43Z
base_branch: master
head_branch: inventory-wrapper
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/484
---

# Pull Request #484: Wrap inventory detail with wrapper

**Author**: @karelhala
**Created**: March 31, 2021 at 01:32 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `inventory-wrapper`

## Description

### Use inventory wrapper

Inventory detail consists of multiple components, and they require direct access to redux. This PR changes the structure of system detail in order to merge inventory redux with patch's one.

---

## Discussion

### Comment by @karelhala on March 31, 2021 at 01:32 PM UTC

@mkholjuraev this PR fixes the issue we just spoke about.

### Comment by @mkholjuraev on March 31, 2021 at 02:00 PM UTC

@karelhala when I run this locally, I still see the issue. Should I do something more? And not only advisory detail, but also systems page also not loading

### Comment by @jiridostal on April 01, 2021 at 11:56 AM UTC

:tada: This PR is included in version 1.14.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.14.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.14.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on March 31, 2021 at 05:49 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/484*
