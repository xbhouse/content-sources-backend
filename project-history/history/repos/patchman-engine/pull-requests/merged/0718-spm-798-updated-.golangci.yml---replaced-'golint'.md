---
type: pull_request
number: 718
title: "SPM-798: updated .golangci.yml - replaced 'golint' with 'revive' "
state: merged
author: josef-hak
created: 2021-06-24T14:48:26Z
updated: 2021-08-26T18:41:52Z
closed: 2021-06-24T15:27:09Z
merged: 2021-06-24T15:27:09Z
base_branch: master
head_branch: lint
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/718
---

# Pull Request #718: SPM-798: updated .golangci.yml - replaced 'golint' with 'revive' 

**Author**: @josef-hak
**Created**: June 24, 2021 at 02:48 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `lint`

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

### Comment by @codecov-commenter on June 24, 2021 at 02:55 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#718](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d71446d) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/68c515c8b0e19859bced767c3621fc55830896cb?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (68c515c) will **not change** coverage.
> The diff coverage is `50.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #718   +/-   ##
=======================================
  Coverage   56.21%   56.21%           
=======================================
  Files          80       80           
  Lines        3693     3693           
=======================================
  Hits         2076     2076           
  Misses       1304     1304           
  Partials      313      313           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.21% <50.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/filter.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9maWx0ZXIuZ28=) | `68.18% <0.00%> (ø)` | |
| [vmaas\_sync/admin\_api.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZG1pbl9hcGkuZ28=) | `0.00% <ø> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `61.76% <ø> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `30.39% <ø> (ø)` | |
| [vmaas\_sync/metrics\_cyndi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9tZXRyaWNzX2N5bmRpLmdv) | `42.85% <ø> (ø)` | |
| [vmaas\_sync/metrics\_db.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9tZXRyaWNzX2RiLmdv) | `59.25% <ø> (ø)` | |
| [vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `73.38% <ø> (ø)` | |
| [vmaas\_sync/refresh.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZWZyZXNoLmdv) | `0.00% <ø> (ø)` | |
| [vmaas\_sync/refresh\_advisory\_caches.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZWZyZXNoX2Fkdmlzb3J5X2NhY2hlcy5nbw==) | `38.88% <ø> (ø)` | |
| [vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `74.54% <ø> (ø)` | |
| ... and [5 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [68c515c...d71446d](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/718?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on June 24, 2021 at 03:13 PM UTC

### Review by @semtexzv - Approved on June 24, 2021 at 03:26 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/718*
