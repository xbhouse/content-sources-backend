---
type: pull_request
number: 1536
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-11-23T04:18:21Z
updated: 2024-11-23T04:23:59Z
closed: 2024-11-23T04:23:55Z
merged: 2024-11-23T04:23:55Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1536
---

# Pull Request #1536: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: November 23, 2024 at 04:18 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `5da8c2f` -> `ebc17bb` |
| quay.io/konflux-ci/tekton-catalog/task-buildah | `bbacdfe` -> `7779f9e` |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `eff773a` -> `0a54211` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `1465898` -> `b4f450f` |
| quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks | `2e37ec3` -> `df8a25a` |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies | `f59a214` -> `47d8d33` |
| quay.io/konflux-ci/tekton-catalog/task-push-dockerfile | `a216178` -> `48bb2ee` |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `0c9667f` -> `28aaf87` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `f165b1c` -> `9fa8acb` |
| quay.io/konflux-ci/tekton-catalog/task-show-sbom | `52f8b96` -> `945a7c9` |
| quay.io/konflux-ci/tekton-catalog/task-source-build | `53a41b0` -> `ddfa1fb` |
| quay.io/konflux-ci/tekton-catalog/task-summary | `b0f049f` -> `870d9a0` |

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on saturday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

👻 **Immortal**: This PR will be recreated if closed unmerged. Get [config help](https://togithub.com/renovatebot/renovate/discussions) if that's undesired.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on November 23, 2024 at 04:18 AM UTC

Commits missing Jira IDs:
298c885b0e2dea40cc6db5ddb6ccf2010dfdeebe


### Comment by @codecov-commenter on November 23, 2024 at 04:23 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1536?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.45%. Comparing base [(`909e0ce`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/909e0cebf086fb6434f96a215803adcf0685272c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`298c885`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/298c885b0e2dea40cc6db5ddb6ccf2010dfdeebe?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1536   +/-   ##
=======================================
  Coverage   63.45%   63.45%           
=======================================
  Files         114      114           
  Lines        9637     9637           
=======================================
  Hits         6115     6115           
  Misses       2980     2980           
  Partials      542      542           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1536/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1536/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.45% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1536?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

----
🚨 Try these New Features:

- [Flaky Tests Detection](https://docs.codecov.com/docs/test-result-ingestion-beta) - Detect and resolve failed and flaky tests

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1536*
