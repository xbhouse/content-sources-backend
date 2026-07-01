---
type: pull_request
number: 1225
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-11-23T08:26:37Z
updated: 2024-11-26T17:28:36Z
closed: 2024-11-26T17:28:36Z
merged: 2024-11-26T17:28:36Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1225
---

# Pull Request #1225: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: November 23, 2024 at 08:26 AM UTC
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

📅 **Schedule**: Branch creation - "after 5am on saturday" (UTC), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

👻 **Immortal**: This PR will be recreated if closed unmerged. Get [config help](https://togithub.com/renovatebot/renovate/discussions) if that's undesired.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @codecov-commenter on November 23, 2024 at 09:17 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1225?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.98%. Comparing base [(`1caa664`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1caa6646b9c104e7d37f139f7105f8bf38f120be?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`e3e13b9`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/e3e13b98bad0d6ab3877bd30baa0225694c9f20d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1225   +/-   ##
=======================================
  Coverage   63.98%   63.98%           
=======================================
  Files         124      124           
  Lines        3210     3210           
  Branches      821      821           
=======================================
  Hits         2054     2054           
  Misses       1156     1156           
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1225?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

----
🚨 Try these New Features:

- [Flaky Tests Detection](https://docs.codecov.com/docs/test-result-ingestion-beta) - Detect and resolve failed and flaky tests
- [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis) - Avoid shipping oversized bundles

### Comment by @hugohezel on November 25, 2024 at 09:02 AM UTC

/retest

### Comment by @hugohezel on November 25, 2024 at 09:37 AM UTC

/retest

### Comment by @hugohezel on November 26, 2024 at 01:33 PM UTC

/retest

---

## Reviews

### Review by @hugohezel - Approved on November 26, 2024 at 05:28 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1225*
