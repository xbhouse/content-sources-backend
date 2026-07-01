---
type: pull_request
number: 1116
title: "fix(ESSNTL-4494): Fix group column and filter"
state: merged
author: gkarat
created: 2023-08-31T20:59:25Z
updated: 2023-09-01T08:44:10Z
closed: 2023-09-01T08:24:20Z
merged: 2023-09-01T08:24:20Z
base_branch: master
head_branch: essntl-4494
labels: ["bug", "enhancement", "released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1116
---

# Pull Request #1116: fix(ESSNTL-4494): Fix group column and filter

**Author**: @gkarat
**Created**: August 31, 2023 at 08:59 PM UTC
**Status**: Merged
**Labels**: `bug`, `enhancement`, `released`
**Base**: `master` ← **Head**: `essntl-4494`

## Description

This removes the custom implementation of group filter and makes sure the column and filter are working according to requirements.

## How to test

Try to do basic filtering with groups, verify the group column: 'N/A' for hosts without any group and a group name for hosts that belong to any of account's groups.

## Screenshots

<img width="1271" alt="Screenshot 2023-08-31 at 22 59 13" src="https://github.com/RedHatInsights/patchman-ui/assets/31385370/59563a86-51ab-4940-839e-b5b1db40bb1a">


---

## Discussion

### Comment by @mkholjuraev on September 01, 2023 at 08:44 AM UTC

:tada: This PR is included in version 1.64.1 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.64.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Approved on August 31, 2023 at 10:24 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1116*
