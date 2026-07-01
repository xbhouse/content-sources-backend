---
type: pull_request
number: 813
title: "fix(SPM-1457): enable up-to-date systems to be operated on"
state: merged
author: mkholjuraev
created: 2022-06-03T08:18:14Z
updated: 2024-04-03T09:21:50Z
closed: 2022-06-07T13:32:25Z
merged: 2022-06-07T13:32:25Z
base_branch: master
head_branch: up-to-date-systems
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/813
---

# Pull Request #813: fix(SPM-1457): enable up-to-date systems to be operated on

**Author**: @mkholjuraev
**Created**: June 03, 2022 at 08:18 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `up-to-date-systems`

## Description

Resolves  [SPM-1457](https://issues.redhat.com/browse/SPM-1457) by enabling up-to-date systems to be bulk selected. NoDataModal is displayed when an up-to-date system is chosen to be remediated. In case, more than one of those systems are passed to Remediation with remediatable systems, that modal does not show up.This is consistent with current Remediation application.

This PR also disables appliying remediation to up-to-date systems on table row actions.

To test:

1.  Navigate to Systems page.
2. Select any one system that is up-to-date. ( you can decide if a system is up-to-date by looking at the table row actions. If 'Apply all applicable advisories' action is disabled, it means that this  system is up-to-date).
3. Click Remediate button on the table toolbar.
4. NoDataModal should show up.

---

## Discussion

### Comment by @codecov-commenter on June 03, 2022 at 10:19 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/813?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#813](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/813?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d497e5f) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/1f12520152242abaa4a103c76d60472842de1681?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1f12520) will **increase** coverage by `0.26%`.
> The diff coverage is `85.71%`.

> :exclamation: Current head d497e5f differs from pull request most recent head 8489e72. Consider uploading reports for the commit 8489e72 to get more accurate results

```diff
@@            Coverage Diff             @@
##           master     #813      +/-   ##
==========================================
+ Coverage   70.69%   70.96%   +0.26%     
==========================================
  Files         102      101       -1     
  Lines        2440     2435       -5     
  Branches      628      629       +1     
==========================================
+ Hits         1725     1728       +3     
+ Misses        647      640       -7     
+ Partials       68       67       -1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/813?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/813/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `73.91% <0.00%> (+3.24%)` | :arrow_up: |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/813/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `68.35% <ø> (-0.40%)` | :arrow_down: |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/813/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `81.81% <100.00%> (+2.87%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/813/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `76.82% <100.00%> (+1.15%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/813?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/813?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [1f12520...8489e72](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/813?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on June 07, 2022 at 01:48 PM UTC

:tada: This PR is included in version 1.48.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gabipodolnikova - Commented on June 07, 2022 at 10:56 AM UTC

### Review by @mkholjuraev - Commented on June 07, 2022 at 11:11 AM UTC

### Review by @gabipodolnikova - Approved on June 07, 2022 at 01:20 PM UTC

LGTM! 😎 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/813*
