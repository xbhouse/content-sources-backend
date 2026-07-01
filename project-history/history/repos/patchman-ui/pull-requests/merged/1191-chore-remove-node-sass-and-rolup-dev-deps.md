---
type: pull_request
number: 1191
title: "chore: remove node-sass and rolup dev deps"
state: merged
author: mkholjuraev
created: 2024-07-24T09:08:39Z
updated: 2024-07-24T12:40:53Z
closed: 2024-07-24T12:28:44Z
merged: 2024-07-24T12:28:44Z
base_branch: master
head_branch: rhineng-530
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1191
---

# Pull Request #1191: chore: remove node-sass and rolup dev deps

**Author**: @mkholjuraev
**Created**: July 24, 2024 at 09:08 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `rhineng-530`

## Description

# Description

Associated Jira ticket: # [(issue)](https://issues.redhat.com/browse/RHINENG-10530)

This PR also fixes the icon colors in CVEs modal, removes unnecessary dev-dependencies
# How to test the PR

1. Navigate to the Advisory All page, search RHSA-2024:3618 and go to the details page of that advisory 
2. Click on cve link 
3. Again click on CVE
4. Observe that CVE modal and whole page functions normally


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

### Comment by @codecov-commenter on July 24, 2024 at 12:09 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1191?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 64.00%. Comparing base [(`1d9863a`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1d9863a4bc186a033e743ccef307caa791c71a65?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`710e78d`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/710e78d9e0dd59555989236984b57e98e57e48af?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1191   +/-   ##
=======================================
  Coverage   64.00%   64.00%           
=======================================
  Files         124      124           
  Lines        3234     3234           
  Branches      830      830           
=======================================
  Hits         2070     2070           
  Misses       1164     1164           
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1191?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on July 24, 2024 at 12:28 PM UTC

IQE tests are failing, but according to logs, they are due to IQE misconfigurations. I am merging this PR.

### Comment by @mkholjuraev on July 24, 2024 at 12:40 PM UTC

:tada: This PR is included in version 1.67.8 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.8)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on July 24, 2024 at 11:44 AM UTC

This fixes the issue, LGTM! :smile: 

Make sure to fix the test

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1191*
