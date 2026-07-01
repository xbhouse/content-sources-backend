---
type: pull_request
number: 990
title: "Decouple step title and sidebar name in Template Wizard"
state: merged
author: leSamo
created: 2023-03-10T23:49:26Z
updated: 2023-05-08T18:16:40Z
closed: 2023-03-13T17:35:37Z
merged: 2023-03-13T17:35:37Z
base_branch: master
head_branch: decouple-titles
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/990
---

# Pull Request #990: Decouple step title and sidebar name in Template Wizard

**Author**: @leSamo
**Created**: March 10, 2023 at 11:49 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `decouple-titles`

## Description

# Description

Associated Jira ticket: [SPM-1893](https://issues.redhat.com/browse/SPM-1893)

- until now the title of the step had to be the same for header label and sidebar label, but in the future we will need them to be different so I decoupled these strings by removing `showTitles: true` and manually adding header to each step
- update System selection step text to the [latest mockups](https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/eKRw04p)
- if user leaves description field empty, review step will state 'No description provided' instead of 'No description available'


---

## Discussion

### Comment by @codecov-commenter on March 10, 2023 at 11:57 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/990?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`66.66`**% and project coverage change: **`-0.03`** :warning:
> Comparison is base [(`d3e57f8`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/d3e57f8eac672dd4fa927fe4a54571aa2dad84fb?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 69.03% compared to head [(`907b784`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/990?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 69.01%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #990      +/-   ##
==========================================
- Coverage   69.03%   69.01%   -0.03%     
==========================================
  Files         115      115              
  Lines        2800     2801       +1     
  Branches      711      712       +1     
==========================================
  Hits         1933     1933              
- Misses        867      868       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/990?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/990?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <ø> (ø)` | |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/990?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `29.03% <ø> (ø)` | |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/990?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `4.65% <ø> (ø)` | |
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/990?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | `7.14% <50.00%> (-0.27%)` | :arrow_down: |
| [...tComponents/PatchSetWizard/steps/ReviewPatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/990?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdQYXRjaFNldC5qcw==) | `100.00% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/990?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 28, 2023 at 08:57 AM UTC

:tada: This PR is included in version 1.62.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on March 13, 2023 at 05:35 PM UTC

LGTM! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/990*
