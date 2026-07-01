---
type: pull_request
number: 1022
title: "Fix adding/removing template not causing refresh in System detail header"
state: merged
author: leSamo
created: 2023-04-12T13:46:22Z
updated: 2023-05-08T18:16:00Z
closed: 2023-04-12T18:50:26Z
merged: 2023-04-12T18:50:26Z
base_branch: master
head_branch: system-detail-template-not-refreshing
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1022
---

# Pull Request #1022: Fix adding/removing template not causing refresh in System detail header

**Author**: @leSamo
**Created**: April 12, 2023 at 01:46 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `system-detail-template-not-refreshing`

## Description

Associated Jira ticket: [SPM-2003](https://issues.redhat.com/browse/SPM-2003)



---

## Discussion

### Comment by @codecov-commenter on April 12, 2023 at 01:56 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1022?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`-0.13`** :warning:
> Comparison is base [(`3e941e0`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/3e941e09b7852c766f6b6e9f3911e4e2de84a58b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.30% compared to head [(`2a2e685`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1022?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.17%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1022      +/-   ##
==========================================
- Coverage   64.30%   64.17%   -0.13%     
==========================================
  Files         116      116              
  Lines        2821     2892      +71     
  Branches      725      763      +38     
==========================================
+ Hits         1814     1856      +42     
- Misses       1007     1036      +29     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1022?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/store/Reducers/SystemDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1022?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbURldGFpbFN0b3JlLmpz) | `69.23% <0.00%> (+12.98%)` | :arrow_up: |

... and [3 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1022/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1022?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on April 17, 2023 at 06:41 PM UTC

:tada: This PR is included in version 1.63.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on April 12, 2023 at 05:11 PM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1022*
