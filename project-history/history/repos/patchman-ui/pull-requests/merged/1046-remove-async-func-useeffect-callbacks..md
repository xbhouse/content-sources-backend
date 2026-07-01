---
type: pull_request
number: 1046
title: "Remove async func useEffect callbacks."
state: merged
author: Hyperkid123
created: 2023-05-09T08:50:00Z
updated: 2023-05-09T16:05:24Z
closed: 2023-05-09T10:39:59Z
merged: 2023-05-09T10:39:59Z
base_branch: master
head_branch: remove-direct-use-effect-async
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1046
---

# Pull Request #1046: Remove async func useEffect callbacks.

**Author**: @Hyperkid123
**Created**: May 09, 2023 at 08:50 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remove-direct-use-effect-async`

## Description

Related ticket: https://issues.redhat.com/browse/RHCLOUD-20494

### Why

TLDR: React only accepts functions as `effect` return value from its callback. Directly used `async` function will always have a return type of `Promise`.

You can read further here: https://github.com/facebook/react/issues/14326


---

## Discussion

### Comment by @codecov-commenter on May 09, 2023 at 09:01 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1046?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00`**% and no project coverage change.
> Comparison is base [(`1855269`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1855269343fc8b8241c9e9968dd4991d4e5c0598?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.09% compared to head [(`efa01ec`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1046?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.09%.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1046   +/-   ##
=======================================
  Coverage   63.09%   63.09%           
=======================================
  Files         119      119           
  Lines        2962     2962           
  Branches      763      763           
=======================================
  Hits         1869     1869           
  Misses       1093     1093           
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1046?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...Components/StatusReports/AdvisoriesStatusReport.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1046?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL0Fkdmlzb3JpZXNTdGF0dXNSZXBvcnQuanM=) | `53.33% <100.00%> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1046?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `90.74% <100.00%> (ø)` | |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1046?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on May 09, 2023 at 04:05 PM UTC

:tada: This PR is included in version 1.63.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on May 09, 2023 at 10:39 AM UTC

LGTM, yup indeed correct. I was young and dummy :)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1046*
