---
type: pull_request
number: 296
title: "Fixes 1580: Add API to list and fetch tasks"
state: merged
author: dpang314
created: 2023-05-30T16:35:14Z
updated: 2023-06-05T19:02:30Z
closed: 2023-06-05T18:32:30Z
merged: 2023-06-05T18:32:30Z
base_branch: main
head_branch: HMS-1580-task-api
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/296
---

# Pull Request #296: Fixes 1580: Add API to list and fetch tasks

**Author**: @dpang314
**Created**: May 30, 2023 at 04:35 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-1580-task-api`

## Description

## Summary

Adds API to list tasks:

`GET /api/content-sources/v1.0/tasks/`

and to fetch a task:

`GET /api/content-sources/v1.0/tasks/:uuid`

It defines a GORM model that is shared between the queue and DAO. The DAO only uses GORM to read in tasks.

## Testing steps

Start an introspection task

Make a GET request to `/api/content-sources/v1.0/tasks/` to see all tasks

Make a GET request to `GET /api/content-sources/v1.0/tasks/:uuid` with the task UUID to the details of the specific task




---

## Discussion

### Comment by @app-sre-bot on May 30, 2023 at 04:35 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on May 30, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-1580

### Comment by @swadeley on May 30, 2023 at 05:10 PM UTC

Hi, rebase please.

### Comment by @swadeley on May 31, 2023 at 01:17 PM UTC

For QE, these will be the command formats for testing:

`app.content_sources.rest_client.tasks_api.list_tasks()`
`app.content_sources.rest_client.tasks_api.get_task('8a625405-<snip>c71e674640')`


### Comment by @swadeley on June 03, 2023 at 03:17 PM UTC

Hi David, please rebase again and ping me when its ready for testing.

### Comment by @swadeley on June 05, 2023 at 06:32 PM UTC

Hi

Looks good:

`$ curl  -u "$BASIC_AUTH" -X GET "https://${URL}/api/content-sources/v1.0/tasks/"
{"data":[{"uuid":"7a3705bd-54d4-497b-a5c9-9c6fa2c8f882","status":"completed","created_at":"2023-06-05T18:19:37Z","ended_at":"2023-06-05T18:19:37Z","error":"","org_id":"3340851"},{"uuid":"dfedeabd-9f26-44fc-b5e2-ce5422f2bf2d","status":"completed","created_at":"2023-06-05T17:59:29Z","ended_at":"2023-06-05T17:59:30Z","error":"","org_id":"3340851"},{"uuid":"35372449-83e8-49c4-abc8-e2854a51ce7f","status":"completed","created_at":"2023-06-05T17:56:56Z","ended_at":"2023-06-05T17:57:02Z","error":"","org_id":"3340851"}],"meta":{"limit":100,"offset":0,"count":3},"links":{"first":"/api/content-sources/v1.0/tasks/?limit=100\u0026offset=0","last":"/api/content-sources/v1.0/tasks/?limit=100\u0026offset=0"}}`


`$ curl  -u "$BASIC_AUTH" -X GET "https://${URL}/api/content-sources/v1.0/tasks/7a3705bd-54d4-497b-a5c9-9c6fa2c8f882"
{"uuid":"7a3705bd-54d4-497b-a5c9-9c6fa2c8f882","status":"completed","created_at":"2023-06-05T18:19:37Z","ended_at":"2023-06-05T18:19:37Z","error":"","org_id":"3340851"}`

---

## Reviews

### Review by @swadeley - Commented on May 31, 2023 at 07:13 AM UTC

### Review by @rverdile - Commented on May 31, 2023 at 02:20 PM UTC

Looks really good overall! A few small comments about things to rename. 

Still need to test and that may inspire some more feedback, but I don't anticipate any huge changes.

### Review by @rverdile - Commented on June 02, 2023 at 05:27 PM UTC

Just have to add the endpoints to here and this is good to go!

https://github.com/content-services/content-sources-backend/blob/main/pkg/middleware/rbac_map.go#L22-L36

### Review by @rverdile - Approved on June 02, 2023 at 08:12 PM UTC

nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/296*
