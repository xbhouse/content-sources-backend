---
type: pull_request
number: 1481
title: "Update Konflux references"
state: merged
author: red-hat-konflux
created: 2024-09-19T12:13:04Z
updated: 2024-09-19T15:35:58Z
closed: 2024-09-19T15:35:58Z
merged: 2024-09-19T15:35:58Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1481
---

# Pull Request #1481: Update Konflux references

**Author**: @red-hat-konflux
**Created**: September 19, 2024 at 12:13 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `8619eab` -> `e487185` |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `abe1bf2` -> `772233e` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `9254f82` -> `412e96e` |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `d1836ac` -> `592bbce` |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies | `fd1fda0` -> `fe7234e` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `4b179c3` -> `60ef0fa` |

---

 - [x] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on September 19, 2024 at 12:13 PM UTC

Commits missing Jira IDs:
5ee23bfdb7627819125f08672c8c9d94167db102


### Comment by @app-sre-bot on September 19, 2024 at 12:13 PM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on September 19, 2024 at 12:18 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1481?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 65.06%. Comparing base [(`4fd71a1`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4fd71a170ad3de7357dd4e685d94138d98a8ce76?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`bba9a67`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bba9a6780ad80b23593ac81e0f37cb6189f203aa?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1481   +/-   ##
=======================================
  Coverage   65.06%   65.06%           
=======================================
  Files         114      114           
  Lines        7877     7877           
=======================================
  Hits         5125     5125           
  Misses       2215     2215           
  Partials      537      537           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1481/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1481/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.06% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1481?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1481*
