---
type: pull_request
number: 1205
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-10-14T04:13:51Z
updated: 2024-10-25T19:12:04Z
closed: 2024-10-14T08:10:52Z
merged: 2024-10-14T08:10:52Z
base_branch: master
head_branch: konflux/references/master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1205
---

# Pull Request #1205: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: October 14, 2024 at 04:13 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-buildah | `e107cfd` -> `2d6e09f` |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `37b9187` -> `28fee4b` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `5ac9b24` -> `1e29eeb` |
| quay.io/konflux-ci/tekton-catalog/task-push-dockerfile | `0d2b6d3` -> `585106f` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `69ae591` -> `c1ea706` |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @app-sre-bot on October 14, 2024 at 04:16 AM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on October 14, 2024 at 04:28 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1205?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 64.01%. Comparing base [(`e81aa1d`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/e81aa1d4c7d76e01ea27d371559cb86a0ca33d5e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`048a8ef`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/048a8ef31d58e9b80bd470169526e3896bb587d7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1205   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1205?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on October 25, 2024 at 07:11 PM UTC

:tada: This PR is included in version 1.68.1 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.68.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Approved on October 14, 2024 at 08:10 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1205*
