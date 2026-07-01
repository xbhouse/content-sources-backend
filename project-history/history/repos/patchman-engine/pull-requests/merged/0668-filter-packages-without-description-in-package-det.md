---
type: pull_request
number: 668
title: "Filter packages without description in package detail endpoint"
state: merged
author: semtexzv
created: 2021-05-07T11:33:40Z
updated: 2021-05-11T08:11:36Z
closed: 2021-05-11T08:11:36Z
merged: 2021-05-11T08:11:36Z
base_branch: master
head_branch: pkg-soft-filter
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/668
---

# Pull Request #668: Filter packages without description in package detail endpoint

**Author**: @semtexzv
**Created**: May 07, 2021 at 11:33 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkg-soft-filter`

## Description

Current implementation grabs whatever package is available (if there is no advisory attached).
New implementation aims to filter out packages which don't have descriptions associated.

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

### Comment by @codecov-commenter on May 07, 2021 at 11:41 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/668?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#668](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/668?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1f49e35) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1b045996f1cf57fd873942cfefba40dd4509788f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1b04599) will **increase** coverage by `0.01%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/668/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/668?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #668      +/-   ##
==========================================
+ Coverage   58.53%   58.54%   +0.01%     
==========================================
  Files          72       72              
  Lines        3364     3365       +1     
==========================================
+ Hits         1969     1970       +1     
  Misses       1111     1111              
  Partials      284      284              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.54% <100.00%> (+0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/668?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/package\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/668/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX2RldGFpbC5nbw==) | `83.33% <100.00%> (+0.35%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/668?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/668?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [1b04599...1f49e35](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/668?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Approved on May 10, 2021 at 08:14 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/668*
