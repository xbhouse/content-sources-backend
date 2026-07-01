---
type: pull_request
number: 1486
title: "RHINENG-12951: fix snyk findings"
state: merged
author: MichaelMraka
created: 2024-09-25T14:34:42Z
updated: 2024-09-26T09:01:17Z
closed: 2024-09-26T09:01:17Z
merged: 2024-09-26T09:01:17Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1486
---

# Pull Request #1486: RHINENG-12951: fix snyk findings

**Author**: @MichaelMraka
**Created**: September 25, 2024 at 02:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [x] Output Encoding
- [x] Authentication and Password Management
- [x] Session Management
- [x] Access Control
- [x] Cryptographic Practices
- [x] Error Handling and Logging
- [x] Data Protection
- [x] Communication Security
- [x] System Configuration
- [x] Database Security
- [x] File Management
- [x] Memory Management
- [x] General Coding Practices


---

## Discussion

### Comment by @jira-linking on September 25, 2024 at 02:34 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-12951


### Comment by @codecov-commenter on September 25, 2024 at 02:40 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1486?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `72.72727%` with `3 lines` in your changes missing coverage. Please review.
> Project coverage is 65.03%. Comparing base [(`cb13c2f`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/cb13c2f866ebdc2c5b503e6297630aeb60362de8?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`8cb59a9`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8cb59a9062de93ae8eccf44edca5b44bee1c6bdb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1486?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/mqueue/mqueue\_impl\_gokafka.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1486?src=pr&el=tree&filepath=base%2Fmqueue%2Fmqueue_impl_gokafka.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | 25.00% | [1 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1486?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1486   +/-   ##
=======================================
  Coverage   65.03%   65.03%           
=======================================
  Files         114      114           
  Lines        7877     7880    +3     
=======================================
+ Hits         5123     5125    +2     
+ Misses       2217     2216    -1     
- Partials      537      539    +2     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1486/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1486/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.03% <72.72%> (+<0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1486?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on September 26, 2024 at 08:44 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1486*
