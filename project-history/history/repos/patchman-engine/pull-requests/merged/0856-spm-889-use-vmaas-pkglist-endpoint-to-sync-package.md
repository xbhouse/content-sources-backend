---
type: pull_request
number: 856
title: "SPM-889: Use vmaas /pkglist endpoint to sync packages"
state: merged
author: josef-hak
created: 2021-11-02T16:08:10Z
updated: 2021-11-08T15:34:15Z
closed: 2021-11-08T15:33:09Z
merged: 2021-11-08T15:33:09Z
base_branch: master
head_branch: pkglist-sync
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/856
---

# Pull Request #856: SPM-889: Use vmaas /pkglist endpoint to sync packages

**Author**: @josef-hak
**Created**: November 02, 2021 at 04:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkglist-sync`

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

### Comment by @codecov-commenter on November 02, 2021 at 04:14 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/856?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#856](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/856?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c729b09) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/24f540e176b7227f55b42d9b4997bc46ae1a130f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (24f540e) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/856/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/856?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #856   +/-   ##
=======================================
  Coverage   58.48%   58.48%           
=======================================
  Files          81       81           
  Lines        4059     4059           
=======================================
  Hits         2374     2374           
  Misses       1351     1351           
  Partials      334      334           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.48% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/856?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/856?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2a1c925...c729b09](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/856?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on November 08, 2021 at 07:31 AM UTC

/retest

### Comment by @josef-hak on November 08, 2021 at 09:33 AM UTC

Yes, but for now I wanted to keep current behaviour by default. Later we can deprecated /pkgtree sync.

### Comment by @josef-hak on November 08, 2021 at 03:07 PM UTC

/retest

### Comment by @josef-hak on November 08, 2021 at 03:15 PM UTC

/retest

### Comment by @josef-hak on November 08, 2021 at 03:33 PM UTC

CI pipeline is failing because of broken x-join. I've tested in ephemeral env and it works. Merging.

---

## Reviews

### Review by @MichaelMraka - Approved on November 08, 2021 at 09:27 AM UTC

lgtm

Just a note: would not be /pkglist sync better default value for future?

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/856*
