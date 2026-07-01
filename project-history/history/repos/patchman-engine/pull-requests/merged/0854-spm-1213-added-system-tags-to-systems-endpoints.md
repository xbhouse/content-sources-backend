---
type: pull_request
number: 854
title: "SPM-1213: Added system tags to /*systems endpoints"
state: merged
author: josef-hak
created: 2021-10-27T21:19:55Z
updated: 2021-11-01T12:35:31Z
closed: 2021-11-01T11:13:38Z
merged: 2021-11-01T11:13:38Z
base_branch: master
head_branch: api-tags
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/854
---

# Pull Request #854: SPM-1213: Added system tags to /*systems endpoints

**Author**: @josef-hak
**Created**: October 27, 2021 at 09:19 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `api-tags`

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

### Comment by @codecov-commenter on November 01, 2021 at 07:53 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#854](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (918d9da) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/90920d6084591e8e0e7c80762f73319d99bf1261?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (90920d6) will **decrease** coverage by `0.15%`.
> The diff coverage is `75.75%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #854      +/-   ##
==========================================
- Coverage   58.64%   58.48%   -0.16%     
==========================================
  Files          81       81              
  Lines        4050     4059       +9     
==========================================
- Hits         2375     2374       -1     
- Misses       1346     1351       +5     
- Partials      329      334       +5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.48% <75.75%> (-0.16%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `82.50% <60.00%> (-9.40%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `82.77% <71.42%> (-0.71%)` | :arrow_down: |
| [base/database/query.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9xdWVyeS5nbw==) | `80.00% <100.00%> (-2.28%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `71.05% <100.00%> (-4.51%)` | :arrow_down: |
| [manager/controllers/advisory\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `48.57% <100.00%> (-1.43%)` | :arrow_down: |
| [manager/controllers/systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2V4cG9ydC5nbw==) | `78.26% <100.00%> (-0.91%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [9cdcb9d...918d9da](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/854?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on November 01, 2021 at 09:18 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on November 01, 2021 at 09:35 AM UTC

LGTM

There's just  `parsin` -> `parsing` typo in commit message:
`SPM-1213: Added tags parsin to /systems_export endpoint and test.`

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/854*
