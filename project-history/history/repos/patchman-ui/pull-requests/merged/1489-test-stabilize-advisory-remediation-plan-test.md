---
type: pull_request
number: 1489
title: "test: stabilize advisory remediation plan test"
state: merged
author: xbhouse
created: 2026-02-05T22:56:37Z
updated: 2026-02-12T16:36:31Z
closed: 2026-02-12T16:36:31Z
merged: 2026-02-12T16:36:31Z
base_branch: master
head_branch: fix-advisory-tests
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1489
---

# Pull Request #1489: test: stabilize advisory remediation plan test

**Author**: @xbhouse
**Created**: February 05, 2026 at 10:56 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-advisory-tests`

## Description

# Description

- Updates the advisory remediation plan test to 
  - only check for the relevant system in the remediation plan 
  - use a separate user to avoid conflicts with other systems
- Refactors the system fixture to allow specifying which user/token to create a system for
- Refactors the test setup/authentication to include [these enhancements](https://github.com/content-services/content-sources-frontend/pull/829)

# How to test the PR

All tests pass

# Before the change


# After the change


# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 05, 2026 at 10:59 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1489?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`87ec4bd`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/87ec4bdb48b1c3c0e5e99802e10cc4341592d4c7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e922f97`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/e922f97cf7876ab2f3fecdb5b9ac0e5ff25a89e4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 5 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1489   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      899      892    -7     
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1489?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @xbhouse on February 06, 2026 at 09:10 PM UTC

i **think** the advisory remediation plan test is failing because the set of systems the plan applies to isn't always stable :thinking: thinking through this a bit more

### Comment by @dominikvagner on February 09, 2026 at 10:42 AM UTC

> i **think** the advisory remediation plan test is failing because the set of systems the plan applies to isn't always stable 🤔 thinking through this a bit more

yeah, it doesn't seem to be stable enough to rely on 💭 
- the amount of systems can change, if other systems with the same advisory exist from other concurrent tests, and the test now expect only its system there, which is an incorrect assumption 🙅🏼 
- the remediations UI seems to show more systems than there are in the table 🤔 the table might respect stale or unavailable systems and count of affected might not 😵‍💫 

we should just check that the system created in that test is there and not care about the count of systems 💭 
what do you think @xbhouse? could you make that fix please 🙏🏼😄 

### Comment by @xbhouse on February 09, 2026 at 02:03 PM UTC

> > i **think** the advisory remediation plan test is failing because the set of systems the plan applies to isn't always stable 🤔 thinking through this a bit more
> 
> yeah, it doesn't seem to be stable enough to rely on 💭
> 
>     * the amount of systems can change, if other systems with the same advisory exist from other concurrent tests, and the test now expect only its system there, which is an incorrect assumption 🙅🏼
> 
>     * the remediations UI seems to show more systems than there are in the table 🤔 the table might respect stale or unavailable systems and count of affected might not 😵‍💫
> 
> 
> we should just check that the system created in that test is there and not care about the count of systems 💭 what do you think @xbhouse? could you make that fix please 🙏🏼😄

yep that makes sense, i'll fix that. there's another issue too where the plan creation fails completely if systems associated with that advisory aren't available in inventory (i guess because they're removed after another test finishes):

```
{
  "errors": [
    {
      "id": "4aeda25b790e4a5d8bca4ad7f7dbef1a",
      "status": 400,
      "code": "UNKNOWN_SYSTEM",
      "title": "One or more requested systems do not exist in Inventory"
    }
  ]
}
```

i'm thinking we could solve this by either removing all systems only at the end of the test suite, or by using a separate user for the advisories tests. i'm leaning towards using a separate user, a global teardown might be a bit more complicated. what do you think @dominikvagner ?

### Comment by @dominikvagner on February 09, 2026 at 03:02 PM UTC

> yep that makes sense, i'll fix that. there's another issue too where the plan creation fails completely if systems associated with that advisory aren't available in inventory (i guess because they're removed after another test finishes):
> 
> ```
> {
>   "errors": [
>     {
>       "id": "4aeda25b790e4a5d8bca4ad7f7dbef1a",
>       "status": 400,
>       "code": "UNKNOWN_SYSTEM",
>       "title": "One or more requested systems do not exist in Inventory"
>     }
>   ]
> }
> ```
> 
> i'm thinking we could solve this by either removing all systems only at the end of the test suite, or by using a separate user for the advisories tests. i'm leaning towards using a separate user, a global teardown might be a bit more complicated. what do you think @dominikvagner ?

oh, didn't see that in the logs at first 🔍
 
yeah, global teardown would mean we would have to change how that entire systems fixture works now so I wouldn't go with that :/
I would either use a separate user, or add a tag for tests we can't/don't want to run in parallel and then run them in a separate project right after all the other UI tests (tags and grep), that might be the simplest solution 💭😅 

_(before we start adding more users we would also want to tweak how we do auth here and use the improved way how we do it in content sources, I could help with that if wanted 😄)_



### Comment by @xbhouse on February 10, 2026 at 10:18 PM UTC

@dominikvagner tried initially to move the advisory test to a separate project to run after the other ui tests, but still saw intermittent 400s related to missing systems in inventory and some 403s related to missing `inventory:read` permissions :face_with_spiral_eyes: 

now i've added a separate user for the advisory tests and added in your auth/setup improvements, which seems to work? i'll run it a few times to make sure

---

## Reviews

### Review by @dominikvagner - Approved on February 12, 2026 at 09:03 AM UTC

looks great! ✨ nice job! 👏🏼 

ack! ✅ 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1489*
