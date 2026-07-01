---
type: pull_request
number: 1061
title: "SPM-1705: Refactor vmaas-sync recalc event"
state: merged
author: michalslomczynski
created: 2022-08-17T09:16:56Z
updated: 2022-08-29T08:46:43Z
closed: 2022-08-29T08:46:43Z
merged: 2022-08-29T08:46:43Z
base_branch: master
head_branch: vmaas-sync-orgid
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1061
---

# Pull Request #1061: SPM-1705: Refactor vmaas-sync recalc event

**Author**: @michalslomczynski
**Created**: August 17, 2022 at 09:16 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `vmaas-sync-orgid`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [x] Output Encoding
- [x] Authentication and Password Management
- [x] Session Management
- [x] Access Control
- [x] Cryptographic Practices
- [x] Error Handling and Logging
- [x] Data Protection
- [x] Communication Security
- [x] System Configuration
- [x] Database Security
- [x] File Management
- [x] Memory Management
- [x] General Coding Practices


---

## Discussion

### Comment by @michalslomczynski on August 17, 2022 at 10:50 AM UTC

Regarding refactor:
- Not a strong preference against previous variable names
- Except very strong preference to get rid of this `groupedData` type and `data` variable names.

### Comment by @psegedy on August 17, 2022 at 05:30 PM UTC

I haven't had time to look at it properly. Wouldn't it be possible to completely remove everything related to `accountInventories` and use only `EvalDataSlice` struct?

### Comment by @codecov-commenter on August 18, 2022 at 08:02 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **61.41**% // Head: **62.37**% // Increases project coverage by **`+0.95%`** :tada:
> Coverage data is based on head [(`d2c4d1a`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`4b2ead3`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4b2ead3ccb5a0dbdaf79148ff0e66187eb017e52?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 93.54% of modified lines in pull request are covered.

> :exclamation: Current head d2c4d1a differs from pull request most recent head daefce9. Consider uploading reports for the commit daefce9 to get more accurate results

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1061      +/-   ##
==========================================
+ Coverage   61.41%   62.37%   +0.95%     
==========================================
  Files          98       98              
  Lines        5401     5379      -22     
==========================================
+ Hits         3317     3355      +38     
+ Misses       1658     1596      -62     
- Partials      426      428       +2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.37% <93.54%> (+0.95%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `10.12% <0.00%> (-0.09%)` | :arrow_down: |
| [base/mqueue/platform\_event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvcGxhdGZvcm1fZXZlbnQuZ28=) | `78.57% <100.00%> (+69.74%)` | :arrow_up: |
| [manager/kafka/kafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9rYWZrYS9rYWZrYS5nbw==) | `68.88% <100.00%> (ø)` | |
| [tasks/vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `70.76% <100.00%> (+0.45%)` | :arrow_up: |
| [tasks/vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `59.25% <100.00%> (+24.64%)` | :arrow_up: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `62.50% <0.00%> (+12.50%)` | :arrow_up: |
| [base/mqueue/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvdGVzdGluZy5nbw==) | `80.00% <0.00%> (+20.00%)` | :arrow_up: |

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1061?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @michalslomczynski on August 18, 2022 at 08:07 AM UTC

@psegedy I replaced it in last commit.

~~I personally prefer `InventoryAccount` over `EvalData` since you do not know what `EvalData` is without checking it's implementation. It is not improved readability if you have to do this IMO. `InventoryAccount` name is more specific. This struct just holds the account of the inventory, and the name suggest exactly that, simple.~~

Reverted back to `EvalDataSlice` to make this PR easier to understand. It was not strong opinion about this one.

### Comment by @MichaelMraka on August 18, 2022 at 08:48 AM UTC

@michalslomczynski, I have hard times to understand this PR. Can you please split it into
- necessary orgId changes,
- independent refactoring PR? 

### Comment by @michalslomczynski on August 18, 2022 at 08:54 AM UTC

Necessary OrgID changes:
[SPM-1705: Add OrgID field to eval recalc messages](https://github.com/RedHatInsights/patchman-engine/pull/1061/commits/7e3101f70d7dd71b82b3ed349f01a640da4c9d78)

### Comment by @michalslomczynski on August 18, 2022 at 09:22 AM UTC

To sum up refactor in case it's hard to track what was changed:
- Removed improper use of Go types
- Removed improper use of pointers
~~- Renamed inventoryAID/inventoryAIDs -> InventoryAccount/InventoryAccounts~~ Reverted this in favor of `EvalDataSlice` Patrick suggested
- Removed one of `MessageData` interface implementation because it was redundant.


### Comment by @michalslomczynski on August 19, 2022 at 09:44 AM UTC

@psegedy I am fine with EvalDataSlice and EvalData names, but we have this weird old variables all over the place, eg. `inventoryAIDs = EvalDataSlice{...}`. I am thinking if we are fine with it.

### Comment by @psegedy on August 22, 2022 at 05:05 PM UTC

@michalslomczynski not a big deal IMO and it would require renaming also some functions. Maybe something like `inventoryAccounts` would be more appropriate now. It used to contain only inventoryID and accountID, but now it also has orgID. `inventoryAIDs` might not be very descriptive, I guess it should stand for `inventory account ids` which is technically still true with orgIDs, I'd let it be as it is.

---

## Reviews

### Review by @psegedy - Commented on August 17, 2022 at 05:38 PM UTC

### Review by @michalslomczynski - Commented on August 18, 2022 at 04:51 AM UTC

### Review by @michalslomczynski - Commented on August 18, 2022 at 05:00 AM UTC

### Review by @michalslomczynski - Commented on August 18, 2022 at 05:15 AM UTC

### Review by @michalslomczynski - Commented on August 18, 2022 at 05:33 AM UTC

### Review by @michalslomczynski - Commented on August 18, 2022 at 07:39 AM UTC

### Review by @psegedy - Commented on August 18, 2022 at 07:54 AM UTC

### Review by @psegedy - Commented on August 18, 2022 at 07:58 AM UTC

### Review by @michalslomczynski - Commented on August 18, 2022 at 08:10 AM UTC

### Review by @MichaelMraka - Commented on August 18, 2022 at 08:36 AM UTC

### Review by @michalslomczynski - Commented on August 19, 2022 at 09:42 AM UTC

### Review by @MichaelMraka - Approved on August 24, 2022 at 12:55 PM UTC

### Review by @psegedy - Approved on August 26, 2022 at 10:37 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1061*
