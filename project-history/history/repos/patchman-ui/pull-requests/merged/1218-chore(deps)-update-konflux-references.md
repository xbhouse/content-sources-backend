---
type: pull_request
number: 1218
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-11-04T04:17:48Z
updated: 2024-11-12T17:57:53Z
closed: 2024-11-04T10:17:43Z
merged: 2024-11-04T10:17:43Z
base_branch: master
head_branch: konflux/references/master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1218
---

# Pull Request #1218: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: November 04, 2024 at 04:17 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-apply-tags | `f485e25` -> `87fd7fc` |
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `327d745` -> `715fa1f` |
| quay.io/konflux-ci/tekton-catalog/task-buildah | `b105a3b` -> `312d829` |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `28fee4b` -> `5948fe1` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `a94b652` -> `747b43a` |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `b4f9599` -> `443ffa8` |
| quay.io/konflux-ci/tekton-catalog/task-init | `092c113` -> `f239f38` |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies | `fe7234e` -> `f53fe54` |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `7aa4d3c` -> `e6a5aa0` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `60ed62a` -> `a1205ae` |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @app-sre-bot on November 04, 2024 at 04:19 AM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on November 04, 2024 at 04:40 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1218?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 64.01%. Comparing base [(`edf5324`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/edf53246a707e45226835930a582f3ef08cbca9c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`182de81`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/182de8120f38f013ffe81e34a864e9a7904646a0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1218   +/-   ##
=======================================
  Coverage   64.01%   64.01%           
=======================================
  Files         124      124           
  Lines        3235     3235           
  Branches      831      831           
=======================================
  Hits         2071     2071           
  Misses       1164     1164           
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1218?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on November 12, 2024 at 05:57 PM UTC

:tada: This PR is included in version 1.69.1 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.69.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @hugohezel - Approved on November 04, 2024 at 09:45 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1218*
