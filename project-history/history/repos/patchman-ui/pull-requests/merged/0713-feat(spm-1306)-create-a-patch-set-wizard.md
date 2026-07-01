---
type: pull_request
number: 713
title: "feat(SPM-1306): create a patch set wizard"
state: merged
author: mkholjuraev
created: 2022-02-02T10:33:32Z
updated: 2024-04-03T09:21:33Z
closed: 2022-02-22T09:51:55Z
merged: 2022-02-22T09:51:55Z
base_branch: master
head_branch: assing-beseline-system
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/713
---

# Pull Request #713: feat(SPM-1306): create a patch set wizard

**Author**: @mkholjuraev
**Created**: February 02, 2022 at 10:33 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `assing-beseline-system`

## Description

We have incomplete UI mockup designs. I would like to improve the wizard once the mockups are complete.

---

## Discussion

### Comment by @codecov-commenter on February 14, 2022 at 10:42 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#713](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (db9fe4f) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/decc51c44e12992b72402ecf05690ae6cf5bdb2b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (decc51c) will **decrease** coverage by `6.54%`.
> The diff coverage is `19.91%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #713      +/-   ##
==========================================
- Coverage   81.87%   75.32%   -6.55%     
==========================================
  Files          84       97      +13     
  Lines        1925     2148     +223     
  Branches      499      545      +46     
==========================================
+ Hits         1576     1618      +42     
- Misses        306      478     +172     
- Partials       43       52       +9     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/store/Reducers/SystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbXNTdG9yZS5qcw==) | `100.00% <ø> (ø)` | |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `5.26% <5.26%> (ø)` | |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `80.89% <9.09%> (-5.41%)` | :arrow_down: |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `9.52% <9.52%> (ø)` | |
| [...Components/PatchSetWizard/steps/RequestProgress.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXF1ZXN0UHJvZ3Jlc3MuanM=) | `10.00% <10.00%> (ø)` | |
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | `10.52% <10.52%> (ø)` | |
| [...nts/PatchSetWizard/InputFields/DescriptionField.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9EZXNjcmlwdGlvbkZpZWxkLmpz) | `11.11% <11.11%> (ø)` | |
| [...Components/PatchSetWizard/InputFields/NameField.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9OYW1lRmllbGQuanM=) | `11.11% <11.11%> (ø)` | |
| [...mponents/PatchSetWizard/InputFields/ToDateField.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9Ub0RhdGVGaWVsZC5qcw==) | `11.11% <11.11%> (ø)` | |
| [...s/PatchSetWizard/InputFields/SelectExistingSets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9TZWxlY3RFeGlzdGluZ1NldHMuanM=) | `14.28% <14.28%> (ø)` | |
| ... and [13 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [decc51c...db9fe4f](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/713?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on February 21, 2022 at 02:16 PM UTC

@bastilian Thanks for the review. I will adopt your suggestions and merge the PR

### Comment by @jiridostal on February 22, 2022 at 10:08 AM UTC

:tada: This PR is included in version 1.39.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.39.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.39.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Commented on February 21, 2022 at 10:33 AM UTC

### Review by @bastilian - Commented on February 21, 2022 at 10:36 AM UTC

### Review by @bastilian - Commented on February 21, 2022 at 10:42 AM UTC

### Review by @bastilian - Commented on February 21, 2022 at 11:38 AM UTC

### Review by @bastilian - Approved on February 21, 2022 at 11:55 AM UTC

Looks good to me! And works as expected! Thank you @mkholjuraev!

### Review by @mkholjuraev - Commented on February 21, 2022 at 02:33 PM UTC

### Review by @mkholjuraev - Commented on February 21, 2022 at 02:33 PM UTC

### Review by @mkholjuraev - Commented on February 21, 2022 at 02:43 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/713*
