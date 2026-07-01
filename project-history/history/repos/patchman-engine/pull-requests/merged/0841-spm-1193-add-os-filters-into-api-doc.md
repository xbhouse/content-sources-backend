---
type: pull_request
number: 841
title: "SPM-1193: add os filters into api doc"
state: merged
author: mkholjuraev
created: 2021-10-07T12:13:52Z
updated: 2021-10-08T08:59:55Z
closed: 2021-10-08T08:59:55Z
merged: 2021-10-08T08:59:55Z
base_branch: master
head_branch: add-os-filter
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/841
---

# Pull Request #841: SPM-1193: add os filters into api doc

**Author**: @mkholjuraev
**Created**: October 07, 2021 at 12:13 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `add-os-filter`

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


Os name, major version, minor version are already sortable, it would be better if it is reflected in the API doc, I think

---

## Discussion

### Comment by @codecov-commenter on October 07, 2021 at 12:22 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#841](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (64d7805) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/55f0a6912d7ad29f574f02e32b5b4c864aed6c71?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (55f0a69) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #841   +/-   ##
=======================================
  Coverage   58.58%   58.58%           
=======================================
  Files          81       81           
  Lines        4035     4035           
=======================================
  Hits         2364     2364           
  Misses       1341     1341           
  Partials      330      330           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.58% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.55% <ø> (ø)` | |
| [manager/controllers/advisory\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `50.00% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `91.89% <ø> (ø)` | |
| [manager/controllers/systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2V4cG9ydC5nbw==) | `79.16% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [129f44f...64d7805](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/841?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Commented on October 07, 2021 at 12:55 PM UTC

Please add a test for os version filtering. It should be similar to TestSystemsWorkloads1()  (in systems_test.go) just modify request to e.g. "/?filter[os_name]=RHEL&filter[os_major]=8&filter[os_minor]=1".

### Review by @josef-hak - Changes Requested on October 08, 2021 at 07:39 AM UTC

### Review by @josef-hak - Approved on October 08, 2021 at 08:59 AM UTC

Now it's perfect, thanks :+1: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/841*
