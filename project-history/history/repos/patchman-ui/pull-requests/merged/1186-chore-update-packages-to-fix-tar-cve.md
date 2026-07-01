---
type: pull_request
number: 1186
title: "chore: Update packages to fix tar CVE"
state: merged
author: leSamo
created: 2024-07-01T11:11:54Z
updated: 2026-04-03T10:43:55Z
closed: 2024-07-09T10:13:40Z
merged: 2024-07-09T10:13:40Z
base_branch: master
head_branch: fix-cve-tar
labels: ["dependencies", "released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1186
---

# Pull Request #1186: chore: Update packages to fix tar CVE

**Author**: @leSamo
**Created**: July 01, 2024 at 11:11 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `released`
**Base**: `master` ← **Head**: `fix-cve-tar`

## Description

https://issues.redhat.com/browse/RHINENG-10828

`tar` versions before 6.2.1 are vulnerable

### Before:
@redhat-cloud-services/frontend-components-inventory-patchman@1.63.5 /home/soleksak/patchman-ui
├─┬ @semantic-release/npm@9.0.2
│ └─┬ npm@8.19.4
│   ├─┬ cacache@16.1.3
│   │ └── tar@6.1.11 deduped
│   ├─┬ libnpmdiff@4.0.5
│   │ └── tar@6.1.11 deduped
│   ├─┬ node-gyp@9.1.0
│   │ └── tar@6.1.11 deduped
│   ├─┬ pacote@13.6.2
│   │ └── tar@6.1.11 deduped
│   └── tar@6.1.11
└─┬ node-sass@7.0.3
  └─┬ node-gyp@8.4.1
    ├─┬ make-fetch-happen@9.1.0
    │ └─┬ cacache@15.3.0
    │   └── tar@6.2.1 deduped
    └── tar@6.2.1


### After:
@redhat-cloud-services/frontend-components-inventory-patchman@1.63.5 /home/soleksak/patchman-ui
├─┬ @semantic-release/npm@11.0.3
│ └─┬ npm@10.8.1
│   ├─┬ cacache@18.0.3
│   │ └── tar@6.2.1 deduped
│   ├─┬ libnpmdiff@6.1.3
│   │ └── tar@6.2.1 deduped
│   ├─┬ node-gyp@10.1.0
│   │ └── tar@6.2.1 deduped
│   ├─┬ pacote@18.0.6
│   │ └── tar@6.2.1 deduped
│   └── tar@6.2.1
└─┬ node-sass@7.0.3
  └─┬ node-gyp@8.4.1
    ├─┬ make-fetch-happen@9.1.0
    │ └─┬ cacache@15.3.0
    │   └── tar@6.2.1 deduped
    └── tar@6.2.1

---

## Discussion

### Comment by @codecov-commenter on July 01, 2024 at 11:25 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1186?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 64.93%. Comparing base ([`bd177e8`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/bd177e867c4bf311837d57cad6c4b7d3b6a1a09d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`44ab6fb`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/44ab6fbd4ff037daaf6cf53e24ef826f3bc5f0af?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 285 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1186      +/-   ##
==========================================
+ Coverage   64.14%   64.93%   +0.79%     
==========================================
  Files         124      124              
  Lines        3207     3411     +204     
  Branches      818      900      +82     
==========================================
+ Hits         2057     2215     +158     
- Misses       1150     1196      +46     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1186?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @leSamo on July 04, 2024 at 03:09 PM UTC

/retest

### Comment by @mkholjuraev on July 24, 2024 at 12:40 PM UTC

:tada: This PR is included in version 1.67.8 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.8)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Approved on July 09, 2024 at 10:13 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1186*
