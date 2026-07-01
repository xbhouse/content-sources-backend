---
type: pull_request
number: 2015
title: "Update Konflux references"
state: merged
author: red-hat-konflux
created: 2026-01-17T04:43:11Z
updated: 2026-01-21T12:46:48Z
closed: 2026-01-21T12:46:48Z
merged: 2026-01-21T12:46:48Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2015
---

# Pull Request #2015: Update Konflux references

**Author**: @red-hat-konflux
**Created**: January 17, 2026 at 04:43 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change | Notes |
|---|---|---|
| quay.io/konflux-ci/tekton-catalog/task-apply-tags | `0.2` -> `0.3` | :warning:[migration](https://redirect.github.com/redhat-appstudio/build-definitions/blob/main/task/apply-tags/0.3/MIGRATION.md):warning: |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `7c2a32d` -> `654b989` |  |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `f3d2d17` -> `b2f2559` |  |
| quay.io/konflux-ci/tekton-catalog/task-coverity-availability-check | `3640087` -> `267d5bc` |  |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `808fe09` -> `1cf21de` |  |
| quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks | `04f7559` -> `b612fd7` |  |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies-oci-ta | `3620777` -> `4c9ff41` |  |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `0b10508` -> `0041778` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-coverity-check-oci-ta | `78f5244` -> `9978b61` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-shell-check-oci-ta | `d44336d` -> `e7a5157` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check-oci-ta | `0eca130` -> `49b7d09` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-unicode-check-oci-ta | `e5a8d3e` -> `8817f50` |  |

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

### Comment by @jira-linking on January 17, 2026 at 04:43 AM UTC

Commits missing Jira IDs:
be5f78648d3ccc21522e1d5811a229c58ba8711e


### Comment by @codecov-commenter on January 17, 2026 at 04:49 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2015?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`506bc9a`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/506bc9a35a8ce23d87c81dff114f8f8e569e4a65?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`be5f786`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/be5f78648d3ccc21522e1d5811a229c58ba8711e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2015   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2015/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2015/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2015?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @TenSt - Approved on January 21, 2026 at 12:46 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2015*
