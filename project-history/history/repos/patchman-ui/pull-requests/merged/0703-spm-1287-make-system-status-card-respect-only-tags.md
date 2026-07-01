---
type: pull_request
number: 703
title: "SPM-1287: make system status card respect only tags, os, systemPlatform"
state: merged
author: mkholjuraev
created: 2022-01-14T11:32:24Z
updated: 2022-05-17T08:50:03Z
closed: 2022-02-01T09:33:33Z
merged: 2022-02-01T09:33:33Z
base_branch: master
head_branch: system-status-card
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/703
---

# Pull Request #703: SPM-1287: make system status card respect only tags, os, systemPlatform

**Author**: @mkholjuraev
**Created**: January 14, 2022 at 11:32 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `system-status-card`

## Description

GlobalFilterReducer usage is enhanced together with the ticket expectation.

---

## Discussion

### Comment by @codecov-commenter on January 24, 2022 at 01:18 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#703](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (466810a) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/e674d2ec3226a5b5376e912676b4f962a6d2e5d1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e674d2e) will **decrease** coverage by `0.06%`.
> The diff coverage is `71.42%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #703      +/-   ##
==========================================
- Coverage   81.98%   81.92%   -0.07%     
==========================================
  Files          84       84              
  Lines        1915     1925      +10     
  Branches      498      498              
==========================================
+ Hits         1570     1577       +7     
- Misses        303      305       +2     
- Partials       42       43       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `86.30% <50.00%> (-0.60%)` | :arrow_down: |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `78.94% <66.66%> (-1.76%)` | :arrow_down: |
| [src/store/Reducers/GlobalFilterStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0dsb2JhbEZpbHRlclN0b3JlLmpz) | `50.00% <66.66%> (+10.00%)` | :arrow_up: |
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | `77.77% <80.00%> (-2.23%)` | :arrow_down: |
| [src/store/ActionTypes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvblR5cGVzLmpz) | `100.00% <100.00%> (ø)` | |
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `89.06% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [e674d2e...466810a](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/703?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on February 22, 2022 at 10:08 AM UTC

:tada: This PR is included in version 1.39.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.39.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.39.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Commented on January 28, 2022 at 07:23 AM UTC

### Review by @mkholjuraev - Commented on January 28, 2022 at 09:19 AM UTC

### Review by @bastilian - Commented on January 31, 2022 at 01:17 PM UTC

### Review by @mkholjuraev - Commented on January 31, 2022 at 01:28 PM UTC

### Review by @bastilian - Commented on January 31, 2022 at 01:45 PM UTC

### Review by @mkholjuraev - Commented on January 31, 2022 at 02:24 PM UTC

### Review by @bastilian - Commented on January 31, 2022 at 03:33 PM UTC

### Review by @bastilian - Approved on February 01, 2022 at 08:22 AM UTC

Works for me! Thank you @mkholjuraev!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/703*
