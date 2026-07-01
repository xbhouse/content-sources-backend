---
type: pull_request
number: 1066
title: "fix(Remediating): SPM-2000 - Selecting all systems selected filtered \u2026"
state: closed
author: AsToNlele
created: 2023-05-30T12:37:14Z
updated: 2024-06-04T11:41:57Z
closed: 2024-06-04T11:41:57Z
base_branch: master
head_branch: SPM-2000
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1066
---

# Pull Request #1066: fix(Remediating): SPM-2000 - Selecting all systems selected filtered …

**Author**: @AsToNlele
**Created**: May 30, 2023 at 12:37 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `SPM-2000`

## Description

# Description

Associated Jira ticket: [SPM-2000](https://issues.redhat.com/browse/SPM-2000)

Selecting all systems while having filters remediates now correct amount of systems


# How to test the PR

1. Go to Patch Systems
2. Apply some filters
3. Select all
4. You can remediate only systems that have installable advisories, so ideally count how many systems have Installable advisories 
5. Remediate
6. Make sure the number of systems matches

# Before the change

When selecting all systems while having filters, wrong system count is remediated

# After the change

Correct amount is remediated

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

### Comment by @codecov-commenter on May 30, 2023 at 12:47 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1066?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`25.00%`** and project coverage change: **`+0.07%`** :tada:
> Comparison is base [(`4fb84ef`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4fb84ef910749985d682da0c051a2d448520ce9d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.15% compared to head [(`c5ed917`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1066?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.23%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1066      +/-   ##
==========================================
+ Coverage   63.15%   63.23%   +0.07%     
==========================================
  Files         120      120              
  Lines        2994     2984      -10     
  Branches      765      761       -4     
==========================================
- Hits         1891     1887       -4     
+ Misses       1103     1097       -6     
```


| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1066?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Advisories/Advisories.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1066?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `100.00% <ø> (ø)` | |
| [src/Utilities/RemediationPairs.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1066?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9SZW1lZGlhdGlvblBhaXJzLmpz) | `0.00% <0.00%> (ø)` | |
| [src/store/Reducers/AdvisoryListStore.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1066?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5TGlzdFN0b3JlLmpz) | `80.00% <ø> (+10.43%)` | :arrow_up: |
| [src/store/Reducers/SystemsStore.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1066?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbXNTdG9yZS5qcw==) | `100.00% <ø> (+17.64%)` | :arrow_up: |
| [src/SmartComponents/Systems/Systems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1066?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `78.78% <100.00%> (-0.63%)` | :arrow_down: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1066?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @leSamo on August 04, 2023 at 02:04 PM UTC

For some reason every time I try to open Status filter and click on a checkbox the whole page freezes and the tab crashes later, and I quite can't put my finger on why it happens, can you add and remove Status filter without page hanging?

### Comment by @AsToNlele on August 21, 2023 at 10:35 AM UTC

/retest

### Comment by @leSamo on August 21, 2023 at 10:23 PM UTC

@AsToNlele were you able to reproduce/fix the issue I mentioned in my previous comment? I believe the last force push was just a rebase, or did it include any changes?

### Comment by @AsToNlele on August 22, 2023 at 06:33 AM UTC

I wasn't able to reproduce, yeah I just rebased :smile: 

### Comment by @AsToNlele on August 22, 2023 at 06:34 AM UTC

/retest

### Comment by @mkholjuraev on March 15, 2024 at 02:09 PM UTC

@AsToNlele seems the issue has been resolved.

### Comment by @mkholjuraev on June 04, 2024 at 11:41 AM UTC

Probably we can close this :)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1066*
