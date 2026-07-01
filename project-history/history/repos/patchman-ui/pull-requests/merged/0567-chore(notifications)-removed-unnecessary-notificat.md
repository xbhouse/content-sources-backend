---
type: pull_request
number: 567
title: "chore(notifications): removed unnecessary notifications and some cleanup"
state: merged
author: mkholjuraev
created: 2021-06-10T08:39:57Z
updated: 2021-08-09T06:56:55Z
closed: 2021-06-11T09:00:26Z
merged: 2021-06-11T09:00:26Z
base_branch: master
head_branch: delete-unused-notifications
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/567
---

# Pull Request #567: chore(notifications): removed unnecessary notifications and some cleanup

**Author**: @mkholjuraev
**Created**: June 10, 2021 at 08:39 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `delete-unused-notifications`

## Description

Notifications middleware is already displaying an error notification every time there is an error. So, duplicate notification was being displayed from Patch.

![image](https://user-images.githubusercontent.com/59481011/121493516-312bb900-c9d8-11eb-9a14-029de6cd768b.png)


---

## Discussion

### Comment by @jiridostal on June 15, 2021 at 08:03 AM UTC

:tada: This PR is included in version 1.20.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.20.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.20.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/567*
