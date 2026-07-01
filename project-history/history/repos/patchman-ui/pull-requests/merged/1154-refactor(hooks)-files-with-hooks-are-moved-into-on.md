---
type: pull_request
number: 1154
title: "refactor(hooks): files with hooks are moved into one folder"
state: merged
author: mkholjuraev
created: 2024-01-05T16:05:59Z
updated: 2024-01-18T10:30:42Z
closed: 2024-01-18T10:10:04Z
merged: 2024-01-18T10:10:04Z
base_branch: master
head_branch: rhineng-6147
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1154
---

# Pull Request #1154: refactor(hooks): files with hooks are moved into one folder

**Author**: @mkholjuraev
**Created**: January 05, 2024 at 04:05 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `rhineng-6147`

## Description

# Description

Associated Jira ticket: # (issue)

As we work on the patch codebase, hooks are being created in new files under /utilities with other types of resources. This leads to an unclean folder. This refactors files with hooks into **/hooks** folder. This also helps to import hooks from a single path. I am planning to create a new hook for RHINENG-6147 and this PR helps to do some cleanup refactor beforehand


# How to test the PR

Please play with the UI in depth, there should be no change in the behavior

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

### Comment by @codecov-commenter on January 05, 2024 at 04:15 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1154?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `2 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`ea62d3d`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ea62d3db8ac9197fa7de867829c621c22a459931?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.21% compared to head [(`62c226b`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1154?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.21%.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1154?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/hooks/useRemediationDataProvider.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1154?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VSZW1lZGlhdGlvbkRhdGFQcm92aWRlci5qcw==) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1154?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1154   +/-   ##
=======================================
  Coverage   63.21%   63.21%           
=======================================
  Files         122      122           
  Lines        3107     3107           
  Branches      809      809           
=======================================
  Hits         1964     1964           
  Misses       1143     1143           
```



</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1154?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on January 09, 2024 at 02:34 PM UTC

/retest

### Comment by @mkholjuraev on January 11, 2024 at 02:49 PM UTC

/retest

### Comment by @mkholjuraev on January 18, 2024 at 09:37 AM UTC

/retest

### Comment by @mkholjuraev on January 18, 2024 at 10:30 AM UTC

:tada: This PR is included in version 1.65.3 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.65.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Commented on January 15, 2024 at 05:03 PM UTC

### Review by @adonispuente - Approved on January 15, 2024 at 05:05 PM UTC

Overall a neat refactor, Id just make the one small change addressed in the comment but overall LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1154*
