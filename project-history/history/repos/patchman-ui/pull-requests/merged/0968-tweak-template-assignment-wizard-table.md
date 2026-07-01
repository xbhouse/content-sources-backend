---
type: pull_request
number: 968
title: "Tweak template assignment wizard table"
state: merged
author: leSamo
created: 2023-02-27T18:30:53Z
updated: 2023-05-08T18:17:04Z
closed: 2023-03-02T10:18:35Z
merged: 2023-03-02T10:18:35Z
base_branch: master
head_branch: tweak-wizard-table
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/968
---

# Pull Request #968: Tweak template assignment wizard table

**Author**: @leSamo
**Created**: February 27, 2023 at 06:30 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `tweak-wizard-table`

## Description

# Description

Mockup: https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/eKRw04p

Tweak template assignment wizard table:
- add sorting to Systems table inside Template creation modal
- display 'N/A' instead of blank cell when OS version in not present for a system
- TableView: Skeleton table now has the same number of columns as the table it's representing

# How to test the PR

Verify the first above 2 bullet points in template assignment wizard.
Verify the last bullet point on all Patchman tables.


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 27, 2023 at 06:32 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/968?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.06**% // Head: **72.06**% // Decreases project coverage by **`-0.01%`** :warning:
> Coverage data is based on head [(`25c4075`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/968?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`3ef909f`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/3ef909f402e617205c86b2e63785d053ee7d499c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 80.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #968      +/-   ##
==========================================
- Coverage   72.06%   72.06%   -0.01%     
==========================================
  Files         110      110              
  Lines        2617     2620       +3     
  Branches      662      665       +3     
==========================================
+ Hits         1886     1888       +2     
- Misses        731      732       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/968?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/968?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `29.03% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/968?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `72.50% <0.00%> (-0.92%)` | :arrow_down: |
| [...rc/PresentationalComponents/TableView/TableView.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/968?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3Lmpz) | `100.00% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/968?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 02, 2023 at 11:16 AM UTC

:tada: This PR is included in version 1.62.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on March 02, 2023 at 10:17 AM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/968*
