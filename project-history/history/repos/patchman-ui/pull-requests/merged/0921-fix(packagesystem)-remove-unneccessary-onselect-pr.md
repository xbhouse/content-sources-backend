---
type: pull_request
number: 921
title: "fix(PackageSystem): remove unneccessary onSelect prop"
state: merged
author: mkholjuraev
created: 2022-11-29T14:08:15Z
updated: 2022-11-30T09:53:56Z
closed: 2022-11-30T09:35:55Z
merged: 2022-11-30T09:35:55Z
base_branch: master
head_branch: master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/921
---

# Pull Request #921: fix(PackageSystem): remove unneccessary onSelect prop

**Author**: @mkholjuraev
**Created**: November 29, 2022 at 02:08 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `master`

## Description

# Description

Associated Jira ticket: No ticket needed
Removes unnecessary onSelect prop from Package systems table. This will prevent unwanted checkboxes on an empty table.

# How to test the PR

1. Visit the package page
2. click on a package to open details
3. search for non existing system to get empty table
4. observe that there are no any checkbox in the table rows


# Before the change
![image](https://user-images.githubusercontent.com/59481011/204550620-43a92eaa-7a91-4cae-9ae8-ea9770c70baf.png)


# After the change

![image](https://user-images.githubusercontent.com/59481011/204550715-bc65a6b8-ce61-45a3-a0fb-6812bafca3a3.png)

# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @jiridostal on November 30, 2022 at 09:53 AM UTC

:tada: This PR is included in version 1.57.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.57.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.57.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on November 30, 2022 at 09:32 AM UTC

LGTM! Thank you @mkholjuraev!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/921*
