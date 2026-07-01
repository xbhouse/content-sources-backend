---
type: pull_request
number: 969
title: "feat: Implement Template assigned systems table"
state: merged
author: leSamo
created: 2023-02-28T16:03:09Z
updated: 2023-05-08T18:17:02Z
closed: 2023-03-02T10:53:47Z
merged: 2023-03-02T10:53:47Z
base_branch: master
head_branch: template-detail-table
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/969
---

# Pull Request #969: feat: Implement Template assigned systems table

**Author**: @leSamo
**Created**: February 28, 2023 at 04:03 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `template-detail-table`

## Description

# Description

Associated Jira ticket: [SPM-1923](https://issues.redhat.com/browse/SPM-1923)

Implement bare Template assigned systems table without filters, sorting working, selection, bulk actions and row actions.

![Screenshot from 2023-02-28 16-59-18](https://user-images.githubusercontent.com/8426204/221908859-895cf6b6-5b25-4154-b1a6-9c8e49d8d5d8.png)

Add 'Edit template' button to Template detail page.

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 28, 2023 at 10:59 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`21.56`**% and project coverage change: **`-1.61`** :warning:
> Comparison is base [(`b8ce309`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/b8ce309d356a31c14e907abe7384fada6dc843c7?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 70.61% compared to head [(`039d694`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 69.01%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #969      +/-   ##
==========================================
- Coverage   70.61%   69.01%   -1.61%     
==========================================
  Files         113      115       +2     
  Lines        2692     2782      +90     
  Branches      684      702      +18     
==========================================
+ Hits         1901     1920      +19     
- Misses        791      862      +71     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...tComponents/PatchSetDetail/PatchSetDetailAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbEFzc2V0cy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `59.09% <0.00%> (ø)` | |
| [src/Utilities/useOnSelect.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91c2VPblNlbGVjdC5qcw==) | `100.00% <ø> (ø)` | |
| [src/store/Reducers/PatchSetDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0RGV0YWlsU3RvcmUuanM=) | `42.85% <0.00%> (ø)` | |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `100.00% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `70.23% <20.00%> (-3.18%)` | :arrow_down: |
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `50.00% <33.33%> (-4.55%)` | :arrow_down: |
| [src/store/Reducers/PatchSetDetailSystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0RGV0YWlsU3lzdGVtc1N0b3JlLmpz) | `35.29% <35.29%> (ø)` | |
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `80.64% <45.45%> (-4.73%)` | :arrow_down: |
| ... and [5 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/969?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @Fewwy on March 02, 2023 at 10:31 AM UTC

LGTM but I don't have any sketches to compare the result to

### Comment by @Fewwy on March 02, 2023 at 10:38 AM UTC

Looking at the sketch:
1. Links are not in the correct order 
![image](https://user-images.githubusercontent.com/62722417/222404505-619dbd69-5fd1-4b94-a146-2589ed0aa353.png)


### Comment by @Fewwy on March 02, 2023 at 10:39 AM UTC

When I'm changing pages in the table the column width changes.
![image](https://user-images.githubusercontent.com/62722417/222405238-9ded2b59-ebe2-4d69-801d-501b639ec0d7.png)

![image](https://user-images.githubusercontent.com/62722417/222405297-016f30f7-90ab-45a8-bd5f-7eec35cc7a22.png)



### Comment by @mkholjuraev on March 02, 2023 at 11:16 AM UTC

:tada: This PR is included in version 1.62.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/969*
