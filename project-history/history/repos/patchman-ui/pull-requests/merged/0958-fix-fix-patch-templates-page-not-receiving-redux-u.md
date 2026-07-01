---
type: pull_request
number: 958
title: "fix: Fix patch templates page not receiving redux updates"
state: merged
author: leSamo
created: 2023-02-14T15:38:45Z
updated: 2026-04-01T19:26:29Z
closed: 2023-02-15T22:59:52Z
merged: 2023-02-15T22:59:52Z
base_branch: master
head_branch: fix-sets-table
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/958
---

# Pull Request #958: fix: Fix patch templates page not receiving redux updates

**Author**: @leSamo
**Created**: February 14, 2023 at 03:38 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-sets-table`

## Description

# Description

Associated Jira ticket: [SPM-1805](https://issues.redhat.com/browse/SPM-1805).

Filter, page and per page update did not work in the Patch templates table due to Redux updates not propagating. Also updated skeleton table to be compact if the loaded table is compact. Sorting does not work yet due to [SPM-1894](https://issues.redhat.com/browse/SPM-1894).

# How to test the PR

Try changing page, per page and filter in the Patch Templates page.

# Before the change
Filter did not always work. Page and per page update did not work at all.

# After the change
Filter, page and per page update should now work.

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 15, 2023 at 12:17 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/958?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `42.85714%` with `4 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 72.36%. Comparing base ([`89bb482`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/89bb482e0410ddb3417b969d4b0379b08944535d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`5f8e7d1`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/5f8e7d131a123990225e41fac5bb70339bb0007c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 492 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/958?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/store/Reducers/PatchSetsReducer.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/958?src=pr&el=tree&filepath=src%2Fstore%2FReducers%2FPatchSetsReducer.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0c1JlZHVjZXIuanM=) | 33.33% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/958?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master     #958      +/-   ##
==========================================
+ Coverage   72.35%   72.36%   +0.01%     
==========================================
  Files         110      110              
  Lines        2615     2616       +1     
  Branches      658      658              
==========================================
+ Hits         1892     1893       +1     
  Misses        723      723              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/958?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mkholjuraev on February 15, 2023 at 11:21 PM UTC

:tada: This PR is included in version 1.59.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.59.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.59.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on February 15, 2023 at 09:41 PM UTC

LGTM! Nice job. 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/958*
