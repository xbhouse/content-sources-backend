---
type: pull_request
number: 1148
title: "SPM-1828: fix response of views apis"
state: merged
author: psegedy
created: 2022-12-07T12:44:00Z
updated: 2022-12-08T12:28:22Z
closed: 2022-12-08T12:28:22Z
merged: 2022-12-08T12:28:22Z
base_branch: master
head_branch: views_api
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1148
---

# Pull Request #1148: SPM-1828: fix response of views apis

**Author**: @psegedy
**Created**: December 07, 2022 at 12:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `views_api`

## Description

- systems with no advisories in DB (no record in system_advisories table) should be included in response
- return empty list instead of list containing empty string [""]

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

### Comment by @jira-linking on December 07, 2022 at 12:44 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1828


### Comment by @codecov-commenter on December 07, 2022 at 12:53 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1148?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.24**% // Head: **62.19**% // Decreases project coverage by **`-0.05%`** :warning:
> Coverage data is based on head [(`b2ddec1`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1148?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`78a2655`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/78a26552f318c41c98082f89684b85a7cc589ef1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 22.22% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1148      +/-   ##
==========================================
- Coverage   62.24%   62.19%   -0.06%     
==========================================
  Files         100      100              
  Lines        5816     5816              
==========================================
- Hits         3620     3617       -3     
- Misses       1732     1733       +1     
- Partials      464      466       +2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.19% <22.22%> (-0.06%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1148?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/middlewares/logger.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1148/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9sb2dnZXIuZ28=) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/systems\_advisories\_view.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1148/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | `75.47% <25.00%> (-4.53%)` | :arrow_down: |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1148/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `72.85% <0.00%> (-4.29%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1148?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on December 07, 2022 at 01:04 PM UTC

/retry

### Comment by @psegedy on December 07, 2022 at 01:04 PM UTC

/retest


---

## Reviews

### Review by @michalslomczynski - Approved on December 07, 2022 at 02:24 PM UTC

### Review by @MichaelMraka - Approved on December 08, 2022 at 10:24 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1148*
