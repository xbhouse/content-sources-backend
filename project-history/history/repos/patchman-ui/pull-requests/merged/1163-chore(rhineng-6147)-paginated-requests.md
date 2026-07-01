---
type: pull_request
number: 1163
title: "chore(RHINENG-6147): Paginated requests"
state: merged
author: mkholjuraev
created: 2024-02-15T20:43:49Z
updated: 2024-03-13T13:32:29Z
closed: 2024-03-05T16:52:51Z
merged: 2024-03-05T16:52:51Z
base_branch: master
head_branch: paginated-requests
labels: ["enhancement", "released", "refactor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1163
---

# Pull Request #1163: chore(RHINENG-6147): Paginated requests

**Author**: @mkholjuraev
**Created**: February 15, 2024 at 08:43 PM UTC
**Status**: Merged
**Labels**: `enhancement`, `released`, `refactor`
**Base**: `master` ← **Head**: `paginated-requests`

## Description

# Description

Associated Jira ticket: # RHINENG-6147

1. Selection of all items was using API calls with limit=-1 across all pages where bulk selection is enabled. This PR converts it to use paginated requests.

2. remediation issues fetching also fetching 1000 items per request, now it is reduced to 100.

3. Unassigning/assigning templates had another usage of limit=-1. Now they are converted to use paginated requests.
# How to test the PR

Please include steps to test your PR.
1. Run the PR locally
2. Observe that bulk selection works as expected in all tables
3. Observe that you can assign/unassign templates as expected and there is no change in functionality
4. Play with the remediation using both in bulk or from the table row actions, and observe the number of issues is correct.

# Before the change


# After the change
Resolves: [RHINENG](https://issues.redhat.com/browse/RHINENG-6147)

# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 16, 2024 at 01:19 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1163?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `2 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`ab9b4d5`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ab9b4d5aa4444da2caf816ef0cf679d694a12871?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.60% compared to head [(`daa004e`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1163?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.69%.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1163?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/RemediationPairs.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1163?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9SZW1lZGlhdGlvblBhaXJzLmpz) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1163?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1163      +/-   ##
==========================================
+ Coverage   63.60%   63.69%   +0.09%     
==========================================
  Files         126      126              
  Lines        3154     3173      +19     
  Branches      817      822       +5     
==========================================
+ Hits         2006     2021      +15     
- Misses       1148     1152       +4     
```



</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1163?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on February 28, 2024 at 11:44 AM UTC

/retest

### Comment by @mkholjuraev on March 05, 2024 at 04:46 PM UTC

@gkarat thank you for the review!

### Comment by @mkholjuraev on March 13, 2024 at 01:31 PM UTC

:tada: This PR is included in version 1.67.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Commented on March 05, 2024 at 11:40 AM UTC

### Review by @mkholjuraev - Commented on March 05, 2024 at 01:36 PM UTC

### Review by @mkholjuraev - Commented on March 05, 2024 at 02:23 PM UTC

### Review by @gkarat - Approved on March 05, 2024 at 03:38 PM UTC

OK, it feels like https://github.com/RedHatInsights/frontend-components/pull/1954 hasn't been released yet for some reason. Let's merge this PR if you feel ready.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1163*
