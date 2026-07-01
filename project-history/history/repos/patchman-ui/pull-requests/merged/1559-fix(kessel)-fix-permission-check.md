---
type: pull_request
number: 1559
title: "fix(Kessel): Fix permission check"
state: merged
author: LightOfHeaven1994
created: 2026-03-26T13:46:37Z
updated: 2026-03-26T15:29:09Z
closed: 2026-03-26T15:29:08Z
merged: 2026-03-26T15:29:08Z
base_branch: master
head_branch: fix-kessel-permission-check
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1559
---

# Pull Request #1559: fix(Kessel): Fix permission check

**Author**: @LightOfHeaven1994
**Created**: March 26, 2026 at 01:46 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-kessel-permission-check`

## Description

`patch_template_edit` is not used anywhere so I removed it to avoid confusion.

This PR fixes permissions check for Kessel. As both permissions we are checking here are [inventory host centric](https://github.com/RedHatInsights/rbac-config/blob/master/configs/stage/schemas/src/patch.ksl) - we need to check these permissions in all workspaces and not in only default one.

I found out about this bug when I was testing other apps with granular inventory permissions. so you do not give `Inventory host view` generic permission but you give an access to read hosts in workspace X only. App was getting ALLOWED_FALSE because we were not checking permissions correctly here

---

## Discussion

### Comment by @codecov-commenter on March 26, 2026 at 01:49 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1559?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `44 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 71.80%. Comparing base ([`53f6193`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/53f619329d3b2cadabdf24e792274d8e48bbe4fa?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`322a900`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/322a900d5805377f54a100387e99b9e318ea68f5?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 6 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1559?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/hooks/useKesselWorkspaceIds.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1559?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FuseKesselWorkspaceIds.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VLZXNzZWxXb3Jrc3BhY2VJZHMuanM=) | 0.00% | [30 Missing and 6 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1559?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/hooks/usePermissionCheck.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1559?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FusePermissionCheck.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VQZXJtaXNzaW9uQ2hlY2suanM=) | 0.00% | [4 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1559?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1559      +/-   ##
==========================================
- Coverage   72.34%   71.80%   -0.54%     
==========================================
  Files          99       99              
  Lines        2408     2426      +18     
  Branches      683      681       -2     
==========================================
  Hits         1742     1742              
- Misses        586      601      +15     
- Partials       80       83       +3     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1559?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @LightOfHeaven1994 - Commented on March 26, 2026 at 02:58 PM UTC

### Review by @dominikvagner - Commented on March 26, 2026 at 03:15 PM UTC

### Review by @dominikvagner - Approved on March 26, 2026 at 03:27 PM UTC

looks good to me 👍🏼
ack ✅ 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1559*
