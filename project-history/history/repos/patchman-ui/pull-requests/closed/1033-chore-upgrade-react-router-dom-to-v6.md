---
type: pull_request
number: 1033
title: "chore: upgrade react-router-dom to v6"
state: closed
author: AsToNlele
created: 2023-04-25T17:26:03Z
updated: 2023-05-17T11:38:25Z
closed: 2023-05-17T11:38:25Z
base_branch: master
head_branch: SPM-1858
labels: ["don't merge"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1033
---

# Pull Request #1033: chore: upgrade react-router-dom to v6

**Author**: @AsToNlele
**Created**: April 25, 2023 at 05:26 PM UTC
**Status**: Closed
**Labels**: `don't merge`
**Base**: `master` ← **Head**: `SPM-1858`

## Description

# Description

Associated Jira ticket: [SPM-1858](https://issues.redhat.com/browse/SPM-1858)

Upgrade of react-router-dom to v6

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on April 27, 2023 at 05:21 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`35.63`**% and project coverage change: **`+0.01`** :tada:
> Comparison is base [(`9649b5f`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/9649b5f2847d9761298e849977e306716a60e1de?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.97% compared to head [(`33b37e5`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.99%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1033      +/-   ##
==========================================
+ Coverage   63.97%   63.99%   +0.01%     
==========================================
  Files         117      118       +1     
  Lines        2859     2866       +7     
  Branches      737      737              
==========================================
+ Hits         1829     1834       +5     
- Misses       1030     1032       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | `0.00% <ø> (ø)` | |
| [src/Router.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlci5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `61.11% <0.00%> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `57.29% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `81.06% <ø> (ø)` | |
| [src/Utilities/SystemsHelpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | `21.73% <ø> (ø)` | |
| [src/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL2luZGV4Lmpz) | `0.00% <0.00%> (ø)` | |
| ... and [11 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

... and [6 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1033?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @AsToNlele on May 17, 2023 at 11:38 AM UTC

Oops messed this up, I'll make a new PR

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1033*
