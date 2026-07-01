---
type: pull_request
number: 1182
title: "fix(RHINENG-9035): bulk select all icon fixed"
state: merged
author: mkholjuraev
created: 2024-05-09T13:58:23Z
updated: 2024-05-13T12:02:02Z
closed: 2024-05-13T11:44:15Z
merged: 2024-05-13T11:44:15Z
base_branch: master
head_branch: rhineng-9035
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1182
---

# Pull Request #1182: fix(RHINENG-9035): bulk select all icon fixed

**Author**: @mkholjuraev
**Created**: May 09, 2024 at 01:58 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `rhineng-9035`

## Description

# Description

Associated Jira ticket: #944 tps://issues.redhat.com/browse/RHINENG-9035

Fixes the loading icon in bulk select by updating the fec-config package. It also includes tiny changes to cleanup bulk select config. The icons on systems tables will be fixed by https://github.com/RedHatInsights/insights-inventory-frontend/pull/2192

# How to test the PR
1. Navigate to Advisories table
2. Select all items in the table
3. Observe that there is now the spinner icon while items are being selected and selection dropdown is disabled

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

### Comment by @mkholjuraev on May 13, 2024 at 09:16 AM UTC

/retest

### Comment by @mkholjuraev on May 13, 2024 at 12:02 PM UTC

:tada: This PR is included in version 1.67.4 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on May 13, 2024 at 11:15 AM UTC

LGTM! Thank you @mkholjuraev!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1182*
