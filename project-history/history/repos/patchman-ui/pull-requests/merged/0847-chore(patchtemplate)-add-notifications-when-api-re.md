---
type: pull_request
number: 847
title: "chore(PatchTemplate): add notifications when api req fails"
state: merged
author: mkholjuraev
created: 2022-07-20T06:44:27Z
updated: 2024-04-03T09:23:13Z
closed: 2022-07-21T11:07:19Z
merged: 2022-07-21T11:07:19Z
base_branch: master
head_branch: template-fail-notification
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/847
---

# Pull Request #847: chore(PatchTemplate): add notifications when api req fails

**Author**: @mkholjuraev
**Created**: July 20, 2022 at 06:44 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `template-fail-notification`

## Description

This adds an error notification when a user tries to create a patch template with a name that already exists.

to reproduce:
1. Click on the Templates nav,
2. Randomly choose an existing template name from the table,
3. Click on the 'Create a template' button on the table toolbar.
4. Create a patch template with the name that you have chosen.
5. After you click on the 'Submit' button on the last step, an error notification should pop up.



---

## Discussion

### Comment by @mkholjuraev on July 21, 2022 at 11:07 AM UTC

@bastilian thank you for the review!

### Comment by @jiridostal on August 01, 2022 at 02:54 PM UTC

:tada: This PR is included in version 1.48.7 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.7)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.7)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on July 21, 2022 at 10:45 AM UTC

LGTM! Thanks @mkholjuraev!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/847*
