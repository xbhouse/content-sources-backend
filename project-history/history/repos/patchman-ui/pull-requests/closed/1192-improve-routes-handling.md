---
type: pull_request
number: 1192
title: "Improve routes handling"
state: closed
author: gkarat
created: 2024-08-19T11:12:33Z
updated: 2025-05-05T09:35:35Z
closed: 2025-05-05T09:35:35Z
base_branch: master
head_branch: improve-routes
labels: ["refactor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1192
---

# Pull Request #1192: Improve routes handling

**Author**: @gkarat
**Created**: August 19, 2024 at 11:12 AM UTC
**Status**: Closed
**Labels**: `refactor`
**Base**: `master` ← **Head**: `improve-routes`

## Description

# Description

Associated Jira ticket: No jira.

This slightly improves the routes handling in Patch. Also removing ErrorComponent prop passing from ZeroState since it's not simply supported.

# How to test the PR

No functionality should change. The routes must work correctly.

# Before the change

The prior routes setup is more error prone or has unnecessary layers.

# After the change

Life should be... uhm... better.

# Dependent work link

-

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on August 19, 2024 at 11:38 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1192?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `0%` with `11 lines` in your changes missing coverage. Please review.
> Project coverage is 63.97%. Comparing base [(`c1e56c0`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/c1e56c0efa157e88f5effef10b730a36e0607f89?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`c7be848`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/c7be848d351860055f620adbdb0c174e381cff81?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1192?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/AccessWrapper.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1192?src=pr&el=tree&filepath=src%2FAccessWrapper.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FjY2Vzc1dyYXBwZXIuanM=) | 0.00% | [8 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Routes.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1192?src=pr&el=tree&filepath=src%2FRoutes.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1192?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1192      +/-   ##
==========================================
- Coverage   64.01%   63.97%   -0.04%     
==========================================
  Files         124      125       +1     
  Lines        3235     3237       +2     
  Branches      831      833       +2     
==========================================
  Hits         2071     2071              
- Misses       1164     1166       +2     
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1192?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @bastilian - Commented on August 19, 2024 at 11:15 AM UTC

### Review by @gkarat - Commented on August 22, 2024 at 10:11 AM UTC

### Review by @bastilian - Commented on August 27, 2024 at 04:41 PM UTC

### Review by @gkarat - Commented on August 28, 2024 at 12:17 PM UTC

### Review by @gkarat - Commented on August 28, 2024 at 12:23 PM UTC

### Review by @bastilian - Commented on August 28, 2024 at 03:30 PM UTC

### Review by @gkarat - Commented on August 29, 2024 at 09:59 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1192*
