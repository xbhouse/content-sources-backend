---
type: pull_request
number: 1065
title: "chore(deps): Update frontend-components"
state: closed
author: AsToNlele
created: 2023-05-25T15:10:30Z
updated: 2023-07-11T12:09:05Z
closed: 2023-07-11T12:09:05Z
base_branch: master
head_branch: fc-3.9.35
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1065
---

# Pull Request #1065: chore(deps): Update frontend-components

**Author**: @AsToNlele
**Created**: May 25, 2023 at 03:10 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `fc-3.9.35`

## Description

# Description

Associated Jira ticket: [SPM-2051](https://issues.redhat.com/browse/SPM-2051)

Updated `Frontend Components` to fix an UI Issue.
With the package update came deprecation of the `<Main>` component so I had to replace it with `section` with proper styling.

# How to test the PR
1. Make sure all filters in all pages have correct styling and don't look like this (mainly the `Tags` and `Operation System` filters):
![image](https://github.com/RedHatInsights/patchman-ui/assets/20592948/678b0ede-b571-4ef5-8200-974c8cb4b3d1)

2. Verify that tables etc. have correct padding in all pages (red)
![image](https://github.com/RedHatInsights/patchman-ui/assets/20592948/98c9c62e-a762-433c-80cc-21f1f644af57)


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

### Comment by @codecov-commenter on May 25, 2023 at 03:20 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch and project coverage have no change.
> Comparison is base [(`cbc9156`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/cbc91561bac895d537de4691de439b440a62aa88?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.80% compared to head [(`0a3b45d`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.80%.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1065   +/-   ##
=======================================
  Coverage   62.80%   62.80%           
=======================================
  Files         119      119           
  Lines        2960     2960           
  Branches      758      758           
=======================================
  Hits         1859     1859           
  Misses       1101     1101           
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...Components/StatusReports/AdvisoriesStatusReport.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL0Fkdmlzb3JpZXNTdGF0dXNSZXBvcnQuanM=) | `53.33% <0.00%> (ø)` | |
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | `90.00% <ø> (ø)` | |
| [src/SmartComponents/Advisories/Advisories.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `100.00% <ø> (ø)` | |
| [...c/SmartComponents/AdvisoryDetail/AdvisoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeURldGFpbC9BZHZpc29yeURldGFpbC5qcw==) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/PackageDetail/PackageDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlRGV0YWlsL1BhY2thZ2VEZXRhaWwuanM=) | `91.30% <ø> (ø)` | |
| [src/SmartComponents/Packages/Packages.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlcy9QYWNrYWdlcy5qcw==) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSet.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <ø> (ø)` | |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <ø> (ø)` | |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `80.30% <ø> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1065?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1065*
