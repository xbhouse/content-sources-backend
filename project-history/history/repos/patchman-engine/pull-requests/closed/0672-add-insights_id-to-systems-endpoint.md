---
type: pull_request
number: 672
title: "add insights_id to /systems endpoint"
state: closed
author: mkholjuraev
created: 2021-05-13T22:49:55Z
updated: 2021-05-18T13:20:36Z
closed: 2021-05-18T13:20:36Z
base_branch: master
head_branch: sync-insights_id
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/672
---

# Pull Request #672: add insights_id to /systems endpoint

**Author**: @mkholjuraev
**Created**: May 13, 2021 at 10:49 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `sync-insights_id`

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

### Comment by @codecov-commenter on May 17, 2021 at 02:14 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#672](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d1d057e) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b06f75eebf0135164f45c2ea5d2d9490957c6021?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (b06f75e) will **increase** coverage by `0.02%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #672      +/-   ##
==========================================
+ Coverage   58.56%   58.59%   +0.02%     
==========================================
  Files          72       72              
  Lines        3367     3369       +2     
==========================================
+ Hits         1972     1974       +2     
  Misses       1111     1111              
  Partials      284      284              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.59% <100.00%> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `27.58% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `89.28% <ø> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.03% <100.00%> (+0.20%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [b06f75e...d1d057e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/672?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on May 18, 2021 at 09:00 AM UTC

I am curious what will happen with already existing systems in the system_platform? Most probably they will not have insights_id added after this PR, right? Because this PR only manages newly uploaded systems. 

---

## Reviews

### Review by @josef-hak - Changes Requested on May 14, 2021 at 11:40 AM UTC

Ok, thanks, some things to update suggested.

### Review by @MichaelMraka - Commented on May 17, 2021 at 10:43 AM UTC

### Review by @mkholjuraev - Commented on May 17, 2021 at 12:26 PM UTC

### Review by @MichaelMraka - Commented on May 17, 2021 at 01:03 PM UTC

### Review by @MichaelMraka - Commented on May 17, 2021 at 01:09 PM UTC

### Review by @MichaelMraka - Changes Requested on May 17, 2021 at 03:04 PM UTC

I've seen a lot of code where people interchanged one kind of id for the other and tests passed but real app (unsurprisingly) the failed :).

### Review by @MichaelMraka - Approved on May 18, 2021 at 08:47 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/672*
