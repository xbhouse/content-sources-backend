---
type: pull_request
number: 1524
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-11-01T22:26:43Z
updated: 2026-04-06T05:22:10Z
closed: 2024-11-07T09:15:06Z
merged: 2024-11-07T09:15:06Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1524
---

# Pull Request #1524: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: November 01, 2024 at 10:26 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `38fdb70` -> `715fa1f` |
| quay.io/konflux-ci/tekton-catalog/task-buildah | `185d52f` -> `fedcfe0` |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `28fee4b` -> `90e371f` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `a94b652` -> `1981b5a` |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `b4f9599` -> `443ffa8` |
| quay.io/konflux-ci/tekton-catalog/task-push-dockerfile | `674e70f` -> `a216178` |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `7aa4d3c` -> `0c9667f` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `60ed62a` -> `8aab0fd` |
| quay.io/konflux-ci/tekton-catalog/task-source-build | `21cb5eb` -> `e2e7aab` |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on November 01, 2024 at 10:26 PM UTC

Commits missing Jira IDs:
1d71d35e6684a933940569919271dd0410cdfd48


### Comment by @codecov-commenter on November 01, 2024 at 10:31 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1524?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 63.45%. Comparing base ([`e6ad77f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e6ad77fc74286ef84a452ca0c27234d58a975284?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`1d71d35`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1d71d35e6684a933940569919271dd0410cdfd48?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1115 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1524      +/-   ##
==========================================
- Coverage   63.50%   63.45%   -0.06%     
==========================================
  Files         114      114              
  Lines        9609     9609              
==========================================
- Hits         6102     6097       -5     
- Misses       2970     2975       +5     
  Partials      537      537              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1524/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1524/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.45% <ø> (-0.06%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1524?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1524*
