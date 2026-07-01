---
type: pull_request
number: 802
title: "fix(SPM-1487): let users create patch sets with zero systems assigned"
state: merged
author: mkholjuraev
created: 2022-05-17T12:48:39Z
updated: 2024-04-03T09:21:47Z
closed: 2022-05-26T06:15:54Z
merged: 2022-05-26T06:15:54Z
base_branch: master
head_branch: remove-systems-validation
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/802
---

# Pull Request #802: fix(SPM-1487): let users create patch sets with zero systems assigned

**Author**: @mkholjuraev
**Created**: May 17, 2022 at 12:48 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remove-systems-validation`

## Description

This resolves: 
https://issues.redhat.com/browse/SPM-1487.

Initially PM asked to prevent users creating patch sets with zero systems. However, during the latest meeting team we agreed to remove the revert this case and remove the validation. Now, wizard lets users go through the step 2 even if no system is selected.


to reproduce: 

1. Go to Patch set page: https://console.stage.redhat.com/beta/insights/patch/patch-set
2. Click 'Create Patch Set' button
3. Fill the name and date, click next.
4. System selection is displayed, 'Next' button is greyed out. There's no way how to complete Patch Set creation.



Similarly, this PR enables patch sets with zero systems assigned to be edited. This bug is tracked: https://issues.redhat.com/browse/SPM-1491. 
To reproduce SPM-1491:

1. Edit a patch set with zero systems
2. Change the e.g. the upto date
3. Go to next page, it's not possible to proceed further if no systems are associated with a patch set.



---

## Discussion

### Comment by @codecov-commenter on May 17, 2022 at 12:59 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/802?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#802](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/802?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (cdd53bd) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/f48fdf1dad3314150b6b9e72bceb22c180baa884?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f48fdf1) will **not change** coverage.
> The diff coverage is `n/a`.

> :exclamation: Current head cdd53bd differs from pull request most recent head d2cdfcc. Consider uploading reports for the commit d2cdfcc to get more accurate results

```diff
@@           Coverage Diff           @@
##           master     #802   +/-   ##
=======================================
  Coverage   70.60%   70.60%           
=======================================
  Files         101      101           
  Lines        2402     2402           
  Branches      619      619           
=======================================
  Hits         1696     1696           
  Misses        638      638           
  Partials       68       68           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/802?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/802/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `43.75% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/802?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/802?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [f48fdf1...d2cdfcc](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/802?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on May 26, 2022 at 05:58 AM UTC

@MichalSajdik thanks for the review. You are always helpful!

### Comment by @jiridostal on May 26, 2022 at 06:31 AM UTC

:tada: This PR is included in version 1.47.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.47.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.47.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @MichalSajdik - Commented on May 19, 2022 at 12:32 PM UTC

### Review by @mkholjuraev - Commented on May 23, 2022 at 03:35 PM UTC

### Review by @MichalSajdik - Approved on May 24, 2022 at 07:22 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/802*
