---
type: pull_request
number: 1430
title: "Test: Fix cleanup to wait until done"
state: merged
author: swadeley
created: 2026-03-27T10:14:52Z
updated: 2026-03-27T13:34:49Z
closed: 2026-03-27T13:34:49Z
merged: 2026-03-27T13:34:49Z
base_branch: main
head_branch: swadeley/extend_cleanup_to_wait_until_done
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1430
---

# Pull Request #1430: Test: Fix cleanup to wait until done

**Author**: @swadeley
**Created**: March 27, 2026 at 10:14 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/extend_cleanup_to_wait_until_done`

## Description

## Summary

Wait until repositories are gone after bulk delete in cleanupRepositories

## Testing steps

tests pass
This issue mostly affected tests using introspection only repos, such as UI/GPG_keys.spec.ts, when run twice in a row.

---

## Reviews

### Review by @TenSt - Approved on March 27, 2026 at 10:28 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1430*
