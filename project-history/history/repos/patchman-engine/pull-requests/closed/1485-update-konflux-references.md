---
type: pull_request
number: 1485
title: "Update Konflux references"
state: closed
author: red-hat-konflux
created: 2024-09-19T16:18:48Z
updated: 2024-09-30T15:15:53Z
closed: 2024-09-30T15:15:53Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1485
---

# Pull Request #1485: Update Konflux references

**Author**: @red-hat-konflux
**Created**: September 19, 2024 at 04:18 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-apply-tags | `e6beb16` -> `f485e25` |
| quay.io/konflux-ci/tekton-catalog/task-buildah | `71d3bb8` -> `e107cfd` |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `772233e` -> `9f4ddaf` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `412e96e` -> `5ac9b24` |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `592bbce` -> `b4f9599` |
| quay.io/konflux-ci/tekton-catalog/task-git-clone | `0bb1be8` -> `2cccdf8` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `60ef0fa` -> `69ae591` |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on September 19, 2024 at 04:18 PM UTC

Commits missing Jira IDs:
ec84a30d53043c9f0251508b7262c552a457a356


### Comment by @app-sre-bot on September 19, 2024 at 04:19 PM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on September 19, 2024 at 04:24 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1485?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 65.03%. Comparing base [(`d0a7214`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d0a7214f88dc42afad0581d6feb72e1d86421e4c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`ec84a30`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ec84a30d53043c9f0251508b7262c552a457a356?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1485   +/-   ##
=======================================
  Coverage   65.03%   65.03%           
=======================================
  Files         114      114           
  Lines        7880     7880           
=======================================
  Hits         5125     5125           
  Misses       2216     2216           
  Partials      539      539           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1485/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1485/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.03% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1485?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1485*
