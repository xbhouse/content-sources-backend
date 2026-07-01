---
type: pull_request
number: 255
title: "Update e2e test template to pull testing image every time"
state: merged
author: tpapaioa
created: 2023-04-21T18:14:47Z
updated: 2023-04-21T18:39:56Z
closed: 2023-04-21T18:39:52Z
merged: 2023-04-21T18:39:52Z
base_branch: main
head_branch: update_e2e_test_image_pull_policy
labels: []
url: https://github.com/content-services/content-sources-backend/pull/255
---

# Pull Request #255: Update e2e test template to pull testing image every time

**Author**: @tpapaioa
**Created**: April 21, 2023 at 06:14 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `update_e2e_test_image_pull_policy`

## Description

## Summary
This PR updates the e2e test template to use an `imagePullPolicy` of `Always` for the testing container image. This is required because we always use the same `hms-integration` image tag to point to the latest build, and by default tags other than `latest` are not re-pulled.

## Testing steps


---

## Reviews

### Review by @swadeley - Approved on April 21, 2023 at 06:27 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/255*
