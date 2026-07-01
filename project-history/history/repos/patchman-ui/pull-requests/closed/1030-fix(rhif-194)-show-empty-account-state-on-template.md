---
type: pull_request
number: 1030
title: "fix(RHIF-194): show empty account state on template details"
state: closed
author: mkholjuraev
created: 2023-04-21T09:21:24Z
updated: 2024-04-03T09:22:33Z
closed: 2023-05-10T08:01:42Z
base_branch: master
head_branch: empty-states
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1030
---

# Pull Request #1030: fix(RHIF-194): show empty account state on template details

**Author**: @mkholjuraev
**Created**: April 21, 2023 at 09:21 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `empty-states`

## Description

# Description

Associated Jira ticket: # https://issues.redhat.com/browse/RHIF-194

Enables rendering **NoRegisteredSystems** on Patch set page.  

# How to test the PR

1. Run the pr locally
2. login to some empty account
3. Navigate to templates page.
4. observer that table is hidden and NoRegisteredSystem component is shown.


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

### Comment by @codecov-commenter on April 21, 2023 at 09:33 AM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1030?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`45.45`**% and project coverage change: **`-0.06`** :warning:
> Comparison is base [(`2b1ed44`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/2b1ed4460f5a572a3782caf553cb0ede8b7ecea0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.22% compared to head [(`72ccba3`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1030?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.17%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1030      +/-   ##
==========================================
- Coverage   64.22%   64.17%   -0.06%     
==========================================
  Files         117      117              
  Lines        2840     2847       +7     
  Branches      735      740       +5     
==========================================
+ Hits         1824     1827       +3     
- Misses       1016     1020       +4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1030?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1030?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [.../PresentationalComponents/Snippets/ErrorHandler.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1030?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FcnJvckhhbmRsZXIuanM=) | `43.47% <50.00%> (+2.56%)` | :arrow_up: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1030?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `89.47% <75.00%> (-1.27%)` | :arrow_down: |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1030?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `84.90% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1030?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on April 27, 2023 at 08:07 AM UTC

I am putting this into draft state as there is a new zero state onboarding components are being created


### Comment by @mkholjuraev on May 10, 2023 at 08:01 AM UTC

Not valid anymore

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1030*
