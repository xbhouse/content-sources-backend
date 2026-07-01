---
type: pull_request
number: 1040
title: "(RHIN-830): Add Zero State to Patch"
state: merged
author: adonispuente
created: 2023-05-02T16:20:39Z
updated: 2023-05-03T14:00:44Z
closed: 2023-05-03T14:00:44Z
merged: 2023-05-03T14:00:44Z
base_branch: master
head_branch: zerostate
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1040
---

# Pull Request #1040: (RHIN-830): Add Zero State to Patch

**Author**: @adonispuente
**Created**: May 02, 2023 at 04:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `zerostate`

## Description

 As in accordance to these mocks, this PR adds zero state to patch,  https://docs.google.com/document/d/1FGuzbL2jyRmn_coYD5Z_o3d0MhOf1vQDT46NQvsfPyA/edit#heading=h.1hvxoje16w3i  https://www.sketch.com/s/550bd962-ca1f-42d8-8d89-d457527a059a/a/qbe9Ol2#Inspect   https://www.sketch.com/s/550bd962-ca1f-42d8-8d89-d457527a059a/prototype/a/AEF47CA9-8F00-40FC-B39B-3B989D3F83FD   To test, sign into an account with 0 systems connected  

---

## Discussion

### Comment by @codecov-commenter on May 02, 2023 at 04:32 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1040?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`-0.38`** :warning:
> Comparison is base [(`84654b7`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/84654b7b0c635b6de03610e147959d50b1414773?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.05% compared to head [(`7884f89`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1040?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.67%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1040      +/-   ##
==========================================
- Coverage   64.05%   63.67%   -0.38%     
==========================================
  Files         116      118       +2     
  Lines        2835     2877      +42     
  Branches      733      739       +6     
==========================================
+ Hits         1816     1832      +16     
- Misses       1019     1045      +26     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1040?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1040?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `0.00% <0.00%> (ø)` | |

... and [25 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1040/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1040?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @mkholjuraev - Changes Requested on May 02, 2023 at 04:31 PM UTC

Patch has a good handler implemented for the zero (empty) state [here](https://github.com/RedHatInsights/patchman-ui/blob/master/src/PresentationalComponents/Snippets/ErrorHandler.js#L53). Patch api adds has_systems boolean to check this senario. 
let's make use of it if you agree by importing the zero-state module inside [NoRegisteredSystems](https://github.com/RedHatInsights/patchman-ui/blob/master/src/PresentationalComponents/Snippets/NoRegisteredSystems.js) component

### Review by @mkholjuraev - Approved on May 03, 2023 at 11:18 AM UTC

As we have a short time, let's merge this and work on refactoring this not to use extra API calls to /hosts later after the code freeze.

Tested locally, LGTM. The only thins that caught my attention is that the global filter is hidden in the mocks but not in this PR.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1040*
