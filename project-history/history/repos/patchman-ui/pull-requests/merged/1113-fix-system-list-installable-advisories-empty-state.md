---
type: pull_request
number: 1113
title: "Fix System list installable advisories empty state text"
state: merged
author: leSamo
created: 2023-08-28T08:28:48Z
updated: 2024-01-08T13:49:57Z
closed: 2023-08-28T09:55:23Z
merged: 2023-08-28T09:55:23Z
base_branch: master
head_branch: fix-no-applicable-adv
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1113
---

# Pull Request #1113: Fix System list installable advisories empty state text

**Author**: @leSamo
**Created**: August 28, 2023 at 08:28 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-no-applicable-adv`

## Description

https://issues.redhat.com/browse/SPM-2085

On the System list page the "Installable advisories" column used to say "No applicable advisories" instead of "No installable advisories" when the system had no advisories.

## Before:
![Screenshot from 2023-08-28 10-08-45](https://github.com/RedHatInsights/patchman-ui/assets/8426204/87c93712-0a19-4149-8a7b-31d9fa469c4f)

## After:
![Screenshot from 2023-08-28 10-09-07](https://github.com/RedHatInsights/patchman-ui/assets/8426204/7e153524-74c7-47ac-b036-f5c5b8191f80)


---

## Discussion

### Comment by @codecov-commenter on August 28, 2023 at 08:37 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1113?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00%`** and no project coverage change.
> Comparison is base [(`4fb84ef`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4fb84ef910749985d682da0c051a2d448520ce9d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.15% compared to head [(`855a490`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1113?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.15%.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1113   +/-   ##
=======================================
  Coverage   63.15%   63.15%           
=======================================
  Files         120      120           
  Lines        2994     2994           
  Branches      765      765           
=======================================
  Hits         1891     1891           
  Misses       1103     1103           
```


| [Files Changed](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1113?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1113?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `60.52% <100.00%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1113?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on August 30, 2023 at 08:07 PM UTC

:tada: This PR is included in version 1.64.0 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.64.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on August 28, 2023 at 08:41 AM UTC

LGTM! 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1113*
