---
type: pull_request
number: 987
title: "Revert \"Revert \"chore: do not wait advisories to display most impactf\u2026"
state: merged
author: mkholjuraev
created: 2023-03-10T09:03:17Z
updated: 2023-03-10T10:19:41Z
closed: 2023-03-10T10:02:33Z
merged: 2023-03-10T10:02:33Z
base_branch: master
head_branch: master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/987
---

# Pull Request #987: Revert "Revert "chore: do not wait advisories to display most impactf…

**Author**: @mkholjuraev
**Created**: March 10, 2023 at 09:03 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `master`

## Description

…ull statistics (#946)""

This reverts commit b0ceadca7885babba8d6ca3ffbd800a724f64610.

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

### Comment by @codecov-commenter on March 10, 2023 at 09:11 AM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/987?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00`**% and project coverage change: **`-1.59`** :warning:
> Comparison is base [(`b8ce309`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/b8ce309d356a31c14e907abe7384fada6dc843c7?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 70.61% compared to head [(`e5dcd8d`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/987?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 69.03%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #987      +/-   ##
==========================================
- Coverage   70.61%   69.03%   -1.59%     
==========================================
  Files         113      115       +2     
  Lines        2692     2800     +108     
  Branches      684      711      +27     
==========================================
+ Hits         1901     1933      +32     
- Misses        791      867      +76     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/987?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/987?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `100.00% <ø> (ø)` | |
| [...Components/StatusReports/AdvisoriesStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/987?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL0Fkdmlzb3JpZXNTdGF0dXNSZXBvcnQuanM=) | `53.33% <100.00%> (+33.33%)` | :arrow_up: |

... and [22 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/987/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/987?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 10, 2023 at 10:19 AM UTC

:tada: This PR is included in version 1.62.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/987*
