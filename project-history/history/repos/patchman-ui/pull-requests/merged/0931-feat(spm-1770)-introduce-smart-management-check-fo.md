---
type: pull_request
number: 931
title: "feat(SPM-1770): introduce smart management check for Templates page"
state: merged
author: mkholjuraev
created: 2022-12-12T21:20:02Z
updated: 2024-04-03T09:21:54Z
closed: 2022-12-13T13:31:30Z
merged: 2022-12-13T13:31:30Z
base_branch: master
head_branch: smart-management
labels: ["released on @fix-semantice-release"]
url: https://github.com/RedHatInsights/patchman-ui/pull/931
---

# Pull Request #931: feat(SPM-1770): introduce smart management check for Templates page

**Author**: @mkholjuraev
**Created**: December 12, 2022 at 09:20 PM UTC
**Status**: Merged
**Labels**: `released on @fix-semantice-release`
**Base**: `master` ← **Head**: `smart-management`

## Description

# Description

Associated Jira ticket: # [SPM-1770](https://issues.redhat.com/browse/SPM-1770)

As a permissioned user who does not have a Smart Management entitlement, when I try to visit Patch templates, I should see that this is Smart Management functionality and be given an opportunity to do something about that, so that I can check it out and decide if I want to buy Smart Management entitlements to keep using this awesome feature. 

# How to test the PR
I will try to create an account without smart management entitlement for testing the PR. I will update once I can create.

- A user with proper permissions, but on an account WITHOUT a Smart Management entitlement, should be see an entitlement blocker page when visiting /insights/patch/templates.

- A user with proper permissions and on an account WITH a Smart Management entitlements, should be able to access /insights/patch/templates as normal.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on December 13, 2022 at 01:20 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/931?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **69.77**% // Head: **72.20**% // Increases project coverage by **`+2.43%`** :tada:
> Coverage data is based on head [(`a882746`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/931?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`c1f00fd`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/c1f00fdad3a4cf207c4fb54c9eed314b0815efcb?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 90.47% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #931      +/-   ##
==========================================
+ Coverage   69.77%   72.20%   +2.43%     
==========================================
  Files         111      111              
  Lines        2607     2623      +16     
  Branches      679      681       +2     
==========================================
+ Hits         1819     1894      +75     
+ Misses        712      658      -54     
+ Partials       76       71       -5     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/931?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/931/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `85.07% <84.61%> (+85.07%)` | :arrow_up: |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/931/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `71.42% <100.00%> (+4.76%)` | :arrow_up: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/931/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `74.50% <100.00%> (+1.04%)` | :arrow_up: |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/931/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `73.41% <0.00%> (+5.06%)` | :arrow_up: |
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/931/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `54.54% <0.00%> (+54.54%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/931?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on December 22, 2022 at 11:38 AM UTC

:tada: This PR is included in version 1.58.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.58.0)
- [npm package (@fix-semantice-release dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.58.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on December 12, 2022 at 10:14 PM UTC

PR works as intended, just has the typo in the message 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/931*
