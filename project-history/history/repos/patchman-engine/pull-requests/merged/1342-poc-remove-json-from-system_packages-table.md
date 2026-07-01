---
type: pull_request
number: 1342
title: "POC: remove json from system_packages table"
state: merged
author: MichaelMraka
created: 2023-11-20T14:06:25Z
updated: 2023-12-05T16:01:34Z
closed: 2023-12-05T16:01:34Z
merged: 2023-12-05T16:01:34Z
base_branch: master
head_branch: poc
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1342
---

# Pull Request #1342: POC: remove json from system_packages table

**Author**: @MichaelMraka
**Created**: November 20, 2023 at 02:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `poc`

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

### Comment by @jira-linking on November 20, 2023 at 02:06 PM UTC

Commits missing Jira IDs:
d1e969892bafe52816428236d801eb0c523272ef
de6563390c58f9e83a91762a55e32206c4b76e0b
7891b844ca55a0671a5b7b5deb83baf520010a1c
1d66f2e93d27d382627bff3493df8e5fbfddd860
a3a89df3529602cf180d997d1e9cf1d10f7d690a
df0c223c8669477e47d9a6d4e5bdb29457551093
f7dcb292fcb91c93c75de7c3fab887e8c11d8a4c
422f99b6be1301480a45a861e612cf04ce9c875c
d4fdba5d6f15201c82543290cacfec155c118f30
a6d797d7a1d3f88502ab78813b1df48c0805f21e
27d235252a462829d72b06a406352a880b4c0ea2
df1344139d93cc6b5d9d3cfbbcc6ef7f70b1b199
2e8602543a2e0b795fef35ef6398a63dde854325
57a4042c7294eae150c5422b610044cda1181e03
e44393b0be7626cc83e276832a0cb7d743ee5bc5
db8859006b2abb41eb2989e300f695d3e504cad8
efbfa9f345ad8e4f3579e2af321a56ee762ec847
e32b92a60d445546a72568a75c71c050e38f3323
ef10a9399abc471a64a5e37bbdbc7cfd29034e7a
09fbea19838699e44c9977280a9d1691abce79aa
30e3faa5aef47011a57a597992bcd451a505f26b
63cf2aedc958297ec9215154c2a4ff7e98bc0ed0
9bb1f692438cc44e16e0a5fe63f0b73be05bcdd8
6c97d8fd6483c30953ac5e6a736d95224b31f745
a086f6bc24c50843ad1bb51d37b6ebd5cd0e188a
da294f2b5481955d506157442bae9a99b7f58e38
5503a1e82b0d2e7f70ec288a58f10520e6327897
59db100b140cb64a49a6c214db0fe66fa5d8110b
3cf8af0d691d74bf88cfc7e7bed524d8fc2fe4bd
b5132b9fb6b21fa24820b81654cdc5c2252556fc


### Comment by @codecov-commenter on November 23, 2023 at 09:27 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `20 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`25f765e`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/25f765ec1438368c4e3d9d1d2d9094362619b9ac?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.82% compared to head [(`b5132b9`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.03%.
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | 85.13% | [6 Missing and 5 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/testing.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | 0.00% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/system\_packages\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXNfZXhwb3J0Lmdv) | 66.66% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1342      +/-   ##
==========================================
+ Coverage   61.82%   62.03%   +0.20%     
==========================================
  Files         108      108              
  Lines        6814     6811       -3     
==========================================
+ Hits         4213     4225      +12     
+ Misses       2062     2052      -10     
+ Partials      539      534       -5     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.03% <81.98%> (+0.20%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1342?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1342*
