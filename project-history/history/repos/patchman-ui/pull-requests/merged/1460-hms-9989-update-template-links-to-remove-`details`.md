---
type: pull_request
number: 1460
title: "HMS-9989: update template links to remove `/details` suffix"
state: merged
author: katarinazaprazna
created: 2026-01-07T23:25:47Z
updated: 2026-01-08T08:21:03Z
closed: 2026-01-08T08:21:03Z
merged: 2026-01-08T08:21:02Z
base_branch: master
head_branch: fix-template-link-in-system-detail
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1460
---

# Pull Request #1460: HMS-9989: update template links to remove `/details` suffix

**Author**: @katarinazaprazna
**Created**: January 07, 2026 at 11:25 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-template-link-in-system-detail`

## Description

# Description

Corrects the redirection logic for template links to ensure users are sent to the specific Template Detail page rather than the general Templates List page.

The previous URL structure included a `/details` segment which was recently deprecated and removed from the Content app's routing. This PR updates the remaining link logic to match the new simplified URL pattern.

Associated Jira ticket: https://issues.redhat.com/browse/HMS-9989

# Before the change
<img width="1821" height="2071" alt="Screenshot From 2026-01-07 16-32-54" src="https://github.com/user-attachments/assets/67d63557-1e15-406b-a23a-d9be1fb4ef9a" />

# How to test the PR
1. Open a System Detail page that has an associated template
2. Click the template link
3. Verify the user is redirected to `/templates/{templateId}` and the Template Detail page renders correctly

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on January 07, 2026 at 11:28 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1460?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `3 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`b98d888`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b98d888b805a79e216602241d58113564c337cf2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`1edefc6`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1edefc64d756b6cf154fda414f47030a49dce0f2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1460?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1460?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsListAssets.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1460?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1460   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      892      899    +7     
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1460?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @dominikvagner - Approved on January 08, 2026 at 08:17 AM UTC

ack ✅ 🚀  thanks! 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1460*
