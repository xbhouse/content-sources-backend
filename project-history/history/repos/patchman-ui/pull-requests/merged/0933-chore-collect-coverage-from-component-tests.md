---
type: pull_request
number: 933
title: "chore: Collect coverage from component tests"
state: merged
author: gkarat
created: 2022-12-20T15:02:49Z
updated: 2023-01-18T15:46:01Z
closed: 2022-12-28T11:08:20Z
merged: 2022-12-28T11:08:20Z
base_branch: master
head_branch: collect-coverage-from-cypress-tests
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/933
---

# Pull Request #933: chore: Collect coverage from component tests

**Author**: @gkarat
**Created**: December 20, 2022 at 03:02 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `collect-coverage-from-cypress-tests`

## Description

This combines coverage reports from jest and cypress together before commiting them to codecov.

# Description

Associated Jira ticket: https://issues.redhat.com/browse/SPM-1826

# How to test the PR

Run tests locally, check that two coverage folders are generated, and then try  to run `npm run coverage`.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on December 20, 2022 at 03:12 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **69.77**% // Head: **72.20**% // Increases project coverage by **`+2.43%`** :tada:
> Coverage data is based on head [(`c0fc974`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`5accc3a`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/5accc3a7bdfc9f7e0854e800b11ddaae45cc9f83?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 90.47% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #933      +/-   ##
==========================================
+ Coverage   69.77%   72.20%   +2.43%     
==========================================
  Files         111      111              
  Lines        2607     2623      +16     
  Branches      679      662      -17     
==========================================
+ Hits         1819     1894      +75     
- Misses        712      729      +17     
+ Partials       76        0      -76     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `90.74% <ø> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `85.07% <84.61%> (+85.07%)` | :arrow_up: |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `71.42% <100.00%> (+4.76%)` | :arrow_up: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `74.50% <100.00%> (+1.04%)` | :arrow_up: |
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL2luZGV4Lmpz) | `0.00% <0.00%> (ø)` | |
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `59.09% <0.00%> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `78.80% <0.00%> (ø)` | |
| ... and [16 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/933?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on January 18, 2023 at 03:45 PM UTC

:tada: This PR is included in version 1.58.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.58.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.58.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on December 21, 2022 at 10:30 AM UTC

LGTM! 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/933*
