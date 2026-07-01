---
type: pull_request
number: 1055
title: "Update zookeeper in docker-compose.test"
state: merged
author: psegedy
created: 2022-08-11T09:58:42Z
updated: 2022-08-11T13:52:54Z
closed: 2022-08-11T13:52:54Z
merged: 2022-08-11T13:52:54Z
base_branch: master
head_branch: update-kafka-test
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1055
---

# Pull Request #1055: Update zookeeper in docker-compose.test

**Author**: @psegedy
**Created**: August 11, 2022 at 09:58 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `update-kafka-test`

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

### Comment by @jira-linking on August 11, 2022 at 09:58 AM UTC

Commits missing Jira IDs:
db17aeca7e0fc26fe1d5e30ea40460006cc47d0a


### Comment by @codecov-commenter on August 11, 2022 at 10:06 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1055?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1055](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1055?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (db17aec) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/c46836e99f12af6e684b35b904357429a071caab?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c46836e) will **decrease** coverage by `0.05%`.
> The diff coverage is `n/a`.

```diff
@@            Coverage Diff             @@
##           master    #1055      +/-   ##
==========================================
- Coverage   61.48%   61.43%   -0.06%     
==========================================
  Files          98       98              
  Lines        5432     5432              
==========================================
- Hits         3340     3337       -3     
- Misses       1665     1668       +3     
  Partials      427      427              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.43% <ø> (-0.06%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1055/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `72.85% <0.00%> (-4.29%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


---

## Reviews

### Review by @MichaelMraka - Approved on August 11, 2022 at 01:52 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1055*
