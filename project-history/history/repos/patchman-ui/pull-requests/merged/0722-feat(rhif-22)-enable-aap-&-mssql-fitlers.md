---
type: pull_request
number: 722
title: "feat(RHIF-22): enable AAP & MSSQL fitlers"
state: merged
author: mkholjuraev
created: 2022-02-25T09:00:38Z
updated: 2022-05-17T08:50:42Z
closed: 2022-04-01T08:19:17Z
merged: 2022-04-01T08:19:17Z
base_branch: master
head_branch: aap-mssql-filters
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/722
---

# Pull Request #722: feat(RHIF-22): enable AAP & MSSQL fitlers

**Author**: @mkholjuraev
**Created**: February 25, 2022 at 09:00 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `aap-mssql-filters`

## Description

resolves: https://issues.redhat.com/browse/RHIF-22

---

## Discussion

### Comment by @codecov-commenter on March 09, 2022 at 11:47 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#722](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f3d143e) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/54839daeaa9815294b486ab2977bf45e3fb597ba?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (54839da) will **increase** coverage by `5.19%`.
> The diff coverage is `71.42%`.

```diff
@@            Coverage Diff             @@
##           master     #722      +/-   ##
==========================================
+ Coverage   66.47%   71.67%   +5.19%     
==========================================
  Files          99      100       +1     
  Lines        2246     2231      -15     
  Branches      562      563       +1     
==========================================
+ Hits         1493     1599     +106     
+ Misses        690      579     -111     
+ Partials       63       53      -10     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `82.33% <75.00%> (+4.33%)` | :arrow_up: |
| [...rtComponents/Remediation/PatchRemediationButton.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9QYXRjaFJlbWVkaWF0aW9uQnV0dG9uLmpz) | `66.66% <0.00%> (-33.34%)` | :arrow_down: |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `85.41% <0.00%> (-9.94%)` | :arrow_down: |
| [src/Utilities/unitTestingUtilities.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91bml0VGVzdGluZ1V0aWxpdGllcy5qcw==) | `18.75% <0.00%> (-9.83%)` | :arrow_down: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `78.84% <0.00%> (-8.26%)` | :arrow_down: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `76.96% <0.00%> (-2.78%)` | :arrow_down: |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `73.07% <0.00%> (ø)` | |
| [...rc/SmartComponents/Remediation/RemediationModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbk1vZGFsLmpz) | | |
| [...c/SmartComponents/Remediation/RemediationWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbldpemFyZC5qcw==) | `66.66% <0.00%> (ø)` | |
| ... and [23 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [54839da...f3d143e](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/722?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 23, 2022 at 02:36 PM UTC

@RedHatInsights/team-interact this PR is ready to be reviewed if you would like to have a look.

### Comment by @mkholjuraev on April 01, 2022 at 08:19 AM UTC

@johnsonm325 thanks for the review.  

Yes, you are right. This was what I wanted to say when I talked about **mapGlobalFilters** in the team-interact channel.

### Comment by @jiridostal on April 01, 2022 at 08:31 AM UTC

:tada: This PR is included in version 1.41.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.41.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.41.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Commented on March 16, 2022 at 01:05 PM UTC

### Review by @johnsonm325 - Approved on March 31, 2022 at 08:29 PM UTC

It all looks complex to me. haha. If I'm correct, you use `mapGlobalFilters` to build a data structure you can use throughout the app. Then then functions in `encodeParams` are used to build the url for hitting the api.

If I'm correct, then it seems to work properly. Nice job.

I think this is something that can be used across apps. Maybe we eventually pull this out into frontend-components.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/722*
