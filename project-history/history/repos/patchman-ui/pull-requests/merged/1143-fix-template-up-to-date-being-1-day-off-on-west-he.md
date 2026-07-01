---
type: pull_request
number: 1143
title: "Fix template up-to date being 1 day off on west hemisphere"
state: merged
author: leSamo
created: 2023-11-14T23:14:54Z
updated: 2023-12-13T11:16:20Z
closed: 2023-11-16T18:13:18Z
merged: 2023-11-16T18:13:18Z
base_branch: master
head_branch: tz-fix
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1143
---

# Pull Request #1143: Fix template up-to date being 1 day off on west hemisphere

**Author**: @leSamo
**Created**: November 14, 2023 at 11:14 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `tz-fix`

## Description

Ticket: https://issues.redhat.com/browse/RHINENG-3122

Fix the timezone bug by setting the up to date time as the next local timezone midnight minus one second (23:59:59).

---

## Discussion

### Comment by @codecov-commenter on November 14, 2023 at 11:21 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1143?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `4 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`d99f4c6`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d99f4c61dcfa2688aa0db10e68d90705f6489a8d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.38% compared to head [(`9f3891b`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1143?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.49%.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1143?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/Hooks.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1143?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | 0.00% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1143?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1143      +/-   ##
==========================================
+ Coverage   63.38%   63.49%   +0.11%     
==========================================
  Files         122      122              
  Lines        3094     3087       -7     
  Branches      802      799       -3     
==========================================
- Hits         1961     1960       -1     
+ Misses       1133     1127       -6     
```



</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1143?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @leSamo on November 16, 2023 at 10:08 AM UTC

/retest

### Comment by @leSamo on November 16, 2023 at 05:14 PM UTC

/retest

### Comment by @mkholjuraev on December 13, 2023 at 11:16 AM UTC

:tada: This PR is included in version 1.65.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.65.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on November 16, 2023 at 01:22 PM UTC

Cool fix :smile: LGTM

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1143*
