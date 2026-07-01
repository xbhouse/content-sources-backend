---
type: pull_request
number: 930
title: "Optimizations"
state: merged
author: mkholjuraev
created: 2022-12-08T11:52:27Z
updated: 2024-04-03T09:21:53Z
closed: 2022-12-22T12:32:28Z
merged: 2022-12-22T12:32:28Z
base_branch: master
head_branch: optimizations
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/930
---

# Pull Request #930: Optimizations

**Author**: @mkholjuraev
**Created**: December 08, 2022 at 11:52 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `optimizations`

## Description

# Description

Associated Jira ticket: # (issue)
for optimization


# How to test the PR

Please play around the UI.

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

### Comment by @codecov-commenter on December 19, 2022 at 12:15 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **69.77**% // Head: **72.21**% // Increases project coverage by **`+2.44%`** :tada:
> Coverage data is based on head [(`2db904d`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`c1f00fd`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/c1f00fdad3a4cf207c4fb54c9eed314b0815efcb?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 48.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #930      +/-   ##
==========================================
+ Coverage   69.77%   72.21%   +2.44%     
==========================================
  Files         111      110       -1     
  Lines        2607     2613       +6     
  Branches      679      679              
==========================================
+ Hits         1819     1887      +68     
+ Misses        712      655      -57     
+ Partials       76       71       -5     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/AppEntry.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcEVudHJ5Lmpz) | `0.00% <ø> (ø)` | |
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [...c/PresentationalComponents/Filters/SearchFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1NlYXJjaEZpbHRlci5qcw==) | `88.88% <75.00%> (-11.12%)` | :arrow_down: |
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `100.00% <100.00%> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `84.31% <100.00%> (+1.96%)` | :arrow_up: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `90.74% <100.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `84.05% <100.00%> (+1.44%)` | :arrow_up: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `74.50% <0.00%> (+1.04%)` | :arrow_up: |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `71.42% <0.00%> (+4.76%)` | :arrow_up: |
| ... and [3 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/930?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on December 22, 2022 at 11:55 AM UTC

@gkarat thank you for the review! I have skimmed through the doc you have linked. Indeed, it has good examples.

While working on the optimizations, I also referred to profiling and measured the re-renders. It seems that most of the re-renders are necessary re-renders together with platform dependant ones. Thus, there is not much saving on re-renders with this PR unfortunately :). But still, there are on average ~15 re-renders on the first load, thus trying to avoid such a tiny function recreation can be a good thing in the end. 

Regarding the **useCallback** in the searchFilter, this is to get rid of keeping a debounced function in a state. This was wrong, the correct way is to store it in a useCallback to have only one instance across re-renders.

---

## Reviews

### Review by @gkarat - Changes Requested on December 14, 2022 at 11:54 AM UTC

### Review by @mkholjuraev - Commented on December 14, 2022 at 12:21 PM UTC

### Review by @mkholjuraev - Commented on December 14, 2022 at 12:24 PM UTC

### Review by @mkholjuraev - Commented on December 14, 2022 at 12:42 PM UTC

### Review by @gkarat - Commented on December 22, 2022 at 09:29 AM UTC

### Review by @gkarat - Commented on December 22, 2022 at 09:31 AM UTC

### Review by @gkarat - Approved on December 22, 2022 at 09:38 AM UTC

this looks good! if it's considered complete, I am OK with merging it.

just a side nit-pick: I would like to make some profiling before/after such PRs to see whether how much this increases the performance. some of the `useCallback`s in tiny, not huge components make me wonder if it really provides any value to add it there. e.g., here is well written article about the different use cases for useCallback: https://dmitripavlutin.com/react-usecallback.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/930*
