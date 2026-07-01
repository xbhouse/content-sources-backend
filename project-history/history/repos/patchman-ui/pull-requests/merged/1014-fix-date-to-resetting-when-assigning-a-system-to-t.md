---
type: pull_request
number: 1014
title: "Fix date-to resetting when assigning a system to template"
state: merged
author: leSamo
created: 2023-04-03T16:46:20Z
updated: 2023-05-08T18:16:30Z
closed: 2023-04-06T21:24:45Z
merged: 2023-04-06T21:24:45Z
base_branch: master
head_branch: fix-date-reset
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1014
---

# Pull Request #1014: Fix date-to resetting when assigning a system to template

**Author**: @leSamo
**Created**: April 03, 2023 at 04:46 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-date-reset`

## Description

There was a bug when you assigned a system to a template in the system list table, the template would have the date-to reset to January 1, year 1. This has now been fixed.

---

## Discussion

### Comment by @codecov-commenter on April 03, 2023 at 04:54 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1014?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`+0.26`** :tada:
> Comparison is base [(`d705f0e`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/d705f0e7d04b79b30dc024dd9e65329bc7503122?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.22% compared to head [(`0405d8b`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1014?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.49%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1014      +/-   ##
==========================================
+ Coverage   64.22%   64.49%   +0.26%     
==========================================
  Files         116      116              
  Lines        2812     2836      +24     
  Branches      723      725       +2     
==========================================
+ Hits         1806     1829      +23     
- Misses       1006     1007       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1014?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1014?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `74.50% <0.00%> (-0.50%)` | :arrow_down: |

... and [2 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1014/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1014?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on April 17, 2023 at 06:41 PM UTC

:tada: This PR is included in version 1.63.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on April 06, 2023 at 03:33 PM UTC

Small nitpick but not completely necessary. LGTM! 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1014*
