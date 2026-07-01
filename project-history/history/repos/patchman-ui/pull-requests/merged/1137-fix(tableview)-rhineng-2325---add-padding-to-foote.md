---
type: pull_request
number: 1137
title: "fix(TableView):  RHINENG-2325 - Add padding to footer via TableToolbar"
state: merged
author: bastilian
created: 2023-10-27T11:43:50Z
updated: 2023-10-30T12:11:17Z
closed: 2023-10-30T11:49:57Z
merged: 2023-10-30T11:49:57Z
base_branch: master
head_branch: RHINENG-2325
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1137
---

# Pull Request #1137: fix(TableView):  RHINENG-2325 - Add padding to footer via TableToolbar

**Author**: @bastilian
**Created**: October 27, 2023 at 11:43 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `RHINENG-2325`

## Description

# Description

Associated Jira ticket: #RHINENG-2325

The padding via the TableToolbar is required to avoid colliding with the "lightbulb"

# How to test the PR

Verify the pagination on any table does not get obscured by the lightbulb

# Before the change
![Screenshot 2023-10-27 at 13 41 34](https://github.com/RedHatInsights/patchman-ui/assets/7757/d221fa6c-3246-4f01-9b87-e9631940d0d4)

# After the change

![Screenshot 2023-10-27 at 13 42 49](https://github.com/RedHatInsights/patchman-ui/assets/7757/fb72bf1d-9f84-4deb-bd38-14333eca4d7a)


---

## Discussion

### Comment by @bastilian on October 27, 2023 at 02:39 PM UTC

/retest

### Comment by @mkholjuraev on October 30, 2023 at 12:10 PM UTC

:tada: This PR is included in version 1.64.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.64.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on October 30, 2023 at 08:17 AM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1137*
