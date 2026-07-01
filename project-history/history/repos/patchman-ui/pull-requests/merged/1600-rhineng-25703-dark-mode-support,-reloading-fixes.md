---
type: pull_request
number: 1600
title: "RHINENG-25703: dark mode support, reloading fixes"
state: merged
author: dominikvagner
created: 2026-05-05T12:02:56Z
updated: 2026-05-13T13:47:39Z
closed: 2026-05-13T13:47:39Z
merged: 2026-05-13T13:47:39Z
base_branch: master
head_branch: fix-dark-mode
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1600
---

# Pull Request #1600: RHINENG-25703: dark mode support, reloading fixes

**Author**: @dominikvagner
**Created**: May 05, 2026 at 12:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-dark-mode`

## Description

## Summary
`Dark Mode Support`: 
This fixes the look of the system detail page and tabs when dark mode is
used.

_Also fixes broken spacing view and removes custom CSS in favor of
PatternFly styling utilities._

`Flickering when revisiting pages`:
This fixes the flickering and unnecessary reloads on repeated navigation
to the Advisories and Packages pages, by using layout effects to force
loading before paint.


Associated Jira ticket: [#RHINENG-25703](https://redhat.atlassian.net/browse/RHINENG-25703)

## How to test the PR
Check that dark mode looks good. _(ignore tag icon in inventory tables, that will be fixed by inventory)_
Verify that Advisories and Packages table don't flicker old data when revisiting them.


---

## Discussion

### Comment by @dominikvagner on May 12, 2026 at 02:00 PM UTC

> this is great, thank you! ✨ the dark mode fixes look good, just a couple icons that need updating on the inventory table 🕶️ and some small linting issues
> 
> i left a couple comments on the filter fixes. if these end up being troublesome, maybe opening a separate PR/ticket for those would be a good idea?

yeah, that's my bad on throwing this all together, I didn't expect the filter changes to be so complicated 😞 
will split that into a different PR as I have discovered few other issues 👍🏼

### Comment by @codecov-commenter on May 13, 2026 at 12:15 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1600?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 76.21%. Comparing base ([`d9dc715`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d9dc715c97b72944640feb0cacc6ca4071226526?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`05ff0a3`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/05ff0a3273d96f8f99ed05a3c6784f50150c09d6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1600      +/-   ##
==========================================
- Coverage   76.34%   76.21%   -0.13%     
==========================================
  Files         103      103              
  Lines        3187     3187              
  Branches      693      693              
==========================================
- Hits         2433     2429       -4     
- Misses        675      678       +3     
- Partials       79       80       +1     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1600?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @xbhouse - Commented on May 07, 2026 at 09:28 PM UTC

this is great, thank you! ✨ the dark mode fixes look good, just a couple icons that need updating on the inventory table 🕶️ and some small linting issues

i left a couple comments on the filter fixes. if these end up being troublesome, maybe opening a separate PR/ticket for those would be a good idea?

### Review by @dominikvagner - Commented on May 12, 2026 at 01:28 PM UTC

### Review by @dominikvagner - Commented on May 12, 2026 at 01:57 PM UTC

### Review by @xbhouse - Approved on May 13, 2026 at 12:32 PM UTC

thank you! lgtm :D 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1600*
