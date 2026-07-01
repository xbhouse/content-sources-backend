---
type: pull_request
number: 1153
title: "fix(SystemsTable): RHINENG-2396 - Prevent state & URL updates when not mounted"
state: merged
author: bastilian
created: 2023-12-12T21:14:59Z
updated: 2026-03-31T16:18:16Z
closed: 2023-12-13T10:53:41Z
merged: 2023-12-13T10:53:41Z
base_branch: master
head_branch: RHINENG-2396
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1153
---

# Pull Request #1153: fix(SystemsTable): RHINENG-2396 - Prevent state & URL updates when not mounted

**Author**: @bastilian
**Created**: December 12, 2023 at 09:14 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `RHINENG-2396`

## Description


# Description

Associated Jira ticket: #RHINENG-2396 

This prevents patch from updating its state and the URL parameters if the component/hook got unmounted while it was fetching systems from the API



# How to test the PR

1) Navigate to the Patch -> Systems page
2) As soon as the page loads and the InventoryTable appears in the loading state navigate away to a different page.


# Before the change

Without preventing the state/URL updates Patch will reappear again once the fetching of systems completed.

# After the change

The patch systems page does not reappear.



# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on December 12, 2023 at 09:25 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1153?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `50.00000%` with `5 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 63.21%. Comparing base ([`b6b9dd0`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b6b9dd053306cedaf49556ae0ae2b669ea66f0cb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`ce2e75c`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ce2e75c6197dbd7b3166e8ba6e1e0afeb2c44d68?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 303 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1153?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/Hooks.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1153?src=pr&el=tree&filepath=src%2FUtilities%2FHooks.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | 50.00% | [5 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1153?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1153      +/-   ##
==========================================
+ Coverage   63.17%   63.21%   +0.03%     
==========================================
  Files         122      122              
  Lines        3101     3107       +6     
  Branches      808      809       +1     
==========================================
+ Hits         1959     1964       +5     
- Misses       1142     1143       +1     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1153?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @bastilian on December 13, 2023 at 10:53 AM UTC

@mkholjuraev Using an AbortController and cancelling the request would have rejected the promise the inventory is waiting for and caused an error to be thrown by this: https://github.com/RedHatInsights/insights-inventory-frontend/blob/master/src/store/inventory-actions.js#L104-L106

### Comment by @mkholjuraev on December 13, 2023 at 11:15 AM UTC

:tada: This PR is included in version 1.65.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.65.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on December 13, 2023 at 10:30 AM UTC

LGTM! Works as expected.  @bastilian thank you!

For my understanding, what lead us to chose refs over AbortController?

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1153*
