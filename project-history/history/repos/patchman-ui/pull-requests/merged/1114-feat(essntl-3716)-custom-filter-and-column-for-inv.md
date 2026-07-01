---
type: pull_request
number: 1114
title: "feat(ESSNTL-3716): custom filter and column for inv groups"
state: merged
author: Fewwy
created: 2023-08-28T17:24:59Z
updated: 2023-08-30T19:42:40Z
closed: 2023-08-30T19:42:40Z
merged: 2023-08-30T19:42:40Z
base_branch: master
head_branch: custom-filter-for-groups
labels: ["enhancement"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1114
---

# Pull Request #1114: feat(ESSNTL-3716): custom filter and column for inv groups

**Author**: @Fewwy
**Created**: August 28, 2023 at 05:24 PM UTC
**Status**: Merged
**Labels**: `enhancement`
**Base**: `master` ← **Head**: `custom-filter-for-groups`

## Description

# Description

Associated Jira ticket: # (issue)
https://issues.redhat.com/browse/ESSNTL-3716

Please include a summary of the change, what this fixes/creates/improves.
1)groups filter
2)groups column
3)filters are added to the url in format "group_name%5D=in%NAMEOFTHEGROUP", if you filter by several groups they should be separated by comma.
4) Use unleash feature flag to check if the filter and column are hidden if the feature flag is off


# How to test the PR
check that filter works and column is populated


# After the change
![image](https://github.com/RedHatInsights/patchman-ui/assets/62722417/8db3f3ec-8fc6-4961-86ec-5cafbc3c4fec)

![image](https://github.com/RedHatInsights/patchman-ui/assets/62722417/7a4dda8f-b9d7-4a7c-8fb8-9aa0b138007b)



---

## Discussion

### Comment by @codecov-commenter on August 30, 2023 at 05:27 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1114?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage is **`51.28%`** of modified lines.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1114?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage |
|---|---|
| [src/Utilities/Hooks.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1114?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `ø` |
| [src/Utilities/useFeatureFlag.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1114?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91c2VGZWF0dXJlRmxhZy5qcw==) | `0.00%` |
| [...c/PresentationalComponents/Filters/GroupsFilter.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1114?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL0dyb3Vwc0ZpbHRlci5qcw==) | `43.75%` |
| [src/Utilities/SystemsHelpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1114?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | `58.33%` |
| [src/Utilities/api.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1114?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `66.66%` |
| [src/SmartComponents/Systems/Systems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1114?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `100.00%` |
| [src/Utilities/constants.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1114?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00%` |


:loudspeaker: Thoughts on this report? [Let us know!](https://about.codecov.io/pull-request-comment-report/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @mkholjuraev - Commented on August 29, 2023 at 06:51 AM UTC

### Review by @Fewwy - Commented on August 29, 2023 at 02:18 PM UTC

### Review by @adonispuente - Commented on August 29, 2023 at 04:54 PM UTC

If I open up the drop down for the groups, If i add two groups and then remove them, I am greeted with no systems/groups in the table when I should see the table filled as it is on page load. It works with reset filters, but not when unclicking items in the dropdown

### Review by @mkholjuraev - Commented on August 30, 2023 at 04:03 AM UTC

A few nitpicks with the issue @adonispuente reported. Then we are good to go!

### Review by @gkarat - Commented on August 30, 2023 at 11:09 AM UTC

### Review by @Fewwy - Commented on August 30, 2023 at 05:10 PM UTC

### Review by @Fewwy - Commented on August 30, 2023 at 05:10 PM UTC

### Review by @Fewwy - Commented on August 30, 2023 at 05:18 PM UTC

### Review by @Fewwy - Commented on August 30, 2023 at 05:19 PM UTC

### Review by @adonispuente - Dismissed on August 30, 2023 at 06:01 PM UTC

great job LGTM!

### Review by @gkarat - Approved on August 30, 2023 at 07:30 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1114*
