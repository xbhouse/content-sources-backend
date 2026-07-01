---
type: pull_request
number: 1092
title: "feat(ESSNTL-3716): Add Inventory groups column and filter"
state: closed
author: Fewwy
created: 2023-07-18T09:37:28Z
updated: 2023-09-01T08:42:31Z
closed: 2023-09-01T08:42:31Z
base_branch: master
head_branch: host-group-intergration
labels: ["enhancement"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1092
---

# Pull Request #1092: feat(ESSNTL-3716): Add Inventory groups column and filter

**Author**: @Fewwy
**Created**: July 18, 2023 at 09:37 AM UTC
**Status**: Closed
**Labels**: `enhancement`
**Base**: `master` ← **Head**: `host-group-intergration`

## Description

# Description
Add group filter and column hidden behind feature flag
Associated Jira ticket: # (issue)
https://issues.redhat.com/browse/ESSNTL-3716
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

### Comment by @codecov-commenter on July 19, 2023 at 11:28 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1092?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`7.69%`** and project coverage change: **`-0.25%`** :warning:
> Comparison is base [(`ba0fe18`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ba0fe18528df24cf52abc1d94296e7ee3e216d8e?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.63% compared to head [(`20a4275`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1092?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.39%.
> Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1092      +/-   ##
==========================================
- Coverage   62.63%   62.39%   -0.25%     
==========================================
  Files         119      119              
  Lines        2987     3002      +15     
  Branches      766      773       +7     
==========================================
+ Hits         1871     1873       +2     
- Misses       1116     1129      +13     
```


| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1092?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Systems/Systems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1092?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `80.59% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1092?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `80.45% <ø> (ø)` | |
| [src/Utilities/Hooks.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1092?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `73.02% <0.00%> (-0.49%)` | :arrow_down: |
| [src/Utilities/SystemsHelpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1092?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | `18.75% <0.00%> (-6.25%)` | :arrow_down: |
| [src/Utilities/constants.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1092?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <100.00%> (ø)` | |

... and [3 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1092/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1092?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on July 19, 2023 at 02:15 PM UTC

/retest

### Comment by @Fewwy on July 20, 2023 at 07:56 AM UTC

/retest

---

## Reviews

### Review by @leSamo - Changes Requested on July 18, 2023 at 11:04 AM UTC

When I try to disable feature flag:
![Screenshot from 2023-07-18 12-01-58](https://github.com/RedHatInsights/patchman-ui/assets/8426204/d3d198a2-58af-433b-a2a9-ad44b8aefc68)

When I try to filter:
[Screencast from 2023-07-18 11-59-08.webm](https://github.com/RedHatInsights/patchman-ui/assets/8426204/4c3a7152-a1d0-4ab5-8388-66e91c65275c)


### Review by @gkarat - Changes Requested on July 25, 2023 at 09:48 AM UTC

When I apply the group filter by choosing a group that has some hosts, I see that data is not filtered. However, the backend already has this filter implemented https://issues.redhat.com/browse/ESSNTL-4817?focusedId=22636009&page=com.atlassian.jira.plugin.system.issuetabpanels%3Acomment-tabpanel#comment-22636009. I also tried to edit and resend the request (to /api/patch/v3/system) with this parameter and it indeed filtered hosts. So let's support it here too within this PR.

The group column is mapped properly and for some hosts I see the link is rendered and it works correct ✅ 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1092*
