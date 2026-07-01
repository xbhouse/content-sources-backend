---
type: pull_request
number: 620
title: "fix(Notifications): do not fire multiple notifications"
state: merged
author: mkholjuraev
created: 2021-08-09T12:34:50Z
updated: 2021-08-12T07:49:39Z
closed: 2021-08-12T07:33:29Z
merged: 2021-08-12T07:33:29Z
base_branch: master
head_branch: do-not-trigger-multiple-notifications
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/620
---

# Pull Request #620: fix(Notifications): do not fire multiple notifications

**Author**: @mkholjuraev
**Created**: August 09, 2021 at 12:34 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `do-not-trigger-multiple-notifications`

## Description

Multiple error notifications were due to multiple requests to the backend and stored notifications in NotificationsReducer. 
1. Now, notificationsReducer is cleared across page changes to prevent notifications 
2. when applied URL params on load, useEffect with dependency **queryParams** was firing 2 times in several pages due to decodeQueryParams decoding integer parameters as string, then which is a different type from origin.
3. In detail pages, we have 2 requests by default one for fetching header data and another one for tablets. Now, I have made changes to first fetch header data and if there is an error, handle the error and do not create an instance of table component. This will result in only one error. But, not displaying the table if there is a problem with a header component may be wrong. Opinions are welcome


---

## Discussion

### Comment by @codecov-commenter on August 11, 2021 at 04:59 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#620](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c23c12e) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/abf56509c0766d90796c4f22ee9707257113f091?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (abf5650) will **increase** coverage by `0.47%`.
> The diff coverage is `58.33%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #620      +/-   ##
==========================================
+ Coverage   52.38%   52.86%   +0.47%     
==========================================
  Files          79       79              
  Lines        1909     1922      +13     
  Branches      417      421       +4     
==========================================
+ Hits         1000     1016      +16     
+ Misses        828      826       -2     
+ Partials       81       80       -1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [.../PresentationalComponents/Snippets/ErrorHandler.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FcnJvckhhbmRsZXIuanM=) | `23.80% <0.00%> (-2.51%)` | :arrow_down: |
| [src/SmartComponents/PackageDetail/PackageDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlRGV0YWlsL1BhY2thZ2VEZXRhaWwuanM=) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Packages/Packages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlcy9QYWNrYWdlcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [...rtComponents/Remediation/PatchRemediationButton.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9QYXRjaFJlbWVkaWF0aW9uQnV0dG9uLmpz) | `100.00% <ø> (+16.66%)` | :arrow_up: |
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `86.00% <66.66%> (+3.02%)` | :arrow_up: |
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `71.87% <66.66%> (+4.68%)` | :arrow_up: |
| [...c/SmartComponents/AdvisoryDetail/AdvisoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeURldGFpbC9BZHZpc29yeURldGFpbC5qcw==) | `91.66% <75.00%> (-3.34%)` | :arrow_down: |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `80.00% <75.00%> (-2.36%)` | :arrow_down: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `77.06% <100.00%> (+0.10%)` | :arrow_up: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `53.19% <0.00%> (+1.06%)` | :arrow_up: |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [db0cfbe...c23c12e](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/620?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on August 12, 2021 at 07:49 AM UTC

:tada: This PR is included in version 1.29.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.29.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.29.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/620*
