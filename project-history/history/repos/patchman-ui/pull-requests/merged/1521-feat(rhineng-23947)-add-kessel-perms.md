---
type: pull_request
number: 1521
title: "feat(RHINENG-23947): Add Kessel perms"
state: merged
author: mtclinton
created: 2026-02-23T19:02:54Z
updated: 2026-03-16T04:54:47Z
closed: 2026-03-16T04:54:47Z
merged: 2026-03-16T04:54:47Z
base_branch: master
head_branch: RHINENG-23947
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1521
---

# Pull Request #1521: feat(RHINENG-23947): Add Kessel perms

**Author**: @mtclinton
**Created**: February 23, 2026 at 07:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-23947`

## Description

Integrate Kessel permission checks behind the
patch-frontend.kessel-enabled feature flag.

When Kessel is enabled, permission checks use the
Kessel SDK (useSelfAccessCheck) against workspace
resources instead of RBAC v1. When disabled,
behavior is unchanged.

Permission mapping (from patch.ksl):
- patch:*:read    -> patch_system_view
- patch:*:*       -> patch_system_edit
- patch:template:write -> patch_template_edit

# Description

RHINENG-23947

Please include a summary of the change, what this fixes/creates/improves.
-Integrate Kessel permission checks behind the
patch-frontend.kessel-enabled feature flag.

- Add AccessCheck.Provider and QueryClientProvider
- Add usePermissionCheck hook toggling RBAC v1/Kessel
- Add useKesselWorkspaces using fetchDefaultWorkspace
  from @project-kessel/react-kessel-access-check SDK
- Use getKesselAccessCheckParams from FEC utilities
- Map patch RBAC v1 permissions to Kessel relations
- Replace usePermissionsWithContext in Routes,
  WithPermission, SystemsTable, PatchSet, and
  PatchSetDetail with usePermissionCheck
- Add useFeatureFlag and useKesselFeatureFlag hooks
- Add test mocks for Kessel SDK and permission hooks




# How to test the PR

Please include steps to test your PR.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [x] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 23, 2026 at 09:54 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `1.85185%` with `53 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 72.30%. Comparing base ([`f4a7c82`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/f4a7c8284aae975cae5809e77ca8ff8c849f2e50?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`4d783ae`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4d783ae9eec0b48c2bae8ceb1b48e1fb7a4529d5?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 5 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/hooks/usePermissionCheck.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FusePermissionCheck.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VQZXJtaXNzaW9uQ2hlY2suanM=) | 0.00% | [18 Missing and 6 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/hooks/useKesselWorkspaces.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FuseKesselWorkspaces.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VLZXNzZWxXb3Jrc3BhY2VzLmpz) | 0.00% | [15 Missing and 3 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...ationalComponents/WithPermission/WithPermission.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?src=pr&el=tree&filepath=src%2FPresentationalComponents%2FWithPermission%2FWithPermission.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9XaXRoUGVybWlzc2lvbi9XaXRoUGVybWlzc2lvbi5qcw==) | 0.00% | [2 Missing and 3 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/index.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?src=pr&el=tree&filepath=src%2Findex.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL2luZGV4Lmpz) | 0.00% | [3 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Routes.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?src=pr&el=tree&filepath=src%2FRoutes.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1521      +/-   ##
==========================================
- Coverage   73.67%   72.30%   -1.37%     
==========================================
  Files          97       99       +2     
  Lines        2359     2405      +46     
  Branches      666      677      +11     
==========================================
+ Hits         1738     1739       +1     
- Misses        551      586      +35     
- Partials       70       80      +10     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1521?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mtclinton on March 12, 2026 at 01:25 AM UTC

Thanks for the review @LightOfHeaven1994 , working on testing this now

### Comment by @mtclinton on March 12, 2026 at 08:09 PM UTC

I pushed the fix to 'useAccessCheckContext must be used within an AccessCheckProvider' that other services were experiencing,
guess i need another review due to IBM source controls noww 

---

## Reviews

### Review by @mtclinton - Commented on February 23, 2026 at 08:53 PM UTC

### Review by @LightOfHeaven1994 - Commented on March 11, 2026 at 09:58 AM UTC

Found few things to fix, otherwise I think it looks good. I wasn't able to test it properly because patchman backend times out

### Review by @LightOfHeaven1994 - Dismissed on March 12, 2026 at 08:57 AM UTC

PR looks good to me 👍 

I don't know why playwright tests are failing. Seems in reality this app checks only read permissions. I see there is `template:write` permission but it's not really used as patch is divided to two GH repositories. 

### Review by @TenSt - Dismissed on March 12, 2026 at 11:26 AM UTC

lgtm

### Review by @LightOfHeaven1994 - Approved on March 12, 2026 at 09:12 PM UTC

LGTM 👍 

### Review by @TenSt - Approved on March 13, 2026 at 02:18 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1521*
