---
type: pull_request
number: 1128
title: "Show \"Managed by Satellite\" in place of Template name"
state: merged
author: leSamo
created: 2023-10-10T16:59:33Z
updated: 2024-01-08T13:49:43Z
closed: 2023-10-11T12:08:20Z
merged: 2023-10-11T12:08:20Z
base_branch: master
head_branch: template-satellite-text
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1128
---

# Pull Request #1128: Show "Managed by Satellite" in place of Template name

**Author**: @leSamo
**Created**: October 10, 2023 at 04:59 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `template-satellite-text`

## Description

https://issues.redhat.com/browse/RHINENG-2485

How to test:

- Make sure "Managed by Satellite" text with tooltip according to the mockup is inside Template column for systems with "satellite_managed: true" attribute in the following tables
  - Systems list
  - Advisory system list
  - Package system list
  - Systems table inside Template Wizard > Apply to systems step


---

## Discussion

### Comment by @codecov-commenter on October 11, 2023 at 09:35 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1128?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `18 lines` in your changes are missing coverage. Please review.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1128?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1128?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `59.79% <0.00%> (-0.63%)` | :arrow_down: |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1128?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `52.17% <5.88%> (-8.36%)` | :arrow_down: |


:loudspeaker: Thoughts on this report? [Let us know!](https://about.codecov.io/pull-request-comment-report/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @niyazRedhat on October 11, 2023 at 11:28 AM UTC

/retest

### Comment by @mkholjuraev on October 30, 2023 at 12:10 PM UTC

:tada: This PR is included in version 1.64.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.64.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on October 11, 2023 at 09:43 AM UTC

LGTM! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1128*
