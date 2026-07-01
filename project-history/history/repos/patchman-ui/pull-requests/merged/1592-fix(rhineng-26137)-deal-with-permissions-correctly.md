---
type: pull_request
number: 1592
title: "fix(RHINENG-26137): deal with permissions correctly based on FF"
state: merged
author: LightOfHeaven1994
created: 2026-04-29T13:35:52Z
updated: 2026-04-30T12:59:25Z
closed: 2026-04-30T12:59:25Z
merged: 2026-04-30T12:59:25Z
base_branch: master
head_branch: RHINENG-26137
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1592
---

# Pull Request #1592: fix(RHINENG-26137): deal with permissions correctly based on FF

**Author**: @LightOfHeaven1994
**Created**: April 29, 2026 at 01:35 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-26137`

## Description

This PR deals with fetching permissions correctly. It will avoid calling both rbac & kessel hooks at the same time and it will be done based on feature flag value.

Also we are properly waiting until feature flag is fetched, for that we use `flagsReady`.


### How to test:

1. Navigate to `Advisories` page for example 
2. Open devtools network tab
3. Refresh the page and check if there are no RBACv1 requests made for patch `access/?application=patch`
4. Check that app loads as usual

---

## Discussion

### Comment by @codecov-commenter on April 29, 2026 at 02:16 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1592?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `25.00000%` with `36 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 76.38%. Comparing base ([`46c9cfe`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/46c9cfe36b6f5cde2136fa986a7f7a092abe4688?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`683e553`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/683e55372ffa9d759ac78b8dd34716859f3c054b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 5 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1592?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Routes.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1592?src=pr&el=tree&filepath=src%2FRoutes.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | 0.00% | [12 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1592?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/PatchPermissionGate.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1592?src=pr&el=tree&filepath=src%2FUtilities%2FPatchPermissionGate.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9QYXRjaFBlcm1pc3Npb25HYXRlLmpz) | 0.00% | [12 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1592?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/hooks/usePermissionCheck.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1592?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FusePermissionCheck.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VQZXJtaXNzaW9uQ2hlY2suanM=) | 0.00% | [5 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1592?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1592      +/-   ##
==========================================
+ Coverage   75.88%   76.38%   +0.50%     
==========================================
  Files         103      103              
  Lines        3164     3185      +21     
  Branches      685      692       +7     
==========================================
+ Hits         2401     2433      +32     
+ Misses        684      674      -10     
+ Partials       79       78       -1     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1592?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @xbhouse - Approved on April 30, 2026 at 12:54 PM UTC

lgtm! thank you ❤️ 

### Review by @dominikvagner - Approved on April 30, 2026 at 12:55 PM UTC

ack! ✅ thx 👏🏼 

_ahh, Bryttanie beat me to it 🐌😄_ 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1592*
