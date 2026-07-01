---
type: pull_request
number: 692
title: "shrink runtime image +100MB by removing dnf cache"
state: merged
author: MichaelMraka
created: 2021-06-03T07:05:08Z
updated: 2021-06-04T12:01:40Z
closed: 2021-06-04T12:01:40Z
merged: 2021-06-04T12:01:40Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/692
---

# Pull Request #692: shrink runtime image +100MB by removing dnf cache

**Author**: @MichaelMraka
**Created**: June 03, 2021 at 07:05 AM UTC
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

### Comment by @codecov-commenter on June 03, 2021 at 07:12 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/692?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#692](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/692?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (62d2d0d) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/3302a2555d095ae635e24aa3cb5cfabf134275ce?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (3302a25) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/692/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/692?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #692   +/-   ##
=======================================
  Coverage   57.84%   57.84%           
=======================================
  Files          73       73           
  Lines        3428     3428           
=======================================
  Hits         1983     1983           
  Misses       1162     1162           
  Partials      283      283           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.84% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/692?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/692?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [3302a25...62d2d0d](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/692?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on June 04, 2021 at 11:12 AM UTC

/retest

---

## Reviews

### Review by @josef-hak - Approved on June 04, 2021 at 12:01 PM UTC

lgtm, thanks

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/692*
