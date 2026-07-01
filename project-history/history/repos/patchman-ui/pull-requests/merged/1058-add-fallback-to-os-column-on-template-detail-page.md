---
type: pull_request
number: 1058
title: "Add fallback to OS column on Template detail page"
state: merged
author: leSamo
created: 2023-05-19T14:04:01Z
updated: 2023-06-05T21:42:05Z
closed: 2023-05-23T12:32:02Z
merged: 2023-05-23T12:32:02Z
base_branch: master
head_branch: fix-os-column-empty
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1058
---

# Pull Request #1058: Add fallback to OS column on Template detail page

**Author**: @leSamo
**Created**: May 19, 2023 at 02:04 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-os-column-empty`

## Description

If template system had OS data unavailable it only showed the blue icon. Correct behaviour is to show "N/A" and no blue icon. 

### Before:
![Screenshot from 2023-05-19 16-02-01](https://github.com/RedHatInsights/patchman-ui/assets/8426204/74c2b963-04f3-4e4d-ae9a-17b21e052cd9)

### After:
![Screenshot from 2023-05-19 16-02-03](https://github.com/RedHatInsights/patchman-ui/assets/8426204/faa02266-481b-4003-81a2-6a5a0dde8084)


### How to test:
- go to https://stage.foo.redhat.com:1337/beta/insights/patch/templates/8264?page=4&perPage=20&sort=os
- observe OS column contents

---

## Discussion

### Comment by @codecov-commenter on May 19, 2023 at 02:14 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1058?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`50.00`**% and project coverage change: **`-0.03`** :warning:
> Comparison is base [(`546dd7e`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/546dd7e53b01ef76706e6f371978f0203c23b437?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.03% compared to head [(`a37f142`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1058?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.01%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1058      +/-   ##
==========================================
- Coverage   63.03%   63.01%   -0.03%     
==========================================
  Files         119      119              
  Lines        2949     2950       +1     
  Branches      754      755       +1     
==========================================
  Hits         1859     1859              
- Misses       1090     1091       +1     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1058?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1058?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `56.70% <0.00%> (-0.60%)` | :arrow_down: |
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1058?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `80.85% <100.00%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1058?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @gkarat on May 23, 2023 at 10:14 AM UTC

tested locally, the issue is resolved, thank you

### Comment by @mkholjuraev on June 05, 2023 at 09:41 PM UTC

:tada: This PR is included in version 1.63.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Approved on May 23, 2023 at 10:13 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1058*
