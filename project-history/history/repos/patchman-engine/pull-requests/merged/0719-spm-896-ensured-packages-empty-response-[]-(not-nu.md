---
type: pull_request
number: 719
title: "SPM-896: Ensured /packages empty response [] (not \"null\")"
state: merged
author: josef-hak
created: 2021-06-24T15:40:57Z
updated: 2021-08-26T18:42:08Z
closed: 2021-06-25T07:36:05Z
merged: 2021-06-25T07:36:05Z
base_branch: master
head_branch: pkg-null
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/719
---

# Pull Request #719: SPM-896: Ensured /packages empty response [] (not "null")

**Author**: @josef-hak
**Created**: June 24, 2021 at 03:40 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkg-null`

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

### Comment by @codecov-commenter on June 24, 2021 at 03:47 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/719?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#719](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/719?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e82b623) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/48a77e42edfd59b8cb8a04dfe1af08ae1ee38061?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (48a77e4) will **not change** coverage.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/719/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/719?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #719   +/-   ##
=======================================
  Coverage   56.21%   56.21%           
=======================================
  Files          80       80           
  Lines        3693     3693           
=======================================
  Hits         2076     2076           
  Misses       1304     1304           
  Partials      313      313           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.21% <100.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/719?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/719/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `89.28% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/719?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/719?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [48a77e4...e82b623](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/719?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @semtexzv - Approved on June 25, 2021 at 06:54 AM UTC

### Review by @MichaelMraka - Approved on June 25, 2021 at 07:34 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/719*
