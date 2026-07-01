---
type: pull_request
number: 433
title: "bug(App.js): change useEffect dependency to location.pathname"
state: merged
author: mkholjuraev
created: 2021-02-11T09:59:00Z
updated: 2026-04-03T11:46:42Z
closed: 2021-02-11T10:30:24Z
merged: 2021-02-11T10:30:24Z
base_branch: master
head_branch: useEffect-dependancy-change
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/433
---

# Pull Request #433: bug(App.js): change useEffect dependency to location.pathname

**Author**: @mkholjuraev
**Created**: February 11, 2021 at 09:59 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `useEffect-dependancy-change`

## Description

Resolves: https://ansible.slack.com/archives/G01NG0Y5RPA/p1613034742012100

---

## Discussion

### Comment by @codecov-io on February 11, 2021 at 10:02 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/433?src=pr&el=h1) Report
> Merging [#433](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/433?src=pr&el=desc) (1fd3e3b) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/714982c751949b5d6d0383017b73e5417c7a9b84?el=desc) (714982c) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/433/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/433?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #433   +/-   ##
=======================================
  Coverage   69.35%   69.35%           
=======================================
  Files          75       75           
  Lines        1315     1315           
  Branches      175      175           
=======================================
  Hits          912      912           
  Misses        337      337           
  Partials       66       66           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/433?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/433/diff?src=pr&el=tree#diff-c3JjL0FwcC5qcw==) | `0.00% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/433?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/433?src=pr&el=footer). Last update [714982c...1fd3e3b](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/433?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on February 19, 2021 at 10:25 AM UTC

:tada: This PR is included in version 1.11.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.11.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.11.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

### Comment by @codecov-commenter on April 03, 2026 at 11:46 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/433?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 69.35%. Comparing base ([`714982c`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/714982c751949b5d6d0383017b73e5417c7a9b84?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`1fd3e3b`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1fd3e3b0755c64664d5baf5df0983e721d39d15e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1053 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master     #433   +/-   ##
=======================================
  Coverage   69.35%   69.35%           
=======================================
  Files          75       75           
  Lines        1315     1315           
  Branches      175      175           
=======================================
  Hits          912      912           
  Misses        337      337           
  Partials       66       66           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/433?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/433*
