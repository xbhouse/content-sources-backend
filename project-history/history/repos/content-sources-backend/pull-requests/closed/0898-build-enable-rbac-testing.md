---
type: pull_request
number: 898
title: "Build: enable rbac testing"
state: closed
author: swadeley
created: 2024-11-19T11:34:15Z
updated: 2025-09-08T07:31:12Z
closed: 2025-01-02T01:42:07Z
base_branch: main
head_branch: swadeley/enable_stage_rbac_testing
labels: []
url: https://github.com/content-services/content-sources-backend/pull/898
---

# Pull Request #898: Build: enable rbac testing

**Author**: @swadeley
**Created**: November 19, 2024 at 11:34 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/enable_stage_rbac_testing`

## Description

## Summary

Testing if RBAC testing works in pr_checks

## Testing steps

check test results


---

## Discussion

### Comment by @swadeley on December 03, 2024 at 11:37 AM UTC

/retest

### Comment by @swadeley on December 04, 2024 at 06:58 AM UTC

Hi, the RBAC test works for backend if `IQE_ENV="ephemeral" `.

 The RBAC test does not test code changes in PRs, its testing backend support for user perms, so it seems not worth it to use in backend PR checks. It will be run in stage post PR merge.

### Comment by @swadeley on December 13, 2024 at 02:07 AM UTC

/retest

### Comment by @jlsherrill on December 18, 2024 at 06:12 PM UTC

https://issues.redhat.com/browse/Build

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/898*
