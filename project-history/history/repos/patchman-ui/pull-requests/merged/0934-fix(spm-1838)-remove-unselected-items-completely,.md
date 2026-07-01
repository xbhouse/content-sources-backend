---
type: pull_request
number: 934
title: "fix(SPM-1838): remove unselected items completely, ignore expanded rows"
state: merged
author: mkholjuraev
created: 2022-12-21T10:49:30Z
updated: 2022-12-23T09:21:40Z
closed: 2022-12-23T09:21:40Z
merged: 2022-12-23T09:21:40Z
base_branch: master
head_branch: fix-selection
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/934
---

# Pull Request #934: fix(SPM-1838): remove unselected items completely, ignore expanded rows

**Author**: @mkholjuraev
**Created**: December 21, 2022 at 10:49 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-selection`

## Description

# Description

Associated Jira ticket: # ([SPM-1838](https://issues.redhat.com/browse/SPM-1838))

Fixes the selection on the System advisories table by removing the unselected item completely from the selected rows object, instead of setting it to selected=falsy. Also, if expanded rows are ignored while turning a row as selected. 


# How to test the PR
1. This touches selection across pages, thus please play with it and try remediating your selected items.

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

### Comment by @codecov-commenter on December 22, 2022 at 09:11 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.20**% // Head: **72.27**% // Increases project coverage by **`+0.06%`** :tada:
> Coverage data is based on head [(`7fc2194`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`a8aaa0b`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/a8aaa0b27b2f5be243b5fd64468351bc89ca3fee?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 59.37% of modified lines in pull request are covered.

> :exclamation: Current head 7fc2194 differs from pull request most recent head de2fd4b. Consider uploading reports for the commit de2fd4b to get more accurate results

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #934      +/-   ##
==========================================
+ Coverage   72.20%   72.27%   +0.06%     
==========================================
  Files         111      110       -1     
  Lines        2623     2615       -8     
  Branches      681      680       -1     
==========================================
- Hits         1894     1890       -4     
+ Misses        658      654       -4     
  Partials       71       71              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/AppEntry.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcEVudHJ5Lmpz) | `0.00% <ø> (ø)` | |
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [...c/PresentationalComponents/Filters/SearchFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1NlYXJjaEZpbHRlci5qcw==) | `88.88% <75.00%> (-11.12%)` | :arrow_down: |
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `100.00% <100.00%> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `84.31% <100.00%> (+1.96%)` | :arrow_up: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `90.74% <100.00%> (ø)` | |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `93.18% <100.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `84.05% <100.00%> (+1.44%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `78.87% <100.00%> (+0.06%)` | :arrow_up: |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/934?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/934*
