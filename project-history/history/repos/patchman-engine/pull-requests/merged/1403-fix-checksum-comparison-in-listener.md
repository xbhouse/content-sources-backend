---
type: pull_request
number: 1403
title: "fix checksum comparison in listener"
state: merged
author: psegedy
created: 2024-03-25T18:04:22Z
updated: 2024-03-26T16:28:53Z
closed: 2024-03-26T16:28:53Z
merged: 2024-03-26T16:28:53Z
base_branch: master
head_branch: checksums
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1403
---

# Pull Request #1403: fix checksum comparison in listener

**Author**: @psegedy
**Created**: March 25, 2024 at 06:04 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `checksums`

## Description

There were a few issues with the current query
- First can be used only on Model
- can't load to map[string]string, gorm showing weird error about sql.Scan called but Next not used
- need to cast jsonb to text for `digest` and encode the hex value

Second issue is that digest can't be used for yum_updates as it contains a hash map and order of fields is not deterministic

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

### Comment by @jira-linking on March 25, 2024 at 06:04 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-9028


### Comment by @codecov-commenter on March 26, 2024 at 01:52 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1403?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `91.66667%` with `1 lines` in your changes are missing coverage. Please review.
> Project coverage is 63.57%. Comparing base [(`ccba24a`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ccba24a45b67a36f4ef6b0d4323682911cb5af4b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`036f012`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1403?dropdown=coverage&src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1403?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1403?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | 91.66% | [0 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1403?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1403      +/-   ##
==========================================
- Coverage   63.59%   63.57%   -0.02%     
==========================================
  Files         112      112              
  Lines        6774     6779       +5     
==========================================
+ Hits         4308     4310       +2     
- Misses       1958     1961       +3     
  Partials      508      508              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1403/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1403/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.57% <91.66%> (-0.02%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1403?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on March 26, 2024 at 02:14 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1403*
