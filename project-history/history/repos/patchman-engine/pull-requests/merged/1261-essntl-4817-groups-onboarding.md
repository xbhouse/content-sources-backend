---
type: pull_request
number: 1261
title: "ESSNTL-4817: Groups onboarding"
state: merged
author: psegedy
created: 2023-07-11T17:58:41Z
updated: 2023-07-14T11:39:37Z
closed: 2023-07-14T11:39:37Z
merged: 2023-07-14T11:39:37Z
base_branch: master
head_branch: inventory_groups
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1261
---

# Pull Request #1261: ESSNTL-4817: Groups onboarding

**Author**: @psegedy
**Created**: July 11, 2023 at 05:58 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `inventory_groups`

## Description

Implemented by parsing groups from rbac in the following structure
```go
var groups = map[string]string{
    rbac.KeyGrouped: '{"[\"id\": \"group_1\"]", "[\"id\": \"group_2\"]", ...}',
    rbac.KeyUngrouped: '[]' // if `nil` is in rbac response
}
```
the `groups` map is then set to `gin.Context`, retrieved from the context by each handler and the following query is added based on the content of the structure above

1. no rbac.KeyGrouped value
```sql
JOIN inventory.hosts ih ON ih.id = sp.inventory_id
```
2. rbac.KeyGrouped exists, rbac.KeyUngrouped is empty
```sql
JOIN inventory.hosts ih ON ih.id = sp.inventory_id
WHERE ih.groups @> ANY ('groups[rbac.KeyGrouped]'::jsonb[])
```
3. both rbac.KeyGrouped and rbac.KeyUngrouped are not empty
```sql
JOIN inventory.hosts ih ON ih.id = sp.inventory_id
WHERE ih.groups @> ANY ('groups[rbac.KeyGrouped]'::jsonb[]) OR ih.groups = '[]'
```

The only missing part is the API filter for group_name
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

### Comment by @jira-linking on July 11, 2023 at 05:58 PM UTC

Commits missing Jira IDs:
2b5c88cbda72f9b344b253a31a89e1dace93f58f
c5eecd23d1aaec530114f7d0cc81221c32cb813e
820544acc83331e158b15210267892ee83dc8840
62f9fdd7704ed1ea96a8a9f064dbf1d09d3684c6
64f75b4dbd383fe64acec2c8a9e097a5fe63cd69
38db27380270b3c6bf5844e6e18d2eee9c988292


### Comment by @codecov-commenter on July 11, 2023 at 06:07 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`84.54`**% and project coverage change: **`+0.18`** :tada:
> Comparison is base [(`c6dca18`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c6dca18aa76c0ebf975e231e024507e9b11818b6?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.34% compared to head [(`38db273`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.52%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1261      +/-   ##
==========================================
+ Coverage   61.34%   61.52%   +0.18%     
==========================================
  Files         106      106              
  Lines        6540     6594      +54     
==========================================
+ Hits         4012     4057      +45     
- Misses       2008     2014       +6     
- Partials      520      523       +3     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.52% <84.54%> (+0.18%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `77.94% <0.00%> (-0.47%)` | :arrow_down: |
| [manager/controllers/advisories\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `46.34% <50.00%> (+1.34%)` | :arrow_up: |
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `6.08% <58.33%> (+6.08%)` | :arrow_up: |
| [manager/middlewares/rbac.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | `77.14% <70.83%> (+0.54%)` | :arrow_up: |
| [manager/controllers/baseline\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `50.96% <80.00%> (+0.46%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `74.54% <100.00%> (+0.47%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `42.85% <100.00%> (+0.75%)` | :arrow_up: |
| [manager/controllers/advisory\_systems\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `41.86% <100.00%> (+1.38%)` | :arrow_up: |
| [manager/controllers/baseline\_systems\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `55.00% <100.00%> (+2.36%)` | :arrow_up: |
| [manager/controllers/baselines.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZXMuZ28=) | `83.78% <100.00%> (-0.22%)` | :arrow_down: |
| ... and [14 more](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1261?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Commented on July 14, 2023 at 08:34 AM UTC

### Review by @MichaelMraka - Approved on July 14, 2023 at 10:53 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1261*
