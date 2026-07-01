---
type: pull_request
number: 653
title: "feat(AdvisoriesReport): SPM-1167 add most impactfull advisory info"
state: merged
author: mkholjuraev
created: 2021-09-20T11:14:34Z
updated: 2022-05-17T08:49:50Z
closed: 2021-09-24T12:28:40Z
merged: 2021-09-24T12:28:40Z
base_branch: master
head_branch: advisories
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/653
---

# Pull Request #653: feat(AdvisoriesReport): SPM-1167 add most impactfull advisory info

**Author**: @mkholjuraev
**Created**: September 20, 2021 at 11:14 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `advisories`

## Description

I have a few suggestions already added into the UI design: 
1. Move advisory name into the title position of the Patternfly card
2. move the number of applicable systems into a separate block GridItem  so that it is not split into 2 lines
3. Move os name to first column Grid item elements so that we have a balanced number of info labels into 2 sides.

![image](https://user-images.githubusercontent.com/59481011/133993936-3708f7e2-529e-4348-823b-03948ec8a01e.png)


I have not put a loader, because I am not sure if we have always more than 4 advisories in the database. If we always have and we know what will be shown, the loader can be put always

---

## Discussion

### Comment by @codecov-commenter on September 20, 2021 at 11:21 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#653](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e9b061e) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/59efb7789999249106e13fc72b273c558103b239?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (59efb77) will **increase** coverage by `0.11%`.
> The diff coverage is `53.33%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #653      +/-   ##
==========================================
+ Coverage   52.10%   52.22%   +0.11%     
==========================================
  Files          81       82       +1     
  Lines        1967     1980      +13     
  Branches      436      441       +5     
==========================================
+ Hits         1025     1034       +9     
- Misses        857      861       +4     
  Partials       85       85              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `88.00% <ø> (+2.00%)` | :arrow_up: |
| [...Components/StatusReports/AdvisoriesStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL0Fkdmlzb3JpZXNTdGF0dXNSZXBvcnQuanM=) | `61.53% <61.53%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [59efb77...e9b061e](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/653?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on September 20, 2021 at 12:20 PM UTC

Suggestions are approved UX!

### Comment by @jiridostal on September 24, 2021 at 12:44 PM UTC

:tada: This PR is included in version 1.33.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.33.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.33.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/653*
