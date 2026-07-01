---
type: pull_request
number: 861
title: "fix(BulkSelect): enable loading state when all entities selected"
state: merged
author: mkholjuraev
created: 2022-08-04T11:26:18Z
updated: 2024-04-03T09:22:15Z
closed: 2022-09-02T15:43:47Z
merged: 2022-09-02T15:43:47Z
base_branch: master
head_branch: bulk-selection-loading
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/861
---

# Pull Request #861: fix(BulkSelect): enable loading state when all entities selected

**Author**: @mkholjuraev
**Created**: August 04, 2022 at 11:26 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `bulk-selection-loading`

## Description

This introduces a loading icon when all entities are selected using the bulk selector.

---

## Discussion

### Comment by @codecov-commenter on August 04, 2022 at 11:34 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/861?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.09**% // Head: **72.16**% // Increases project coverage by **`+0.06%`** :tada:
> Coverage data is based on head [(`c83ebc2`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/861?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`fb81f43`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/fb81f43be4273d6f4cce01f4d4847d0f5a2ffeaa?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 100.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #861      +/-   ##
==========================================
+ Coverage   72.09%   72.16%   +0.06%     
==========================================
  Files         109      109              
  Lines        2498     2504       +6     
  Branches      635      637       +2     
==========================================
+ Hits         1801     1807       +6     
  Misses        637      637              
  Partials       60       60              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/861?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/861/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `77.90% <100.00%> (+0.79%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/861?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @AsToNlele on August 05, 2022 at 12:10 PM UTC

Thank you for your PR!
I've found in some really rare occasions there are two labels visible, but it would show up randomly and very few times
![image](https://user-images.githubusercontent.com/20592948/183069987-01428594-767d-4630-beb9-0f182a49a693.png)
Looks great otherwise!


### Comment by @mkholjuraev on August 05, 2022 at 12:26 PM UTC

@AsToNlele great catch! Can you tell in what cases, does this occur?

### Comment by @AsToNlele on August 05, 2022 at 01:59 PM UTC

I'm afraid its completely random, It happened to me on the first try and then after like 10 tries.

### Comment by @mkholjuraev on August 19, 2022 at 10:10 AM UTC

@AsToNlele the issue is fixed

### Comment by @adonispuente on September 01, 2022 at 01:18 PM UTC

The ' Select All’ option seems to be working perfectly fine, but the loading icon doesnt pop up when just clicking the empty square, even though clicking the empty square does the same action. I believe we should add this icon when clicking this button

### Comment by @mkholjuraev on September 02, 2022 at 10:09 AM UTC

@adonispuente thanks for the suggestion. I have enabled the loader for that action as well. Please have a look.

### Comment by @jiridostal on September 02, 2022 at 03:59 PM UTC

:tada: This PR is included in version 1.49.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.49.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.49.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on September 02, 2022 at 02:29 PM UTC

great PR, LGTM! 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/861*
