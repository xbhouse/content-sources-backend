---
type: pull_request
number: 1141
title: "fix(InventoryTable): RHINENG-2912 - Remove custom OS filter handling"
state: merged
author: bastilian
created: 2023-11-03T11:43:34Z
updated: 2023-11-13T12:50:45Z
closed: 2023-11-13T12:30:00Z
merged: 2023-11-13T12:30:00Z
base_branch: master
head_branch: RHINENG-2912
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1141
---

# Pull Request #1141: fix(InventoryTable): RHINENG-2912 - Remove custom OS filter handling

**Author**: @bastilian
**Created**: November 03, 2023 at 11:43 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `RHINENG-2912`

## Description

# Description

Associated Jira ticket: #RHINENG-2912

Removes the custom handling for the OS filter for InventoryTables

# Before the change

https://github.com/RedHatInsights/patchman-ui/assets/7757/2c301da9-6150-41df-8690-42022b7de260

# After the change

https://github.com/RedHatInsights/patchman-ui/assets/7757/1a348481-62b6-4c56-8d4e-15933b5590b8


---

## Discussion

### Comment by @codecov-commenter on November 03, 2023 at 11:52 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1141?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `4 lines` in your changes are missing coverage. Please review.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1141?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1141?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `82.35% <ø> (-0.34%)` | :arrow_down: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1141?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `90.74% <ø> (-0.17%)` | :arrow_down: |
| [src/SmartComponents/Systems/Systems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1141?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `80.30% <100.00%> (-0.30%)` | :arrow_down: |
| [src/Utilities/SystemsHelpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1141?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | `72.97% <100.00%> (ø)` | |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1141?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/Hooks.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1141?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `73.33% <50.00%> (+0.96%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1141?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `82.50% <83.33%> (+0.44%)` | :arrow_up: |


:loudspeaker: Thoughts on this report? [Let us know!](https://about.codecov.io/pull-request-comment-report/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


### Comment by @bastilian on November 07, 2023 at 11:59 AM UTC

@mkholjuraev Thank you for reviewing and testing it! The filtering should now work as expected again.

### Comment by @mkholjuraev on November 13, 2023 at 12:50 PM UTC

:tada: This PR is included in version 1.65.1 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.65.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Changes Requested on November 06, 2023 at 03:58 PM UTC

Codewise looks clean. While testing, I found out that appliying rhel 8.8 does not filter the table correctly. I see that there are rhel 8.8 systems, but applying the filter results in an empty result.  It seems filtering syntax got affected, it has incorrect value : **in:in:RHEL 8.8**

### Review by @mkholjuraev - Approved on November 08, 2023 at 03:29 PM UTC

LGTM! I see that the new function is covered with tests. We can enhance it more by writing tests that ensure that actvie OS filter indeed goes to the API request as expected.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1141*
