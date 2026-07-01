---
type: pull_request
number: 640
title: "Create batch upload eval events not single event for every system"
state: merged
author: MichaelMraka
created: 2021-04-27T13:09:58Z
updated: 2021-05-04T08:58:21Z
closed: 2021-05-04T08:58:21Z
merged: 2021-05-04T08:58:21Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/640
---

# Pull Request #640: Create batch upload eval events not single event for every system

**Author**: @MichaelMraka
**Created**: April 27, 2021 at 01:09 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

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

### Comment by @codecov-commenter on May 04, 2021 at 08:01 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#640](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (49edc30) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/eda0d6a1c4d83db621aaf54129a9c1534d54e48f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (eda0d6a) will **decrease** coverage by `0.46%`.
> The diff coverage is `43.58%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #640      +/-   ##
==========================================
- Coverage   58.76%   58.30%   -0.47%     
==========================================
  Files          72       72              
  Lines        3303     3312       +9     
==========================================
- Hits         1941     1931      -10     
- Misses       1081     1099      +18     
- Partials      281      282       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.30% <43.58%> (-0.47%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `27.58% <0.00%> (-33.96%)` | :arrow_down: |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `34.61% <57.14%> (-24.91%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `77.72% <78.57%> (-0.60%)` | :arrow_down: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `70.37% <100.00%> (+1.13%)` | :arrow_up: |
| [vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `71.42% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [eda0d6a...49edc30](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/640?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Approved on May 04, 2021 at 08:58 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/640*
