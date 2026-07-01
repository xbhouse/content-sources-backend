---
type: pull_request
number: 821
title: "fix(TableView): wrap table toolbar with ErrorHandler"
state: merged
author: mkholjuraev
created: 2022-06-15T11:16:32Z
updated: 2024-04-03T09:23:23Z
closed: 2022-06-21T09:13:19Z
merged: 2022-06-21T09:13:19Z
base_branch: master
head_branch: move-error-handler
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/821
---

# Pull Request #821: fix(TableView): wrap table toolbar with ErrorHandler

**Author**: @mkholjuraev
**Created**: June 15, 2022 at 11:16 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `move-error-handler`

## Description

This will hide the table toolbar when there is an some error.

Before: 

![image](https://user-images.githubusercontent.com/59481011/173809261-ffba1b7e-ee70-4aaa-9583-86c21151efb4.png)




After:
![image](https://user-images.githubusercontent.com/59481011/173808940-22c21d3a-e04f-4c7c-b0cd-786a11c6324e.png)


---

## Discussion

### Comment by @mkholjuraev on June 21, 2022 at 06:30 AM UTC

@AsToNlele thanks for the review. I will apply your suggestion and then merge!

### Comment by @jiridostal on June 21, 2022 at 09:43 AM UTC

:tada: This PR is included in version 1.48.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Commented on June 20, 2022 at 03:03 PM UTC

### Review by @AsToNlele - Approved on June 20, 2022 at 03:06 PM UTC

Sometimes there's ``status.hasError`` and sometimes ``hasError`` those are just some nitpicks :smile:
Other than that great job! Sorry for the delay

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/821*
