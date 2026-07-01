---
type: pull_request
number: 1142
title: "Do not allow template application for Satellite managed systems"
state: merged
author: leSamo
created: 2023-11-06T13:45:11Z
updated: 2023-11-09T11:01:01Z
closed: 2023-11-07T11:52:53Z
merged: 2023-11-07T11:52:53Z
base_branch: master
head_branch: 2489-satellite
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1142
---

# Pull Request #1142: Do not allow template application for Satellite managed systems

**Author**: @leSamo
**Created**: November 06, 2023 at 01:45 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `2489-satellite`

## Description

Ticket: https://issues.redhat.com/browse/RHINENG-2489
Mockups: https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/P19w2K1


- Disable row action kebab items related to templates on System list table, but do not disable row checkbox (since Remediation is allowed)
- Disable row action kebab items related to templates in System detail header
- Update Template assignment and removal modals to handle some Satellite managed systems being selected.
- Make sure Systems table in Template wizard contains only Systems not managed by Satellite


---

## Discussion

### Comment by @codecov-commenter on November 06, 2023 at 01:47 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1142?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `16 lines` in your changes are missing coverage. Please review.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1142?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Modals/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1142?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvSGVscGVycy5qcw==) | `100.00% <100.00%> (ø)` | |
| [src/SmartComponents/Modals/UnassignSystemsModal.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1142?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvVW5hc3NpZ25TeXN0ZW1zTW9kYWwuanM=) | `96.15% <100.00%> (+0.32%)` | :arrow_up: |
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1142?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `44.44% <ø> (ø)` | |
| [...c/SmartComponents/Modals/useUnassignSystemsHook.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1142?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvdXNlVW5hc3NpZ25TeXN0ZW1zSG9vay5qcw==) | `81.81% <33.33%> (-18.19%)` | :arrow_down: |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1142?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `4.54% <0.00%> (ø)` | |
| [src/SmartComponents/Modals/AssignSystemsModal.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1142?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvQXNzaWduU3lzdGVtc01vZGFsLmpz) | `51.11% <52.00%> (+3.49%)` | :arrow_up: |

... and [2 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1142/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


:loudspeaker: Thoughts on this report? [Let us know!](https://about.codecov.io/pull-request-comment-report/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on November 09, 2023 at 11:00 AM UTC

:tada: This PR is included in version 1.65.0 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.65.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on November 07, 2023 at 11:12 AM UTC

LGTM! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1142*
