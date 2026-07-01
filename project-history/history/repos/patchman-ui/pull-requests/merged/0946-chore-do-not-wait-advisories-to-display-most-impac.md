---
type: pull_request
number: 946
title: "chore: do not wait advisories to display most impactfull statistics"
state: merged
author: mkholjuraev
created: 2023-01-12T20:31:21Z
updated: 2023-01-18T15:46:04Z
closed: 2023-01-12T22:00:02Z
merged: 2023-01-12T22:00:02Z
base_branch: master
head_branch: master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/946
---

# Pull Request #946: chore: do not wait advisories to display most impactfull statistics

**Author**: @mkholjuraev
**Created**: January 12, 2023 at 08:31 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `master`

## Description

# Description

Associated Jira ticket: # (issue)

Please include a summary of the change, what this fixes/creates/improves.


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

### Comment by @codecov-commenter on January 12, 2023 at 09:51 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/946?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.30**% // Head: **72.44**% // Increases project coverage by **`+0.14%`** :tada:
> Coverage data is based on head [(`c4fc9c7`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/946?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`235cdd0`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/235cdd0894b7a4c06c6f74c5a149b02f146d06ca?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 100.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #946      +/-   ##
==========================================
+ Coverage   72.30%   72.44%   +0.14%     
==========================================
  Files         110      110              
  Lines        2618     2617       -1     
  Branches      662      661       -1     
==========================================
+ Hits         1893     1896       +3     
+ Misses        725      721       -4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/946?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/946?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `100.00% <ø> (ø)` | |
| [...Components/StatusReports/AdvisoriesStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/946?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL0Fkdmlzb3JpZXNTdGF0dXNSZXBvcnQuanM=) | `53.33% <100.00%> (+33.33%)` | :arrow_up: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/946?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `58.18% <0.00%> (-0.91%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/946?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on January 18, 2023 at 03:45 PM UTC

:tada: This PR is included in version 1.58.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.58.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.58.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/946*
