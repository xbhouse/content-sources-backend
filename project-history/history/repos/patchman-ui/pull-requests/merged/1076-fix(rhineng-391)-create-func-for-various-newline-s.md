---
type: pull_request
number: 1076
title: "fix(RHINENG-391): Create func for various newline scenarios"
state: merged
author: johnsonm325
created: 2023-06-20T21:04:45Z
updated: 2023-07-03T16:56:48Z
closed: 2023-07-03T16:34:53Z
merged: 2023-07-03T16:34:53Z
base_branch: master
head_branch: rhineng-391
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1076
---

# Pull Request #1076: fix(RHINENG-391): Create func for various newline scenarios

**Author**: @johnsonm325
**Created**: June 20, 2023 at 09:04 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `rhineng-391`

## Description

To see what's broken check the 2 following advisories: RHBA-2023:2052
RHSA-2023:3246

We have 2 different ways that descriptions come in for Advisories. The most common way is to separate every line with `\n\n`. This would create double space between each line. To remedy this, we have a `preserveNewlines` function in `Helpers.js` that removes the extra `\n`.

But, the new variation of description correctly provides `\n\n` for paragraph separations and only `\n` for bullet separations. But with the current `preserveNewlines` functionality, it removes the `\n` for the bullets, making them appear as a single line (see RHBA-2023:2052).

This PR adds an additional function that takes the substring of the description that includes the bullets and replaces any `\n\n` with a single `\n`, ignoring the final `\n\n` because it is for the paragraph that comes after the bullets.

I expect this to fix most, if not all, of our current descriptions. If a new layout for descriptions comes through at some point, it could be broken. The best fix is to standardize the descriptions for all advisories.

# Description

Associated Jira ticket: # (issue)

Please include a summary of the change, what this fixes/creates/improves.


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

### Comment by @codecov-commenter on June 20, 2023 at 09:13 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1076?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00`**% and project coverage change: **`+0.07`** :tada:
> Comparison is base [(`24222b6`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/24222b6c7aff689b6ddab7ff23aceb6228056665?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.80% compared to head [(`0a2a0fb`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1076?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.87%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1076      +/-   ##
==========================================
+ Coverage   62.80%   62.87%   +0.07%     
==========================================
  Files         119      119              
  Lines        2960     2971      +11     
  Branches      758      763       +5     
==========================================
+ Hits         1859     1868       +9     
- Misses       1101     1103       +2     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1076?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...ationalComponents/AdvisoryHeader/AdvisoryHeader.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1076?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeUhlYWRlci9BZHZpc29yeUhlYWRlci5qcw==) | `83.33% <100.00%> (+0.72%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1076?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `80.45% <100.00%> (-0.41%)` | :arrow_down: |

... and [3 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1076/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1076?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on July 03, 2023 at 04:56 PM UTC

:tada: This PR is included in version 1.63.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on June 20, 2023 at 11:01 PM UTC

small nitpick, tested with multiple advisories, looks good on my end. LGTM

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1076*
