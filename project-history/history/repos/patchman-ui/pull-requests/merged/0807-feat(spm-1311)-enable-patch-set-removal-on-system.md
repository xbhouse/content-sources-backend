---
type: pull_request
number: 807
title: "feat(SPM-1311): enable patch set removal on System details page"
state: merged
author: mkholjuraev
created: 2022-05-31T12:05:14Z
updated: 2024-04-03T09:21:49Z
closed: 2022-06-06T13:20:47Z
merged: 2022-06-06T13:20:47Z
base_branch: master
head_branch: assign-upto-date-systems
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/807
---

# Pull Request #807: feat(SPM-1311): enable patch set removal on System details page

**Author**: @mkholjuraev
**Created**: May 31, 2022 at 12:05 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `assign-upto-date-systems`

## Description

Enables system removal from a patch set on system details page (https://issues.redhat.com/browse/SPM-1311).

To verify, visit a system details page

1.  when that system has a set assigned, you should see 'Remove from patch set' action enabled on the right corner dropdown.

- the action should open a modal with confirmation purpose
- system should get unassigned, when user  hits on 'Remove' button on the modal.
- page should refresh after successful removal

2. when that system has no set assigned, you should see  'Remove from patch set' action disabled on the right corner dropdown
 

---

## Discussion

### Comment by @codecov-commenter on June 01, 2022 at 02:23 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/807?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#807](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/807?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2e472f8) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/1f12520152242abaa4a103c76d60472842de1681?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1f12520) will **increase** coverage by `0.12%`.
> The diff coverage is `83.33%`.

```diff
@@            Coverage Diff             @@
##           master     #807      +/-   ##
==========================================
+ Coverage   70.69%   70.81%   +0.12%     
==========================================
  Files         102      102              
  Lines        2440     2440              
  Branches      628      628              
==========================================
+ Hits         1725     1728       +3     
+ Misses        647      644       -3     
  Partials       68       68              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/807?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/807/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `54.54% <50.00%> (ø)` | |
| [src/SmartComponents/Modals/UnassignSystemsModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/807/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvVW5hc3NpZ25TeXN0ZW1zTW9kYWwuanM=) | `100.00% <100.00%> (ø)` | |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/807/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `83.33% <100.00%> (+10.00%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/807?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/807?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [1f12520...2e472f8](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/807?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on June 03, 2022 at 10:24 AM UTC

@RedHatInsights/team-interact I would appreciate a look on this PR. 

### Comment by @jiridostal on June 06, 2022 at 01:42 PM UTC

:tada: This PR is included in version 1.48.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gabipodolnikova - Approved on June 06, 2022 at 01:17 PM UTC

Looking good! 👍  

### Review by @gkarat - Commented on June 06, 2022 at 01:21 PM UTC

### Review by @mkholjuraev - Commented on June 06, 2022 at 01:32 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/807*
