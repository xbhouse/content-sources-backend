---
type: pull_request
number: 752
title: "SPM-1055: Improve package-systems endpoint sql query"
state: merged
author: josef-hak
created: 2021-07-30T13:54:42Z
updated: 2021-08-26T18:41:58Z
closed: 2021-08-02T12:14:41Z
merged: 2021-08-02T12:14:41Z
base_branch: master
head_branch: name-id
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/752
---

# Pull Request #752: SPM-1055: Improve package-systems endpoint sql query

**Author**: @josef-hak
**Created**: July 30, 2021 at 01:54 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `name-id`

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

### Comment by @codecov-commenter on July 30, 2021 at 02:02 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#752](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8913729) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ca7796ee14754408a938f55c2a3e4a3d3933545c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ca7796e) will **increase** coverage by `0.18%`.
> The diff coverage is `86.20%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #752      +/-   ##
==========================================
+ Coverage   57.03%   57.22%   +0.18%     
==========================================
  Files          81       81              
  Lines        3833     3857      +24     
==========================================
+ Hits         2186     2207      +21     
- Misses       1329     1331       +2     
- Partials      318      319       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.22% <86.20%> (+0.18%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb3JlLmdv) | `11.29% <0.00%> (-0.19%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `77.31% <66.66%> (-0.52%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `73.52% <100.00%> (+0.80%)` | :arrow_up: |
| [manager/controllers/package\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXNfZXhwb3J0Lmdv) | `44.82% <100.00%> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.50% <100.00%> (+2.22%)` | :arrow_up: |
| [manager/controllers/system\_advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllc19leHBvcnQuZ28=) | `55.55% <100.00%> (+4.04%)` | :arrow_up: |
| [manager/controllers/system\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGVsZXRlLmdv) | `55.88% <100.00%> (+4.26%)` | :arrow_up: |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `88.88% <100.00%> (+1.38%)` | :arrow_up: |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `69.44% <100.00%> (+2.77%)` | :arrow_up: |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [814d508...8913729](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/752?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on July 30, 2021 at 02:18 PM UTC

/retest

---

## Reviews

### Review by @semtexzv - Approved on July 30, 2021 at 02:27 PM UTC

### Review by @MichaelMraka - Approved on August 02, 2021 at 12:00 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/752*
