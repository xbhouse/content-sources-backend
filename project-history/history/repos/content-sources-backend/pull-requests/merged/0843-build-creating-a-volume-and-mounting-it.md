---
type: pull_request
number: 843
title: "Build: Creating a volume and mounting it"
state: merged
author: mayurilahane
created: 2024-10-10T12:17:34Z
updated: 2024-12-09T13:50:11Z
closed: 2024-12-09T13:50:10Z
merged: 2024-12-09T13:50:10Z
base_branch: main
head_branch: mlahane/empty_dir
labels: []
url: https://github.com/content-services/content-sources-backend/pull/843
---

# Pull Request #843: Build: Creating a volume and mounting it

**Author**: @mayurilahane
**Created**: October 10, 2024 at 12:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `mlahane/empty_dir`

## Description

Need to create and mount an empty directory inside the pod using a container. 
So that you can upload any files from that directory or can use it as storage while testing if you download anything.

---

## Discussion

### Comment by @jlsherrill on October 10, 2024 at 12:43 PM UTC

test failure doesn't look related?

### Comment by @swadeley on October 10, 2024 at 07:11 PM UTC

> test failure doesn't look related?

No it is not and that should be fixed now, I'll try a retest 

### Comment by @swadeley on October 10, 2024 at 07:12 PM UTC

/retest

### Comment by @swadeley on October 10, 2024 at 07:15 PM UTC

Hi @mayurilahane, I see you added this to `deployments/testing-integration.yaml`

 is this upload test only going to run in integration testing?

To use this in content-sources testing you need to add to `deployments/deployment.yaml`

---

## Reviews

### Review by @swadeley - Approved on December 09, 2024 at 01:50 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/843*
