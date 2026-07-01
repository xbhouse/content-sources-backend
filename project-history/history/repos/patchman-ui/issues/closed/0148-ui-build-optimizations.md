---
type: issue
number: 148
title: "UI build optimizations"
state: closed
author: Hyperkid123
created: 2020-06-04T09:16:34Z
updated: 2020-06-08T11:20:49Z
closed: 2020-06-05T11:41:17Z
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/issues/148
---

# Issue #148: UI build optimizations

**Author**: @Hyperkid123
**Created**: June 04, 2020 at 09:16 AM UTC
**Status**: Closed
**Labels**: `released`

## Description

@karelhala is currently working optimizing the inventory component and he noticed that the build size of the patchman UI is very large even though the number of used components fairly small.

I would like to update some configuration files in this project which should considerably decrease the final build size output, improve performance, and allow @karelehala to do some more optimization in the inventory component.

### What would the changes be?
- change the way how certain modules (PF, FCE) are imported
  - either directly change import paths in code
  - or transform these imports via babel (preferred option)
- fiddle around with webpack config to split the node_modules chunk
  - there is no need to consolidate all node_modules in one massive chunk (it increases the initial entry point load and evaluation time)
  - there can be multiple vendors chunk for each lazy-loaded part of the app

cc @jiridostal
 

---

## Discussion

### Comment by @jiridostal on June 08, 2020 at 11:20 AM UTC

:tada: This issue has been resolved in version 0.12.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.12.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.12.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/issues/148*
