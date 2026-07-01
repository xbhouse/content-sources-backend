---
type: pull_request
number: 1942
title: "Update Konflux references"
state: merged
author: red-hat-konflux
created: 2025-11-22T04:37:59Z
updated: 2025-11-29T04:43:41Z
closed: 2025-11-29T04:42:43Z
merged: 2025-11-29T04:42:43Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1942
---

# Pull Request #1942: Update Konflux references

**Author**: @red-hat-konflux
**Created**: November 22, 2025 at 04:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change | Notes |
|---|---|---|
| quay.io/konflux-ci/tekton-catalog/task-apply-tags | `4c2b0a2` -> `a61d8a6` |  |
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `0.1` -> `0.2` | :warning:[migration](https://redirect.github.com/redhat-appstudio/build-definitions/blob/main/task/build-image-index/0.2/MIGRATION.md):warning: |
| quay.io/konflux-ci/tekton-catalog/task-buildah-oci-ta | `0.6` -> `0.7` | :warning:[migration](https://redirect.github.com/redhat-appstudio/build-definitions/blob/main/task/buildah-oci-ta/0.7/MIGRATION.md):warning: |
| quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks | `9568c51` -> `04f7559` |  |
| quay.io/konflux-ci/tekton-catalog/task-git-clone-oci-ta | `3dc39ea` -> `ea64f5b` |  |
| quay.io/konflux-ci/tekton-catalog/task-init | `3ca52e1` -> `4072de8` |  |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies-oci-ta | `5946ca5` -> `3fa0204` |  |
| quay.io/konflux-ci/tekton-catalog/task-push-dockerfile-oci-ta | `13633d5` -> `08bba4a` |  |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `1b6c20a` -> `13cf619` |  |
| quay.io/konflux-ci/tekton-catalog/task-source-build-oci-ta | `282cb5a` -> `4abb2db` |  |

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

### Comment by @codecov-commenter on November 29, 2025 at 04:42 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1942?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.83%. Comparing base ([`b97b652`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b97b6527474e6fc7c7ad9c4ce007c6225ff5dfbb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`988deed`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/988deed28fab5c399883e22279605057b62ff002?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1942   +/-   ##
=======================================
  Coverage   58.83%   58.83%           
=======================================
  Files         131      131           
  Lines        8407     8407           
=======================================
  Hits         4946     4946           
  Misses       2927     2927           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1942/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1942/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.83% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1942?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1942*
