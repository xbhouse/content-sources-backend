---
type: pull_request
number: 1183
title: "fix: double icons in  display name column on package systems table"
state: merged
author: mkholjuraev
created: 2024-05-10T09:37:01Z
updated: 2024-05-22T10:44:22Z
closed: 2024-05-22T10:20:44Z
merged: 2024-05-22T10:20:44Z
base_branch: master
head_branch: fix-double-icons
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1183
---

# Pull Request #1183: fix: double icons in  display name column on package systems table

**Author**: @mkholjuraev
**Created**: May 10, 2024 at 09:37 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-double-icons`

## Description

# Description

Associated Jira ticket: No ticket needed

A small fix for the double icons on the package systems page. As the page completely uses the inventory table and also there were extra column definition, double icons appeared. Removing the extra one fixes it.

I am not sure even if this icon should be displayed in package systems table. I will consult it with the ux.

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

### Comment by @codecov-commenter on May 13, 2024 at 09:23 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1183?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.83%. Comparing base [(`42e843c`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/42e843cd45f2fa8abd04d8ccd7a64b008913e1a6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`4ce945f`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1183?dropdown=coverage&src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.


<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1183      +/-   ##
==========================================
- Coverage   63.84%   63.83%   -0.02%     
==========================================
  Files         124      124              
  Lines        3225     3224       -1     
  Branches      826      826              
==========================================
- Hits         2059     2058       -1     
  Misses       1166     1166              
```



</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1183?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on May 22, 2024 at 10:44 AM UTC

:tada: This PR is included in version 1.67.6 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.6)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Commented on May 10, 2024 at 02:51 PM UTC

### Review by @mkholjuraev - Commented on May 13, 2024 at 09:13 AM UTC

### Review by @mkholjuraev - Commented on May 13, 2024 at 09:13 AM UTC

### Review by @gkarat - Approved on May 22, 2024 at 09:47 AM UTC

LGTM, thanks @mkholjuraev 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1183*
