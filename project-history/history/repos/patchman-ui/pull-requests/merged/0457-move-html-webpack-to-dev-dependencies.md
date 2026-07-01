---
type: pull_request
number: 457
title: "Move html webpack to dev dependencies"
state: merged
author: Hyperkid123
created: 2021-03-09T13:19:25Z
updated: 2021-03-11T09:59:31Z
closed: 2021-03-09T13:56:38Z
merged: 2021-03-09T13:56:38Z
base_branch: master
head_branch: fix-html-webpack-dependency
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/457
---

# Pull Request #457: Move html webpack to dev dependencies

**Author**: @Hyperkid123
**Created**: March 09, 2021 at 01:19 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-html-webpack-dependency`

## Description

Including the `html-webpack-plugin` in dependencies can cause issues with applications using v5 of this plugin. Moving it to dev dependencies will prevent the duplicate package from installing.

We also have to release a new version of the inventory component as well. Thanks!

cc @karelhala @jiridostal 

---

## Discussion

### Comment by @jiridostal on March 11, 2021 at 09:59 AM UTC

:tada: This PR is included in version 1.12.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.12.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.12.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on March 09, 2021 at 01:56 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/457*
