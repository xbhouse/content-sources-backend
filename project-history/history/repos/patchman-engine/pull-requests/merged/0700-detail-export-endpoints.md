---
type: pull_request
number: 700
title: "Detail export endpoints"
state: merged
author: semtexzv
created: 2021-06-08T12:01:01Z
updated: 2021-06-15T12:28:43Z
closed: 2021-06-15T12:28:43Z
merged: 2021-06-15T12:28:43Z
base_branch: master
head_branch: detail-endpoints
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/700
---

# Pull Request #700: Detail export endpoints

**Author**: @semtexzv
**Created**: June 08, 2021 at 12:01 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `detail-endpoints`

## Description

Add export endpoints for detail lists.

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

### Comment by @codecov-commenter on June 11, 2021 at 12:53 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#700](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4a95f2b) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/d5cf47f8c47487f0f7c4c2ff1dbdb8329e879ca9?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d5cf47f) will **decrease** coverage by `0.08%`.
> The diff coverage is `52.05%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #700      +/-   ##
==========================================
- Coverage   56.55%   56.46%   -0.09%     
==========================================
  Files          75       77       +2     
  Lines        3517     3588      +71     
==========================================
+ Hits         1989     2026      +37     
- Misses       1245     1266      +21     
- Partials      283      296      +13     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.46% <52.05%> (-0.09%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.55% <50.00%> (ø)` | |
| [manager/controllers/advisory\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `50.00% <50.00%> (ø)` | |
| [manager/controllers/system\_advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllc19leHBvcnQuZ28=) | `54.28% <54.28%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [d5cf47f...4a95f2b](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/700?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Changes Requested on June 09, 2021 at 07:38 AM UTC

We use `go_test.sh` script in github PR tests where it's better to see all failing tests at once.

Maybe you can introduce some env variable for people who prefer to run it with `-fastfail` in local deployment.

### Review by @MichaelMraka - Approved on June 14, 2021 at 08:14 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/700*
