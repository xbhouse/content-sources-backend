---
type: pull_request
number: 1106
title: "AdvisoriesPage - Remove unused 'page' parameter from queryParams in the AdvisoryListStore"
state: merged
author: LiorKGOW
created: 2023-08-02T10:24:23Z
updated: 2023-08-07T14:35:23Z
closed: 2023-08-07T14:16:00Z
merged: 2023-08-07T14:16:00Z
base_branch: master
head_branch: advisories-remove-page-from-url-params
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1106
---

# Pull Request #1106: AdvisoriesPage - Remove unused 'page' parameter from queryParams in the AdvisoryListStore

**Author**: @LiorKGOW
**Created**: August 02, 2023 at 10:24 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `advisories-remove-page-from-url-params`

## Description

# Description

Associated Jira ticket: [#SPM-1623](https://issues.redhat.com/browse/SPM-1623)

Current page is always set to 1 in the URL on Advisories Page, even though the current presented page value it is being handled by a PF4 Pagination component inside the page 
(the 'page' parameter from the queryParams in the AdvisoryListStore is not used).
It was suggested on the ticket by Jiri that the page parameter should be removed completely from the URL
After a conversation with Sebastian, we have decided to also remove page_size from queryParam, because we already handle offset & limit in the URL which have the same purposes of page & page_size

# How to test the PR
Insure the component is working correctly, even without the missing 'page' & 'page_size' parameters from the queryParams.

# Screenshots Before the change
Old URL of the 1st page, with 'page' & 'page_size' query parameters (page is 1 (pink), page_size (orange) & offset is 0 (yellow)): 
![old url with page & page_size queryParams (on page 1)](https://github.com/RedHatInsights/patchman-ui/assets/93318917/6fa7b6c1-6b70-46cb-b9f1-ea23f4ff47e8)

Old URL of the 2nd page, with 'page' query parameter which is still 1 (page is still 1 (pink) & offset is 20 (yellow)): 
![old url with the page queryParam (on page 2)](https://github.com/RedHatInsights/patchman-ui/assets/93318917/37a83db6-a79f-4155-a233-63539cb92935)


Whole Advisories Page on the 1st page (pink: page parameter in the URL is 1 as it should be, green: the 'page' in the page is also 1, yellow: offset is 0 (in the URL), red: the 1st page is being presented, items 1-20 are being shown in the table, insuring we are on the 1st page):
![before - whole page on page 1](https://github.com/RedHatInsights/patchman-ui/assets/93318917/a50af8fc-9e5c-4aad-a2ae-ab230d3d7238)


Whole Advisories Page on the 2nd page (pink: page parameter in the URL is 1 (this is the bug), green: the 'page' in the page is 2, yellow: offset is 20 (in the URL), red: the 2nd page is being presented, items 21-40 are being shown in the table, insuring we are on the 2nd page):
![before - whole page on page 2 (page in the URL is still 1)](https://github.com/RedHatInsights/patchman-ui/assets/93318917/c8abdc93-d16b-41ad-8135-968b76c58c0e)


# Screenshots After the change

New URL of the 1st page, without 'page' & 'page_size' in the query parameters, 'limit' has now swapped 'page_size' ('limit' swapped 'page_size' (pink) & 'offset' is 0 (yellow)): 
![new url without the page & page_size queryParams (on page 1)](https://github.com/RedHatInsights/patchman-ui/assets/93318917/1f76d36c-0454-476a-8131-b8941b7304c2)

New URL of the 2nd page, without 'page' & 'page_size' in the query parameters ('limit' swapped 'page_size' (pink) & 'offset' is 20 (yellow)): 
![new url without page   page_size in queryParams & limit is visible(on page 2)](https://github.com/RedHatInsights/patchman-ui/assets/93318917/7b2a77f7-f6fb-4664-a5c2-ed5708dbabb7)


Whole Advisories Page on the 1st page ('page' parameter is not shown in the URL (the fix), green: the 'page' in the page is 1, yellow: 'offset' is 0, orange: 'limit' has now swapped 'page_size', red: the 1st page is being presented, items 1-20 are being shown in the table, insuring we are on the 1st page):
![after - whole page on page 1 (page was removed from the url   page_size was switched with limit)](https://github.com/RedHatInsights/patchman-ui/assets/93318917/29807ee7-c18b-4f63-a06a-71f67ff69061)

Whole Advisories Page on the 2nd page (the 'page' query parameter in the URL is not shown (the fix), green: the 'page' in the page is 2, yellow: offset is 20, orange: 'limit' has now swapped 'page_size', red: the 2nd page is being presented, items 21-40 are being shown in the table, insuring we are on the 2nd page):
![after - whole page on page 2 (offset is still being handled in the url)](https://github.com/RedHatInsights/patchman-ui/assets/93318917/532b48a8-04b1-403e-8962-6a96fa215527)

# Checklist:

- [X] The commit message has the Jira ticket linked
- [X] PR has a short description
- [X] Screenshots before and after the change are added
- [X] Tests for the changes have been adjusted accordingly 
- [X] README.md is updated if necessary -> no need for an update
- [X] Needs additional dependent work -> no need 

---

## Discussion

### Comment by @codecov-commenter on August 02, 2023 at 11:06 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1106?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00%`** and project coverage change: **`+0.18%`** :tada:
> Comparison is base [(`d22968d`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d22968d8ef0bd435547730012ff23fc4a0d21fd6?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.57% compared to head [(`ed16b06`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1106?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.76%.
> Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1106      +/-   ##
==========================================
+ Coverage   62.57%   62.76%   +0.18%     
==========================================
  Files         119      120       +1     
  Lines        2993     3470     +477     
  Branches      769      931     +162     
==========================================
+ Hits         1873     2178     +305     
- Misses       1120     1292     +172     
```


| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1106?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/store/Reducers/AdvisoryListStore.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1106?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5TGlzdFN0b3JlLmpz) | `69.56% <100.00%> (+1.38%)` | :arrow_up: |

... and [21 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1106/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1106?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @LiorKGOW on August 07, 2023 at 10:56 AM UTC

> Works well, I do not see those redundant params in the URL.
> 
> Note: Probably those params can be removed from [here](https://github.com/RedHatInsights/patchman-ui/blob/ed16b06ddee2729db03a787a9a408da6a1357f20/src/Utilities/constants.js#L46) completely to make all pages behave same

I tried at first removing them from the file you specified, but that broke lots of tests, including ones that are not related to the Advisories Page & others who are related to components that do not use the AdvisoryListStore
( like: Packages.js.test , SystemAdvisoryListStore.test.js , SystemPackageListStore.test.js )

They are all using the same constant 'storeListDefaults' from Utilities/constants in their initialization

After consulting with @bastilian and the rest of the team, we decided to go for this fix (just removing page & page_size parameters from QueryParams) to align the PR with the issue

Do you think we should open a different issue on this matter ? in order to keep all the other pages aligned ? 
@mkholjuraev 

### Comment by @mkholjuraev on August 07, 2023 at 02:35 PM UTC

:tada: This PR is included in version 1.63.11 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.11)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on August 07, 2023 at 10:51 AM UTC

Works well, I do not see those redundant params in the URL. 

Note: Probably those params can be removed from [here](https://github.com/RedHatInsights/patchman-ui/blob/ed16b06ddee2729db03a787a9a408da6a1357f20/src/Utilities/constants.js#L46) completely to make all pages behave same

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1106*
