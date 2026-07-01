---
type: pull_request
number: 1093
title: "support for skipping runs/deploy in pr_check in ephemeral"
state: merged
author: tkasparek
created: 2022-09-01T13:43:41Z
updated: 2022-09-02T08:50:45Z
closed: 2022-09-02T08:50:45Z
merged: 2022-09-02T08:50:45Z
base_branch: master
head_branch: pr_check_labels
labels: ["skip-deploy"]
url: https://github.com/RedHatInsights/patchman-engine/pull/1093
---

# Pull Request #1093: support for skipping runs/deploy in pr_check in ephemeral

**Author**: @tkasparek
**Created**: September 01, 2022 at 01:43 PM UTC
**Status**: Merged
**Labels**: `skip-deploy`
**Base**: `master` ← **Head**: `pr_check_labels`

## Description

*No description provided*

---

## Discussion

### Comment by @psegedy on September 01, 2022 at 01:44 PM UTC

😲 

### Comment by @tkasparek on September 02, 2022 at 06:45 AM UTC

![image](https://user-images.githubusercontent.com/6339153/188075199-ffbd58da-de2c-426f-9130-1e24322a1a8d.png)
No test run
![image](https://user-images.githubusercontent.com/6339153/188075245-6885dc44-d424-41c4-8457-c79a6821dbda.png)
Everything took 7 minutes 26 seconds.
![image](https://user-images.githubusercontent.com/6339153/188075338-5b126f3b-3655-4cfd-a391-9574eadbe219.png)
Deployment to ephemeral skipped.

@psegedy I'd say it looks sane. Do you agree?

---

## Reviews

### Review by @psegedy - Approved on September 01, 2022 at 01:46 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1093*
