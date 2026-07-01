---
type: pull_request
number: 965
title: "feat: Template detail page"
state: merged
author: leSamo
created: 2023-02-21T01:06:11Z
updated: 2023-05-08T18:17:03Z
closed: 2023-02-28T09:59:53Z
merged: 2023-02-28T09:59:53Z
base_branch: master
head_branch: patch-detail-page
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/965
---

# Pull Request #965: feat: Template detail page

**Author**: @leSamo
**Created**: February 21, 2023 at 01:06 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `patch-detail-page`

## Description

# Description

Associated Jira tickets: [SPM-1510](https://issues.redhat.com/browse/SPM-1510), [SPM-1901](https://issues.redhat.com/browse/SPM-1901)

Please include a summary of the change, what this fixes/creates/improves.


# How to test the PR

Please include steps to test your PR.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 21, 2023 at 01:13 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.37**% // Head: **71.72**% // Decreases project coverage by **`-0.65%`** :warning:
> Coverage data is based on head [(`0f719ad`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`41668f2`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/41668f27e9a0a506ad08fe24c2bde5487e428661?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 25.96% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #965      +/-   ##
==========================================
- Coverage   72.37%   71.72%   -0.65%     
==========================================
  Files         110      113       +3     
  Lines        2617     2794     +177     
  Branches      658      724      +66     
==========================================
+ Hits         1894     2004     +110     
- Misses        723      790      +67     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `54.54% <ø> (ø)` | |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `72.50% <ø> (-0.92%)` | :arrow_down: |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `100.00% <ø> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `59.09% <35.00%> (+0.90%)` | :arrow_up: |
| [src/store/Reducers/PatchSetDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0RGV0YWlsU3RvcmUuanM=) | `42.85% <42.85%> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `83.33% <50.00%> (-1.75%)` | :arrow_down: |
| [src/SmartComponents/Modals/DeleteSetModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvRGVsZXRlU2V0TW9kYWwuanM=) | `60.00% <60.00%> (ø)` | |
| [src/PresentationalComponents/Header/Header.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9IZWFkZXIvSGVhZGVyLmpz) | `100.00% <100.00%> (ø)` | |
| ... and [12 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/965?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on February 28, 2023 at 10:14 AM UTC

:tada: This PR is included in version 1.61.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.61.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.61.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Commented on February 27, 2023 at 11:00 AM UTC

### Review by @leSamo - Commented on February 27, 2023 at 11:04 AM UTC

### Review by @leSamo - Commented on February 27, 2023 at 11:10 AM UTC

### Review by @leSamo - Commented on February 27, 2023 at 12:30 PM UTC

### Review by @leSamo - Commented on February 27, 2023 at 02:06 PM UTC

### Review by @leSamo - Commented on February 27, 2023 at 02:17 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/965*
