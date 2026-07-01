---
type: pull_request
number: 1029
title: "reword duplicate baseline name error message"
state: merged
author: mkholjuraev
created: 2022-07-20T06:58:06Z
updated: 2022-07-21T08:12:56Z
closed: 2022-07-21T08:12:55Z
merged: 2022-07-21T08:12:55Z
base_branch: master
head_branch: rename-baseline
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1029
---

# Pull Request #1029: reword duplicate baseline name error message

**Author**: @mkholjuraev
**Created**: July 20, 2022 at 06:58 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `rename-baseline`

## Description

This rewords the duplicate baseline name error message from 'baseline name already exists' to 'patch template name already exists'. 


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

### Comment by @jira-linking on July 20, 2022 at 06:58 AM UTC

Commits missing Jira IDs:
9e6e1d35e4401fdd4ac2d0de630afc71a1faa837


### Comment by @codecov-commenter on July 20, 2022 at 07:06 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1029?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1029](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1029?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d605595) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/fd466045239b9c4ed8afae2e6294685d32eee94a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (fd46604) will **not change** coverage.
> The diff coverage is `100.00%`.

```diff
@@           Coverage Diff           @@
##           master    #1029   +/-   ##
=======================================
  Coverage   61.25%   61.25%           
=======================================
  Files          95       95           
  Lines        5413     5413           
=======================================
  Hits         3316     3316           
  Misses       1673     1673           
  Partials      424      424           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.25% <100.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1029?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/baseline\_create.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1029/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | `75.80% <100.00%> (ø)` | |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1029/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `81.31% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1029?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1029?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [b6ea455...d605595](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1029?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Changes Requested on July 20, 2022 at 11:58 AM UTC

please create a `const` in any file of baseline package, similarly as error for missing name in `manager/controllers/baseline_create.go`:
```
const BaselineMissingNameErr = "missing required parameter 'name'"
```


### Review by @psegedy - Approved on July 21, 2022 at 07:58 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1029*
