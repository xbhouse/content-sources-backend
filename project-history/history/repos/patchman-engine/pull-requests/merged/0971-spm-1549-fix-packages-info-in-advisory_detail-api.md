---
type: pull_request
number: 971
title: "SPM-1549: fix packages info in advisory_detail api"
state: merged
author: psegedy
created: 2022-06-07T18:08:18Z
updated: 2022-06-13T12:36:48Z
closed: 2022-06-13T12:36:48Z
merged: 2022-06-13T12:36:48Z
base_branch: master
head_branch: spm-1549
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/971
---

# Pull Request #971: SPM-1549: fix packages info in advisory_detail api

**Author**: @psegedy
**Created**: June 07, 2022 at 06:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `spm-1549`

## Description

- add advisory detail v2 api returning packages in `map[string][]string` format, instead of `map[string]string`
- expose v2 api for all other endpoints
- swaggo/swag does not make it easy to make multiple openapi.json files, ~it creates 1 big openapi.json with all paths and I split it into multiple files in `docs/docs.go:filterOpenAPI` and exposed them separately as `api/patch/v1/openapi.json` and `/api/patch/v2/openapi.json`. I'm open to suggestions.~ 
- move current `docs/openapi.json` to `docs/v1/openapi.json`
- generate only new spec in `docs/v2/openapi.json` and expose both specs on `api/patch/v1/openapi.json` and `/api/patch/v2/openapi.json`

TODO:
- [x] test in ephemeral
- [x] make sure unit tests are passing

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

### Comment by @codecov-commenter on June 07, 2022 at 06:27 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#971](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (3ca6734) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b0df3bc65747157d0dc3ba53dad581a824aa171a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (b0df3bc) will **increase** coverage by `0.20%`.
> The diff coverage is `76.22%`.

```diff
@@            Coverage Diff             @@
##           master     #971      +/-   ##
==========================================
+ Coverage   60.63%   60.84%   +0.20%     
==========================================
  Files          94       94              
  Lines        5134     5197      +63     
==========================================
+ Hits         3113     3162      +49     
- Misses       1614     1624      +10     
- Partials      407      411       +4     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.84% <76.22%> (+0.20%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `93.84% <ø> (ø)` | |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `75.00% <ø> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `72.50% <ø> (ø)` | |
| [manager/controllers/advisory\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `50.00% <ø> (ø)` | |
| [manager/controllers/baseline\_create.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | `75.40% <ø> (ø)` | |
| [manager/controllers/baseline\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZWxldGUuZ28=) | `71.87% <ø> (ø)` | |
| [manager/controllers/baseline\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZXRhaWwuZ28=) | `88.09% <ø> (ø)` | |
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `73.91% <ø> (ø)` | |
| [manager/controllers/baseline\_systems\_remove.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zX3JlbW92ZS5nbw==) | `51.72% <ø> (ø)` | |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `81.31% <ø> (ø)` | |
| ... and [21 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [14cbc0e...3ca6734](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/971?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on June 08, 2022 at 10:56 AM UTC

/retest

### Comment by @psegedy on June 08, 2022 at 03:19 PM UTC

looks like I need to wait till https://github.com/RedHatInsights/patchman-engine/pull/966 is merged as there might be some confilcts

### Comment by @digitronik on June 09, 2022 at 07:57 AM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on June 13, 2022 at 11:44 AM UTC

### Review by @MichaelMraka - Approved on June 13, 2022 at 11:47 AM UTC

### Review by @psegedy - Commented on June 13, 2022 at 12:02 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/971*
