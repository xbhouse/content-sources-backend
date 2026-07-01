---
type: pull_request
number: 787
title: "Integrate patch set"
state: merged
author: mkholjuraev
created: 2022-04-21T19:42:11Z
updated: 2022-04-26T08:29:15Z
closed: 2022-04-26T07:58:36Z
merged: 2022-04-26T07:58:36Z
base_branch: master
head_branch: integrate-patch-set
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/787
---

# Pull Request #787: Integrate patch set

**Author**: @mkholjuraev
**Created**: April 21, 2022 at 07:42 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `integrate-patch-set`

## Description

This PR will add:

-  patch set info on Inventory detail header.
-  status and patch status filters into Review systems table in the patch set wizard.
-  default filters into Review systems table  in the patch set wizard.
-  patch set column into Review systems table  in the patch set wizard.
- remove feature flags for the Patch set table.

---

## Discussion

### Comment by @codecov-commenter on April 22, 2022 at 10:54 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#787](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (79abf33) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/cfe1af5b793485acc3fccc99233840d69ed42fde?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (cfe1af5) will **increase** coverage by `0.11%`.
> The diff coverage is `55.55%`.

> :exclamation: Current head 79abf33 differs from pull request most recent head c8dd097. Consider uploading reports for the commit c8dd097 to get more accurate results

```diff
@@            Coverage Diff             @@
##           master     #787      +/-   ##
==========================================
+ Coverage   71.53%   71.65%   +0.11%     
==========================================
  Files         101      101              
  Lines        2396     2392       -4     
  Branches      620      617       -3     
==========================================
  Hits         1714     1714              
+ Misses        615      611       -4     
  Partials       67       67              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `31.42% <0.00%> (+1.69%)` | :arrow_up: |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `53.33% <ø> (ø)` | |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `4.54% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `69.44% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `76.47% <ø> (+4.24%)` | :arrow_up: |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `68.75% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `75.51% <ø> (+0.17%)` | :arrow_up: |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `85.18% <80.00%> (ø)` | |
| [src/store/Reducers/SystemDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbURldGFpbFN0b3JlLmpz) | `100.00% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [cfe1af5...c8dd097](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/787?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on April 25, 2022 at 04:57 PM UTC

@gabipodolnikova one more review would be great, I have just applied your suggestion, no new changes further.

### Comment by @mkholjuraev on April 26, 2022 at 07:58 AM UTC

@gabipodolnikova Thanks for the review!

I will definately discuss it. Can you please describe sorting issue? is it Patch set column sort?

### Comment by @jiridostal on April 26, 2022 at 08:29 AM UTC

:tada: This PR is included in version 1.47.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.47.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.47.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gabipodolnikova - Commented on April 25, 2022 at 07:56 AM UTC

Found just one tiny thing and these are just a few mostly UX thoughts:
In the Inventory detail header, when there is no patch set name or it is undefined, you see a blank space or undefined - I would consider some label (or something) to have it clearly stated that it is not a mistake that there is a blank space.
The same applies to the System table in the Create patch set wizard. There is a blank column.
Other thing in this table is that I would consider some default sorting, because when you select a system and click Next and then go Back you still have a selected system, but because there is no sorting, the selected system is in a different spot and I was confused.

### Review by @mkholjuraev - Commented on April 25, 2022 at 04:51 PM UTC

### Review by @gabipodolnikova - Approved on April 26, 2022 at 07:55 AM UTC

Looking good!
Thanks for explaining the missing name, I think it would be great to talk about it with UX, also mention the sorting thing in the table (in the wizard) please, it bothered me a little bit :D 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/787*
