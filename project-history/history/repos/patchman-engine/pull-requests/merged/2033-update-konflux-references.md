---
type: pull_request
number: 2033
title: "Update Konflux references"
state: merged
author: red-hat-konflux
created: 2026-01-31T04:46:31Z
updated: 2026-01-31T04:52:21Z
closed: 2026-01-31T04:52:08Z
merged: 2026-01-31T04:52:08Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2033
---

# Pull Request #2033: Update Konflux references

**Author**: @red-hat-konflux
**Created**: January 31, 2026 at 04:46 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change | Notes |
|---|---|---|
| quay.io/konflux-ci/tekton-catalog/task-buildah-oci-ta | `0.7` -> `0.8` | :warning:[migration](https://redirect.github.com/redhat-appstudio/build-definitions/blob/main/task/buildah-oci-ta/0.8/MIGRATION.md):warning: |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `654b989` -> `dadfea7` |  |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `78f0349` -> `76efc01` |  |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `1cf21de` -> `e3a55cc` |  |
| quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks | `329b149` -> `204fd39` |  |
| quay.io/konflux-ci/tekton-catalog/task-init | `b349d24` -> `ebf0677` |  |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies-oci-ta | `c651d76` -> `c664a6d` |  |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `f2df541` -> `fb6c97a` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-coverity-check-oci-ta | `9978b61` -> `9d0bc70` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check-oci-ta | `c5e640a` -> `a70272a` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-unicode-check-oci-ta | `0ca0203` -> `1818a5b` |  |

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

### Comment by @jira-linking on January 31, 2026 at 04:46 AM UTC

Commits missing Jira IDs:
477cee4da8990b19d56b369d1f29bd02d1d08f6d


### Comment by @codecov-commenter on January 31, 2026 at 04:52 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2033?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.25%. Comparing base ([`ddbdbdd`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ddbdbdd35c3509584cd2b242e4f116a4f1cb53a9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`477cee4`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/477cee4da8990b19d56b369d1f29bd02d1d08f6d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2033   +/-   ##
=======================================
  Coverage   59.25%   59.25%           
=======================================
  Files         134      134           
  Lines        8615     8615           
=======================================
  Hits         5105     5105           
  Misses       2967     2967           
  Partials      543      543           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2033/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2033/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.25% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2033?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2033*
