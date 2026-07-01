---
type: pull_request
number: 1266
title: "RHINENG-1291: use kafka KRaft mode in docker-compose"
state: merged
author: psegedy
created: 2023-07-20T15:32:26Z
updated: 2023-07-21T09:41:26Z
closed: 2023-07-21T09:41:26Z
merged: 2023-07-21T09:41:26Z
base_branch: master
head_branch: kafka
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1266
---

# Pull Request #1266: RHINENG-1291: use kafka KRaft mode in docker-compose

**Author**: @psegedy
**Created**: July 20, 2023 at 03:32 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `kafka`

## Description

https://docs.confluent.io/platform/current/kafka-metadata/kraft.html

@MichaelMraka please try this on your system. I'm not sure how kafka SSL should work and if it still works with CONTROLLER, do you know how to try this? (maybe it works, maybe we need to add CONTROLLER:SSL instead of CONTROLLER:PLAINTEXT)

I removed all kafka metrics since they are only part of the cp-enterprise-kafka, do we use it somewhere? We can take a look whether enterprise container supports KRaft mode then
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

### Comment by @jira-linking on July 20, 2023 at 03:32 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-1291


### Comment by @codecov-commenter on July 20, 2023 at 03:41 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`23.52`**% and project coverage change: **`-0.17`** :warning:
> Comparison is base [(`6d5b7b6`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/6d5b7b6906028fb189ef52d642006189d504056f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.40% compared to head [(`5e9b3bf`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.23%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1266      +/-   ##
==========================================
- Coverage   61.40%   61.23%   -0.17%     
==========================================
  Files         106      106              
  Lines        6651     6653       +2     
==========================================
- Hits         4084     4074      -10     
- Misses       2034     2040       +6     
- Partials      533      539       +6     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.23% <23.52%> (-0.17%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `46.34% <0.00%> (ø)` | |
| [manager/controllers/advisory\_systems\_v3.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX3YzLmdv) | `72.34% <0.00%> (-3.75%)` | :arrow_down: |
| [manager/controllers/baseline\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `49.03% <0.00%> (-1.93%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `60.43% <0.00%> (-2.20%)` | :arrow_down: |
| [manager/controllers/packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `63.23% <0.00%> (ø)` | |
| [manager/controllers/system\_detail.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `60.41% <0.00%> (-3.42%)` | :arrow_down: |
| [manager/controllers/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.92% <16.66%> (-1.24%)` | :arrow_down: |
| [manager/controllers/systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `71.42% <71.42%> (+2.85%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `74.54% <100.00%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1266?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on July 21, 2023 at 09:40 AM UTC

Works on podman /RHEL9. Let's skip SSL for now.

---

## Reviews

### Review by @MichaelMraka - Approved on July 21, 2023 at 09:40 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1266*
