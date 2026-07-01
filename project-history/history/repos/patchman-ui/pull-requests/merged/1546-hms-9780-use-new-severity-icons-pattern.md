---
type: pull_request
number: 1546
title: "HMS-9780: use new severity icons pattern"
state: merged
author: Dugowitch
created: 2026-03-16T18:29:42Z
updated: 2026-03-26T16:10:25Z
closed: 2026-03-26T16:10:20Z
merged: 2026-03-26T16:10:20Z
base_branch: master
head_branch: hms9780
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1546
---

# Pull Request #1546: HMS-9780: use new severity icons pattern

**Author**: @Dugowitch
**Created**: March 16, 2026 at 06:29 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `hms9780`

## Description

# Description
Replaces the advisory severity icon with a set of new icons. Simplifies the code too.

Associated Jira ticket: #HMS-9780

# How to test the PR
tests pass, only icons are changed in the UI

# Before the change
<img width="1578" height="731" alt="image" src="https://github.com/user-attachments/assets/8f25f696-1632-48f0-a2da-7b35a7bda7e4" />
<img width="1594" height="668" alt="image" src="https://github.com/user-attachments/assets/990f2e97-7f1c-46aa-8aa2-4588fa6ffba7" />
<img width="1822" height="174" alt="image" src="https://github.com/user-attachments/assets/fd031d51-dfd9-49df-9d39-1aedcea06976" />

# After the change
<img width="1574" height="725" alt="image" src="https://github.com/user-attachments/assets/ec479296-2c96-4089-8343-5c6923cdf0e8" />
<img width="1584" height="668" alt="image" src="https://github.com/user-attachments/assets/18603f2b-a3a8-4470-b90f-26dee32ee4f0" />
<img width="1822" height="174" alt="image" src="https://github.com/user-attachments/assets/5cd3bdd8-ab76-4a2b-a7a2-ee858ef89b2a" />

# Checklist:
- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on March 16, 2026 at 06:51 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1546?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `81.03448%` with `11 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 71.94%. Comparing base ([`0454fd6`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/0454fd61846c7c27bd47a0840b90a6dadf02294b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`79f3a47`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/79f3a47298a58a0d62462e463ac8eb2bbecd3378?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1546?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...nalComponents/AdvisorySeverity/AdvisorySeverity.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1546?src=pr&el=tree&filepath=src%2FPresentationalComponents%2FAdvisorySeverity%2FAdvisorySeverity.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeVNldmVyaXR5L0Fkdmlzb3J5U2V2ZXJpdHkuanM=) | 66.66% | [6 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1546?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1546?src=pr&el=tree&filepath=src%2FUtilities%2FDataMappers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | 63.63% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1546?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...tationalComponents/Snippets/DescriptionWithLink.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1546?src=pr&el=tree&filepath=src%2FPresentationalComponents%2FSnippets%2FDescriptionWithLink.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9EZXNjcmlwdGlvbldpdGhMaW5rLmpz) | 85.71% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1546?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1546      +/-   ##
==========================================
+ Coverage   71.80%   71.94%   +0.13%     
==========================================
  Files          99       98       -1     
  Lines        2426     2445      +19     
  Branches      681      686       +5     
==========================================
+ Hits         1742     1759      +17     
- Misses        601      603       +2     
  Partials       83       83              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1546?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @Dugowitch on March 20, 2026 at 11:37 AM UTC

Thank you for the review! Please have a look at my response and let me know which one would you prefer

### Comment by @Dugowitch on March 26, 2026 at 03:52 PM UTC

Needed to rebase again, because a PR got merged

---

## Reviews

### Review by @katarinazaprazna - Commented on March 18, 2026 at 10:53 PM UTC

### Review by @katarinazaprazna - Commented on March 18, 2026 at 10:53 PM UTC

### Review by @katarinazaprazna - Commented on March 18, 2026 at 11:15 PM UTC

### Review by @Dugowitch - Commented on March 19, 2026 at 12:22 PM UTC

### Review by @Dugowitch - Commented on March 19, 2026 at 12:22 PM UTC

### Review by @katarinazaprazna - Commented on March 19, 2026 at 12:47 PM UTC

### Review by @katarinazaprazna - Dismissed on March 19, 2026 at 12:55 PM UTC

Thanks for the PR! Great work here. I just have one small note on the `AdvisorySeverity` component composition for you to check out

### Review by @Dugowitch - Commented on March 19, 2026 at 01:33 PM UTC

### Review by @katarinazaprazna - Commented on March 20, 2026 at 05:15 PM UTC

### Review by @katarinazaprazna - Commented on March 20, 2026 at 05:20 PM UTC

### Review by @katarinazaprazna - Approved on March 26, 2026 at 02:24 PM UTC

Thank you!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1546*
