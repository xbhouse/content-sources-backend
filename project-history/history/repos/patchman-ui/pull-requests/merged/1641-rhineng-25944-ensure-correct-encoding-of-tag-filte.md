---
type: pull_request
number: 1641
title: "RHINENG-25944: ensure correct encoding of tag filters"
state: merged
author: dominikvagner
created: 2026-05-27T07:59:01Z
updated: 2026-06-08T09:49:30Z
closed: 2026-06-08T09:49:28Z
merged: 2026-06-08T09:49:28Z
base_branch: master
head_branch: tags-filter-fixes
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1641
---

# Pull Request #1641: RHINENG-25944: ensure correct encoding of tag filters

**Author**: @dominikvagner
**Created**: May 27, 2026 at 07:59 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `tags-filter-fixes`

## Description

## Summary
This makes sure the global tag filters are correctly URI encoded, only a
single encode pass. Previously they could be double encoded which made
the API always return no systems.

Associated Jira ticket: [#RHINENG-25944](https://redhat.atlassian.net/browse/RHINENG-25944)

> Also includes a QOL fix for local development, _that sets the maxWorkers 
> in the Jest config to prevent accidental computer lock-ups/freezes when running
> tests, because the default is using the maximum of available resources._

## Testing steps
Test out filtering with both the "local" tag filters (in the table toolbar) and with the global filters at the top.
Both of them should now work.



---

## Discussion

### Comment by @codecov-commenter on May 27, 2026 at 08:03 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1641?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `83.33333%` with `3 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 77.57%. Comparing base ([`6480ecc`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/6480ecc8f7c4e7f9f75c5108b99e7bfda1d23acc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`48e84cc`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/48e84cc688b99c438398ebe047811185d1ed5775?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1641?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1641?src=pr&el=tree&filepath=src%2FUtilities%2FHelpers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | 83.33% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1641?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1641      +/-   ##
==========================================
- Coverage   77.58%   77.57%   -0.01%     
==========================================
  Files         103      103              
  Lines        3266     3278      +12     
  Branches      728      730       +2     
==========================================
+ Hits         2534     2543       +9     
- Misses        655      658       +3     
  Partials       77       77              
```
</details>

[:umbrella: View full report in Codecov by Harness](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1641?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @xbhouse on June 04, 2026 at 02:47 PM UTC

FYI the "Advisories Tests › Create a new remediation plan" UI test will likely fail until culled systems can be deleted in patch again. this [PR](https://github.com/RedHatInsights/patchman-engine/pull/2229) should fix that :crossed_fingers: 

---

## Reviews

### Review by @TenSt - Approved on May 27, 2026 at 01:56 PM UTC

lgtm

### Review by @xbhouse - Commented on June 04, 2026 at 03:05 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1641*
