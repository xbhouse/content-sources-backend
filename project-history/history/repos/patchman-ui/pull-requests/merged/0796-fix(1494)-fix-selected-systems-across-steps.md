---
type: pull_request
number: 796
title: "fix(1494): fix selected systems across steps"
state: merged
author: mkholjuraev
created: 2022-05-12T08:40:19Z
updated: 2024-04-03T09:21:43Z
closed: 2022-05-16T14:56:32Z
merged: 2022-05-16T14:56:32Z
base_branch: master
head_branch: rename-patch-set-title-wizard
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/796
---

# Pull Request #796: fix(1494): fix selected systems across steps

**Author**: @mkholjuraev
**Created**: May 12, 2022 at 08:40 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `rename-patch-set-title-wizard`

## Description

This resolves the ticket: https://issues.redhat.com/browse/SPM-1494

To reproduce:

1. Create new patch set using wizard
2. On 2nd page select e.g. 3 systems, deselect 2 of them so 1 is selected.
3. Go to Final page of the wizard, the wizard shows 1 system in patch set
4. Go back to system selection, 3 systems are selected although no action was taken.
5. Go next to final page of the wizard, it shows 3 systems now although only actions taken were going to previous (2nd page) and immediately to 3rd page.

---

## Discussion

### Comment by @codecov-commenter on May 12, 2022 at 08:46 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/796?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#796](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/796?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c46fda0) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/03ef9f193bf4ec4e2dc500ed03b32e10bf967fec?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (03ef9f1) will **decrease** coverage by `0.22%`.
> The diff coverage is `13.33%`.

```diff
@@            Coverage Diff             @@
##           master     #796      +/-   ##
==========================================
- Coverage   70.83%   70.60%   -0.23%     
==========================================
  Files         101      101              
  Lines        2393     2402       +9     
  Branches      618      619       +1     
==========================================
+ Hits         1695     1696       +1     
- Misses        631      638       +7     
- Partials       67       68       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/796?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/796/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `43.75% <0.00%> (-2.92%)` | :arrow_down: |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/796/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `4.54% <0.00%> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/796/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `73.66% <20.00%> (-1.68%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/796?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/796?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [03ef9f1...c46fda0](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/796?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on May 16, 2022 at 02:56 PM UTC

@adonispuente sounds great!

### Comment by @jiridostal on May 16, 2022 at 03:12 PM UTC

:tada: This PR is included in version 1.47.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.47.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.47.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on May 16, 2022 at 02:14 PM UTC

I think this fix works great, I don't see any issues!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/796*
