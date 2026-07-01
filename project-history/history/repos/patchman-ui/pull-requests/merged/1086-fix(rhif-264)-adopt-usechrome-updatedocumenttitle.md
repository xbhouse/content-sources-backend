---
type: pull_request
number: 1086
title: "fix(RHIF-264):  adopt useChrome updateDocumentTitle hook"
state: merged
author: adonispuente
created: 2023-07-03T21:42:32Z
updated: 2023-07-18T14:33:24Z
closed: 2023-07-18T14:14:54Z
merged: 2023-07-18T14:14:54Z
base_branch: master
head_branch: RHIF-264
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1086
---

# Pull Request #1086: fix(RHIF-264):  adopt useChrome updateDocumentTitle hook

**Author**: @adonispuente
**Created**: July 03, 2023 at 09:42 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `RHIF-264`

## Description

# Description
A partial cause of https://issues.redhat.com/browse/RHIF-207 is that the document title isnt updated properly. This PR introduces the chrome hook.
Associated Jira ticket: # (RHIF-266)

# How to test the PR
Navigate to all pages and verify that the document changes appropriately 


# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on July 03, 2023 at 09:51 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`69.56`**% and project coverage change: **`-0.03`** :warning:
> Comparison is base [(`0784bf8`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/0784bf8a4d2b94b033a2ce6ff6596e2901702698?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.84% compared to head [(`17d74fe`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.81%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1086      +/-   ##
==========================================
- Coverage   62.84%   62.81%   -0.03%     
==========================================
  Files         119      119              
  Lines        2971     2977       +6     
  Branches      763      765       +2     
==========================================
+ Hits         1867     1870       +3     
- Misses       1104     1107       +3     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/Hooks.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `73.50% <ø> (-0.52%)` | :arrow_down: |
| [...c/SmartComponents/AdvisoryDetail/AdvisoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeURldGFpbC9BZHZpc29yeURldGFpbC5qcw==) | `95.65% <75.00%> (-4.35%)` | :arrow_down: |
| [src/SmartComponents/Advisories/Advisories.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `100.00% <100.00%> (ø)` | |
| [src/SmartComponents/PackageDetail/PackageDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlRGV0YWlsL1BhY2thZ2VEZXRhaWwuanM=) | `91.66% <100.00%> (+0.36%)` | :arrow_up: |
| [src/SmartComponents/Packages/Packages.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlcy9QYWNrYWdlcy5qcw==) | `100.00% <100.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `80.59% <100.00%> (+0.29%)` | :arrow_up: |
| [src/Utilities/constants.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <100.00%> (ø)` | |

... and [2 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1086?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @adonispuente on July 12, 2023 at 03:21 PM UTC

@leSamo I agree, the dependencies should be consistent and intl doesnt really need to be a watcher. I've sent up the changes, thanks for taking a look at this!

### Comment by @mkholjuraev on July 18, 2023 at 02:33 PM UTC

:tada: This PR is included in version 1.63.8 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.8)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Changes Requested on July 07, 2023 at 10:22 AM UTC

Thank you, I have only two small nitpicks -- one is in a separate comment, and the other one is that the dependency arrays of `useEffect` hook are a bit inconsistent, sometimes they include `intl`, sometimes not, I think it doesn't need to include  `intl` only `chrome`. But if it's consistent across all `useEffect` hooks related to titles it's okay with me.

### Review by @leSamo - Approved on July 18, 2023 at 02:14 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1086*
