---
type: pull_request
number: 1000
title: "Fix template up-to date being off by a day in negative UTC timezones"
state: merged
author: leSamo
created: 2023-03-24T10:50:52Z
updated: 2023-04-17T18:42:10Z
closed: 2023-04-03T18:14:15Z
merged: 2023-04-03T18:14:15Z
base_branch: master
head_branch: fix-template-date
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1000
---

# Pull Request #1000: Fix template up-to date being off by a day in negative UTC timezones

**Author**: @leSamo
**Created**: March 24, 2023 at 10:50 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-template-date`

## Description

- Fix template up-to date being off by a day in negative UTC timezones
- Add hasGutter to header so there's margin between long template names and dropdown, after screenshot:
![Screenshot from 2023-03-24 06-43-08](https://user-images.githubusercontent.com/8426204/227501325-593206fe-42aa-423d-b6a1-2c5743b8f610.png)

### How to test:
1. Set your timezone to some place with negative UTC offset, for example New York. 
2. Create a template with some date.
3. In the last (review) step of the template creation check if the date is the same as you set.
4. Go to detail page of your just created template.
5. Check if the date in the header is the same as you set.
6. Don't forget to set your timezone back :grinning: 

---

## Discussion

### Comment by @codecov-commenter on April 03, 2023 at 06:01 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1000?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`90.00`**% and project coverage change: **`+0.40`** :tada:
> Comparison is base [(`d705f0e`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/d705f0e7d04b79b30dc024dd9e65329bc7503122?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.22% compared to head [(`bac9681`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1000?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.62%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1000      +/-   ##
==========================================
+ Coverage   64.22%   64.62%   +0.40%     
==========================================
  Files         116      116              
  Lines        2812     2844      +32     
  Branches      723      726       +3     
==========================================
+ Hits         1806     1838      +32     
  Misses       1006     1006              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1000?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/PresentationalComponents/Header/Header.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1000?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9IZWFkZXIvSGVhZGVyLmpz) | `100.00% <ø> (ø)` | |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1000?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...tComponents/PatchSetWizard/steps/ReviewPatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1000?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdQYXRjaFNldC5qcw==) | `100.00% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1000?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `81.06% <100.00%> (+0.58%)` | :arrow_up: |

... and [2 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1000/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1000?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on April 17, 2023 at 06:41 PM UTC

:tada: This PR is included in version 1.63.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on April 03, 2023 at 04:52 PM UTC

LGTM! @leSamo thank you.

### Review by @leSamo - Commented on April 03, 2023 at 05:26 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1000*
