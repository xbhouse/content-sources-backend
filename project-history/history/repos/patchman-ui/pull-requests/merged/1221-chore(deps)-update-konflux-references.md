---
type: pull_request
number: 1221
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-11-18T04:16:53Z
updated: 2024-11-18T09:00:28Z
closed: 2024-11-18T09:00:28Z
merged: 2024-11-18T09:00:28Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1221
---

# Pull Request #1221: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: November 18, 2024 at 04:16 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-buildah | `5a92177` -> `bbacdfe` |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `b824db3` -> `eff773a` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `59a538a` -> `1465898` |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `4eb168d` -> `5a1a165` |
| quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks | `9918d8d` -> `2e37ec3` |
| quay.io/konflux-ci/tekton-catalog/task-git-clone | `2cccdf8` -> `d091a9e` |
| quay.io/konflux-ci/tekton-catalog/task-init | `5efc5c7` -> `0523b51` |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies | `39ec983` -> `f59a214` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `9dfcd4e` -> `f165b1c` |
| quay.io/konflux-ci/tekton-catalog/task-summary | `2ed7b9e` -> `b0f049f` |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @app-sre-bot on November 18, 2024 at 04:19 AM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on November 18, 2024 at 04:32 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1221?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.98%. Comparing base [(`ec0323a`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ec0323a1bfaa9d1df373a815fe2c66c4401e468a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`2bd9f1e`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/2bd9f1e34f05f8355b60c1625b956b708a8e414d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1221   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1221?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @LightOfHeaven1994 - Approved on November 18, 2024 at 09:00 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1221*
