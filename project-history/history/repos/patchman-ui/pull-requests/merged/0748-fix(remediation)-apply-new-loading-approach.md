---
type: pull_request
number: 748
title: "fix(Remediation): apply new loading approach"
state: merged
author: mkholjuraev
created: 2022-03-10T19:39:30Z
updated: 2022-05-17T08:50:39Z
closed: 2022-03-14T22:45:21Z
merged: 2022-03-14T22:45:21Z
base_branch: master
head_branch: update-remediations
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/748
---

# Pull Request #748: fix(Remediation): apply new loading approach

**Author**: @mkholjuraev
**Created**: March 10, 2022 at 07:39 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `update-remediations`

## Description

TICKET: https://issues.redhat.com/browse/SPM-1220.

This PR affets all remediation buttons, modals used in Patch UI. It is intended to load remediations using **AsyncComponents**.

---

## Discussion

### Comment by @codecov-commenter on March 10, 2022 at 07:45 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#748](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (83764dd) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/54839daeaa9815294b486ab2977bf45e3fb597ba?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (54839da) will **increase** coverage by `4.84%`.
> The diff coverage is `56.81%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #748      +/-   ##
==========================================
+ Coverage   66.47%   71.31%   +4.84%     
==========================================
  Files          99      100       +1     
  Lines        2246     2214      -32     
  Branches      562      555       -7     
==========================================
+ Hits         1493     1579      +86     
+ Misses        690      579     -111     
+ Partials       63       56       -7     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...Components/StatusReports/AdvisoriesStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL0Fkdmlzb3JpZXNTdGF0dXNSZXBvcnQuanM=) | `53.33% <ø> (+3.33%)` | :arrow_up: |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `73.07% <ø> (ø)` | |
| [...rtComponents/Remediation/AsyncRemediationButton.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9Bc3luY1JlbWVkaWF0aW9uQnV0dG9uLmpz) | `33.33% <33.33%> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `85.41% <37.50%> (-9.94%)` | :arrow_down: |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `66.66% <42.85%> (+66.66%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `77.22% <50.00%> (-0.78%)` | :arrow_down: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `54.00% <50.00%> (+0.46%)` | :arrow_up: |
| [src/Utilities/unitTestingUtilities.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91bml0VGVzdGluZ1V0aWxpdGllcy5qcw==) | `18.75% <50.00%> (-9.83%)` | :arrow_down: |
| [...c/SmartComponents/Remediation/RemediationWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbldpemFyZC5qcw==) | `66.66% <66.66%> (ø)` | |
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | `90.00% <100.00%> (+90.00%)` | :arrow_up: |
| ... and [28 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [a610d4e...83764dd](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/748?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 14, 2022 at 10:19 PM UTC

> Works for me! I did notice some minor things
> 
> When the button or the wizard are loaded they show the spinner which looks a bit odd.

The loading spinner is on purpose. Patch had before to show that system issues are being loaded from backend. We had to fetch all issues per system and it took a lot of time.

>  Kapture.2022-03-14.at.13.45.14.mp4 
> In Compliance has a disabled "Remediate" button as the fallback, so the table/toolbar don't jump as much. Maybe for the wizard a empty span or something would make sense?

Spinner did help to get rid of the jitter. Thanks @bastilian for the suggestion.



### Comment by @jiridostal on March 14, 2022 at 10:58 PM UTC

:tada: This PR is included in version 1.40.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.40.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.40.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Commented on March 14, 2022 at 12:37 PM UTC

### Review by @bastilian - Approved on March 14, 2022 at 12:50 PM UTC

Works for me! I did notice some minor things

When the button or the wizard are loaded they show the spinner which looks a bit odd.

https://user-images.githubusercontent.com/7757/158174966-315e3318-2a45-4a76-9f40-1ea301af25e6.mp4

 In Compliance has a disabled "Remediate" button as the fallback, so the table/toolbar don't jump as much. Maybe for the wizard a empty span or something would make sense?

### Review by @mkholjuraev - Commented on March 14, 2022 at 01:15 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/748*
