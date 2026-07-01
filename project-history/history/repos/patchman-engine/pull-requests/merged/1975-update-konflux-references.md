---
type: pull_request
number: 1975
title: "Update Konflux references"
state: merged
author: red-hat-konflux
created: 2025-12-13T04:39:40Z
updated: 2025-12-13T04:45:33Z
closed: 2025-12-13T04:45:03Z
merged: 2025-12-13T04:45:03Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1975
---

# Pull Request #1975: Update Konflux references

**Author**: @red-hat-konflux
**Created**: December 13, 2025 at 04:39 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-apply-tags | `a61d8a6` -> `e4017ec` |
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `803ae1e` -> `39561ac` |
| quay.io/konflux-ci/tekton-catalog/task-buildah-oci-ta | `916781b` -> `2de614f` |
| quay.io/konflux-ci/tekton-catalog/task-init | `d6a1010` -> `75b88ee` |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `90c2b32` -> `20eb21c` |

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

### Comment by @jira-linking on December 13, 2025 at 04:39 AM UTC

Commits missing Jira IDs:
fc94ad93bfd6e32bda311b5e1d7964d4293533d7


### Comment by @codecov-commenter on December 13, 2025 at 04:45 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1975?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`ccef583`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ccef583f87a3350575cdb810caae79496bb2ee92?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`fc94ad9`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/fc94ad93bfd6e32bda311b5e1d7964d4293533d7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1975   +/-   ##
=======================================
  Coverage   59.01%   59.01%           
=======================================
  Files         131      131           
  Lines        8493     8493           
=======================================
  Hits         5012     5012           
  Misses       2947     2947           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1975/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1975/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1975?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1975*
