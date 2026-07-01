---
type: pull_request
number: 2020
title: "Update Konflux references"
state: merged
author: red-hat-konflux
created: 2026-01-24T04:37:59Z
updated: 2026-01-24T04:43:36Z
closed: 2026-01-24T04:43:27Z
merged: 2026-01-24T04:43:26Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2020
---

# Pull Request #2020: Update Konflux references

**Author**: @red-hat-konflux
**Created**: January 24, 2026 at 04:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change | Notes |
|---|---|---|
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `985d1ef` -> `05d3d8a` |  |
| quay.io/konflux-ci/tekton-catalog/task-buildah-oci-ta | `7d5818e` -> `c422f1e` |  |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `b2f2559` -> `78f0349` |  |
| quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks | `b612fd7` -> `329b149` |  |
| quay.io/konflux-ci/tekton-catalog/task-git-clone-oci-ta | `56f65a1` -> `306b69e` |  |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies-oci-ta | `4c9ff41` -> `c651d76` |  |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `0041778` -> `f2df541` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check-oci-ta | `49b7d09` -> `c5e640a` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-unicode-check-oci-ta | `0.3` -> `0.4` | :warning:[migration](https://redirect.github.com/redhat-appstudio/build-definitions/blob/main/task/sast-unicode-check-oci-ta/0.4/MIGRATION.md):warning: |
| quay.io/konflux-ci/tekton-catalog/task-show-sbom | `beb0616` -> `e2c1b4e` |  |

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on January 24, 2026 at 04:38 AM UTC

Commits missing Jira IDs:
25b8034d4e3f5676e9d901719160779cf8ec6fc8


### Comment by @codecov-commenter on January 24, 2026 at 04:43 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2020?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.18%. Comparing base ([`d6ac7ee`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d6ac7ee256d5b2817dbf9676350aa05659a83852?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`25b8034`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/25b8034d4e3f5676e9d901719160779cf8ec6fc8?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2020   +/-   ##
=======================================
  Coverage   59.18%   59.18%           
=======================================
  Files         133      133           
  Lines        8599     8599           
=======================================
  Hits         5089     5089           
  Misses       2967     2967           
  Partials      543      543           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2020/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2020/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.18% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2020?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2020*
