---
type: pull_request
number: 637
title: "feat(StatusReport): SPM-1043 create System status report"
state: merged
author: mkholjuraev
created: 2021-08-20T09:59:23Z
updated: 2021-08-24T07:46:13Z
closed: 2021-08-24T07:30:20Z
merged: 2021-08-24T07:30:20Z
base_branch: master
head_branch: advisories
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/637
---

# Pull Request #637: feat(StatusReport): SPM-1043 create System status report

**Author**: @mkholjuraev
**Created**: August 20, 2021 at 09:59 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `advisories`

## Description

There is a spacing difference between the top and bottom of the status bar. I liked that :). Because the status bar and the table are one group and it makes sense to put it a bit far from the header.

![image](https://user-images.githubusercontent.com/59481011/130216628-1518807f-62e7-498e-b40c-96ba03a7df2d.png)


---

## Discussion

### Comment by @codecov-commenter on August 20, 2021 at 10:05 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#637](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (478df86) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/303ae0700d96a40debc3eb3b1ca4f2f1892045d8?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (303ae07) will **decrease** coverage by `0.33%`.
> The diff coverage is `15.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #637      +/-   ##
==========================================
- Coverage   52.94%   52.60%   -0.34%     
==========================================
  Files          78       79       +1     
  Lines        1921     1937      +16     
  Branches      420      423       +3     
==========================================
+ Hits         1017     1019       +2     
- Misses        824      836      +12     
- Partials       80       82       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <0.00%> (ø)` | |
| [src/store/Reducers/SystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbXNTdG9yZS5qcw==) | `40.00% <0.00%> (-10.00%)` | :arrow_down: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `85.00% <50.00%> (-0.62%)` | :arrow_down: |
| [src/store/ActionTypes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvblR5cGVzLmpz) | `100.00% <100.00%> (ø)` | |
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `72.30% <100.00%> (+0.43%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [303ae07...478df86](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/637?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on August 20, 2021 at 01:20 PM UTC

Looks great! A few notes:

- Can we add a spinner or indicate that the number is loading?
- Numbers should be clickable - not only blue - on click it should filter the list

### Comment by @jiridostal on August 23, 2021 at 02:26 PM UTC

The title should not be "Most impactful advisories". :-)
The space beween the icon and number is too big, make it smaller please.
Other than that, instant merge!

### Comment by @mkholjuraev on August 23, 2021 at 02:48 PM UTC

Yeah, I have to live with stupid mistakes :)

I have noticed that in the mockup there is not a title at all. https://marvelapp.com/prototype/15116dj4/screen/80040041

This is how UI looks now

![image](https://user-images.githubusercontent.com/59481011/130468005-3b7c869a-944f-4114-91c3-873240d8c71b.png)


### Comment by @jiridostal on August 24, 2021 at 07:46 AM UTC

:tada: This PR is included in version 1.31.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.31.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.31.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/637*
