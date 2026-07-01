---
type: pull_request
number: 611
title: "feat(webpack): SPM-1062 Switch to webpack proxy and enable auto reload"
state: merged
author: jiridostal
created: 2021-07-30T09:42:42Z
updated: 2022-11-16T08:15:29Z
closed: 2021-08-02T10:04:54Z
merged: 2021-08-02T10:04:53Z
base_branch: master
head_branch: webpack-proxy
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/611
---

# Pull Request #611: feat(webpack): SPM-1062 Switch to webpack proxy and enable auto reload

**Author**: @jiridostal
**Created**: July 30, 2021 at 09:42 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `webpack-proxy`

## Description

Please make sure it works, it may break the dev setup. Works for me.

---

## Discussion

### Comment by @mkholjuraev on July 30, 2021 at 01:46 PM UTC

Autorefresh sounds good! But when I run it locally, I have some proxy issues.
![image](https://user-images.githubusercontent.com/59481011/127662319-a71e8d94-dad2-4489-aade-ebe15676af78.png)


### Comment by @jiridostal on July 31, 2021 at 09:03 PM UTC

turn off proxy and run the app with "npm run start:proxy:beta" and try again

### Comment by @jiridostal on August 02, 2021 at 10:20 AM UTC

:tada: This PR is included in version 1.28.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.28.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.28.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on August 02, 2021 at 10:04 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/611*
