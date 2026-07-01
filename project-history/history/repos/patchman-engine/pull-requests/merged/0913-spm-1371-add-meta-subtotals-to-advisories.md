---
type: pull_request
number: 913
title: "SPM-1371: Add meta subtotals to advisories"
state: merged
author: michalslomczynski
created: 2022-02-28T13:24:04Z
updated: 2022-03-03T12:13:25Z
closed: 2022-03-02T16:01:12Z
merged: 2022-03-02T16:01:12Z
base_branch: master
head_branch: advisories-subtotals
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/913
---

# Pull Request #913: SPM-1371: Add meta subtotals to advisories

**Author**: @michalslomczynski
**Created**: February 28, 2022 at 01:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `advisories-subtotals`

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

### Comment by @MichaelMraka on March 01, 2022 at 10:05 AM UTC

Please remove the `fix(..)` from the first commit message we do not use this kind of annotation in patch.
Actually it will be better to squash the two commits together before merging.

### Comment by @MichaelMraka on March 01, 2022 at 10:06 AM UTC

Please add some test for the sub totals.

### Comment by @codecov-commenter on March 02, 2022 at 10:29 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#913](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ae9cd64) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/40f5c7fc802b0e783b9996307aaa8469a5b35bed?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (40f5c7f) will **increase** coverage by `0.11%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #913      +/-   ##
==========================================
+ Coverage   59.24%   59.35%   +0.11%     
==========================================
  Files          90       90              
  Lines        4696     4709      +13     
==========================================
+ Hits         2782     2795      +13     
  Misses       1528     1528              
  Partials      386      386              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `59.35% <100.00%> (+0.11%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/test\_utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZXN0X3V0aWxzLmdv) | `100.00% <ø> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `93.84% <100.00%> (+1.11%)` | :arrow_up: |
| [manager/controllers/baseline\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZXRhaWwuZ28=) | `87.17% <0.00%> (ø)` | |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `74.68% <0.00%> (+0.32%)` | :arrow_up: |
| [manager/controllers/baseline\_create.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | `69.49% <0.00%> (+0.52%)` | :arrow_up: |
| [manager/controllers/baseline\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZWxldGUuZ28=) | `71.87% <0.00%> (+0.90%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [c5a5a1b...ae9cd64](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/913?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @michalslomczynski on March 02, 2022 at 11:06 AM UTC

@MichaelMraka I added test to subtotals, what do you think?

---

## Reviews

### Review by @MichaelMraka - Changes Requested on February 28, 2022 at 01:44 PM UTC

### Review by @MichaelMraka - Commented on March 01, 2022 at 10:02 AM UTC

### Review by @MichaelMraka - Changes Requested on March 02, 2022 at 02:12 PM UTC

### Review by @michalslomczynski - Commented on March 02, 2022 at 02:16 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/913*
