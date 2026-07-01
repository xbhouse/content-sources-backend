---
type: pull_request
number: 1181
title: "refactor(TemplateWizard): Remove obsolete code"
state: merged
author: leSamo
created: 2024-04-24T22:28:33Z
updated: 2024-05-13T13:41:45Z
closed: 2024-05-13T13:24:23Z
merged: 2024-05-13T13:24:23Z
base_branch: master
head_branch: remove-template-code
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1181
---

# Pull Request #1181: refactor(TemplateWizard): Remove obsolete code

**Author**: @leSamo
**Created**: April 24, 2024 at 10:28 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remove-template-code`

## Description

- removed commented in docs template link that was planned, but later scrapped
- removed dead code in `ConfigurationStepFields` which was unreachable due to `shouldShowRadioButtons = false;`, this code was only used for migrations between two development versions of Template wizard

---

## Discussion

### Comment by @codecov-commenter on April 24, 2024 at 10:40 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1181?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `50.00000%` with `1 lines` in your changes are missing coverage. Please review.
> Project coverage is 64.15%. Comparing base [(`42e843c`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/42e843cd45f2fa8abd04d8ccd7a64b008913e1a6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`d0112e6`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1181?dropdown=coverage&src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1181?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1181?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2Fsteps%2FConfigurationStepFields.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | 50.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1181?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1181      +/-   ##
==========================================
+ Coverage   63.84%   64.15%   +0.30%     
==========================================
  Files         124      124              
  Lines        3225     3208      -17     
  Branches      826      818       -8     
==========================================
- Hits         2059     2058       -1     
+ Misses       1166     1150      -16     
```



</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1181?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @LiorKGOW on April 25, 2024 at 11:26 AM UTC

/retest

### Comment by @leSamo on April 25, 2024 at 11:51 AM UTC

@LiorKGOW I don't think we will be able to get the PR check passing since the QA tests are faulty I believe

### Comment by @leSamo on April 29, 2024 at 12:45 PM UTC

/retest

### Comment by @niyazRedhat on May 02, 2024 at 01:30 PM UTC

/retest



### Comment by @mkholjuraev on May 13, 2024 at 01:41 PM UTC

:tada: This PR is included in version 1.67.5 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.5)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on May 13, 2024 at 09:32 AM UTC

LGTM! tested locally, no issues I encountered

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1181*
