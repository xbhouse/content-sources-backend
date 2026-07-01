---
type: pull_request
number: 888
title: "feat(SPM-1707): paginate remediation pair fetching"
state: merged
author: mkholjuraev
created: 2022-10-13T20:33:55Z
updated: 2024-04-03T09:21:58Z
closed: 2022-10-17T10:52:48Z
merged: 2022-10-17T10:52:48Z
base_branch: master
head_branch: fix-remediations
labels: ["enhancement", "released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/888
---

# Pull Request #888: feat(SPM-1707): paginate remediation pair fetching

**Author**: @mkholjuraev
**Created**: October 13, 2022 at 08:33 PM UTC
**Status**: Merged
**Labels**: `enhancement`, `released`
**Base**: `master` ← **Head**: `fix-remediations`

## Description

This will enable paginating the request to fetch remediation pairs on Systems and Advisories pages.

To test:
1. Visit the advisories page
2. bulk select a page/all rows
3. click on Remediate
4. Observe that there are multiple requests made to fetch paginated result
5. Observe that those results are being aggregated correctly
6. Observe that Remediation modal opens up as expected and you can successfully remediate
7. Please do so on the Systems page as well.

---

## Discussion

### Comment by @codecov-commenter on October 15, 2022 at 06:35 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/888?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **71.70**% // Head: **71.66**% // Decreases project coverage by **`-0.04%`** :warning:
> Coverage data is based on head [(`0c1655f`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/888?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`d38a3f1`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/d38a3f166562a661e5de063e51263298e3a2e91b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 100.00% of modified lines in pull request are covered.

> :exclamation: Current head 0c1655f differs from pull request most recent head 993c188. Consider uploading reports for the commit 993c188 to get more accurate results

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #888      +/-   ##
==========================================
- Coverage   71.70%   71.66%   -0.05%     
==========================================
  Files         108      108              
  Lines        2520     2527       +7     
  Branches      650      649       -1     
==========================================
+ Hits         1807     1811       +4     
- Misses        650      653       +3     
  Partials       63       63              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/888?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/useRemediationDataProvider.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/888/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91c2VSZW1lZGlhdGlvbkRhdGFQcm92aWRlci5qcw==) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/888/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `54.31% <0.00%> (-2.59%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/888?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on October 17, 2022 at 10:24 AM UTC

@gkarat your comment regarding the Advisor page selection does not comply with the UX guidelines. We explicitly implemented bulk selection in such a way that after all items are selected in the table, you can not select a page. 

### Comment by @gkarat on October 17, 2022 at 10:30 AM UTC

ok, then PR is clean, LGTM

### Comment by @mkholjuraev on October 17, 2022 at 10:50 AM UTC

@gkarat thank you for the review!

### Comment by @jiridostal on October 17, 2022 at 11:07 AM UTC

:tada: This PR is included in version 1.53.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.53.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.53.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Approved on October 17, 2022 at 08:40 AM UTC

Looks OK to me, tried locally and the requests are chunked as expected on both pages. But there is a couple of things I want you to improve, please see the comments.

---
Also one thing, rather a minor bug:
1. Go to the Advisories page and select all items in the table.
2. Then try to select only the current page (20 items by default).

![image](https://user-images.githubusercontent.com/31385370/196136882-ff080915-cd47-458e-962a-50992db4f6cf.png)

You won't be able to do that and even after you select only the page items, the bulk select still displays that all systems are selected. You can fix it by clicking "-" in the checkbox, and then select the page. But I would expect you can also achieve this with both ways.

### Review by @mkholjuraev - Commented on October 17, 2022 at 10:25 AM UTC

### Review by @mkholjuraev - Commented on October 17, 2022 at 10:26 AM UTC

### Review by @mkholjuraev - Commented on October 17, 2022 at 10:28 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/888*
