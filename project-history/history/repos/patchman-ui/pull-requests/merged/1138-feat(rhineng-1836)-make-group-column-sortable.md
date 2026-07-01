---
type: pull_request
number: 1138
title: "feat(RHINENG-1836): Make group column sortable"
state: merged
author: gkarat
created: 2023-11-01T16:09:08Z
updated: 2023-11-09T11:01:07Z
closed: 2023-11-09T10:40:05Z
merged: 2023-11-09T10:40:05Z
base_branch: master
head_branch: rhineng-1836
labels: ["enhancement", "released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1138
---

# Pull Request #1138: feat(RHINENG-1836): Make group column sortable

**Author**: @gkarat
**Created**: November 01, 2023 at 04:09 PM UTC
**Status**: Merged
**Labels**: `enhancement`, `released`
**Base**: `master` ← **Head**: `rhineng-1836`

## Description

## Description

Associated Jira ticket: https://issues.redhat.com/browse/RHINENG-1836.

This makes the "Group" column sortable on two views: /patch/systems and /patch/advisories/%id.

## How to test the PR

The Inventory change is not yet merged to ensure that the tenant apps support sorting first and avoid breaking change.

1. Pull this Inventory PR locally https://github.com/RedHatInsights/insights-inventory-frontend/pull/2070.
2. Run Patch with this PR `npm run start -- --port=8004`
3. Run Inventory with `LOCAL_APPS=patch:8004~http npm run start:proxy`
4. Navigate to the updated views (in stage-stable environment) and make sure you can sort by system groups.
5. Check the URL search parameters too: while sorting the table, and then refreshing the page - the same sorting must be applied.

## After the change

<img width="1512" alt="Screenshot 2023-11-01 at 17 03 55" src="https://github.com/RedHatInsights/patchman-ui/assets/31385370/3fc3f8ec-e38d-416e-aee0-4b4656bdee58">

<img width="1512" alt="Screenshot 2023-11-01 at 17 04 10" src="https://github.com/RedHatInsights/patchman-ui/assets/31385370/27cd0b64-2dcd-406c-b726-746dfc4497e4">

## Dependent work link

https://issues.redhat.com/browse/RHINENG-1658

## Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [x] README.md is updated if necessary
- [x] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on November 01, 2023 at 04:18 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1138?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1138?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/SystemsHelpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1138?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | `72.97% <100.00%> (+44.84%)` | :arrow_up: |


:loudspeaker: Thoughts on this report? [Let us know!](https://about.codecov.io/pull-request-comment-report/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @gkarat on November 02, 2023 at 12:32 PM UTC

Looks like sort parameter from URL is not applied when refreshing

### Comment by @Fewwy on November 09, 2023 at 10:22 AM UTC

I don't think sorting works correctly
![image](https://github.com/RedHatInsights/patchman-ui/assets/62722417/e87d22c3-316b-4ea2-a7b7-f364cd6cc380)
![image](https://github.com/RedHatInsights/patchman-ui/assets/62722417/f610d2ce-5121-4f0a-8544-ec62e3f177d1)


### Comment by @mkholjuraev on November 09, 2023 at 11:00 AM UTC

:tada: This PR is included in version 1.65.0 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.65.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @Fewwy - Approved on November 09, 2023 at 10:26 AM UTC

Other than this issue - LGTM

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1138*
