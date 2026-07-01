---
type: pull_request
number: 915
title: "Build: set longer pipeline timeout"
state: closed
author: swadeley
created: 2024-12-09T03:14:53Z
updated: 2025-09-08T07:31:13Z
closed: 2024-12-09T11:38:34Z
base_branch: main
head_branch: swadeley/set_longer_pipeline_timeout
labels: []
url: https://github.com/content-services/content-sources-backend/pull/915
---

# Pull Request #915: Build: set longer pipeline timeout

**Author**: @swadeley
**Created**: December 09, 2024 at 03:14 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/set_longer_pipeline_timeout`

## Description

## Summary

    IQE Merge Request smoke-test job needs more than the default one hour to build the image and run the tests.
    This doc suggested adding a longer timeout for the whole pipeline as a start: https://tekton.dev/docs/pipelines/pipelineruns/#configuring-a-failure-timeout


## Testing steps

 pr checks pass



---

## Discussion

### Comment by @swadeley on December 09, 2024 at 05:05 AM UTC

/retest

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/915*
