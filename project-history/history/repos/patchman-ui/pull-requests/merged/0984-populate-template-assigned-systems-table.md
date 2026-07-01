---
type: pull_request
number: 984
title: "Populate template assigned systems table"
state: merged
author: leSamo
created: 2023-03-08T22:40:27Z
updated: 2023-05-08T18:16:44Z
closed: 2023-03-09T08:48:40Z
merged: 2023-03-09T08:48:40Z
base_branch: master
head_branch: template-detail-columns
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/984
---

# Pull Request #984: Populate template assigned systems table

**Author**: @leSamo
**Created**: March 08, 2023 at 10:40 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `template-detail-columns`

## Description

# Description

Associated Jira ticket: [SPM-1930](https://issues.redhat.com/browse/SPM-1930)

Mockup: https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/DP7Ka9y

It's still missing tags column and OS column is not sortable, but we will add it as a part of another ticket.

---

## Discussion

### Comment by @codecov-commenter on March 08, 2023 at 10:46 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`50.00`**% and project coverage change: **`-0.04`** :warning:
> Comparison is base [(`a718b0d`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/a718b0d4d35648b8084a34e17a5d828ab220836f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 68.94% compared to head [(`52d4edc`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 68.90%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #984      +/-   ##
==========================================
- Coverage   68.94%   68.90%   -0.04%     
==========================================
  Files         115      115              
  Lines        2798     2801       +3     
  Branches      709      712       +3     
==========================================
+ Hits         1929     1930       +1     
- Misses        869      871       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...tComponents/PatchSetDetail/PatchSetDetailAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbEFzc2V0cy5qcw==) | `0.00% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `68.18% <0.00%> (-1.59%)` | :arrow_down: |
| [src/Utilities/SystemsHelpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | `26.66% <0.00%> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `81.84% <100.00%> (ø)` | |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `75.00% <100.00%> (ø)` | |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/984?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 09, 2023 at 09:08 AM UTC

:tada: This PR is included in version 1.62.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on March 09, 2023 at 08:42 AM UTC

Codewise looks good!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/984*
