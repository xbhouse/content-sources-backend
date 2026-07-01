---
type: pull_request
number: 1207
title: "chore(deps): update konflux references"
state: merged
author: red-hat-konflux
created: 2024-10-21T04:17:11Z
updated: 2024-10-25T19:12:06Z
closed: 2024-10-23T14:59:27Z
merged: 2024-10-23T14:59:27Z
base_branch: master
head_branch: konflux/references/master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1207
---

# Pull Request #1207: chore(deps): update konflux references

**Author**: @red-hat-konflux
**Created**: October 21, 2024 at 04:17 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `e487185` -> `327d745` |
| quay.io/konflux-ci/tekton-catalog/task-buildah | `2d6e09f` -> `504e29e` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `1e29eeb` -> `a94b652` |
| quay.io/konflux-ci/tekton-catalog/task-push-dockerfile | `585106f` -> `674e70f` |
| quay.io/konflux-ci/tekton-catalog/task-show-sbom | `9bfc6b9` -> `52f8b96` |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @app-sre-bot on October 21, 2024 at 04:19 AM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on October 21, 2024 at 04:50 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1207?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 64.01%. Comparing base [(`d81c9d4`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d81c9d48fcd3b728c4a9f545fe98e4ec416bfa6f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`def8294`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/def8294126b5bdb2c60bb53777f9419df3b2c3e0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1207   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1207?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @LightOfHeaven1994 on October 22, 2024 at 01:41 PM UTC

/retest

### Comment by @LightOfHeaven1994 on October 22, 2024 at 02:13 PM UTC

/retest

### Comment by @LightOfHeaven1994 on October 22, 2024 at 09:33 PM UTC

/retest

### Comment by @mkholjuraev on October 25, 2024 at 07:12 PM UTC

:tada: This PR is included in version 1.68.1 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.68.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @LightOfHeaven1994 - Approved on October 23, 2024 at 02:59 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1207*
