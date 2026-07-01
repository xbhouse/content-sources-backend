---
type: pull_request
number: 1074
title: "fix(SPM-2055): Select all on disabled advisories"
state: merged
author: johnsonm325
created: 2023-06-02T14:53:54Z
updated: 2023-06-05T21:41:44Z
closed: 2023-06-02T19:06:39Z
merged: 2023-06-02T19:06:39Z
base_branch: master
head_branch: spm-2055
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1074
---

# Pull Request #1074: fix(SPM-2055): Select all on disabled advisories

**Author**: @johnsonm325
**Created**: June 02, 2023 at 02:53 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `spm-2055`

## Description

Steps to Reproduce:
Navigate to the Advosry details page
Apply status filter-select Applicable
then select all() from filter

This PR correctly avoids selecting systems with the status, "Applicable", which should be disabled. Also check this functionality by going to Inventory -> Systems -> choose a system -> Click Patch tab. This should show the same selection functionality.

This PR also fixes a bug not mentioned in the Jira card. There is a bug on Patch -> Systems where if you select a system and then you go to any page after page 1 and `Select all`, it will only select advisories on that page and after. It will not select any items on previous pages.

# Description

Associated Jira ticket: # (issue)

Please include a summary of the change, what this fixes/creates/improves.


# How to test the PR

Please include steps to test your PR.

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

### Comment by @mkholjuraev on June 05, 2023 at 09:41 PM UTC

:tada: This PR is included in version 1.63.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on June 02, 2023 at 04:39 PM UTC

codewise LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1074*
