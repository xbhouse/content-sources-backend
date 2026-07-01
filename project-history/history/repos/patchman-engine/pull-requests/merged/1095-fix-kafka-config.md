---
type: pull_request
number: 1095
title: "fix kafka config"
state: merged
author: psegedy
created: 2022-09-05T09:16:35Z
updated: 2022-09-06T13:48:26Z
closed: 2022-09-06T13:48:25Z
merged: 2022-09-06T13:48:25Z
base_branch: master
head_branch: kafkakafkakafka
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1095
---

# Pull Request #1095: fix kafka config

**Author**: @psegedy
**Created**: September 05, 2022 at 09:16 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `kafkakafkakafka`

## Description

there was an error:
```
{"@timestamp":"2022-09-01T19:17:31.548Z","level":"error","message":"Unable to establish connection to consumer group coordinator for group patchman: [33] Unsupported SASL Mechanism: the broker does not support the requested SASL mechanism","type":"kafka"}
693
{"@timestamp":"2022-09-01T19:17:31.548Z","level":"error","message":"[33] Unsupported SASL Mechanism: the broker does not support the requested SASL mechanism","type":"kafka"}
```

I don't know how to test it, ephemeral is passing and stage is currently using local kafka instead of managed kafka

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

### Comment by @jira-linking on September 05, 2022 at 09:16 AM UTC

Commits missing Jira IDs:
766d586d07ee93b63d17622bdc0595d48bb22b88


### Comment by @codecov-commenter on September 05, 2022 at 10:26 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **62.65**% // Head: **62.86**% // Increases project coverage by **`+0.20%`** :tada:
> Coverage data is based on head [(`766d586`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`19754c9`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/19754c9d8805f1bdc6cf029dd7919251f16f1632?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 60.33% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1095      +/-   ##
==========================================
+ Coverage   62.65%   62.86%   +0.20%     
==========================================
  Files          97       97              
  Lines        5463     5534      +71     
==========================================
+ Hits         3423     3479      +56     
- Misses       1603     1606       +3     
- Partials      437      449      +12     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.86% <60.33%> (+0.20%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/database.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9kYXRhYmFzZS5nbw==) | `41.17% <0.00%> (ø)` | |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `63.85% <0.00%> (-1.58%)` | :arrow_down: |
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `73.33% <ø> (ø)` | |
| [manager/controllers/baselines.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZXMuZ28=) | `90.69% <ø> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `84.37% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `69.44% <ø> (ø)` | |
| [manager/middlewares/logger.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9sb2dnZXIuZ28=) | `0.00% <0.00%> (ø)` | |
| [manager/middlewares/prometheus.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9wcm9tZXRoZXVzLmdv) | `0.00% <0.00%> (ø)` | |
| [tasks/system\_culling/culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvc3lzdGVtX2N1bGxpbmcvY3VsbGluZy5nbw==) | `0.00% <0.00%> (ø)` | |
| [tasks/system\_culling/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvc3lzdGVtX2N1bGxpbmcvbWV0cmljcy5nbw==) | `0.00% <0.00%> (ø)` | |
| ... and [14 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1095?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on September 06, 2022 at 09:26 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1095*
