---
type: pull_request
number: 367
title: "feat(internationalization): translate smart components"
state: merged
author: mkholjuraev
created: 2020-12-17T09:26:28Z
updated: 2026-04-03T12:25:56Z
closed: 2020-12-18T10:44:22Z
merged: 2020-12-18T10:44:22Z
base_branch: master
head_branch: translate-smart-compponents
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/367
---

# Pull Request #367: feat(internationalization): translate smart components

**Author**: @mkholjuraev
**Created**: December 17, 2020 at 09:26 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `translate-smart-compponents`

## Description

resolves: https://issues.redhat.com/browse/SPM-606

---

## Discussion

### Comment by @codecov-io on December 17, 2020 at 10:28 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367?src=pr&el=h1) Report
> Merging [#367](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367?src=pr&el=desc) (2b3b499) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/5d86a7c62d31b8992649e7f10ba48f0c5c086f74?el=desc) (5d86a7c) will **increase** coverage by `0.33%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #367      +/-   ##
==========================================
+ Coverage   73.20%   73.53%   +0.33%     
==========================================
  Files          69       69              
  Lines        1209     1209              
  Branches      159      159              
==========================================
+ Hits          885      889       +4     
+ Misses        271      267       -4     
  Partials       53       53              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `95.23% <ø> (ø)` | |
| [...c/SmartComponents/AdvisoryDetail/AdvisoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeURldGFpbC9BZHZpc29yeURldGFpbC5qcw==) | `100.00% <ø> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/PackageDetail/PackageDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlRGV0YWlsL1BhY2thZ2VEZXRhaWwuanM=) | `0.00% <ø> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Packages/Packages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlcy9QYWNrYWdlcy5qcw==) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/SystemDetail/InventoryPage.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5UGFnZS5qcw==) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/SystemDetail/SystemDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvU3lzdGVtRGV0YWlsLmpz) | `100.00% <ø> (ø)` | |
| [...c/SmartComponents/SystemPackages/SystemPackages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1QYWNrYWdlcy9TeXN0ZW1QYWNrYWdlcy5qcw==) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `100.00% <ø> (ø)` | |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367?src=pr&el=footer). Last update [5d86a7c...2b3b499](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/367?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on December 18, 2020 at 10:53 AM UTC

:tada: This PR is included in version 1.3.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.3.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.3.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

### Comment by @codecov-commenter on April 03, 2026 at 12:25 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/367?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 73.53%. Comparing base ([`5d86a7c`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/5d86a7c62d31b8992649e7f10ba48f0c5c086f74?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`2b3b499`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/2b3b499e85fe864cb96daf41f8c31afa5fbb29d9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1120 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master     #367      +/-   ##
==========================================
+ Coverage   73.20%   73.53%   +0.33%     
==========================================
  Files          69       69              
  Lines        1209     1209              
  Branches      159      159              
==========================================
+ Hits          885      889       +4     
+ Misses        271      267       -4     
  Partials       53       53              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/367?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/367*
