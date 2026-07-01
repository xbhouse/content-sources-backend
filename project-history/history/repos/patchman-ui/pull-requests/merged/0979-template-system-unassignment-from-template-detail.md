---
type: pull_request
number: 979
title: "Template system unassignment from Template detail page"
state: merged
author: leSamo
created: 2023-03-03T14:21:11Z
updated: 2023-05-08T18:17:01Z
closed: 2023-03-07T13:14:11Z
merged: 2023-03-07T13:14:10Z
base_branch: master
head_branch: unassign-system-from-template
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/979
---

# Pull Request #979: Template system unassignment from Template detail page

**Author**: @leSamo
**Created**: March 03, 2023 at 02:21 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `unassign-system-from-template`

## Description

# Description

Associated Jira ticket: [SPM-1904](https://issues.redhat.com/browse/SPM-1904), [SPM-1907](https://issues.redhat.com/browse/SPM-1907), [SPM-1927](https://issues.redhat.com/browse/SPM-1927)

From System list page:
It was not possible to remove template from stale systems. This is now fixed.

From Template detail page:
Add table row actions accessible with patch:template:write permission, these actions contain action to unassign system from a template.

## Testing instructions

**a) System list page:**
Find a system that has template assigned.

i) Unassign system from template using row actions kebab.
ii) Unassign multiple systems from template using bulk action kebab in the table toolbar. Make sure to try to select some systems that do have template and some that don't - modal should then point out to you that some systems don't have a template and won't be affected.

**b) Template detail page:**
Unassign system from template using row actions kebab.
Bulk action is not currently implemented.

---

## Discussion

### Comment by @codecov-commenter on March 03, 2023 at 02:28 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/979?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`50.00`**% and project coverage change: **`-0.11`** :warning:
> Comparison is base [(`058b068`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/058b068f9a734efcd3bc7e590a0dbd456b478f66?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 69.02% compared to head [(`d2099ab`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/979?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 68.91%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #979      +/-   ##
==========================================
- Coverage   69.02%   68.91%   -0.11%     
==========================================
  Files         115      115              
  Lines        2786     2796      +10     
  Branches      705      708       +3     
==========================================
+ Hits         1923     1927       +4     
- Misses        863      869       +6     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/979?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/979?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/979?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `69.76% <ø> (ø)` | |
| [src/SmartComponents/Modals/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/979?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvSGVscGVycy5qcw==) | `100.00% <100.00%> (ø)` | |
| [src/SmartComponents/Modals/UnassignSystemsModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/979?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvVW5hc3NpZ25TeXN0ZW1zTW9kYWwuanM=) | `95.83% <100.00%> (+0.59%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/979?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 09, 2023 at 09:08 AM UTC

:tada: This PR is included in version 1.62.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on March 07, 2023 at 12:51 PM UTC

Works as expected! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/979*
