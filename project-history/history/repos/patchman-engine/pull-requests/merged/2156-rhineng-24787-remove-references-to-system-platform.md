---
type: pull_request
number: 2156
title: "RHINENG-24787: remove references to system platform"
state: merged
author: TenSt
created: 2026-04-15T13:39:53Z
updated: 2026-04-16T10:20:50Z
closed: 2026-04-16T10:20:50Z
merged: 2026-04-16T10:20:50Z
base_branch: master
head_branch: stepan/RHINENG-24787-remove-references-to-system-platform
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2156
---

# Pull Request #2156: RHINENG-24787: remove references to system platform

**Author**: @TenSt
**Created**: April 15, 2026 at 01:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `stepan/RHINENG-24787-remove-references-to-system-platform`

## Description

This PR:
- updates services to persist and query host data on the physical system_inventory and system_patch tables (joined where a full “system” row is needed) 
- introduces SystemPlatformV2 as the in-memory aggregate (Inventory + Patch)
- removes the monolithic SystemPlatform model
- adds a dedicated JOIN loader (evaluator/system_platform_v2_load.go + tests)
- threads the new types through evaluation, listener upload, manager APIs, tasks, notifications, remediations, and test helpers
- updates VMaaS JSON parsing to take SystemInventory
- updates "Deploy / FloorPlan SQL" in deploy/clowdapp.yaml to use system_inventory and system_patch
- updates Grafana panel queries to track DB size for system_inventory_* and system_patch_* partitions
- updates Docs (docs/md/architecture.md, docs/md/database.md) to describe inventory vs patch and the optional VIEW.



---

## Discussion

### Comment by @github-actions on April 15, 2026 at 01:40 PM UTC

<!-- sc-environment-impact-check -->
## SC Environment Impact Assessment

**Overall Impact:** 🟠 **HIGH**

<details>
<summary>View full report</summary>

### Summary

- **Total Issues:** 2
- 🟠 High: 1
- 🟢 Low: 1

### Detailed Findings

#### 🟠 HIGH Impact

**ClowdApp configuration change**
- File: `deploy/clowdapp.yaml`
- Category: `clowdapp_config`
- **Recommendation:** Verify all resource limits, environment variables, and dependencies are compatible with SC Environment. Check for any public cloud-specific features.

#### 🟢 LOW Impact

**Environment configuration change detected**
- File: `evaluator/system_platform_v2_load_test.go`
- Category: `environment_config`
- Details:
  - Found `Environment` in `evaluator/system_platform_v2_load_test.go` at line 24
- **Recommendation:** Review environment-specific settings to ensure SC Environment is properly configured.

### Required Actions

- [ ] Review all findings above
- [ ] Verify SC Environment compatibility for all detected changes
- [ ] Update deployment documentation if needed
- [ ] Coordinate with ROSA Core team or deployment timeline

</details>

---
*This assessment was automatically generated. Please review carefully and consult with the ROSA Core team for critical/high impact changes.*

### Comment by @codecov-commenter on April 15, 2026 at 02:45 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `74.25968%` with `113 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 59.29%. Comparing base ([`b1f4464`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b1f4464b172a4c2a15416e0282ff0e1a609b2ba4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`c336470`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c336470bd25d9951453e89975e69d3c19c9f7a9d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/database/testing.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&filepath=base%2Fdatabase%2Ftesting.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | 0.00% | [44 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&filepath=evaluator%2Fevaluate.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | 79.74% | [10 Missing and 6 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&filepath=base%2Fdatabase%2Futils.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | 8.33% | [11 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/system\_detail.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&filepath=manager%2Fcontrollers%2Fsystem_detail.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | 9.09% | [10 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/models/models.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&filepath=base%2Fmodels%2Fmodels.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tb2RlbHMvbW9kZWxzLmdv) | 0.00% | [8 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/utils/vmaas.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&filepath=base%2Futils%2Fvmaas.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | 0.00% | [6 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&filepath=listener%2Fupload.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | 93.75% | [1 Missing and 3 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/notification/notification.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&filepath=base%2Fnotification%2Fnotification.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9ub3RpZmljYXRpb24vbm90aWZpY2F0aW9uLmdv) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&filepath=manager%2Fcontrollers%2Futils.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | 50.00% | [2 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate\_advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&filepath=evaluator%2Fevaluate_advisories.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | 86.66% | [0 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| ... and [4 more](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #2156      +/-   ##
==========================================
+ Coverage   59.21%   59.29%   +0.08%     
==========================================
  Files         134      135       +1     
  Lines        8637     8748     +111     
==========================================
+ Hits         5114     5187      +73     
- Misses       2988     3022      +34     
- Partials      535      539       +4     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.29% <74.25%> (+0.08%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2156?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @sourcery-ai - Commented on April 15, 2026 at 01:39 PM UTC

Sorry @TenSt, your pull request is larger than the review limit of 150000 diff characters

### Review by @sourcery-ai - Commented on April 15, 2026 at 03:30 PM UTC

Sorry @TenSt, your pull request is larger than the review limit of 150000 diff characters

### Review by @MichaelMraka - Commented on April 15, 2026 at 08:45 PM UTC

### Review by @MichaelMraka - Commented on April 15, 2026 at 09:11 PM UTC

### Review by @MichaelMraka - Commented on April 15, 2026 at 09:23 PM UTC

### Review by @MichaelMraka - Commented on April 15, 2026 at 09:38 PM UTC

### Review by @MichaelMraka - Commented on April 15, 2026 at 09:43 PM UTC

### Review by @MichaelMraka - Commented on April 15, 2026 at 09:45 PM UTC

### Review by @MichaelMraka - Commented on April 15, 2026 at 10:06 PM UTC

### Review by @TenSt - Commented on April 16, 2026 at 08:03 AM UTC

### Review by @TenSt - Commented on April 16, 2026 at 08:09 AM UTC

### Review by @TenSt - Commented on April 16, 2026 at 08:16 AM UTC

### Review by @TenSt - Commented on April 16, 2026 at 08:26 AM UTC

### Review by @MichaelMraka - Commented on April 16, 2026 at 08:29 AM UTC

### Review by @TenSt - Commented on April 16, 2026 at 08:44 AM UTC

### Review by @MichaelMraka - Commented on April 16, 2026 at 09:25 AM UTC

### Review by @MichaelMraka - Approved on April 16, 2026 at 09:29 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2156*
