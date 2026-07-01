---
type: pull_request
number: 1320
title: "Test: Add retry logic for pulp clients"
state: closed
author: swadeley
created: 2025-12-01T10:53:27Z
updated: 2026-01-13T06:51:56Z
closed: 2025-12-01T14:54:18Z
base_branch: main
head_branch: swadeley/add_retry_logic_for_pulp_clients
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1320
---

# Pull Request #1320: Test: Add retry logic for pulp clients

**Author**: @swadeley
**Created**: December 01, 2025 at 10:53 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/add_retry_logic_for_pulp_clients`

## Description

## Summary

  Retry connections with no response from pulp.

## Testing steps

All tests pass

---

## Discussion

### Comment by @swadeley on December 01, 2025 at 02:54 PM UTC

Closing as this approach will cause more load to Pulp (retrying too soon). We should consider a retry from the UI with a gap or wait time longer than the backend's 60s wait time.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1320*
