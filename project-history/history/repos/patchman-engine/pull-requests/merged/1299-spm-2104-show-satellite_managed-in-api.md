---
type: pull_request
number: 1299
title: "SPM-2104: show satellite_managed in API"
state: merged
author: Mischulee
created: 2023-08-25T10:28:52Z
updated: 2023-08-25T16:14:35Z
closed: 2023-08-25T16:14:35Z
merged: 2023-08-25T16:14:35Z
base_branch: master
head_branch: up
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1299
---

# Pull Request #1299: SPM-2104: show satellite_managed in API

**Author**: @Mischulee
**Created**: August 25, 2023 at 10:28 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `up`

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

### Comment by @jira-linking on August 25, 2023 at 10:28 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-2104


### Comment by @codecov-commenter on August 25, 2023 at 10:39 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1299?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch and project coverage have no change.
> Comparison is base [(`99e4b5b`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/99e4b5bb1172326afa34f946059a250ae0fe368e?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.01% compared to head [(`cb4ab3d`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1299?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.01%.
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1299   +/-   ##
=======================================
  Coverage   62.01%   62.01%           
=======================================
  Files         106      106           
  Lines        6674     6674           
=======================================
  Hits         4139     4139           
  Misses       2003     2003           
  Partials      532      532           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1299/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1299/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1299?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1299?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `42.85% <ø> (ø)` | |
| [manager/controllers/advisory\_systems\_v3.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1299?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX3YzLmdv) | `72.34% <ø> (ø)` | |
| [manager/controllers/systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1299?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `71.42% <ø> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1299?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on August 25, 2023 at 04:14 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1299*
