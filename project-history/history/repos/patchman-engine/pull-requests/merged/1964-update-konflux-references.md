---
type: pull_request
number: 1964
title: "Update Konflux references"
state: merged
author: red-hat-konflux
created: 2025-12-06T04:39:52Z
updated: 2025-12-06T04:45:48Z
closed: 2025-12-06T04:45:19Z
merged: 2025-12-06T04:45:19Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1964
---

# Pull Request #1964: Update Konflux references

**Author**: @red-hat-konflux
**Created**: December 06, 2025 at 04:39 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `8ec7d7b` -> `ee558db` |
| quay.io/konflux-ci/tekton-catalog/task-git-clone-oci-ta | `ea64f5b` -> `0a89e1a` |
| quay.io/konflux-ci/tekton-catalog/task-init | `4072de8` -> `d6a1010` |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `13cf619` -> `90c2b32` |

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

### Comment by @jira-linking on December 06, 2025 at 04:39 AM UTC

Commits missing Jira IDs:
892df85f4be2fca7203f69a5361bb63d05e12d04


### Comment by @codecov-commenter on December 06, 2025 at 04:45 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1964?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.84%. Comparing base ([`7aa3456`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/7aa3456cb08941eb925bfc8a9ef1999ac0ce411d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`892df85`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/892df85f4be2fca7203f69a5361bb63d05e12d04?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1964   +/-   ##
=======================================
  Coverage   58.84%   58.84%           
=======================================
  Files         131      131           
  Lines        8436     8436           
=======================================
  Hits         4964     4964           
  Misses       2937     2937           
  Partials      535      535           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1964/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1964/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.84% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1964?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1964*
