---
type: pull_request
number: 1245
title: "Build: chore(IQE-2762): Update IQE selenium image"
state: closed
author: akhil-jha
created: 2025-10-22T07:33:53Z
updated: 2025-11-04T15:16:50Z
closed: 2025-11-04T15:16:50Z
base_branch: main
head_branch: update_sel_image
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1245
---

# Pull Request #1245: Build: chore(IQE-2762): Update IQE selenium image

**Author**: @akhil-jha
**Created**: October 22, 2025 at 07:33 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `update_sel_image`

## Description

Image update

---

## Discussion

### Comment by @akhil-jha on October 22, 2025 at 07:41 AM UTC

@mshriver 

### Comment by @swadeley on October 23, 2025 at 08:05 AM UTC

Hi @akhil-jha , please rebase to clear the format check error

### Comment by @marusak on October 25, 2025 at 03:28 PM UTC

Where do we use this image? @swadeley 

### Comment by @swadeley on October 27, 2025 at 08:16 PM UTC

> Where do we use this image? @swadeley

Hi @marusak 

We still run IQE tests in Konflux and automated merge requests in app-interface also run IQE tests.
These tests are configured to run API tests only, so this image should in theory not be required as its for UI, but we would need to test want happens if we just remove the image as it might break deployment.

Justin and Ryan are working to disable these tests however Justin still wants an ephemeral deployment to take place so I offered to just disable the IQE test run for the moment.

### Comment by @swadeley on October 27, 2025 at 08:30 PM UTC

Hi

Testing removing Selenium image here:

Test: Stop deploying selenium image #1260

### Comment by @swadeley on November 04, 2025 at 03:16 PM UTC

Hi

closing in favour of:  Test: Stop deploying IQE integration tests #1283 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1245*
