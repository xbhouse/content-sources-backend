---
type: pull_request
number: 1390
title: "RHINENG-8716: fix dynamically set limit in c.Request.URL"
state: merged
author: psegedy
created: 2024-03-05T17:33:49Z
updated: 2024-03-06T14:21:47Z
closed: 2024-03-06T14:21:24Z
merged: 2024-03-06T14:21:24Z
base_branch: master
head_branch: views_limit
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1390
---

# Pull Request #1390: RHINENG-8716: fix dynamically set limit in c.Request.URL

**Author**: @psegedy
**Created**: March 05, 2024 at 05:33 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `views_limit`

## Description

`c.Query(key)` claims to be a shortcut for `c.Request.URL.Query().Get(key)` by `gin` documentation. The reality is that `c.Query()` tries to get the value from QueryCache if cache exists, The cache is created with first c.Query() call on given context (might be created with other calls as well).

The problem is with /views apis. The query is added dynamically to `c.Request.URL.RawQuery` there. My guess is that it worked previously because we modified the context `c` with `apiver`... it looks like it stopped working after removal of `apiver` decisions and removal of v2 structs (it works in prod but not in stage) but I haven't tried to trace the exact commit.

IMO it should be safer to use `c.Request.URL.Query().Get(key)`
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

### Comment by @jira-linking on March 05, 2024 at 05:33 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-8716


### Comment by @codecov-commenter on March 06, 2024 at 02:15 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1390?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `0%` with `1 lines` in your changes are missing coverage. Please review.
> Project coverage is 63.68%. Comparing base [(`242afaf`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/242afaf85ff5608629a75e41ffb93249347595c0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`47c7f51`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1390?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1390?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/utils/gin.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1390?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1390?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1390      +/-   ##
==========================================
+ Coverage   63.65%   63.68%   +0.03%     
==========================================
  Files         107      107              
  Lines        6504     6504              
==========================================
+ Hits         4140     4142       +2     
+ Misses       1875     1873       -2     
  Partials      489      489              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1390/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1390/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.68% <0.00%> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1390?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on March 06, 2024 at 02:21 PM UTC

quay is acting up, this is the error from the ci
```
15:14:16 ++ podman push quay.io/cloudservices/patchman-engine-app:pr-1390-47c7f51
15:14:17 Getting image source signatures
15:14:17 Error: trying to reuse blob sha256:0427f9bc962101ecc89f99131f22d62365ee3ddc45b57ff7b27754d500ce4ecc at destination: can't talk to a V1 container registry
```
merging it since only a comment is added with the last force push

---

## Reviews

### Review by @MichaelMraka - Approved on March 06, 2024 at 01:13 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1390*
