---
type: pull_request
number: 1038
title: "Add ReadHeaderTimeout"
state: merged
author: michalslomczynski
created: 2022-07-27T08:06:13Z
updated: 2022-07-27T13:57:00Z
closed: 2022-07-27T13:57:00Z
merged: 2022-07-27T13:57:00Z
base_branch: master
head_branch: header-tout
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1038
---

# Pull Request #1038: Add ReadHeaderTimeout

**Author**: @michalslomczynski
**Created**: July 27, 2022 at 08:06 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `header-tout`

## Description

```
base/utils/gin.go:59:9: G112: Potential Slowloris Attack because ReadHeaderTimeout is not configured in the http.Server (gosec)
        srv := http.Server{Addr: addr, Handler: handler}
```
I've set same value as is default in nginx

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

### Comment by @jira-linking on July 27, 2022 at 08:06 AM UTC

Commits missing Jira IDs:
b130143b708e825d1adc74c90c081ebbfe35f0ff


### Comment by @codecov-commenter on July 27, 2022 at 08:14 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1038?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1038](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1038?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (b130143) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ba8d228d4a20d4402182068f745348ca5112a1fa?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ba8d228) will **not change** coverage.
> The diff coverage is `100.00%`.

```diff
@@           Coverage Diff           @@
##           master    #1038   +/-   ##
=======================================
  Coverage   61.25%   61.25%           
=======================================
  Files          95       95           
  Lines        5448     5448           
=======================================
  Hits         3337     3337           
  Misses       1682     1682           
  Partials      429      429           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.25% <100.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1038?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1038/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | `22.85% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1038?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1038?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [864844e...b130143](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1038?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on July 27, 2022 at 01:45 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1038*
