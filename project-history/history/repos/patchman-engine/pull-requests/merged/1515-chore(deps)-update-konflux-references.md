---
type: pull_request
number: 1515
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-10-21T16:13:18Z
updated: 2024-10-24T08:25:01Z
closed: 2024-10-24T08:25:01Z
merged: 2024-10-24T08:25:01Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1515
---

# Pull Request #1515: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: October 21, 2024 at 04:13 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-buildah | `67f0290` -> `504e29e` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `1e29eeb` -> `a94b652` |
| quay.io/konflux-ci/tekton-catalog/task-show-sbom | `9bfc6b9` -> `52f8b96` |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on October 21, 2024 at 04:13 PM UTC

Commits missing Jira IDs:
fe425da76b8f550f48d8a1629a2037ab0cc2bf53


### Comment by @codecov-commenter on October 21, 2024 at 04:19 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1515?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.50%. Comparing base [(`8e614a9`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8e614a996fc3ff7e664e7264e90e0d7a99033115?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`fe425da`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/fe425da76b8f550f48d8a1629a2037ab0cc2bf53?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1515   +/-   ##
=======================================
  Coverage   63.50%   63.50%           
=======================================
  Files         114      114           
  Lines        9609     9609           
=======================================
  Hits         6102     6102           
  Misses       2970     2970           
  Partials      537      537           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1515/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1515/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.50% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1515?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on October 24, 2024 at 07:36 AM UTC

/retest

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1515*
