---
type: pull_request
number: 798
title: "fix(SPM-1490, SPM-1493): wizard texts updated"
state: merged
author: mkholjuraev
created: 2022-05-12T11:54:51Z
updated: 2024-04-03T09:21:46Z
closed: 2022-05-17T14:27:31Z
merged: 2022-05-17T14:27:31Z
base_branch: master
head_branch: wizard-title
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/798
---

# Pull Request #798: fix(SPM-1490, SPM-1493): wizard texts updated

**Author**: @mkholjuraev
**Created**: May 12, 2022 at 11:54 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `wizard-title`

## Description

Resolves: 

1. [Editing a patch set created wizard named 'Create patch set'](https://issues.redhat.com/browse/SPM-1490). Wizard title is renamed to 'Edit patch set'
2. [Assigning system to a patch set created wizard named 'Create patch set'](https://issues.redhat.com/browse/SPM-1493) Wizard title is renamed to 'Assign system(s) to a patch set'.
3. ['Assign to patch sets' action should be 'Assign to a patch set' (singular form).](https://issues.redhat.com/browse/SPM-1489) (all occurrances in Systems page)

---

## Discussion

### Comment by @codecov-commenter on May 12, 2022 at 02:40 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/798?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#798](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/798?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (accd6e3) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/03ef9f193bf4ec4e2dc500ed03b32e10bf967fec?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (03ef9f1) will **decrease** coverage by `0.35%`.
> The diff coverage is `7.14%`.

```diff
@@            Coverage Diff             @@
##           master     #798      +/-   ##
==========================================
- Coverage   70.83%   70.47%   -0.36%     
==========================================
  Files         101      101              
  Lines        2393     2405      +12     
  Branches      618      623       +5     
==========================================
  Hits         1695     1695              
- Misses        631      641      +10     
- Partials       67       69       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/798?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/798/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <0.00%> (-0.31%)` | :arrow_down: |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/798/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `76.92% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/798/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `69.44% <ø> (ø)` | |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/798/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `26.92% <7.69%> (-19.75%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/798?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/798?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [03ef9f1...accd6e3](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/798?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @adonispuente on May 16, 2022 at 07:02 PM UTC

When pulling down the PR and running it , though I have the code changes (shown in the vs code on the left) the changes dont seem to have taken affect. 

https://user-images.githubusercontent.com/60629070/168662989-2aa5e474-a053-42a0-8f86-b1ce4b7e6f35.mp4


https://user-images.githubusercontent.com/60629070/168662994-bdaa2a40-cbe3-4cfd-9582-6d033b3801f3.mp4


https://user-images.githubusercontent.com/60629070/168662996-afd25961-424b-441e-9a0b-a8ae2b44929e.mp4



### Comment by @mkholjuraev on May 17, 2022 at 06:59 AM UTC

@adonispuente  maybe a hard refresh does the trick? I can see the effect of the PR when I run this. I am ready to jump to a short call if needed.

https://user-images.githubusercontent.com/59481011/168748833-ceef6910-2ef1-483d-9b78-90e97a582af0.mp4

 

### Comment by @adonispuente on May 17, 2022 at 01:15 PM UTC

Seems it was a problem on my local set up specifically, I redownloaded the repo and can see all changes. Everything looks good!

### Comment by @mkholjuraev on May 17, 2022 at 02:27 PM UTC

@adonispuente Awesome!

### Comment by @jiridostal on May 17, 2022 at 02:43 PM UTC

:tada: This PR is included in version 1.47.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.47.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.47.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on May 17, 2022 at 01:15 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/798*
