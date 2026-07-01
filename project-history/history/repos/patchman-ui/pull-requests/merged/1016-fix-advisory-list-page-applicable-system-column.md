---
type: pull_request
number: 1016
title: "Fix Advisory list page applicable system column"
state: merged
author: leSamo
created: 2023-04-03T18:57:47Z
updated: 2023-05-08T18:16:29Z
closed: 2023-04-06T21:29:26Z
merged: 2023-04-06T21:29:26Z
base_branch: master
head_branch: fix-advisory-list
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1016
---

# Pull Request #1016: Fix Advisory list page applicable system column

**Author**: @leSamo
**Created**: April 03, 2023 at 06:57 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-advisory-list`

## Description

See comment in [SPM-1948](https://issues.redhat.com/browse/SPM-1948)

- rename "Applicable systems" to "Affected systems" in table and dashbar
- update `advisories/` endpoint version from v2 to v3 to fix incorrect count of systems in applicable/affected systems column

---

## Discussion

### Comment by @codecov-commenter on April 03, 2023 at 07:03 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1016?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch and project coverage have no change.
> Comparison is base [(`b664598`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/b664598ee2017b56334bda703444afb350f65a90?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.32% compared to head [(`188001c`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1016?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.32%.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1016   +/-   ##
=======================================
  Coverage   64.32%   64.32%           
=======================================
  Files         116      116           
  Lines        2820     2820           
  Branches      724      724           
=======================================
  Hits         1814     1814           
  Misses       1006     1006           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1016?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...sentationalComponents/TableView/TableViewAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1016?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3QXNzZXRzLmpz) | `100.00% <ø> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1016?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `56.36% <0.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1016?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @adonispuente - Approved on April 06, 2023 at 03:29 PM UTC

LGTM!
Just that small nitpick. 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1016*
