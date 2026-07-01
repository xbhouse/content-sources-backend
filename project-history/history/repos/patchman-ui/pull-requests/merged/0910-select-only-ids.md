---
type: pull_request
number: 910
title: "Select only ids"
state: merged
author: mkholjuraev
created: 2022-11-09T13:11:18Z
updated: 2024-04-03T09:21:55Z
closed: 2022-11-18T21:04:26Z
merged: 2022-11-18T21:04:26Z
base_branch: master
head_branch: select-only-ids
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/910
---

# Pull Request #910: Select only ids

**Author**: @mkholjuraev
**Created**: November 09, 2022 at 01:11 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `select-only-ids`

## Description

Resolves: [SPM-1783](https://issues.redhat.com/browse/SPM-1783).

Generally, this PR introduces using new endpoints to select only ids instead of full data. This has optimized the time to fetch those ids. Apart from that, the PR has some refactoring job that is essentially moving useOnSelect hook into a separate file, improving readability and structure.  

---

## Discussion

### Comment by @codecov-commenter on November 09, 2022 at 03:04 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.04**% // Head: **72.15**% // Increases project coverage by **`+0.10%`** :tada:
> Coverage data is based on head [(`55e87df`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`531979b`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/531979b8f7363aea1d9194fcf09dc714ead0e6a8?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 93.58% of modified lines in pull request are covered.

> :exclamation: Current head 55e87df differs from pull request most recent head 99a468e. Consider uploading reports for the commit 99a468e to get more accurate results

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #910      +/-   ##
==========================================
+ Coverage   72.04%   72.15%   +0.10%     
==========================================
  Files         110      111       +1     
  Lines        2572     2571       -1     
  Branches      655      658       +3     
==========================================
+ Hits         1853     1855       +2     
+ Misses        657      652       -5     
- Partials       62       64       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `4.65% <0.00%> (+0.20%)` | :arrow_up: |
| [src/SmartComponents/Systems/SystemsHelpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNIZWxwZXJzLmpz) | `100.00% <ø> (ø)` | |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `73.04% <ø> (-5.96%)` | :arrow_down: |
| [src/Utilities/RawDataForTesting.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9SYXdEYXRhRm9yVGVzdGluZy5qcw==) | `100.00% <ø> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `56.77% <50.00%> (+2.46%)` | :arrow_up: |
| [src/Utilities/useOnSelect.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91c2VPblNlbGVjdC5qcw==) | `98.48% <98.48%> (ø)` | |
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `100.00% <100.00%> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `82.00% <100.00%> (-1.34%)` | :arrow_down: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `78.00% <100.00%> (-1.63%)` | :arrow_down: |
| ... and [6 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/910?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on November 17, 2022 at 10:45 AM UTC

@adonispuente thank you for the review. Good catch, fixed.

### Comment by @jiridostal on November 18, 2022 at 09:22 PM UTC

:tada: This PR is included in version 1.57.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.57.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.57.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Changes Requested on November 15, 2022 at 04:31 PM UTC

Currently if you add filters to the table and begin to select all, it ignores the filters and still selects all available systems. The useOnSelect is ignoring filters
Ex. If I have 1000 total systems and filter down to 200, when I 'select all', it will select all 1000

### Review by @adonispuente - Approved on November 18, 2022 at 04:02 PM UTC

PR works as intended, LGTM! Just gotta fix the merge conflict

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/910*
