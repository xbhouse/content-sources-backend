---
type: pull_request
number: 963
title: "fix: Display 'No template' instead of blank cell in template column"
state: merged
author: leSamo
created: 2023-02-15T23:27:29Z
updated: 2023-05-08T18:17:05Z
closed: 2023-02-27T13:31:55Z
merged: 2023-02-27T13:31:55Z
base_branch: master
head_branch: template-column-fallback
labels: ["bug", "released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/963
---

# Pull Request #963: fix: Display 'No template' instead of blank cell in template column

**Author**: @leSamo
**Created**: February 15, 2023 at 11:27 PM UTC
**Status**: Merged
**Labels**: `bug`, `released`
**Base**: `master` ← **Head**: `template-column-fallback`

## Description

# Description

Associated Jira ticket: [SPM-1892](https://issues.redhat.com/browse/SPM-1892)
Mockup: https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/DP7Ka9y

Display 'No template' instead of blank cell in template column in System list table

# How to test the PR

Verify correspondence with the mockup.

# Before the change
![Screenshot from 2023-02-16 00-14-45](https://user-images.githubusercontent.com/8426204/219215666-d7ddbadc-174c-4e59-8436-801e9c46b71f.png)

# After the change
![Screenshot from 2023-02-16 00-14-24](https://user-images.githubusercontent.com/8426204/219215671-ee5b1ed7-71ec-4d9b-a4a2-9d495fc68669.png)

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @mkholjuraev on February 27, 2023 at 01:49 PM UTC

:tada: This PR is included in version 1.59.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.59.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.59.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on February 27, 2023 at 10:34 AM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/963*
