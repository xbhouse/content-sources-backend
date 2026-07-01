---
type: pull_request
number: 1283
title: "Test: Stop deploying IQE integration tests"
state: merged
author: swadeley
created: 2025-11-04T10:28:46Z
updated: 2026-01-13T06:53:49Z
closed: 2025-11-04T15:15:20Z
merged: 2025-11-04T15:15:20Z
base_branch: main
head_branch: swadeley/delete_deployments_testing-integration
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1283
---

# Pull Request #1283: Test: Stop deploying IQE integration tests

**Author**: @swadeley
**Created**: November 04, 2025 at 10:28 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/delete_deployments_testing-integration`

## Description

## Summary

Stop deploying IQE integration tests
We have already stopped running IQE UI tests and now are ready to retire the IQE API tests


## Testing steps
Playwright tests still pass



---

## Discussion

### Comment by @swadeley on November 04, 2025 at 11:27 AM UTC

```
 [3]:	E: Sub-process /usr/bin/dpkg returned an error code (2)
[3]:	dpkg: error: cannot stat pathname '/tmp/apt-dpkg-install-mcrrCa/000-fonts-ipafont-gothic_00303-21ubuntu1_all.deb': No such file or directory
[3]:	
[3]:	Failed to install browsers
[3]:	Error: Installation process exited with code: 100
[3]:	error Command failed with exit code 1.
```

### Comment by @swadeley on November 04, 2025 at 11:28 AM UTC

/retest

---

## Reviews

### Review by @TenSt - Approved on November 04, 2025 at 10:42 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1283*
