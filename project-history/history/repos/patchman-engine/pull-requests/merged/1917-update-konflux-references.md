---
type: pull_request
number: 1917
title: "Update Konflux references"
state: merged
author: red-hat-konflux
created: 2025-11-08T04:17:51Z
updated: 2025-11-15T04:54:39Z
closed: 2025-11-15T04:54:35Z
merged: 2025-11-15T04:54:34Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1917
---

# Pull Request #1917: Update Konflux references

**Author**: @red-hat-konflux
**Created**: November 08, 2025 at 04:17 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-apply-tags | `f44be1b` -> `4c2b0a2` |
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `d4c07e2` -> `31197f4` |
| quay.io/konflux-ci/tekton-catalog/task-buildah-oci-ta | `220bb6f` -> `27b04ea` |
| quay.io/konflux-ci/tekton-catalog/task-coverity-availability-check | `5623e48` -> `3640087` |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `f59175d` -> `462baed` |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies-oci-ta | `0503f93` -> `5946ca5` |
| quay.io/konflux-ci/tekton-catalog/task-sast-coverity-check-oci-ta | `ae62d14` -> `78f5244` |
| quay.io/konflux-ci/tekton-catalog/task-sast-shell-check-oci-ta | `1f0fcba` -> `d44336d` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check-oci-ta | `60f2dac` -> `8ad28b7` |
| quay.io/konflux-ci/tekton-catalog/task-sast-unicode-check-oci-ta | `1833c61` -> `e5a8d3e` |

---

### Configuration

📅 **Schedule**: Branch creation - Between 05:00 AM and 11:59 PM, only on Saturday ( * 5-23 * * 6 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

👻 **Immortal**: This PR will be recreated if closed unmerged. Get [config help](https://redirect.github.com/renovatebot/renovate/discussions) if that's undesired.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @codecov-commenter on November 08, 2025 at 04:24 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1917?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.90%. Comparing base ([`37c1027`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/37c102798b4e17b685bf403a126d80b455d618d7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`406a57a`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/406a57aa9a5dc025cc9b3e9d8ba3073f056a1219?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1917   +/-   ##
=======================================
  Coverage   58.90%   58.90%           
=======================================
  Files         131      131           
  Lines        8421     8421           
=======================================
  Hits         4960     4960           
  Misses       2927     2927           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1917/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1917/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.90% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1917?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1917*
