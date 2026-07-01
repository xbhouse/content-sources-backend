---
type: pull_request
number: 1541
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-12-07T04:23:02Z
updated: 2024-12-07T04:32:33Z
closed: 2024-12-07T04:32:33Z
merged: 2024-12-07T04:32:33Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1541
---

# Pull Request #1541: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: December 07, 2024 at 04:23 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change | Notes |
|---|---|---|
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `7b2c5ab` -> `a89c141` |  |
| quay.io/konflux-ci/tekton-catalog/task-buildah | `7d3f090` -> `2388313` |  |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `0.1` -> `0.2` | :warning:[migration](https://redirect.github.com/redhat-appstudio/build-definitions/blob/main/task/clamav-scan/0.2/MIGRATION.md):warning: |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies | `d3d8a8b` -> `53fc6d8` |  |
| quay.io/konflux-ci/tekton-catalog/task-source-build | `957f765` -> `bacd55a` |  |

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on saturday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

👻 **Immortal**: This PR will be recreated if closed unmerged. Get [config help](https://redirect.github.com/renovatebot/renovate/discussions) if that's undesired.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC4xMzIuMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOC4xMzIuMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @jira-linking on December 07, 2024 at 04:23 AM UTC

Commits missing Jira IDs:
bc7b858bcc626684fbd6a302ab135320894fc14c


### Comment by @codecov-commenter on December 07, 2024 at 04:28 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1541?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.33%. Comparing base [(`926aa02`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/926aa02f254f16afbfec8ac182c30c1c702b228a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`bc7b858`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bc7b858bcc626684fbd6a302ab135320894fc14c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1541   +/-   ##
=======================================
  Coverage   63.33%   63.33%           
=======================================
  Files         115      115           
  Lines        9673     9673           
=======================================
  Hits         6126     6126           
  Misses       3003     3003           
  Partials      544      544           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1541/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1541/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.33% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1541?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1541*
