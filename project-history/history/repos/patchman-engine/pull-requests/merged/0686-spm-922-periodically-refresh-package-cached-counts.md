---
type: pull_request
number: 686
title: "SPM-922 Periodically refresh package cached counts"
state: merged
author: semtexzv
created: 2021-05-25T11:56:30Z
updated: 2021-06-07T11:35:07Z
closed: 2021-06-07T11:35:06Z
merged: 2021-06-07T11:35:06Z
base_branch: master
head_branch: refresh
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/686
---

# Pull Request #686: SPM-922 Periodically refresh package cached counts

**Author**: @semtexzv
**Created**: May 25, 2021 at 11:56 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `refresh`

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

### Comment by @josef-hak on May 27, 2021 at 03:46 AM UTC

@semtexzv could you add SPM-922 to commit message too?

### Comment by @codecov-commenter on May 27, 2021 at 09:18 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/686?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#686](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/686?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c1f8a4e) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8a02d2cfeb75ed15d4823db3a7e0e3e22687cbd0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8a02d2c) will **decrease** coverage by `0.18%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/686/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/686?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #686      +/-   ##
==========================================
- Coverage   58.13%   57.95%   -0.19%     
==========================================
  Files          73       74       +1     
  Lines        3404     3415      +11     
==========================================
  Hits         1979     1979              
- Misses       1141     1152      +11     
  Partials      284      284              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.95% <0.00%> (-0.19%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/686?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/refresh.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/686/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZWZyZXNoLmdv) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/686/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `40.51% <0.00%> (-0.36%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/686?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/686?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [8a02d2c...c1f8a4e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/686?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Changes Requested on May 26, 2021 at 07:31 AM UTC

- Could you please mark commit and PR title with SPM-XYZ? Jira will detect and assign this PR automatically to the ticket.
- Tests fails on lint:
~~~
vmaas_sync/refresh.go:27: unnecessary trailing newline (whitespace)
~~~

### Review by @josef-hak - Changes Requested on May 27, 2021 at 03:44 AM UTC

vmaas_sync/refresh.go:1:1: don't use an underscore in package name (golint)
package vmaas_sync
^
Note: to quickly check for lint issues run `golangci-lint run` command in your local patchman-engine folder. You will avoid many reviews iterations.

### Review by @josef-hak - Changes Requested on May 27, 2021 at 09:24 AM UTC

### Review by @josef-hak - Changes Requested on May 31, 2021 at 10:41 AM UTC

### Review by @josef-hak - Changes Requested on June 02, 2021 at 02:12 PM UTC

Almost done. Please, just one thing, thanks.

### Review by @semtexzv - Commented on June 07, 2021 at 10:32 AM UTC

### Review by @josef-hak - Approved on June 07, 2021 at 11:28 AM UTC

lgtm, thanks

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/686*
