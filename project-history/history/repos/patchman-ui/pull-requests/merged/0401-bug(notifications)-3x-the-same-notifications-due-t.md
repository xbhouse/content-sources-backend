---
type: pull_request
number: 401
title: "bug(Notifications): 3x the same notifications due to multiple backend\u2026"
state: merged
author: mkholjuraev
created: 2021-01-25T07:50:52Z
updated: 2021-08-09T06:54:58Z
closed: 2021-01-25T09:41:03Z
merged: 2021-01-25T09:41:03Z
base_branch: master
head_branch: 3x-notification
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/401
---

# Pull Request #401: bug(Notifications): 3x the same notifications due to multiple backend…

**Author**: @mkholjuraev
**Created**: January 25, 2021 at 07:50 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `3x-notification`

## Description

resolves: https://issues.redhat.com/browse/SPM-701

There was 3-2c calls to backend on load and this was causing extra notifications

---

## Discussion

### Comment by @jiridostal on January 25, 2021 at 01:21 PM UTC

:tada: This PR is included in version 1.8.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.8.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.8.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/401*
