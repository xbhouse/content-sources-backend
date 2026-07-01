---
type: pull_request
number: 895
title: "Support component tests"
state: merged
author: gkarat
created: 2022-11-01T15:40:49Z
updated: 2022-11-10T13:08:49Z
closed: 2022-11-07T09:44:18Z
merged: 2022-11-07T09:44:18Z
base_branch: master
head_branch: support-cypress
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/895
---

# Pull Request #895: Support component tests

**Author**: @gkarat
**Created**: November 01, 2022 at 03:40 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `support-cypress`

## Description

This PR installs cypress-related dependencies and makes sure we can run component tests. The cypress bundle is now also included to the Test stage in Travis CI (this doesn't modify pr_check since we are not changing `verify` npm script).

![изображение](https://user-images.githubusercontent.com/31385370/199274967-6bcacf1b-230a-47dc-b964-d3739b32ae49.png)


---

## Discussion

### Comment by @codecov-commenter on November 01, 2022 at 03:49 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/895?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.26**% // Head: **72.04**% // Decreases project coverage by **`-0.22%`** :warning:
> Coverage data is based on head [(`664dfb5`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/895?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`daae4a7`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/daae4a7abe4c099ffefb14dc7c0fdb41c53633e3?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 0.00% of modified lines in pull request are covered.

> :exclamation: Current head 664dfb5 differs from pull request most recent head 3ba7b19. Consider uploading reports for the commit 3ba7b19 to get more accurate results

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #895      +/-   ##
==========================================
- Coverage   72.26%   72.04%   -0.23%     
==========================================
  Files         109      110       +1     
  Lines        2564     2572       +8     
  Branches      655      655              
==========================================
  Hits         1853     1853              
- Misses        649      657       +8     
  Partials       62       62              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/895?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...ntationalComponents/Header/HeaderBreadcrumbs.cy.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/895/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9IZWFkZXIvSGVhZGVyQnJlYWRjcnVtYnMuY3kuanM=) | `0.00% <0.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/895?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on November 10, 2022 at 01:08 PM UTC

:tada: This PR is included in version 1.55.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.55.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.55.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on November 04, 2022 at 12:36 PM UTC

LGTM! @gkarat thank you!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/895*
