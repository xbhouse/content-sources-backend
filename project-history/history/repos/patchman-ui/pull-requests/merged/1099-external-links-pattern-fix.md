---
type: pull_request
number: 1099
title: "External Links Pattern Fix"
state: merged
author: LiorKGOW
created: 2023-07-31T14:54:08Z
updated: 2023-08-07T14:35:29Z
closed: 2023-07-31T16:40:38Z
merged: 2023-07-31T16:40:38Z
base_branch: master
head_branch: external-links-pattern-fix
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1099
---

# Pull Request #1099: External Links Pattern Fix

**Author**: @LiorKGOW
**Created**: July 31, 2023 at 02:54 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `external-links-pattern-fix`

## Description

# Description
Changing the external link pattern on Patchman to be aligned with Vulnerability.
External links in Patchman now use ExternalLinkAltIcon PF4 icon component and have link text in front of the icon. 

New icon used: 
![ExternalLinkAltIcon](https://github.com/RedHatInsights/patchman-ui/assets/93318917/c434b05a-17de-4ff6-95f0-de102b2f0680)

Old icon: 
![Old icon](https://github.com/RedHatInsights/patchman-ui/assets/93318917/694e381c-37a4-44c6-80c2-abd1947577b0)


Associated Jira ticket: [#SPM-1761](https://issues.redhat.com/browse/SPM-1761)

# How to test the PR
Insure the correct icon is being presented on the screen
This is the new icon used: [ExternalLinkAltIcon](https://www.patternfly.org/2023.02/guidelines/icons/#all-icons)

# Before the change
![Before - whole page](https://github.com/RedHatInsights/patchman-ui/assets/93318917/5084818a-0dfa-4cd3-ae5d-863d2f2629dd)

![Before - just the link](https://github.com/RedHatInsights/patchman-ui/assets/93318917/abb31125-92f9-4bb9-ae26-fcda59492043)


# After the change
![After - whole page](https://github.com/RedHatInsights/patchman-ui/assets/93318917/1f8c42c5-45ab-4c92-b0b0-675c37043f00)

![After - just the link](https://github.com/RedHatInsights/patchman-ui/assets/93318917/085ef904-107f-410e-ae1f-0ec348273215)


# Checklist:

- [X] The commit message has the Jira ticket linked
- [X] PR has a short description
- [X] Screenshots before and after the change are added
- [x] Tests for the changes have been added -> no need for tests, this is a small change
- [X] README.md is updated if necessary
- [x] Needs additional dependent work -> no


---

## Discussion

### Comment by @LiorKGOW on July 31, 2023 at 03:04 PM UTC

I hope this PR is organized as expected, please let me know in case something is missing and I will add it :) 
I am not sure why, but when I pushed the PR, it automatically added a new commit with the package.json changes to it. Is this OK ? 

### Comment by @codecov-commenter on July 31, 2023 at 03:04 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1099?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch and project coverage have no change.
> Comparison is base [(`747e04d`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/747e04dc644a2e6c313f4f1960acec08f22c16de?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.57% compared to head [(`105f8ee`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1099?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.57%.
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1099   +/-   ##
=======================================
  Coverage   62.57%   62.57%           
=======================================
  Files         119      119           
  Lines        2993     2993           
  Branches      769      769           
=======================================
  Hits         1873     1873           
  Misses       1120     1120           
```


| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1099?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [.../PresentationalComponents/Snippets/ExternalLink.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1099?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FeHRlcm5hbExpbmsuanM=) | `100.00% <ø> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1099?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @LiorKGOW on July 31, 2023 at 03:35 PM UTC

> Hey, this PR looks perfect, thanks for the icon fix, can you please remove the commit which updates the Patternfly package? It makes template wizard crash. After that is done, PR can be merged. You can now change the associated ticket's status to "Code review".

Thank you for the quick response @leSamo 
I'll fix it and ping you when it is done 👍🏼 

### Comment by @LiorKGOW on July 31, 2023 at 04:13 PM UTC

@leSamo I have removed the extra commit 👍🏼 
This PR can be merged now

### Comment by @mkholjuraev on August 07, 2023 at 02:35 PM UTC

:tada: This PR is included in version 1.63.11 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.11)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Commented on July 31, 2023 at 03:07 PM UTC

Hey, this PR looks perfect, thanks for the icon fix, can you please remove the commit which updates the Patternfly package? It makes template wizard crash. After that is done, PR can be merged. You can now change the associated ticket's status to "Code review".

### Review by @leSamo - Approved on July 31, 2023 at 04:40 PM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1099*
