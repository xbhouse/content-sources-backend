---
type: pull_request
number: 1024
title: "feat(ESSNTL-4496): Replace SmartManager w/ satellite"
state: merged
author: adonispuente
created: 2023-04-12T15:59:23Z
updated: 2023-04-17T18:41:40Z
closed: 2023-04-17T18:18:44Z
merged: 2023-04-17T18:18:44Z
base_branch: master
head_branch: ESSNTL-4496
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1024
---

# Pull Request #1024: feat(ESSNTL-4496): Replace SmartManager w/ satellite

**Author**: @adonispuente
**Created**: April 12, 2023 at 03:59 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `ESSNTL-4496`

## Description

# Description
Associated Jira ticket: # ([ESSNTL-4496)](https://issues.redhat.com/browse/ESSNTL-4496)
This replaces all references of Red Hat Smart Manager with Satellite. I've left entitlements as is for the time being as thats user schema.

# How to test the PR
Just skim the code and UI for any references to smart manager and verify there are none the user can see.

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work

---

## Discussion

### Comment by @codecov-commenter on April 12, 2023 at 04:10 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1024?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`14.28`**% and project coverage change: **`+0.06`** :tada:
> Comparison is base [(`84654b7`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/84654b7b0c635b6de03610e147959d50b1414773?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.05% compared to head [(`039a220`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1024?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.12%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1024      +/-   ##
==========================================
+ Coverage   64.05%   64.12%   +0.06%     
==========================================
  Files         116      116              
  Lines        2835     2832       -3     
  Branches      733      733              
==========================================
  Hits         1816     1816              
+ Misses       1019     1016       -3     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1024?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1024?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1024?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1024?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `72.22% <100.00%> (ø)` | |

... and [2 files with indirect coverage changes](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1024/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1024?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @adonispuente on April 17, 2023 at 02:21 PM UTC

Thanks for taking a look @leSamo , new change is up

### Comment by @mkholjuraev on April 17, 2023 at 06:40 PM UTC

:tada: This PR is included in version 1.63.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Changes Requested on April 17, 2023 at 11:39 AM UTC

Thank you, looks good, I only have one small request.

### Review by @leSamo - Approved on April 17, 2023 at 02:26 PM UTC

Perfect

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1024*
