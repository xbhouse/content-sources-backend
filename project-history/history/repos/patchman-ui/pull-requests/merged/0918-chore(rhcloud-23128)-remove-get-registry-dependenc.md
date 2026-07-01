---
type: pull_request
number: 918
title: "chore(RHCLOUD-23128): Remove get registry dependency and use create store directly"
state: merged
author: karelhala
created: 2022-11-23T14:26:44Z
updated: 2022-11-28T19:59:04Z
closed: 2022-11-28T10:30:59Z
merged: 2022-11-28T10:30:59Z
base_branch: master
head_branch: remove-get-registry
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/918
---

# Pull Request #918: chore(RHCLOUD-23128): Remove get registry dependency and use create store directly

**Author**: @karelhala
**Created**: November 23, 2022 at 02:26 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remove-get-registry`

## Description

# Description

Associated Jira ticket: #RHCLOUD-23128

The function `getRegistry` is no longer needed, on contrary with module federation it can bring additonal problems because of incorrect access of a store instance. This PR removes the usage of said function and replaces it with proper `store.replaceReducer` function.

# How to test the PR

Run patch locally and test if everything works as before (pay closer attention to pages with inventory in them).

---

## Discussion

### Comment by @jiridostal on November 28, 2022 at 07:58 PM UTC

:tada: This PR is included in version 1.57.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.57.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.57.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Commented on November 23, 2022 at 03:00 PM UTC

### Review by @mkholjuraev - Commented on November 23, 2022 at 03:07 PM UTC

### Review by @karelhala - Commented on November 23, 2022 at 03:09 PM UTC

### Review by @gkarat - Commented on November 23, 2022 at 03:19 PM UTC

### Review by @gkarat - Commented on November 23, 2022 at 03:21 PM UTC

### Review by @mkholjuraev - Commented on November 23, 2022 at 03:32 PM UTC

### Review by @karelhala - Commented on November 23, 2022 at 04:10 PM UTC

### Review by @mkholjuraev - Approved on November 28, 2022 at 10:30 AM UTC

LGTM! Works as expected. I do not see any issues while testing. Just to be sure, I checked against recent reload issues, and I did not see it happening. Thanks

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/918*
