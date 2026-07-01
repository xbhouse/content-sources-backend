---
type: pull_request
number: 339
title: "Fixes 2224: admin task fetch 404"
state: merged
author: dpang314
created: 2023-07-27T14:48:48Z
updated: 2023-08-02T19:39:37Z
closed: 2023-08-02T19:39:37Z
merged: 2023-08-02T19:39:37Z
base_branch: main
head_branch: HMS-2224-admin-task-fetch-404
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/339
---

# Pull Request #339: Fixes 2224: admin task fetch 404

**Author**: @dpang314
**Created**: July 27, 2023 at 02:48 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `HMS-2224-admin-task-fetch-404`

## Description

## Summary

Improves the "not found" error message when fetching an admin task using a uuid with invalid syntax.

Previously, when making a GET request to `/admin/tasks/:uuid` with a uuid that had invalid syntax, e.g. "invalid-uuid", the error message would be `(ERROR: invalid input syntax for type uuid: "invalid-uuid" (SQLSTATE 22P02))`

The PR changes the error message to `Could not find task with UUID invalid-uuid` which is consistent with other fetch APIs

## Testing steps

```curl localhost:8000/api/content-sources/v1.0/admin/tasks/bad-uuid  -H "`./scripts/header.sh 1 adminUser`"```
The "admin_tasks" config feature must be enabled. Either both "admin_tasks.accounts" and "admin_tasks.users" must  be nil to allow any user to access it, or "adminUser" must be added to "admin_tasks.users" to access the endpoint.

This can also be tested through the frontend in [this PR](https://github.com/content-services/content-sources-frontend/pull/128) by navigating to `/beta/insights/content/admin-tasks/bad-uuid` You should be redirected to `/beta/insights/content/admin-tasks` and the error will pop up as a notification.

---

## Discussion

### Comment by @jlsherrill on July 27, 2023 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-2224

---

## Reviews

### Review by @jlsherrill - Approved on July 31, 2023 at 02:04 PM UTC

Nice work!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/339*
