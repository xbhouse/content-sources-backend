---
type: pull_request
number: 720
title: "SPM-915: Added \"package_versions\" endpoint."
state: merged
author: josef-hak
created: 2021-06-29T19:38:15Z
updated: 2021-08-26T18:42:12Z
closed: 2021-07-12T11:29:30Z
merged: 2021-07-12T11:29:30Z
base_branch: master
head_branch: pkg-ver
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/720
---

# Pull Request #720: SPM-915: Added "package_versions" endpoint.

**Author**: @josef-hak
**Created**: June 29, 2021 at 07:38 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkg-ver`

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

### Comment by @codecov-commenter on June 29, 2021 at 07:45 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/720?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#720](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/720?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1967223) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/02a7ce22917b46ef00e9aea94be750d7d8a8f35a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (02a7ce2) will **increase** coverage by `0.05%`.
> The diff coverage is `61.76%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/720/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/720?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #720      +/-   ##
==========================================
+ Coverage   56.21%   56.26%   +0.05%     
==========================================
  Files          80       81       +1     
  Lines        3693     3727      +34     
==========================================
+ Hits         2076     2097      +21     
- Misses       1304     1312       +8     
- Partials      313      318       +5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.26% <61.76%> (+0.05%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/720?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/720/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `72.72% <ø> (ø)` | |
| [manager/controllers/package\_versions.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/720/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3ZlcnNpb25zLmdv) | `61.76% <61.76%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/720?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/720?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [02a7ce2...1967223](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/720?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @semtexzv - Approved on July 01, 2021 at 04:17 AM UTC

One question, otherwise looks good

### Review by @josef-hak - Commented on July 12, 2021 at 09:51 AM UTC

### Review by @josef-hak - Commented on July 12, 2021 at 11:29 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/720*
