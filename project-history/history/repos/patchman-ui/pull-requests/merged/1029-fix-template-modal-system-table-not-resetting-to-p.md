---
type: pull_request
number: 1029
title: "Fix template modal system table not resetting to page 1"
state: merged
author: leSamo
created: 2023-04-19T23:10:37Z
updated: 2023-05-08T18:14:45Z
closed: 2023-04-25T10:38:34Z
merged: 2023-04-25T10:38:34Z
base_branch: master
head_branch: fix-review-step-page-1
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1029
---

# Pull Request #1029: Fix template modal system table not resetting to page 1

**Author**: @leSamo
**Created**: April 19, 2023 at 11:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-review-step-page-1`

## Description

See https://github.com/RedHatInsights/patchman-ui/pull/1028#pullrequestreview-1391792579

Bug that was fixed happened in Template modal -> System review step if you navigated to page 2 and then applied filter which returned less than 20 rows; calling page 2 on this data caused 400 error. Fixed by resetting to page 1 on every filter change, which is behaviour of all other Patchman tables (and Vulnerabilities as well).

---

## Discussion

### Comment by @codecov-commenter on April 19, 2023 at 11:18 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1029?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`+0.39`** :tada:
> Comparison is base [(`a8070d7`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/a8070d7bc43da92bdc99429f86652677ab333021?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.12% compared to head [(`b5da50b`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1029?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.51%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1029      +/-   ##
==========================================
+ Coverage   64.12%   64.51%   +0.39%     
==========================================
  Files         116      117       +1     
  Lines        2832     2996     +164     
  Branches      733      808      +75     
==========================================
+ Hits         1816     1933     +117     
- Misses       1016     1063      +47     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1029?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1029?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `4.65% <ø> (ø)` | |

... and [7 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1029/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1029?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @AsToNlele - Approved on April 25, 2023 at 10:13 AM UTC

LGTM good find :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1029*
