---
type: pull_request
number: 1556
title: "WIP: logical replica publication/subscription"
state: closed
author: psegedy
created: 2025-01-15T17:45:34Z
updated: 2025-01-21T16:14:19Z
closed: 2025-01-21T16:14:19Z
base_branch: master
head_branch: rds_pub_sub
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1556
---

# Pull Request #1556: WIP: logical replica publication/subscription

**Author**: @psegedy
**Created**: January 15, 2025 at 05:45 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `rds_pub_sub`

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

### Comment by @jira-linking on January 15, 2025 at 05:45 PM UTC

Commits missing Jira IDs:
88183f9bd000c2f8c8335cd590d2fa9090b23fe8


### Comment by @codecov-commenter on January 15, 2025 at 05:50 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `5.10204%` with `93 lines` in your changes missing coverage. Please review.
> Project coverage is 57.49%. Comparing base [(`f4c3bfe`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f4c3bfef73ecd11b2d41113d801dada294eccef9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`88183f9`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/88183f9bd000c2f8c8335cd590d2fa9090b23fe8?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [tasks/pubsub/pubsub.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?src=pr&el=tree&filepath=tasks%2Fpubsub%2Fpubsub.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvcHVic3ViL3B1YnN1Yi5nbw==) | 0.00% | [70 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/setup.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?src=pr&el=tree&filepath=base%2Fdatabase%2Fsetup.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | 28.57% | [8 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [tasks/common.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?src=pr&el=tree&filepath=tasks%2Fcommon.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3MvY29tbW9uLmdv) | 0.00% | [7 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/utils/config.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?src=pr&el=tree&filepath=base%2Futils%2Fconfig.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | 20.00% | [3 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?src=pr&el=tree&filepath=base%2Fdatabase%2Futils.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1556      +/-   ##
==========================================
- Coverage   57.96%   57.49%   -0.48%     
==========================================
  Files         143      144       +1     
  Lines       11131    11224      +93     
==========================================
+ Hits         6452     6453       +1     
- Misses       4100     4189      +89     
- Partials      579      582       +3     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.49% <5.10%> (-0.48%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1556?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on January 21, 2025 at 04:14 PM UTC

will be done by AWS blue/green deployments

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1556*
