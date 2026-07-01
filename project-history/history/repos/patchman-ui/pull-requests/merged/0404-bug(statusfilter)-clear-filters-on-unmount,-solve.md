---
type: pull_request
number: 404
title: "bug(StatusFilter): clear filters on unmount, solve 2x upgradable label"
state: merged
author: mkholjuraev
created: 2021-01-26T08:49:35Z
updated: 2021-08-09T06:54:57Z
closed: 2021-02-02T13:56:51Z
merged: 2021-02-02T13:56:51Z
base_branch: master
head_branch: packages-filters
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/404
---

# Pull Request #404: bug(StatusFilter): clear filters on unmount, solve 2x upgradable label

**Author**: @mkholjuraev
**Created**: January 26, 2021 at 08:49 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `packages-filters`

## Description

Resolves: https://issues.redhat.com/browse/SPM-651 and https://issues.redhat.com/browse/SPM-697

---

## Discussion

### Comment by @mkholjuraev on January 26, 2021 at 09:03 AM UTC

Better to merge after https://issues.redhat.com/browse/SPM-703 is resolved.

### Comment by @mkholjuraev on February 01, 2021 at 11:23 PM UTC

The final solution is kind of hackish. SystemPackages and packages pages need to send their own filter object name and types.
I think it would be better if we have a short discussion about it later. 

In this commit, I removed clearing PackagesListStore as we have discussed.


### Comment by @jiridostal on February 02, 2021 at 10:15 AM UTC

Package detail page does not load with an error:
![image](https://user-images.githubusercontent.com/6286045/106585585-c6f59a00-6547-11eb-88a7-c75bb13a4a3a.png)

Also, isn't it nicer to have two separate filters than merging these two cases into one?

### Comment by @jiridostal on February 02, 2021 at 09:19 PM UTC

:tada: This PR is included in version 1.9.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.9.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.9.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/404*
