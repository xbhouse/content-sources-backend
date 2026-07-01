---
type: pull_request
number: 1019
title: "chore(RHIF-225): Update /beta to /preview"
state: merged
author: johnsonm325
created: 2023-04-05T19:56:19Z
updated: 2023-04-20T20:20:34Z
closed: 2023-04-20T20:20:34Z
merged: 2023-04-20T20:20:34Z
base_branch: master
head_branch: rhif-225
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1019
---

# Pull Request #1019: chore(RHIF-225): Update /beta to /preview

**Author**: @johnsonm325
**Created**: April 05, 2023 at 07:56 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `rhif-225`

## Description

# Description

Associated Jira ticket: # (issue)

Please include a summary of the change, what this fixes/creates/improves.


# How to test the PR

To test, you can run via /beta or /preview. We previously had issues with the local app being served. To ensure that it is properly being served, run the app via beta and make a change in the code or add a console log and ensure that you're getting local changes.

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

### Comment by @codecov-commenter on April 06, 2023 at 07:09 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1019?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`-0.21`** :warning:
> Comparison is base [(`3c97b53`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/3c97b53b1054dbfd4b06a3621d375de6d7818f73?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.32% compared to head [(`5946adf`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1019?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.12%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1019      +/-   ##
==========================================
- Coverage   64.32%   64.12%   -0.21%     
==========================================
  Files         116      116              
  Lines        2820     2832      +12     
  Branches      724      733       +9     
==========================================
+ Hits         1814     1816       +2     
- Misses       1006     1016      +10     
```


[see 6 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1019/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1019?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @johnsonm325 on April 13, 2023 at 05:54 PM UTC

@mkholjuraev this will need another test now that it is being properly served.

---

## Reviews

### Review by @mkholjuraev - Commented on April 06, 2023 at 05:33 PM UTC

Codewise, LGTM! It seems tests will pass with updated snapshots. 

### Review by @mkholjuraev - Approved on April 19, 2023 at 04:10 PM UTC

LGTM! I went through all links&routes and did not face issues.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1019*
