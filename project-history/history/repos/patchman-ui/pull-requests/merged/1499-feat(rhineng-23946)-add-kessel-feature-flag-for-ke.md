---
type: pull_request
number: 1499
title: "feat(RHINENG-23946): Add kessel feature flag for Kessel integration"
state: merged
author: mtclinton
created: 2026-02-11T02:00:13Z
updated: 2026-02-17T08:58:00Z
closed: 2026-02-16T18:18:26Z
merged: 2026-02-16T18:18:26Z
base_branch: master
head_branch: RHINENG-23946
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1499
---

# Pull Request #1499: feat(RHINENG-23946): Add kessel feature flag for Kessel integration

**Author**: @mtclinton
**Created**: February 11, 2026 at 02:00 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-23946`

## Description

Add useKesselFeatureFlag hook to gate Kessel authorization integration behind the patch-frontend.kessel-enabled feature flag.

This is part of the Kessel integration effort for patchman-ui. The feature flag allows us to incrementally roll out Kessel-based access checks without impacting existing functionality.

Ref: https://issues.redhat.com/browse/RHINENG-23946

# Description

Associated Jira ticket: # (issue)

Please include a summary of the change, what this fixes/creates/improves.


# How to test the PR

Please include steps to test your PR.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 11, 2026 at 02:07 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1499?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `3 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.50%. Comparing base ([`87ec4bd`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/87ec4bdb48b1c3c0e5e99802e10cc4341592d4c7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a42000d`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a42000dad8ced63bd65fdc2900bb184a742b6c78?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1499?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/hooks/useFeatureFlag.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1499?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FuseFeatureFlag.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VGZWF0dXJlRmxhZy5qcw==) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1499?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1499      +/-   ##
==========================================
- Coverage   62.53%   62.50%   -0.04%     
==========================================
  Files         127      127              
  Lines        3310     3312       +2     
  Branches      899      892       -7     
==========================================
  Hits         2070     2070              
- Misses       1120     1122       +2     
  Partials      120      120              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1499?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @rverdile on February 16, 2026 at 06:12 PM UTC

I assume this will be followed up https://issues.redhat.com/browse/RHINENG-23947 ? Feel free to reach out if you have any questions!

### Comment by @mtclinton on February 16, 2026 at 06:17 PM UTC

@rverdile exactly, yep I plan to work on that issue next week. Thanks for the review!


### Comment by @katarinazaprazna on February 17, 2026 at 08:57 AM UTC

Hi @mtclinton, thanks for the contribution! If you run into any questions regarding Kessel or the backend, feel free to give @rverdile or @MichaelMraka a shout. They’ll be happy to help :)

---

## Reviews

### Review by @rverdile - Approved on February 12, 2026 at 03:26 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1499*
