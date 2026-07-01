---
type: pull_request
number: 328
title: "fix(ConditionalFilter): radio type filter empty value is fixed"
state: merged
author: mkholjuraev
created: 2020-11-12T09:16:43Z
updated: 2021-08-09T06:55:45Z
closed: 2020-11-13T10:25:35Z
merged: 2020-11-13T10:25:35Z
base_branch: master
head_branch: fix-filter-empty-string
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/328
---

# Pull Request #328: fix(ConditionalFilter): radio type filter empty value is fixed

**Author**: @mkholjuraev
**Created**: November 12, 2020 at 09:16 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-filter-empty-string`

## Description

not a complete solution [SPM-614](https://issues.redhat.com/browse/SPM-614) 

But using this approach we could avoid 
`    // Empty string value is not supported by PF Radio at the moment
    if (currentValue === '' || !currentValue) {
        currentValue = '0';
    }`

---

## Discussion

### Comment by @jiridostal on November 13, 2020 at 10:32 AM UTC

:tada: This PR is included in version 1.0.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.0.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.0.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/328*
