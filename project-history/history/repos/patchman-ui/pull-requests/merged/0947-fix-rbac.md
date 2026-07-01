---
type: pull_request
number: 947
title: "Fix rbac"
state: merged
author: mkholjuraev
created: 2023-01-16T21:35:06Z
updated: 2024-04-03T09:22:49Z
closed: 2023-01-18T15:32:17Z
merged: 2023-01-18T15:32:17Z
base_branch: master
head_branch: fix-rbac
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/947
---

# Pull Request #947: Fix rbac

**Author**: @mkholjuraev
**Created**: January 16, 2023 at 09:35 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-rbac`

## Description

# Description

Associated Jira ticket: # [(SPM-1852)](https://issues.redhat.com/browse/SPM-1852)

Fixes RBAC issue. Previously, a user was required to have both admins and 'parch read' access. Now read review is removed.


# How to test the PR

Steps to Reproduce:
1. Navigate to 'https://console.redhat.com/settings/rbac/groups'
5. Edit corresponding group to add Patch Administrator role and remove Patch Viewer role

Actual results:
Navigate to 'https://console.redhat.com/insights/patch/advisories' which will result in below error:
You do not have access to patch

Expected results:
Access to Patch with only the Patch Administrator role (without the Patch Viewer also needed)

Additional info:
The other Insights apps are working and can be used as a cross reference/comparison

# Before the change


# After the change


# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @Fewwy on January 17, 2023 at 04:23 PM UTC

Tested it locally, LGTM

### Comment by @gkarat on January 18, 2023 at 02:42 PM UTC

/retest

### Comment by @mkholjuraev on January 18, 2023 at 03:45 PM UTC

:tada: This PR is included in version 1.58.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.58.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.58.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @Fewwy - Approved on January 17, 2023 at 04:23 PM UTC

### Review by @gkarat - Approved on January 18, 2023 at 02:38 PM UTC

LGTM! tested on my account with 2 users: one org admin, one default basic one. by moving roles here and there, it turns out to be working well now. Patch admin role is now sufficient to see Patch

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/947*
