---
type: pull_request
number: 1056
title: "chore: enable warning duplicate shared packages"
state: merged
author: mkholjuraev
created: 2023-05-18T14:56:04Z
updated: 2023-06-05T21:41:56Z
closed: 2023-05-19T10:04:52Z
merged: 2023-05-19T10:04:52Z
base_branch: master
head_branch: align-packages
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1056
---

# Pull Request #1056: chore: enable warning duplicate shared packages

**Author**: @mkholjuraev
**Created**: May 18, 2023 at 02:56 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `align-packages`

## Description

# Description

Associated Jira ticket: # (issue)

This enables debugging shared scope and to get info on what app is loaded multiple times. The information is displayed in the warnings

# How to test the PR

Please include steps to test your PR.

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

### Comment by @codecov-commenter on May 18, 2023 at 03:06 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1056?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`-0.05`** :warning:
> Comparison is base [(`4ba61fc`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4ba61fc53b504993da9c699eccd5add27d66379b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.08% compared to head [(`a260f52`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1056?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.03%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1056      +/-   ##
==========================================
- Coverage   63.08%   63.03%   -0.05%     
==========================================
  Files         119      119              
  Lines        2947     2949       +2     
  Branches      753      754       +1     
==========================================
  Hits         1859     1859              
- Misses       1088     1090       +2     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1056?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1056?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1056?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on June 05, 2023 at 09:41 PM UTC

:tada: This PR is included in version 1.63.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on May 18, 2023 at 08:34 PM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1056*
