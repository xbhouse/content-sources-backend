---
type: pull_request
number: 1299
title: "chore(build): use shared Dockerfile"
state: merged
author: AsToNlele
created: 2025-02-27T09:03:58Z
updated: 2025-02-27T10:00:10Z
closed: 2025-02-27T10:00:10Z
merged: 2025-02-27T10:00:10Z
base_branch: master
head_branch: dockerfile
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1299
---

# Pull Request #1299: chore(build): use shared Dockerfile

**Author**: @AsToNlele
**Created**: February 27, 2025 at 09:03 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `dockerfile`

## Description

# Description

Associated Jira ticket: [RHINENG-16159](https://issues.redhat.com/browse/RHINENG-16159)

# Checklist:

- [x] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 27, 2025 at 09:12 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1299?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 62.04%. Comparing base [(`911f74b`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/911f74b1353d1ce4f56b051cd2646770e62ac1ef?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`3cf6275`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/3cf627589066037e81f05a0a87d496b7fc447ad6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1299   +/-   ##
=======================================
  Coverage   62.04%   62.04%           
=======================================
  Files         126      126           
  Lines        3367     3367           
  Branches      868      868           
=======================================
  Hits         2089     2089           
  Misses       1278     1278           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1299/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1299/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.04% <ø> (ø)` | |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1299/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (?)` | |
| [jest](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1299/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.04% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1299?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @LightOfHeaven1994 on February 27, 2025 at 10:00 AM UTC

I saw this PR check failures for other PRs, I think it's related to VMAAS service being broken

---

## Reviews

### Review by @LightOfHeaven1994 - Approved on February 27, 2025 at 09:59 AM UTC

LGTM, great work!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1299*
