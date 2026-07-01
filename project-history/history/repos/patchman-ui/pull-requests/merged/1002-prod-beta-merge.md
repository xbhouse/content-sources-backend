---
type: pull_request
number: 1002
title: "Prod beta merge"
state: merged
author: mkholjuraev
created: 2023-03-28T08:48:12Z
updated: 2024-04-03T09:22:35Z
closed: 2023-03-28T10:17:58Z
merged: 2023-03-28T10:17:58Z
base_branch: prod-beta
head_branch: prod-beta-merge
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1002
---

# Pull Request #1002: Prod beta merge

**Author**: @mkholjuraev
**Created**: March 28, 2023 at 08:48 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `prod-beta` ← **Head**: `prod-beta-merge`

## Description

# Description

This essentially aligns master and stage-stable branches. Also, it promotes useChrome commits:[RHIF-128](https://issues.redhat.com/browse/RHIF-128)

---

## Discussion

### Comment by @codecov-commenter on March 28, 2023 at 09:31 AM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`27.72`**% and project coverage change: **`-4.74`** :warning:
> Comparison is base [(`a718b0d`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/a718b0d4d35648b8084a34e17a5d828ab220836f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 68.94% compared to head [(`b347554`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.21%.

<details><summary>Additional details and impacted files</summary>


```diff
@@              Coverage Diff              @@
##           prod-beta    #1002      +/-   ##
=============================================
- Coverage      68.94%   64.21%   -4.74%     
=============================================
  Files            115      116       +1     
  Lines           2798     2811      +13     
  Branches         709      722      +13     
=============================================
- Hits            1929     1805     -124     
- Misses           869     1006     +137     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...sentationalComponents/AdvisoryType/AdvisoryType.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeVR5cGUvQWR2aXNvcnlUeXBlLmpz) | `100.00% <ø> (ø)` | |
| [.../PresentationalComponents/Filters/CreatorFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL0NyZWF0b3JGaWx0ZXIuanM=) | `0.00% <0.00%> (ø)` | |
| [src/PresentationalComponents/Header/Header.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9IZWFkZXIvSGVhZGVyLmpz) | `100.00% <ø> (ø)` | |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `72.22% <ø> (ø)` | |
| [src/PresentationalComponents/Snippets/Error.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FcnJvci5qcw==) | `100.00% <ø> (ø)` | |
| [.../PresentationalComponents/Snippets/ExternalLink.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FeHRlcm5hbExpbmsuanM=) | `100.00% <ø> (ø)` | |
| [...rc/PresentationalComponents/TableView/TableView.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3Lmpz) | `100.00% <ø> (ø)` | |
| [...sentationalComponents/TableView/TableViewAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3QXNzZXRzLmpz) | `100.00% <ø> (ø)` | |
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `0.00% <0.00%> (ø)` | |
| ... and [26 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

... and [4 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1002?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1002*
