---
type: pull_request
number: 2044
title: "Update Konflux references"
state: merged
author: red-hat-konflux
created: 2026-02-07T04:49:01Z
updated: 2026-02-07T04:57:07Z
closed: 2026-02-07T04:57:07Z
merged: 2026-02-07T04:57:07Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2044
---

# Pull Request #2044: Update Konflux references

**Author**: @red-hat-konflux
**Created**: February 07, 2026 at 04:49 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change | Notes |
|---|---|---|
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `05d3d8a` -> `8c422a5` |  |
| quay.io/konflux-ci/tekton-catalog/task-buildah-oci-ta | `8a98418` -> `ba95646` |  |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `dadfea7` -> `3ff4d1c` |  |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `76efc01` -> `4f5ccf2` |  |
| quay.io/konflux-ci/tekton-catalog/task-coverity-availability-check | `267d5bc` -> `a24d8f3` |  |
| quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks | `204fd39` -> `33b7133` |  |
| quay.io/konflux-ci/tekton-catalog/task-init | `0.2` -> `0.3` | :warning:[migration](https://redirect.github.com/redhat-appstudio/build-definitions/blob/main/task/init/0.3/MIGRATION.md):warning: |
| quay.io/konflux-ci/tekton-catalog/task-push-dockerfile-oci-ta | `08bba4a` -> `6fb61be` |  |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `fb6c97a` -> `a99d8fd` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-coverity-check-oci-ta | `9d0bc70` -> `e8c6357` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-shell-check-oci-ta | `e7a5157` -> `f475b4b` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check-oci-ta | `a70272a` -> `0c2ab8c` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-unicode-check-oci-ta | `1818a5b` -> `b38140b` |  |

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

### Comment by @jira-linking on February 07, 2026 at 04:49 AM UTC

Commits missing Jira IDs:
03511fc5b2c521c4ce7152c65416f94170239c9e


### Comment by @codecov-commenter on February 07, 2026 at 04:54 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2044?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.39%. Comparing base ([`9783708`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/97837089756913646fd9d15cf1b244aeae52ad0e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`03511fc`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/03511fc5b2c521c4ce7152c65416f94170239c9e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2044   +/-   ##
=======================================
  Coverage   59.39%   59.39%           
=======================================
  Files         134      134           
  Lines        8678     8678           
=======================================
  Hits         5154     5154           
  Misses       2977     2977           
  Partials      547      547           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2044/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2044/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.39% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2044?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2044*
