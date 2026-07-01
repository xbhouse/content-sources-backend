---
type: pull_request
number: 661
title: "SPM-866: speed up package loading"
state: merged
author: MichaelMraka
created: 2021-05-05T11:49:24Z
updated: 2021-07-12T11:55:02Z
closed: 2021-07-12T11:55:02Z
merged: 2021-07-12T11:55:02Z
base_branch: master
head_branch: pr3
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/661
---

# Pull Request #661: SPM-866: speed up package loading

**Author**: @MichaelMraka
**Created**: May 05, 2021 at 11:49 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr3`

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

### Comment by @semtexzv on May 12, 2021 at 02:28 PM UTC

@Josca I don't think so. The array could be resized, therefore pointers invalidated. Take a look at the method for adding new entries. It accepts a pointer, and just copies the value out of this pointer. 

### Comment by @codecov-commenter on July 08, 2021 at 12:04 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#661](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8b9757a) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/2e072717ea9abe20119e20d63c4750d2337c8e6b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2e07271) will **increase** coverage by `0.80%`.
> The diff coverage is `87.13%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #661      +/-   ##
==========================================
+ Coverage   56.12%   56.93%   +0.80%     
==========================================
  Files          80       80              
  Lines        3695     3773      +78     
==========================================
+ Hits         2074     2148      +74     
- Misses       1308     1312       +4     
  Partials      313      313              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.93% <87.13%> (+0.80%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `12.06% <0.00%> (-0.90%)` | :arrow_down: |
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `74.28% <85.88%> (+5.36%)` | :arrow_up: |
| [evaluator/package\_cache.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3BhY2thZ2VfY2FjaGUuZ28=) | `93.15% <93.15%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.50% <100.00%> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `62.04% <100.00%> (+0.27%)` | :arrow_up: |
| [vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `73.60% <100.00%> (+0.21%)` | :arrow_up: |
| [vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `75.00% <100.00%> (+0.45%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2e07271...8b9757a](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/661?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on July 09, 2021 at 10:55 AM UTC

/retest

---

## Reviews

### Review by @semtexzv - Commented on May 12, 2021 at 12:01 PM UTC

### Review by @josef-hak - Approved on May 12, 2021 at 01:25 PM UTC

Well done. Thanks @MichaelMraka , great update. I added some suggestions and comments but basically lgtm :wink: . Of course it has to be well tested as there are many changed things.

### Review by @MichaelMraka - Commented on May 13, 2021 at 07:34 AM UTC

### Review by @semtexzv - Commented on May 13, 2021 at 07:40 AM UTC

### Review by @MichaelMraka - Commented on May 13, 2021 at 07:41 AM UTC

### Review by @MichaelMraka - Commented on May 14, 2021 at 07:31 AM UTC

### Review by @MichaelMraka - Commented on May 14, 2021 at 07:32 AM UTC

### Review by @MichaelMraka - Commented on May 14, 2021 at 07:32 AM UTC

### Review by @semtexzv - Commented on May 18, 2021 at 08:57 AM UTC

### Review by @josef-hak - Changes Requested on June 18, 2021 at 09:18 AM UTC

We have some conflicts here.

### Review by @josef-hak - Approved on July 12, 2021 at 11:45 AM UTC

lgtm

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/661*
