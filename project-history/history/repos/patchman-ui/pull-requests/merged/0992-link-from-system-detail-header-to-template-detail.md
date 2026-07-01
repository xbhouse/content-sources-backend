---
type: pull_request
number: 992
title: "Link from System detail header to template detail"
state: merged
author: leSamo
created: 2023-03-14T19:49:21Z
updated: 2023-05-08T18:16:36Z
closed: 2023-03-17T16:33:08Z
merged: 2023-03-17T16:33:08Z
base_branch: master
head_branch: system-detail-template-link
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/992
---

# Pull Request #992: Link from System detail header to template detail

**Author**: @leSamo
**Created**: March 14, 2023 at 07:49 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `system-detail-template-link`

## Description

# Description

Associated Jira ticket: [SPM-1945](https://issues.redhat.com/browse/SPM-1945)

In System detail header
- update template name to link to Template detail page
- fix race condition where Inventory API response could overwrite `hasThirdPartyRepo` and `patchSetName` parameters and header would show up as having no template and no third party repo
- fix tags not displaying



---

## Discussion

### Comment by @codecov-commenter on March 14, 2023 at 07:54 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/992?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> :exclamation: No coverage uploaded for pull request base (`master@b26f944`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#section-missing-base-commit).
> Patch coverage: 71.42% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff            @@
##             master     #992   +/-   ##
=========================================
  Coverage          ?   68.90%           
=========================================
  Files             ?      116           
  Lines             ?     2811           
  Branches          ?      715           
=========================================
  Hits              ?     1937           
  Misses            ?      874           
  Partials          ?        0           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/992?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/store/Reducers/SystemDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/992?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbURldGFpbFN0b3JlLmpz) | `56.25% <0.00%> (ø)` | |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/992?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `85.18% <80.00%> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/992?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `58.18% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/992?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 28, 2023 at 08:57 AM UTC

:tada: This PR is included in version 1.62.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on March 17, 2023 at 03:06 PM UTC

LGTM! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/992*
