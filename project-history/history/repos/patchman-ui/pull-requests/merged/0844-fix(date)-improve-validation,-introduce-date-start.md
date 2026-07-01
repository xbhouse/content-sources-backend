---
type: pull_request
number: 844
title: "fix(Date): improve validation, introduce date start limit"
state: merged
author: mkholjuraev
created: 2022-07-07T10:16:56Z
updated: 2024-04-03T09:23:16Z
closed: 2022-07-07T21:28:29Z
merged: 2022-07-07T21:28:29Z
base_branch: master
head_branch: fix-to-date
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/844
---

# Pull Request #844: fix(Date): improve validation, introduce date start limit

**Author**: @mkholjuraev
**Created**: July 07, 2022 at 10:16 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-to-date`

## Description

This is to improve date field validation and introduce start limit.

To check locally:

- make sure that the date input does not let users select a date prior to 1990.01.01
- make sure that input field validates: wrong date, date input format


---

## Discussion

### Comment by @mkholjuraev on July 07, 2022 at 09:28 PM UTC

@adonispuente thank you!

### Comment by @jiridostal on July 07, 2022 at 09:45 PM UTC

:tada: This PR is included in version 1.48.6 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.6)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.6)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on July 07, 2022 at 09:06 PM UTC

This change works great, LGTM!
I didn't know the comparison operators '>' would work with dates like that 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/844*
