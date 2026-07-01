---
type: pull_request
number: 960
title: "feat(Templates): Update templates page header"
state: merged
author: leSamo
created: 2023-02-15T20:47:16Z
updated: 2023-05-08T18:17:08Z
closed: 2023-02-15T22:59:40Z
merged: 2023-02-15T22:59:40Z
base_branch: master
head_branch: update-templates-header
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/960
---

# Pull Request #960: feat(Templates): Update templates page header

**Author**: @leSamo
**Created**: February 15, 2023 at 08:47 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `update-templates-header`

## Description

# Description

Associated Jira ticket: [SPM-1548](https://issues.redhat.com/browse/SPM-1548)
Associated mockup: https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/ELywAxO

- update header title in Patch Templates page from 'Patch template' to 'Templates'.
- add popover next to header title
- update `frontend.yml` with current information

# How to test the PR

Verify correspondence of Template list page header with the [mockup](https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/ELywAxO).

# Mockup
![Screenshot from 2023-02-15 21-14-58](https://user-images.githubusercontent.com/8426204/219152444-d09f36ba-1038-47bb-862c-8688297ec427.png)

# Before the change
![Screenshot from 2023-02-15 21-13-32](https://user-images.githubusercontent.com/8426204/219152494-10b8d857-bbc6-47a2-a8f5-175da255dfe1.png)

# After the change
![Screenshot from 2023-02-15 21-07-51](https://user-images.githubusercontent.com/8426204/219152520-54f55a6e-7bc2-497a-848e-d23110628f1e.png)

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @mkholjuraev on February 15, 2023 at 11:21 PM UTC

:tada: This PR is included in version 1.59.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.59.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.59.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on February 15, 2023 at 09:40 PM UTC

LGTM! Thank you @leSamo .

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/960*
