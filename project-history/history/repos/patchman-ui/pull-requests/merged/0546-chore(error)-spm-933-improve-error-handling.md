---
type: pull_request
number: 546
title: "chore(Error): SPM-933 improve error handling"
state: merged
author: mkholjuraev
created: 2021-05-26T12:06:50Z
updated: 2021-08-09T06:56:57Z
closed: 2021-06-10T07:28:47Z
merged: 2021-06-10T07:28:47Z
base_branch: master
head_branch: improve-error-handling
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/546
---

# Pull Request #546: chore(Error): SPM-933 improve error handling

**Author**: @mkholjuraev
**Created**: May 26, 2021 at 12:06 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `improve-error-handling`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on June 08, 2021 at 11:11 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#546](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (fc7f3e4) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/f744d1e667d9280d15f80b9e7facb2dd6e52b5ba?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f744d1e) will **decrease** coverage by `0.36%`.
> The diff coverage is `39.47%`.

> :exclamation: Current head fc7f3e4 differs from pull request most recent head eee92d7. Consider uploading reports for the commit eee92d7 to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #546      +/-   ##
==========================================
- Coverage   54.02%   53.65%   -0.37%     
==========================================
  Files          76       77       +1     
  Lines        1777     1791      +14     
  Branches      378      385       +7     
==========================================
+ Hits          960      961       +1     
- Misses        745      756      +11     
- Partials       72       74       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `84.61% <ø> (-0.30%)` | :arrow_down: |
| [...c/SmartComponents/AdvisoryDetail/AdvisoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeURldGFpbC9BZHZpc29yeURldGFpbC5qcw==) | `95.00% <0.00%> (-5.00%)` | :arrow_down: |
| [src/SmartComponents/AdvisoryDetail/CvesModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeURldGFpbC9DdmVzTW9kYWwuanM=) | `5.00% <ø> (+0.12%)` | :arrow_up: |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/PackageDetail/PackageDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlRGV0YWlsL1BhY2thZ2VEZXRhaWwuanM=) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Packages/Packages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlcy9QYWNrYWdlcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/axiosInterceptors.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9heGlvc0ludGVyY2VwdG9ycy5qcw==) | `60.60% <0.00%> (-8.36%)` | :arrow_down: |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <ø> (ø)` | |
| ... and [6 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [f744d1e...eee92d7](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/546?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on June 15, 2021 at 08:03 AM UTC

:tada: This PR is included in version 1.20.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.20.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.20.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/546*
