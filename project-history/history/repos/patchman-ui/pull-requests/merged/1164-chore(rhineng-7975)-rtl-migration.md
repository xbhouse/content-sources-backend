---
type: pull_request
number: 1164
title: "chore(RHINENG-7975): RTL migration"
state: merged
author: mkholjuraev
created: 2024-02-22T13:04:47Z
updated: 2024-03-13T13:32:51Z
closed: 2024-02-28T12:00:10Z
merged: 2024-02-28T12:00:10Z
base_branch: master
head_branch: rtl-migration
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1164
---

# Pull Request #1164: chore(RHINENG-7975): RTL migration

**Author**: @mkholjuraev
**Created**: February 22, 2024 at 01:04 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `rtl-migration`

## Description

# Description

Associated Jira ticket: # [(issue)](https://issues.redhat.com/browse/RHINENG-7975)

- cleans up the tests by removing unnecessarily mocked redux selectors, components, functions
- rewrites enzyme tests to RTL
- tests for inventory tables are commented out, for now, to unblock the migration of PF5 and react. Also, we need to find a meaningful way of testing fed-modules.


# How to test the PR

Please include steps to test your PR.

# Before the change


# After the change


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

### Comment by @codecov-commenter on February 22, 2024 at 01:14 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1164?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `96.00000%` with `2 lines` in your changes are missing coverage. Please review.
> Project coverage is 62.86%. Comparing base [(`ab9b4d5`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ab9b4d5aa4444da2caf816ef0cf679d694a12871?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`c1786e3`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1164?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1164?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/TestingUtilities.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1164?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9UZXN0aW5nVXRpbGl0aWVzLmpz) | 95.65% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1164?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1164      +/-   ##
==========================================
- Coverage   63.60%   62.86%   -0.74%     
==========================================
  Files         126      127       +1     
  Lines        3154     3199      +45     
  Branches      817      820       +3     
==========================================
+ Hits         2006     2011       +5     
- Misses       1148     1188      +40     
```



</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1164?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on February 22, 2024 at 02:30 PM UTC

/retest

### Comment by @gkarat on February 28, 2024 at 10:30 AM UTC

@mkholjuraev, thanks for addressing my comments! Let's finish on InventoryTable part... other than that LGTM.

### Comment by @mkholjuraev on February 28, 2024 at 11:16 AM UTC

> @mkholjuraev, thanks for addressing my comments! Let's finish on InventoryTable part... other than that LGTM.

@gkarat thank you for another look. Maybe, we could merge this PR and work in parallel with you? I will continue with the tests for the Inventory pages and you can start your work on the migration.

### Comment by @mkholjuraev on February 28, 2024 at 11:32 AM UTC

/retest

### Comment by @mkholjuraev on February 28, 2024 at 11:58 AM UTC

the pr_check is failing due to vmass sync. Ignoring it for now. 

### Comment by @mkholjuraev on March 13, 2024 at 01:32 PM UTC

:tada: This PR is included in version 1.67.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Commented on February 26, 2024 at 11:16 AM UTC

### Review by @bastilian - Commented on February 26, 2024 at 11:17 AM UTC

### Review by @bastilian - Commented on February 26, 2024 at 11:25 AM UTC

### Review by @mkholjuraev - Commented on February 26, 2024 at 03:40 PM UTC

### Review by @mkholjuraev - Commented on February 26, 2024 at 03:41 PM UTC

### Review by @mkholjuraev - Commented on February 26, 2024 at 03:43 PM UTC

### Review by @bastilian - Commented on February 27, 2024 at 11:18 AM UTC

### Review by @bastilian - Commented on February 27, 2024 at 11:29 AM UTC

### Review by @gkarat - Commented on February 27, 2024 at 11:35 AM UTC

### Review by @gkarat - Commented on February 27, 2024 at 11:43 AM UTC

### Review by @gkarat - Commented on February 27, 2024 at 11:58 AM UTC

### Review by @mkholjuraev - Commented on February 27, 2024 at 12:11 PM UTC

### Review by @mkholjuraev - Commented on February 27, 2024 at 12:15 PM UTC

### Review by @mkholjuraev - Commented on February 27, 2024 at 12:15 PM UTC

### Review by @gkarat - Commented on February 27, 2024 at 12:16 PM UTC

### Review by @gkarat - Commented on February 27, 2024 at 12:17 PM UTC

### Review by @mkholjuraev - Commented on February 27, 2024 at 12:19 PM UTC

### Review by @mkholjuraev - Commented on February 27, 2024 at 01:26 PM UTC

### Review by @mkholjuraev - Commented on February 27, 2024 at 09:18 PM UTC

### Review by @gkarat - Commented on February 28, 2024 at 10:28 AM UTC

### Review by @gkarat - Approved on February 28, 2024 at 11:27 AM UTC

Agree, let's merge it. Thanks for this huge chunk of tests improved

### Review by @mkholjuraev - Commented on February 28, 2024 at 11:54 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1164*
