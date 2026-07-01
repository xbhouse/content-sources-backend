---
type: pull_request
number: 1220
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-11-11T04:10:05Z
updated: 2024-11-13T12:50:38Z
closed: 2024-11-13T12:50:38Z
merged: 2024-11-13T12:50:38Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1220
---

# Pull Request #1220: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: November 11, 2024 at 04:10 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `715fa1f` -> `5da8c2f` |
| quay.io/konflux-ci/tekton-catalog/task-buildah | `312d829` -> `5a92177` |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `5948fe1` -> `b824db3` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `747b43a` -> `59a538a` |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `443ffa8` -> `4eb168d` |
| quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks | `5131cce` -> `9918d8d` |
| quay.io/konflux-ci/tekton-catalog/task-init | `f239f38` -> `5efc5c7` |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies | `f53fe54` -> `39ec983` |
| quay.io/konflux-ci/tekton-catalog/task-push-dockerfile | `674e70f` -> `a216178` |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `e6a5aa0` -> `0c9667f` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `a1205ae` -> `9dfcd4e` |
| quay.io/konflux-ci/tekton-catalog/task-source-build | `21cb5eb` -> `53a41b0` |
| quay.io/konflux-ci/tekton-catalog/task-summary | `d97c04a` -> `2ed7b9e` |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @app-sre-bot on November 11, 2024 at 04:10 AM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on November 11, 2024 at 04:28 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1220?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.98%. Comparing base [(`55cd68f`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/55cd68f6dd72799991da69bf780b3a4fc15769cb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`6fea993`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/6fea9936a7c84077e7760a488bd081ad3f2a64e2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1220   +/-   ##
=======================================
  Coverage   63.98%   63.98%           
=======================================
  Files         124      124           
  Lines        3210     3210           
  Branches      821      821           
=======================================
  Hits         2054     2054           
  Misses       1156     1156           
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1220?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @LightOfHeaven1994 on November 12, 2024 at 10:35 PM UTC

/retest

### Comment by @LightOfHeaven1994 on November 12, 2024 at 10:42 PM UTC

/retest

---

## Reviews

### Review by @LightOfHeaven1994 - Approved on November 13, 2024 at 12:50 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1220*
