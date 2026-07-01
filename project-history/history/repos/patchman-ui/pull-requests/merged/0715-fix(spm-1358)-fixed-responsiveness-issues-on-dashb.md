---
type: pull_request
number: 715
title: "fix(SPM-1358): Fixed responsiveness issues on dashbars"
state: merged
author: AsToNlele
created: 2022-02-11T13:59:03Z
updated: 2022-03-01T13:47:19Z
closed: 2022-03-01T08:45:51Z
merged: 2022-03-01T08:45:51Z
base_branch: master
head_branch: SPM-1358
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/715
---

# Pull Request #715: fix(SPM-1358): Fixed responsiveness issues on dashbars

**Author**: @AsToNlele
**Created**: February 11, 2022 at 01:59 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `SPM-1358`

## Description

Advisories:
- Before
![1-before](https://user-images.githubusercontent.com/20592948/153603894-664e2fb1-c387-4d8a-8316-ad6f615dfb3b.png)

- After
![1-after](https://user-images.githubusercontent.com/20592948/153604024-a0a341f4-ab50-4481-9b23-742997164537.png)

Systems:
- Before
![2-before](https://user-images.githubusercontent.com/20592948/153604167-f6bf512a-9526-4b7d-9c42-d021c72292ea.png)
- After
![2-after](https://user-images.githubusercontent.com/20592948/153604229-dbdd8dfe-e145-4c30-8a93-d743fe8a9aef.png)



---

## Discussion

### Comment by @mkholjuraev on February 16, 2022 at 06:58 PM UTC

@AsToNlele Can you please modify commit names so that they follow semantic release conventions. You can have a look at previos commit names. We use something like:
1. 'feat(AFFECTED_COMPONENT_NAME): your commit message' for new features,
2. 'fix(AFFECTED_COMPONENT_NAME): your commit message' for a bug fix.
3. 'chore(AFFECTED_COMPONENT_NAME): your commit message' for commits not intended to trigger semantic release build. For example a minor changes, cleanups... 

### Comment by @AsToNlele on February 17, 2022 at 09:41 AM UTC

@mkholjuraev Sorry about that! Are the commit names correct now?

### Comment by @codecov-commenter on February 17, 2022 at 09:47 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/715?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#715](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/715?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (bf99e8b) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/5fa217c597df2941d2526eb3e4f42832f31396b1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5fa217c) will **not change** coverage.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/715/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/715?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #715   +/-   ##
=======================================
  Coverage   81.87%   81.87%           
=======================================
  Files          84       84           
  Lines        1925     1925           
  Branches      499      499           
=======================================
  Hits         1576     1576           
  Misses        306      306           
  Partials       43       43           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/715?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...Components/StatusReports/AdvisoriesStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/715/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL0Fkdmlzb3JpZXNTdGF0dXNSZXBvcnQuanM=) | `53.33% <0.00%> (ø)` | |
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/715/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | `77.77% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/715?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/715?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [5fa217c...bf99e8b](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/715?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on February 21, 2022 at 10:29 AM UTC

@AsToNlele the names look good. please merge the PR it looks good to go.

### Comment by @jiridostal on March 01, 2022 at 01:47 PM UTC

:tada: This PR is included in version 1.40.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.40.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.40.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on February 15, 2022 at 05:30 AM UTC

Nicely done! Thank you @AsToNlele!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/715*
