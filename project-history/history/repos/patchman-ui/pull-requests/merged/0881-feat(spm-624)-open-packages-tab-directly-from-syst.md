---
type: pull_request
number: 881
title: "feat(SPM-624): open packages tab directly from systes page."
state: merged
author: mkholjuraev
created: 2022-10-03T08:25:30Z
updated: 2024-04-03T09:21:57Z
closed: 2022-11-10T12:50:51Z
merged: 2022-11-10T12:50:51Z
base_branch: master
head_branch: advisor-package-tab
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/881
---

# Pull Request #881: feat(SPM-624): open packages tab directly from systes page.

**Author**: @mkholjuraev
**Created**: October 03, 2022 at 08:25 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `advisor-package-tab`

## Description

This enables to open the Packages tab directly by clicking the package number on the Systems page.

To test:
1. Visit the Systems page,
2. click on a package number in the table
3. Observe that the packages tab is opened by default.
4. In any other case, the advisor tab is active by default.
5. When systems detail page is opened using URL, advisor tab should be active, page should not break.

---

## Discussion

### Comment by @mkholjuraev on October 06, 2022 at 07:45 PM UTC

@leSamo thank you for the review. I have fixed the regression hopefully.

### Comment by @mkholjuraev on November 04, 2022 at 12:32 PM UTC

@leSamo another review, please?


### Comment by @mkholjuraev on November 09, 2022 at 12:08 PM UTC

@RedHatInsights/team-interact it would be great if this gets an attention

### Comment by @mkholjuraev on November 10, 2022 at 12:42 PM UTC

I have reviewed the PR, looks like there are no regressions. Merging, ask the QA to verify in the associated ticket.

### Comment by @jiridostal on November 10, 2022 at 01:08 PM UTC

:tada: This PR is included in version 1.55.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.55.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.55.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Changes Requested on October 06, 2022 at 09:48 AM UTC

Hmmm point of this PR works well, but it looks like Advisories table has a regression:
1. Go to System detail page, Advisories tab
2. Navigate to the next page in table
3. Notice the current page did not get changed and `page=1&page_size=20` parameters in the URL are being replicated.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/881*
