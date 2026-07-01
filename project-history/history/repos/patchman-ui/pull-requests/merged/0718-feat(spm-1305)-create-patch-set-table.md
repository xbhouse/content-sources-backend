---
type: pull_request
number: 718
title: "feat(SPM-1305): create patch set table"
state: merged
author: mkholjuraev
created: 2022-02-21T10:16:56Z
updated: 2024-04-03T09:21:34Z
closed: 2022-02-24T11:01:49Z
merged: 2022-02-24T11:01:49Z
base_branch: master
head_branch: patch-set-table
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/718
---

# Pull Request #718: feat(SPM-1305): create patch set table

**Author**: @mkholjuraev
**Created**: February 21, 2022 at 10:16 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `patch-set-table`

## Description

This enables  to view Patch set list. Buttons, actions will be integrated in the next ticket https://issues.redhat.com/browse/SPM-1380

---

## Discussion

### Comment by @codecov-commenter on February 22, 2022 at 12:20 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#718](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (bc50f05) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/0d470b192ae92f83ad18a521c9ac510e0aecfe44?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0d470b1) will **decrease** coverage by `2.35%`.
> The diff coverage is `12.50%`.

> :exclamation: Current head bc50f05 differs from pull request most recent head 1bf29e6. Consider uploading reports for the commit 1bf29e6 to get more accurate results

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #718      +/-   ##
==========================================
- Coverage   75.32%   72.97%   -2.36%     
==========================================
  Files          97       99       +2     
  Lines        2148     2231      +83     
  Branches      545      562      +17     
==========================================
+ Hits         1618     1628      +10     
- Misses        478      548      +70     
- Partials       52       55       +3     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `0.00% <0.00%> (ø)` | |
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | `10.52% <0.00%> (ø)` | |
| [src/store/Reducers/PatchSetsReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0c1JlZHVjZXIuanM=) | `29.41% <0.00%> (-33.09%)` | :arrow_down: |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `68.75% <11.11%> (-7.31%)` | :arrow_down: |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `66.66% <50.00%> (-8.34%)` | :arrow_down: |
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `31.42% <50.00%> (+1.12%)` | :arrow_up: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `56.56% <50.00%> (ø)` | |
| [...rc/PresentationalComponents/TableView/TableView.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3Lmpz) | `100.00% <100.00%> (ø)` | |
| [src/store/ActionTypes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvblR5cGVzLmpz) | `100.00% <100.00%> (ø)` | |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [0d470b1...1bf29e6](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/718?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on March 01, 2022 at 01:47 PM UTC

:tada: This PR is included in version 1.40.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.40.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.40.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Commented on February 21, 2022 at 10:50 AM UTC

### Review by @mkholjuraev - Commented on February 21, 2022 at 12:18 PM UTC

### Review by @mkholjuraev - Commented on February 21, 2022 at 12:18 PM UTC

### Review by @bastilian - Commented on February 23, 2022 at 01:05 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/718*
