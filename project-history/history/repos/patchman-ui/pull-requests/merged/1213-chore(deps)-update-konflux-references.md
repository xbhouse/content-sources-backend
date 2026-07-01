---
type: pull_request
number: 1213
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-10-28T04:10:41Z
updated: 2024-10-30T09:48:31Z
closed: 2024-10-28T09:15:32Z
merged: 2024-10-28T09:15:32Z
base_branch: master
head_branch: konflux/references/master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1213
---

# Pull Request #1213: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: October 28, 2024 at 04:10 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change | Notes |
|---|---|---|
| quay.io/konflux-ci/tekton-catalog/task-buildah | `504e29e` -> `b105a3b` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `0.2` -> `0.3` | :warning:[migration](https://togithub.com/redhat-appstudio/build-definitions/blob/main/task/sast-snyk-check/0.3/MIGRATION.md):warning: |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @app-sre-bot on October 28, 2024 at 04:10 AM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on October 28, 2024 at 04:31 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1213?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 64.01%. Comparing base [(`5a2efdb`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/5a2efdb6575429875a90fed1adf30254e639ac61?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`dd7d713`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/dd7d7132d8a41d3ef5a4c2ffb843df3ede824cbb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1213   +/-   ##
=======================================
  Coverage   64.01%   64.01%           
=======================================
  Files         124      124           
  Lines        3235     3235           
  Branches      831      831           
=======================================
  Hits         2071     2071           
  Misses       1164     1164           
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1213?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on October 30, 2024 at 09:48 AM UTC

:tada: This PR is included in version 1.68.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.68.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @LightOfHeaven1994 - Approved on October 28, 2024 at 09:15 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1213*
