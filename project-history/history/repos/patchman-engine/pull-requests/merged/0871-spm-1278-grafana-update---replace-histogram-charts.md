---
type: pull_request
number: 871
title: "SPM-1278: Grafana update - replace histogram charts"
state: merged
author: josef-hak
created: 2021-12-07T16:07:46Z
updated: 2021-12-07T17:22:37Z
closed: 2021-12-07T17:21:48Z
merged: 2021-12-07T17:21:48Z
base_branch: master
head_branch: grafana
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/871
---

# Pull Request #871: SPM-1278: Grafana update - replace histogram charts

**Author**: @josef-hak
**Created**: December 07, 2021 at 04:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `grafana`

## Description

- histogram charts replaced with standard charts (easier to understand)
- removed some less interesting charts to keep the board clear

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

### Comment by @codecov-commenter on December 07, 2021 at 04:18 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/871?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#871](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/871?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (acbac74) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0bb4c8dbc70ff4c8eb980b523b869ceffd890b12?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0bb4c8d) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/871/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/871?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #871   +/-   ##
=======================================
  Coverage   58.77%   58.77%           
=======================================
  Files          81       81           
  Lines        4245     4245           
=======================================
  Hits         2495     2495           
  Misses       1396     1396           
  Partials      354      354           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.77% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/871?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/871?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [b447b55...acbac74](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/871?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on December 07, 2021 at 04:26 PM UTC

/retest

### Comment by @josef-hak on December 07, 2021 at 05:03 PM UTC

/retest

### Comment by @josef-hak on December 07, 2021 at 05:19 PM UTC

/retest

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/871*
