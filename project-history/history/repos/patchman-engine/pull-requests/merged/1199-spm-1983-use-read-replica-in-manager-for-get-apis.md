---
type: pull_request
number: 1199
title: "SPM-1983: use read replica in manager for GET APIs"
state: merged
author: psegedy
created: 2023-03-22T16:03:21Z
updated: 2023-03-23T13:21:48Z
closed: 2023-03-23T13:21:47Z
merged: 2023-03-23T13:21:47Z
base_branch: master
head_branch: read_replica
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1199
---

# Pull Request #1199: SPM-1983: use read replica in manager for GET APIs

**Author**: @psegedy
**Created**: March 22, 2023 at 04:03 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `read_replica`

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

### Comment by @codecov-commenter on March 22, 2023 at 04:48 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1199?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`34.37`**% and project coverage change: **`-0.15`** :warning:
> Comparison is base [(`b42e944`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b42e9449d607af037aa9e0a69bce60a436608fbb?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.11% compared to head [(`09688a3`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1199?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.97%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1199      +/-   ##
==========================================
- Coverage   62.11%   61.97%   -0.15%     
==========================================
  Files         103      103              
  Lines        6205     6232      +27     
==========================================
+ Hits         3854     3862       +8     
- Misses       1859     1874      +15     
- Partials      492      496       +4     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.97% <34.37%> (-0.15%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1199?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/middlewares/db.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1199?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9kYi5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/database/setup.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1199?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | `68.42% <41.66%> (-6.21%)` | :arrow_down: |
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1199?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `61.33% <46.15%> (-1.98%)` | :arrow_down: |

... and [1 file with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1199/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1199?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on March 23, 2023 at 12:21 PM UTC

@MichaelMraka https://github.com/RedHatInsights/patchman-engine/pull/1199/files#diff-7f68e4ce34116e0e20004477ac345d1b97b835778becfff67c60aaea3f4d5d7eR35, so if `DBReadReplicaEnabled` is enabled (which is only set in manager) and it will be only used for `http.MethodGet` method

---

## Reviews

### Review by @MichaelMraka - Commented on March 23, 2023 at 10:19 AM UTC

I don't see/understand how/where is configured which APIs can connect to read replica and which have to go to primary db.

### Review by @MichaelMraka - Approved on March 23, 2023 at 12:49 PM UTC

Thanks Patrik, that's nice and clever solution.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1199*
