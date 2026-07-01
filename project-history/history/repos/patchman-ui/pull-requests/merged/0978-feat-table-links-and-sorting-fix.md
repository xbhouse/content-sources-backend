---
type: pull_request
number: 978
title: "feat: Table links and sorting fix"
state: merged
author: leSamo
created: 2023-03-02T18:43:50Z
updated: 2023-05-08T18:16:59Z
closed: 2023-03-07T14:28:31Z
merged: 2023-03-07T14:28:31Z
base_branch: master
head_branch: various-tweaks
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/978
---

# Pull Request #978: feat: Table links and sorting fix

**Author**: @leSamo
**Created**: March 02, 2023 at 06:43 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `various-tweaks`

## Description

# Description
- Link from Advisory detail to Template detail page
- Fix sorting on Template detail page
- Link from Template detail to Systems detail page

## Testing instructions
a) Go to Advisory detail page and in the table find a system that has a template. Verify you can navigate to Template detail page by clicking on the system name inside the table.
b) In Template detail page you should be able to correctly sort the table by display name. All other columns should have sorting indicators disabled.
c) Go to Template detail page and click on any system name inside the first column. You should be correctly navigated to Systems detail page of clicked system.



---

## Discussion

### Comment by @codecov-commenter on March 02, 2023 at 06:49 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`16.66`**% and project coverage change: **`+0.03`** :tada:
> Comparison is base [(`058b068`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/058b068f9a734efcd3bc7e590a0dbd456b478f66?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 69.02% compared to head [(`30b70a0`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 69.05%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #978      +/-   ##
==========================================
+ Coverage   69.02%   69.05%   +0.03%     
==========================================
  Files         115      115              
  Lines        2786     2789       +3     
  Branches      705      706       +1     
==========================================
+ Hits         1923     1926       +3     
  Misses        863      863              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...tComponents/PatchSetDetail/PatchSetDetailAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbEFzc2V0cy5qcw==) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `72.41% <0.00%> (-5.37%)` | :arrow_down: |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `69.76% <ø> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `59.09% <0.00%> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `84.61% <50.00%> (+0.30%)` | :arrow_up: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `75.00% <0.00%> (+0.65%)` | :arrow_up: |
| [src/Utilities/unitTestingUtilities.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91bml0VGVzdGluZ1V0aWxpdGllcy5qcw==) | `32.43% <0.00%> (+2.70%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/978?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 09, 2023 at 09:08 AM UTC

:tada: This PR is included in version 1.62.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on March 07, 2023 at 12:46 PM UTC

LGTM! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/978*
