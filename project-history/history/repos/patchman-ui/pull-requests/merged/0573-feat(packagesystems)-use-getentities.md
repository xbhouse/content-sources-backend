---
type: pull_request
number: 573
title: "feat(PackageSystems): use getEntities"
state: merged
author: mkholjuraev
created: 2021-06-24T07:27:29Z
updated: 2026-04-03T21:17:06Z
closed: 2021-06-29T08:16:40Z
merged: 2021-06-29T08:16:40Z
base_branch: master
head_branch: getEntities-package-details
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/573
---

# Pull Request #573: feat(PackageSystems): use getEntities

**Author**: @mkholjuraev
**Created**: June 24, 2021 at 07:27 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `getEntities-package-details`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on June 24, 2021 at 07:31 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `18.75000%` with `26 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 54.41%. Comparing base ([`1850af6`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1850af60bb305669e16f87728523f744954ebf62?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`5207840`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/5207840d3da939327e4e3c55982cd1a5ee73b1ba?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 908 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?src=pr&el=tree&filepath=src%2FSmartComponents%2FPackageSystems%2FPackageSystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | 0.00% | [9 Missing and 6 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/Hooks.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?src=pr&el=tree&filepath=src%2FUtilities%2FHooks.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | 16.66% | [4 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?src=pr&el=tree&filepath=src%2FUtilities%2FHelpers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | 25.00% | [2 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/api.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?src=pr&el=tree&filepath=src%2FUtilities%2Fapi.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?src=pr&el=tree&filepath=src%2FSmartComponents%2FAdvisorySystems%2FAdvisorySystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master     #573      +/-   ##
==========================================
+ Coverage   53.72%   54.41%   +0.68%     
==========================================
  Files          77       76       -1     
  Lines        1798     1757      -41     
  Branches      380      370      -10     
==========================================
- Hits          966      956      -10     
+ Misses        761      727      -34     
- Partials       71       74       +3     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/573?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @jiridostal on June 29, 2021 at 08:25 AM UTC

:tada: This PR is included in version 1.22.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.22.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.22.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/573*
