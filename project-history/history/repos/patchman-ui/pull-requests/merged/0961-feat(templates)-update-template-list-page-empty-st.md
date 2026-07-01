---
type: pull_request
number: 961
title: "feat(Templates): Update Template list page empty state"
state: merged
author: leSamo
created: 2023-02-15T22:43:59Z
updated: 2023-05-08T18:17:06Z
closed: 2023-02-27T17:00:05Z
merged: 2023-02-27T17:00:05Z
base_branch: master
head_branch: patch-list-empty-state
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/961
---

# Pull Request #961: feat(Templates): Update Template list page empty state

**Author**: @leSamo
**Created**: February 15, 2023 at 10:43 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `patch-list-empty-state`

## Description

# Description

Associated Jira ticket: [SPM-1892](https://issues.redhat.com/browse/SPM-1892)
Mockup: https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/25O3lOp

Update Template list page empty state according to the mockup.

# How to test the PR
Verify correspondence of Templates page when there are 0 templates present to the mockup.

# Mockup
![Screenshot from 2023-02-15 23-43-26](https://user-images.githubusercontent.com/8426204/219205080-98662354-6a72-4971-92e5-d4bf3365c294.png)

# Before the change
![Screenshot from 2023-02-15 23-09-54](https://user-images.githubusercontent.com/8426204/219205102-b36cf24f-6c6c-4e3f-a127-ffc27f6e6066.png)

# After the change
![Screenshot from 2023-02-15 23-08-19](https://user-images.githubusercontent.com/8426204/219205127-e76dce89-9ff1-49d5-9ea0-4176d18ef65d.png)

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 15, 2023 at 11:11 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/961?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.37**% // Head: **72.39**% // Increases project coverage by **`+0.02%`** :tada:
> Coverage data is based on head [(`05f0b18`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/961?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`41668f2`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/41668f27e9a0a506ad08fe24c2bde5487e428661?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 64.28% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #961      +/-   ##
==========================================
+ Coverage   72.37%   72.39%   +0.02%     
==========================================
  Files         110      110              
  Lines        2617     2619       +2     
  Branches      658      659       +1     
==========================================
+ Hits         1894     1896       +2     
  Misses        723      723              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/961?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/961?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `74.35% <0.00%> (+0.94%)` | :arrow_up: |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/961?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `73.33% <66.66%> (+1.90%)` | :arrow_up: |
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/961?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `84.05% <87.50%> (-1.02%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/961?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on February 27, 2023 at 05:15 PM UTC

:tada: This PR is included in version 1.60.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.60.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.60.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on February 27, 2023 at 10:40 AM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/961*
