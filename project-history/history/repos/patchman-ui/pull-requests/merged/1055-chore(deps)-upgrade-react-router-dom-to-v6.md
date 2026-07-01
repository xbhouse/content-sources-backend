---
type: pull_request
number: 1055
title: "chore(deps): upgrade react-router-dom to v6"
state: merged
author: AsToNlele
created: 2023-05-17T11:53:46Z
updated: 2023-08-07T14:35:17Z
closed: 2023-08-02T10:12:06Z
merged: 2023-08-02T10:12:05Z
base_branch: master
head_branch: SPM-1858v2
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1055
---

# Pull Request #1055: chore(deps): upgrade react-router-dom to v6

**Author**: @AsToNlele
**Created**: May 17, 2023 at 11:53 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `SPM-1858v2`

## Description

# Description

Associated Jira ticket: [SPM-1858](https://issues.redhat.com/browse/SPM-1858)

How to test:
1. Make sure all the links redirect correctly
2. Approve
3. Profit?

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on May 19, 2023 at 04:06 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`44.28`**% and project coverage change: **`-0.06`** :warning:
> Comparison is base [(`cddb13e`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/cddb13e7622aeb37928a39d6c15cb06f519c320d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.79% compared to head [(`630b22a`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.73%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1055      +/-   ##
==========================================
- Coverage   62.79%   62.73%   -0.06%     
==========================================
  Files         119      120       +1     
  Lines        2973     2979       +6     
  Branches      763      763              
==========================================
+ Hits         1867     1869       +2     
- Misses       1106     1110       +4     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | `0.00% <ø> (ø)` | |
| [src/Routes.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSet.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `66.66% <0.00%> (ø)` | |
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `56.70% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `80.45% <0.00%> (ø)` | |
| [src/Utilities/NavigateToSystem.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9OYXZpZ2F0ZVRvU3lzdGVtLmpz) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/SystemsHelpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | `19.23% <0.00%> (-2.51%)` | :arrow_down: |
| ... and [12 more](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1055?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @AsToNlele on June 27, 2023 at 11:17 AM UTC

After Frontend components have been upgraded, there is one less step :tada: 

### Comment by @bastilian on June 27, 2023 at 05:46 PM UTC

@AsToNlele Just out of precaution, we should get a second approval from @RedHatInsights/team-interact before merging.

### Comment by @AsToNlele on June 30, 2023 at 12:09 PM UTC

/retest

### Comment by @AsToNlele on June 30, 2023 at 01:22 PM UTC

/retest

### Comment by @AsToNlele on July 04, 2023 at 11:58 AM UTC

/retest

### Comment by @mkholjuraev on August 07, 2023 at 02:34 PM UTC

:tada: This PR is included in version 1.63.11 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.11)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Commented on June 27, 2023 at 05:20 PM UTC

### Review by @bastilian - Commented on June 27, 2023 at 05:36 PM UTC

### Review by @bastilian - Approved on June 27, 2023 at 05:45 PM UTC

Nice job! Thank you @AsToNlele!

### Review by @bastilian - Changes Requested on June 27, 2023 at 06:05 PM UTC

Sorry, but I did find issues doublechecking on `/preview`.
There everything appears to fall apart. :/

### Review by @bastilian - Commented on June 28, 2023 at 06:50 PM UTC

### Review by @bastilian - Commented on June 28, 2023 at 06:54 PM UTC

### Review by @johnsonm325 - Approved on July 17, 2023 at 09:47 PM UTC

This seems to work okay.

There is a bug when you apply an os filter to the systems table and then try to refresh the page. But it exists in stage, so isn't related to this PR. I'm creating a card.

### Review by @bastilian - Commented on July 25, 2023 at 10:50 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1055*
