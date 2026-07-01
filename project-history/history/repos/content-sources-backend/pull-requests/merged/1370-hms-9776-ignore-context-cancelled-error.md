---
type: pull_request
number: 1370
title: "HMS-9776: ignore context cancelled error"
state: merged
author: TenSt
created: 2026-01-21T12:27:53Z
updated: 2026-01-22T11:18:40Z
closed: 2026-01-22T11:18:40Z
merged: 2026-01-22T11:18:40Z
base_branch: main
head_branch: stepan/HMS-9776-ignore-context-cancelled-error-in-snapshot-handler
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1370
---

# Pull Request #1370: HMS-9776: ignore context cancelled error

**Author**: @TenSt
**Created**: January 21, 2026 at 12:27 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `stepan/HMS-9776-ignore-context-cancelled-error-in-snapshot-handler`

## Description

## Summary

This PR adds check for the `context cancelled` error in the snapshot handler. The context can be cancelled only in case the worker is shutting off, so the task will be re-processed by the next worker, there is no need to log this as error.

Plus minor refactor by removing duplicate nil error checks in another place.

## Testing steps



---

## Discussion

### Comment by @xbhouse on January 21, 2026 at 12:30 PM UTC

https://issues.redhat.com/browse/HMS-9776

---

## Reviews

### Review by @rverdile - Approved on January 21, 2026 at 07:09 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1370*
