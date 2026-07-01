---
type: pull_request
number: 1080
title: "Handle 504 errors with correct empty state"
state: merged
author: leSamo
created: 2023-06-30T14:48:47Z
updated: 2023-08-14T13:21:55Z
closed: 2023-06-30T15:46:46Z
merged: 2023-06-30T15:46:46Z
base_branch: master
head_branch: handle-504
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1080
---

# Pull Request #1080: Handle 504 errors with correct empty state

**Author**: @leSamo
**Created**: June 30, 2023 at 02:48 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `handle-504`

## Description

Ticket: https://issues.redhat.com/browse/SPM-2131

Show correct empty state when 504 (Gateway timeout) is received in API response.

## Before:
![Screenshot from 2023-06-30 16-47-51](https://github.com/RedHatInsights/patchman-ui/assets/8426204/4c534b50-b4eb-48ca-a495-ffdf25ecbbda)

## After:
![Screenshot from 2023-06-30 16-47-39](https://github.com/RedHatInsights/patchman-ui/assets/8426204/0f312ff8-6c9d-452c-b27c-23431ec3b36b)


---

## Discussion

### Comment by @codecov-commenter on June 30, 2023 at 02:56 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1080?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`-0.03`** :warning:
> Comparison is base [(`ab52aba`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ab52aba51d27a8d7ba89c9a0f28749402b136fc3?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.87% compared to head [(`f3dd95a`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1080?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.85%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1080      +/-   ##
==========================================
- Coverage   62.87%   62.85%   -0.03%     
==========================================
  Files         119      119              
  Lines        2966     2967       +1     
  Branches      761      762       +1     
==========================================
  Hits         1865     1865              
- Misses       1101     1102       +1     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1080?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [.../PresentationalComponents/Snippets/ErrorHandler.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1080?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FcnJvckhhbmRsZXIuanM=) | `39.13% <0.00%> (-1.78%)` | :arrow_down: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1080?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on July 03, 2023 at 04:56 PM UTC

:tada: This PR is included in version 1.63.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on June 30, 2023 at 03:20 PM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1080*
