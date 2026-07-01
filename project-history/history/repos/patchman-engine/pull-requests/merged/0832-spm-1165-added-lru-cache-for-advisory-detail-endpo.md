---
type: pull_request
number: 832
title: "SPM-1165: Added LRU cache for advisory detail endpoint"
state: merged
author: josef-hak
created: 2021-09-16T14:07:51Z
updated: 2021-09-17T11:16:36Z
closed: 2021-09-17T08:41:51Z
merged: 2021-09-17T08:41:51Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/832
---

# Pull Request #832: SPM-1165: Added LRU cache for advisory detail endpoint

**Author**: @josef-hak
**Created**: September 16, 2021 at 02:07 PM UTC
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

### Comment by @codecov-commenter on September 16, 2021 at 02:31 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/832?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#832](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/832?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (481d361) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f751bd683c3255afae8e9df77ca312463cc8e670?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f751bd6) will **increase** coverage by `0.12%`.
> The diff coverage is `66.66%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/832/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/832?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #832      +/-   ##
==========================================
+ Coverage   57.95%   58.08%   +0.12%     
==========================================
  Files          81       81              
  Lines        3908     3941      +33     
==========================================
+ Hits         2265     2289      +24     
- Misses       1325     1328       +3     
- Partials      318      324       +6     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.08% <66.66%> (+0.12%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/832?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/832/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `70.52% <66.66%> (+1.17%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/832?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/832?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [f751bd6...481d361](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/832?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Changes Requested on September 17, 2021 at 07:29 AM UTC

### Review by @josef-hak - Commented on September 17, 2021 at 07:52 AM UTC

### Review by @josef-hak - Commented on September 17, 2021 at 07:52 AM UTC

### Review by @MichaelMraka - Approved on September 17, 2021 at 07:55 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/832*
