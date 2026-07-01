---
type: pull_request
number: 738
title: "Integrate patch set"
state: merged
author: mkholjuraev
created: 2022-03-03T22:12:56Z
updated: 2022-04-08T15:05:41Z
closed: 2022-04-08T14:48:26Z
merged: 2022-04-08T14:48:26Z
base_branch: master
head_branch: integrate-patch-set
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/738
---

# Pull Request #738: Integrate patch set

**Author**: @mkholjuraev
**Created**: March 03, 2022 at 10:12 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `integrate-patch-set`

## Description

The PR is intended to resolve: https://issues.redhat.com/browse/SPM-1380.

You can create/edit, delete a patch set in the patch set page,

---

## Discussion

### Comment by @codecov-commenter on March 03, 2022 at 10:20 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#738](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (3535f8e) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/83aa7aba6e4548e78d8a069186f0cecb4017d625?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (83aa7ab) will **decrease** coverage by `0.14%`.
> The diff coverage is `33.77%`.

```diff
@@            Coverage Diff             @@
##           master     #738      +/-   ##
==========================================
- Coverage   71.01%   70.86%   -0.15%     
==========================================
  Files         100      101       +1     
  Lines        2225     2334     +109     
  Branches      559      590      +31     
==========================================
+ Hits         1580     1654      +74     
- Misses        586      614      +28     
- Partials       59       66       +7     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `0.00% <0.00%> (ø)` | |
| [...tComponents/PatchSetWizard/steps/ReviewPatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdQYXRjaFNldC5qcw==) | `16.66% <0.00%> (-3.34%)` | :arrow_down: |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `5.26% <0.00%> (ø)` | |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `76.04% <0.00%> (-0.93%)` | :arrow_down: |
| [src/store/Reducers/PatchSetsReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0c1JlZHVjZXIuanM=) | `27.77% <0.00%> (-1.64%)` | :arrow_down: |
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | `7.69% <7.69%> (-2.84%)` | :arrow_down: |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `8.00% <8.33%> (-1.53%)` | :arrow_down: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `73.81% <18.75%> (-3.41%)` | :arrow_down: |
| [src/store/Reducers/SpecificPatchSetReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1NwZWNpZmljUGF0Y2hTZXRSZWR1Y2VyLmpz) | `26.31% <26.31%> (ø)` | |
| ... and [19 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [83aa7ab...3535f8e](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/738?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on April 08, 2022 at 03:05 PM UTC

:tada: This PR is included in version 1.42.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.42.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.42.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/738*
