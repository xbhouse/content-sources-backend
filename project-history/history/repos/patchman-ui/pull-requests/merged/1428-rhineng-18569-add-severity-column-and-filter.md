---
type: pull_request
number: 1428
title: "RHINENG-18569: add severity column and filter"
state: merged
author: katarinazaprazna
created: 2025-11-03T13:44:09Z
updated: 2025-11-04T14:13:55Z
closed: 2025-11-04T14:13:55Z
merged: 2025-11-04T14:13:55Z
base_branch: master
head_branch: add-severity-column-to-advisories-table
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1428
---

# Pull Request #1428: RHINENG-18569: add severity column and filter

**Author**: @katarinazaprazna
**Created**: November 03, 2025 at 01:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `add-severity-column-to-advisories-table`

## Description

# Description

This PR implements a severity column and associated filter to advisory views.

* Tables Updated: The `Advisories` and `SystemAdvisories` tables now include the new column and filter.
* Advisor App: No changes were required in the separate Advisor repository, as the `SystemAdvisories` table is automatically proxied there.
* New Component: A new `<AdvisorySeverity>` component was created to standardize the display of the severity icon and label across the UI.

We're now displaying an icon for the `severity=null` state to signify that "None" is a valid state as opposed to omitting the icon, which could be mistaken for a bug. Showing the icon maintains consistency since all other severity levels have one. Let me know your thoughts.

<img width="2239" height="1399" alt="Screenshot From 2025-11-03 14-02-55" src="https://github.com/user-attachments/assets/958c4ede-fc7f-44ff-bc5f-8b64c1261a3e" />

<img width="2239" height="1399" alt="Screenshot From 2025-11-03 22-01-59" src="https://github.com/user-attachments/assets/37e60d77-0909-4fe7-b7eb-af6c2a0c498e" />

**Heads up:** Just wanted to flag that not all files in Patch were following the ESLint rules, which might make the code review feel cumbersome. Let me know if you'd like me to undo the inconsistent formatting to streamline the process for you.

# How to test the PR

- Verify that the app builds without any errors.
- Ensure that all automated tests pass.

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [x] README.md is updated if necessary
- [x] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on November 03, 2025 at 01:46 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1428?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `71.60494%` with `69 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.93%. Comparing base ([`486146d`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/486146d5d2faf740a23a8be78d95be068bbc1169?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`952a283`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/952a283a75249e48adc6e5ef6f15491660561b49?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1428?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1428?src=pr&el=tree&filepath=src%2FUtilities%2FHelpers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | 72.22% | [38 Missing and 7 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1428?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1428?src=pr&el=tree&filepath=src%2FUtilities%2FDataMappers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | 53.06% | [22 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1428?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...ationalComponents/AdvisoryHeader/AdvisoryHeader.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1428?src=pr&el=tree&filepath=src%2FPresentationalComponents%2FAdvisoryHeader%2FAdvisoryHeader.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeUhlYWRlci9BZHZpc29yeUhlYWRlci5qcw==) | 66.66% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1428?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1428      +/-   ##
==========================================
+ Coverage   62.66%   62.93%   +0.26%     
==========================================
  Files         126      128       +2     
  Lines        3305     3399      +94     
  Branches      863      906      +43     
==========================================
+ Hits         2071     2139      +68     
- Misses       1113     1141      +28     
+ Partials      121      119       -2     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1428?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @katarinazaprazna on November 03, 2025 at 08:42 PM UTC

https://issues.redhat.com/browse/RHINENG-18569

### Comment by @katarinazaprazna on November 03, 2025 at 08:50 PM UTC

I found that when users try to filter advisories by including both a null (no severity) and specific integer severities e.g., [2, null], they receive a confusing error.

<img width="2239" height="1399" alt="Screenshot From 2025-11-03 18-58-10" src="https://github.com/user-attachments/assets/0d9b34d4-4a99-4d26-b7d6-ea6d7cd7a8b9" />

This happens because the backend ignores `IN (...)` lists that contain `NULL`. Since users might encounter this regularly, I plan to create a new PR to introduce a more user-friendly error message.

### Comment by @katarinazaprazna on November 04, 2025 at 01:01 PM UTC

Hey @LightOfHeaven1994, thanks for the feedback! That's a great point. I'll look into that

---

## Reviews

### Review by @LightOfHeaven1994 - Approved on November 04, 2025 at 11:43 AM UTC

@katarinazaprazna Looks good to me, thanks for adding that! Maybe one suggestion is when we do table export I see that severity is displayed there as 1,2,3 or 4. It might be better to show the same thing as we do on UI - Low, Medium, etc. WDYT?
Other than that looks good

### Review by @TenSt - Approved on November 04, 2025 at 12:19 PM UTC

lgtm, great job!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1428*
