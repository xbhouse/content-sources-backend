---
type: pull_request
number: 1472
title: "RHINENG-12299: fix index out of range error"
state: merged
author: psegedy
created: 2024-09-02T09:08:52Z
updated: 2024-09-02T13:17:33Z
closed: 2024-09-02T13:17:33Z
merged: 2024-09-02T13:17:33Z
base_branch: master
head_branch: outofrange
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1472
---

# Pull Request #1472: RHINENG-12299: fix index out of range error

**Author**: @psegedy
**Created**: September 02, 2024 at 09:08 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `outofrange`

## Description

```
panic: runtime error: index out of range [0] with length 0
goroutine 261071 [running]:
app/base/database.GetBaselineConfig(0xc03a4e1040)
/go/src/app/base/database/baseline.go:31 +0x445
app/evaluator.limitVmaasToBaseline(0x0?, 0xc03a565d40)
/go/src/app/evaluator/evaluate_baseline.go:12 +0x2a
app/evaluator.evaluateWithVmaas(0x19f48b8?, 0xc03a4e1040, 0xc03a4e1040?)
/go/src/app/evaluator/evaluate.go:285 +0x4c
app/evaluator.evaluateInDatabase({0x19f48b8, 0xc00082e320}, 0xc0004d5d50, {0xc015ce2a50, 0x24})
/go/src/app/evaluator/evaluate.go:230 +0xa6
app/evaluator.Evaluate({0x19f48b8, 0xc00082e320}, 0xc023e47550?, {0xc015ce2a50, 0x24}, {0xc00004a011, 0x6})
/go/src/app/evaluator/evaluate.go:146 +0x21c
app/evaluator.runEvaluate.func1()
/go/src/app/evaluator/evaluate.go:186 +0x6c
created by app/evaluator.runEvaluate in goroutine 206
/go/src/app/evaluator/evaluate.go:185 +0x273 
```

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

### Comment by @jira-linking on September 02, 2024 at 09:08 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-12299


### Comment by @codecov-commenter on September 02, 2024 at 09:13 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1472?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 65.01%. Comparing base [(`2e46900`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/2e46900759a25ba09158546f4e182d6f0dcec159?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`4266f9c`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4266f9cf42451fec65c409ace32f89d08d71742c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1472   +/-   ##
=======================================
  Coverage   65.01%   65.01%           
=======================================
  Files         114      114           
  Lines        7823     7823           
=======================================
  Hits         5086     5086           
  Misses       2207     2207           
  Partials      530      530           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1472/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1472/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.01% <100.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1472?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on September 02, 2024 at 09:21 AM UTC

### Review by @Dugowitch - Approved on September 02, 2024 at 09:32 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1472*
