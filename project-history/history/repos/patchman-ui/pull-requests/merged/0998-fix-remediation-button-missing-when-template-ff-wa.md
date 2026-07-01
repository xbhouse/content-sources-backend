---
type: pull_request
number: 998
title: "Fix Remediation button missing when Template FF was false"
state: merged
author: leSamo
created: 2023-03-23T12:38:27Z
updated: 2023-05-08T18:16:32Z
closed: 2023-03-23T16:16:28Z
merged: 2023-03-23T16:16:28Z
base_branch: master
head_branch: fix-remediation
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/998
---

# Pull Request #998: Fix Remediation button missing when Template FF was false

**Author**: @leSamo
**Created**: March 23, 2023 at 12:38 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-remediation`

## Description

https://redhat-internal.slack.com/archives/CPS5FH07Q/p1679499017978129

---

## Discussion

### Comment by @codecov-commenter on March 23, 2023 at 01:14 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/998?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00`**% and project coverage change: **`+0.01`** :tada:
> Comparison is base [(`385a1ba`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/385a1bafab1c26fadd384e27874978a4c6285176?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.91% compared to head [(`91ee333`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/998?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.92%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #998      +/-   ##
==========================================
+ Coverage   64.91%   64.92%   +0.01%     
==========================================
  Files         116      116              
  Lines        2813     2814       +1     
  Branches      722      723       +1     
==========================================
+ Hits         1826     1827       +1     
  Misses        987      987              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/998?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/998?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `84.05% <100.00%> (+0.23%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/998?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 28, 2023 at 08:56 AM UTC

:tada: This PR is included in version 1.62.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on March 23, 2023 at 04:00 PM UTC

LGTM! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/998*
