---
type: pull_request
number: 1098
title: "Split installable and applicable version in System detail >\u00a0Packages table"
state: merged
author: leSamo
created: 2023-07-22T01:00:56Z
updated: 2023-08-14T13:21:48Z
closed: 2023-08-01T14:40:22Z
merged: 2023-08-01T14:40:22Z
base_branch: master
head_branch: system-detail-packages-update
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1098
---

# Pull Request #1098: Split installable and applicable version in System detail > Packages table

**Author**: @leSamo
**Created**: July 22, 2023 at 01:00 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `system-detail-packages-update`

## Description

https://issues.redhat.com/browse/SPM-1946

- Replace "Latest version" column with two columns – "Installable version" and "Latest applicable version"
- Disable row checkbox if the the package status is not installable
- Update "Status" filter options (new options are "Up to date", "Applicable", "Installable")

Explanation:
Packages have 3 different versions displayed -- currently installed version, latest version that is available (so called "Applicable" version) and "Installable" version. If system is not in a template, Applicable and Installable versions are the same. But if a system is part of template, the template restricts available versions up to a date, for example if template date is 2022/12/31 it will only show versions released before that date in Installable column.

What to test:
- Try filtering by Status
- Check that installable version is never greater than applicable version
- If a system does not have a template installable and applicable versions should be the same
- If package installable version is greater that current version, checkbox for remediation should be enabled, otherwise disabled.

---

## Discussion

### Comment by @codecov-commenter on July 22, 2023 at 01:10 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`86.66`**% and project coverage change: **`-0.01`** :warning:
> Comparison is base [(`747e04d`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/747e04dc644a2e6c313f4f1960acec08f22c16de?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.57% compared to head [(`338bd37`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.57%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1098      +/-   ##
==========================================
- Coverage   62.57%   62.57%   -0.01%     
==========================================
  Files         119      119              
  Lines        2993     2998       +5     
  Branches      769      770       +1     
==========================================
+ Hits         1873     1876       +3     
- Misses       1120     1122       +2     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...sentationalComponents/TableView/TableViewAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3QXNzZXRzLmpz) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `59.45% <ø> (ø)` | |
| [src/Utilities/RawDataForTesting.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9SYXdEYXRhRm9yVGVzdGluZy5qcw==) | `100.00% <ø> (ø)` | |
| [src/Utilities/constants.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `80.00% <75.00%> (-0.46%)` | :arrow_down: |
| [...c/PresentationalComponents/Filters/StatusFilter.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1N0YXR1c0ZpbHRlci5qcw==) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `58.00% <100.00%> (+0.85%)` | :arrow_up: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1098?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @leSamo on July 22, 2023 at 11:24 AM UTC

/retest

### Comment by @leSamo on July 24, 2023 at 10:11 AM UTC

/retest

### Comment by @leSamo on July 25, 2023 at 12:29 PM UTC

/retest

### Comment by @mkholjuraev on August 07, 2023 at 02:34 PM UTC

:tada: This PR is included in version 1.63.11 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.11)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on August 01, 2023 at 02:28 PM UTC

LGTM, good job! :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1098*
