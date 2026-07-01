---
type: pull_request
number: 592
title: "Use no content response when account is not found"
state: merged
author: semtexzv
created: 2021-04-06T06:53:45Z
updated: 2021-04-08T09:45:25Z
closed: 2021-04-08T09:45:25Z
merged: 2021-04-08T09:45:25Z
base_branch: master
head_branch: no-content
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/592
---

# Pull Request #592: Use no content response when account is not found

**Author**: @semtexzv
**Created**: April 06, 2021 at 06:53 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `no-content`

## Description

Failure to find an account is not an authentication failure. The account could be valid. Fixes SPM-815

---

## Discussion

### Comment by @codecov-io on April 06, 2021 at 07:19 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592?src=pr&el=h1) Report
> Merging [#592](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592?src=pr&el=desc) (bd91684) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/33b2098d8e213d7897edd9526c1078fc71e27678?el=desc) (33b2098) will **decrease** coverage by `0.07%`.
> The diff coverage is `55.55%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #592      +/-   ##
==========================================
- Coverage   60.33%   60.26%   -0.08%     
==========================================
  Files          67       68       +1     
  Lines        2945     2982      +37     
==========================================
+ Hits         1777     1797      +20     
- Misses        913      926      +13     
- Partials      255      259       +4     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.26% <55.55%> (-0.08%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `61.53% <ø> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.55% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `89.28% <ø> (ø)` | |
| [manager/middlewares/authentication.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree#diff-bWFuYWdlci9taWRkbGV3YXJlcy9hdXRoZW50aWNhdGlvbi5nbw==) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/admin\_api.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZG1pbl9hcGkuZ28=) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/repo\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9yZXBvX3N5bmMuZ28=) | `33.33% <33.33%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `34.40% <50.00%> (+1.46%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.25% <63.15%> (+0.42%)` | :arrow_up: |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `68.42% <100.00%> (+3.71%)` | :arrow_up: |
| ... and [3 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592?src=pr&el=footer). Last update [56bb665...bd91684](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/592?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on April 06, 2021 at 10:41 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/592*
