---
type: pull_request
number: 1160
title: "HMS-8894: Test to search rpm in non added popular repo"
state: merged
author: mayurilahane
created: 2025-07-29T05:05:19Z
updated: 2025-08-06T13:19:39Z
closed: 2025-08-06T13:19:39Z
merged: 2025-08-06T13:19:39Z
base_branch: main
head_branch: mlahane/HMS-8894
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1160
---

# Pull Request #1160: HMS-8894: Test to search rpm in non added popular repo

**Author**: @mayurilahane
**Created**: July 29, 2025 at 05:05 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `mlahane/HMS-8894`

## Description

## Summary
migrate 
https://gitlab.cee.redhat.com/insights-qe/iqe-content-sources-plugin/-/blob/master/iqe_content_sources/tests/test_repository_api_only.py?ref_type=heads#L1510


## Testing steps

peer review



---

## Discussion

### Comment by @rverdile on July 29, 2025 at 02:16 PM UTC

I believe this would also need a change like this: https://github.com/content-services/content-sources-frontend/blob/main/.github/workflows/playwright-actions.yml#L224-L226 (without the "working_directory"). EPEL isn't currently introspected as part of the backend playwright CI

### Comment by @jlsherrill on July 31, 2025 at 06:00 AM UTC

https://issues.redhat.com/browse/HMS-8894

---

## Reviews

### Review by @mayurilahane - Commented on July 31, 2025 at 04:48 AM UTC

### Review by @rverdile - Commented on July 31, 2025 at 01:49 PM UTC

### Review by @mayurilahane - Commented on August 01, 2025 at 03:30 AM UTC

### Review by @mayurilahane - Commented on August 01, 2025 at 03:47 AM UTC

### Review by @rverdile - Commented on August 01, 2025 at 12:58 PM UTC

### Review by @rverdile - Commented on August 01, 2025 at 12:58 PM UTC

### Review by @mayurilahane - Commented on August 01, 2025 at 04:00 PM UTC

### Review by @rverdile - Commented on August 05, 2025 at 07:26 PM UTC

### Review by @mayurilahane - Commented on August 05, 2025 at 07:36 PM UTC

### Review by @rverdile - Approved on August 06, 2025 at 01:05 PM UTC

nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1160*
