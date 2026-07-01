---
type: pull_request
number: 949
title: "chore(RHIF-160): decouple advisory detail header and table"
state: merged
author: mkholjuraev
created: 2023-01-20T14:17:43Z
updated: 2024-04-03T09:22:48Z
closed: 2023-01-24T20:31:12Z
merged: 2023-01-24T20:31:12Z
base_branch: master
head_branch: decouple-adv-header-table
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/949
---

# Pull Request #949: chore(RHIF-160): decouple advisory detail header and table

**Author**: @mkholjuraev
**Created**: January 20, 2023 at 02:17 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `decouple-adv-header-table`

## Description

# Description

Associated Jira ticket: # [RHIF-160](https://issues.redhat.com/browse/RHIF-160)

The affected systems table has to wait for the header. There is no need to have this approach as error handling is done in the affected systems table as well.

# How to test the PR

This should not change functionality overall, just run the PR locally and verify that there is nothing weird in Advisory detail page.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on January 20, 2023 at 02:25 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/949?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.44**% // Head: **72.37**% // Decreases project coverage by **`-0.08%`** :warning:
> Coverage data is based on head [(`7ea6236`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/949?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`c9dac6d`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/c9dac6d3919658d9e99625d016d012a314c3535d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 33.33% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #949      +/-   ##
==========================================
- Coverage   72.44%   72.37%   -0.08%     
==========================================
  Files         110      110              
  Lines        2617     2617              
  Branches      661      658       -3     
==========================================
- Hits         1896     1894       -2     
- Misses        721      723       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/949?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/949?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/AdvisoryDetail/AdvisoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/949?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeURldGFpbC9BZHZpc29yeURldGFpbC5qcw==) | `100.00% <ø> (ø)` | |
| [...ationalComponents/AdvisoryHeader/AdvisoryHeader.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/949?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeUhlYWRlci9BZHZpc29yeUhlYWRlci5qcw==) | `82.60% <50.00%> (-3.11%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/949?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on January 25, 2023 at 12:42 PM UTC

:tada: This PR is included in version 1.58.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.58.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.58.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @Fewwy - Approved on January 24, 2023 at 04:20 PM UTC

Tested it locally as you suggested. Couldn't find any issues. LGTM

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/949*
