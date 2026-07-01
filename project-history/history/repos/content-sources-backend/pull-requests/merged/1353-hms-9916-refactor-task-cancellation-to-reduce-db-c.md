---
type: pull_request
number: 1353
title: "HMS-9916: refactor task cancellation to reduce db connections"
state: merged
author: rverdile
created: 2026-01-05T21:47:33Z
updated: 2026-01-30T15:38:03Z
closed: 2026-01-30T15:37:58Z
merged: 2026-01-30T15:37:58Z
base_branch: main
head_branch: cancel-refactor
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1353
---

# Pull Request #1353: HMS-9916: refactor task cancellation to reduce db connections

**Author**: @rverdile
**Created**: January 05, 2026 at 09:47 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `cancel-refactor`

## Description

## Summary
This PR mainly refactors task cancellation so that there is no longer 1 connection used per worker in the worker pool. Instead, the worker pool uses 1 connection to handle task cancellation for every worker. 

## Testing steps
These testing steps will make it so you can simulate the workers getting stuck from the cancel listeners using all the connections in the pool.

It's just one example of how to test this PR. This PR also changes enqueue/dequeue to use PgListener, so the tasking system in general should be tested.

1. Switch to main
2. Apply this patch
```
diff --git a/pkg/tasks/introspect.go b/pkg/tasks/introspect.go
index 1a868ca..cd241b0 100644
--- a/pkg/tasks/introspect.go
+++ b/pkg/tasks/introspect.go
@@ -5,6 +5,7 @@ import (
        "encoding/json"
        "errors"
        "fmt"
+       "time"
 
        "github.com/content-services/content-sources-backend/pkg/config"
        "github.com/content-services/content-sources-backend/pkg/dao"
@@ -24,6 +25,7 @@ func IntrospectHandler(ctx context.Context, task *models.TaskInfo, _ *queue.Queu
        if err := json.Unmarshal(task.Payload, &p); err != nil {
                return fmt.Errorf("payload incorrect type for IntrospectHandler: %w", err)
        }
+       time.Sleep(1 * time.Second)
        intro := IntrospectionTask{
                URL:    p.Url,
                daoReg: dao.GetDaoRegistry(db.DB),
 ```
3. in config.yaml the worker_count to 7 and the pool_limit to 7
4. start the task worker `go run cmd/content-sources/main.go consumer`
5. create 7 introspect tasks. you can use [this script](https://gist.github.com/rverdile/7c96d898d6d38e1ee179d0a3eb502bd6) to help you i.e. `NUM_TASKS=7 manage-test-tasks.sh create`
6. the tasks should all get stuck (dequeued but never finished)
7. stop the server and remove the tasks `manage-test-tasks.sh remove`
8. switch to this PR and repeat the above steps, this time the tasks should all finish (but fail because the repositories don't exist)


---

## Discussion

### Comment by @xbhouse on January 05, 2026 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-9916

### Comment by @xbhouse on January 16, 2026 at 06:51 PM UTC

nice work on this! doing some testing and i see a couple things:

1. when trying to shut down the server while tasks are running, it looks like some tasks continue to run and the server never exits: 

```
^C1:25PM INF failed task requeue listener shutting down severity=info
1:25PM INF Caught context done, closing api server. severity=info
1:25PM INF Stopping workers severity=info
1:25PM INF instrumentation stopped severity=info
1:25PM INF apiServer stopped severity=info
1:25PM INF Stopping metrics collector go routine severity=info
1:25PM INF custom metrics stopped severity=info
1:25PM INF cancellation listener shutting down severity=info
1:25PM INF heartbeat listener shutting down severity=info
1:25PM INF /home/bhouse/Repos/content-sources/content-sources-backend/pkg/db/db_logger.go:173
[1.997ms] [rows:1] UPDATE repository_configurations SET failed_snapshot_count = failed_snapshot_count + 1  WHERE uuid = '5ab4db67-4277-4393-8fbc-8f9e6567a853' AND repository_configurations.deleted_at IS NULL severity=info
1:25PM INF /home/bhouse/Repos/content-sources/content-sources-backend/pkg/db/db_logger.go:173
[1.584ms] [rows:1] UPDATE repository_configurations SET failed_snapshot_count = failed_snapshot_count + 1  WHERE uuid = '040d1743-c01d-48a0-af45-9035870b57fa' AND repository_configurations.deleted_at IS NULL severity=info
1:26PM INF /home/bhouse/Repos/content-sources/content-sources-backend/pkg/db/db_logger.go:173
[1.704ms] [rows:1] UPDATE repository_configurations SET failed_snapshot_count = failed_snapshot_count + 1  WHERE uuid = 'c7447cce-4993-49ad-85f8-2df0db9231a6' AND repository_configurations.deleted_at IS NULL severity=info
1:26PM INF /home/bhouse/Repos/content-sources/content-sources-backend/pkg/db/db_logger.go:173
[2.193ms] [rows:1] UPDATE repository_configurations SET failed_snapshot_count = failed_snapshot_count + 1  WHERE uuid = '3337916a-00f8-40de-9bbe-8ea01d897b05' AND repository_configurations.deleted_at IS NULL severity=info
1:26PM INF /home/bhouse/Repos/content-sources/content-sources-backend/pkg/db/db_logger.go:173
[1.357ms] [rows:1] UPDATE repository_configurations SET failed_snapshot_count = failed_snapshot_count + 1  WHERE uuid = 'ecf6b01d-07ab-45a5-9c19-a71447a3b8b8' AND repository_configurations.deleted_at IS NULL severity=info
1:26PM INF /home/bhouse/Repos/content-sources/content-sources-backend/pkg/db/db_logger.go:173
[1.404ms] [rows:1] UPDATE repository_configurations SET failed_snapshot_count = failed_snapshot_count + 1  WHERE uuid = '6ac035e2-d9ae-4432-b560-49811d50f6dc' AND repository_configurations.deleted_at IS NULL severity=info
1:26PM INF /home/bhouse/Repos/content-sources/content-sources-backend/pkg/db/db_logger.go:173
[1.556ms] [rows:1] UPDATE repository_configurations SET failed_snapshot_count = failed_snapshot_count + 1  WHERE uuid = 'e8b82b3d-490f-498d-ae64-27325229bd87' AND repository_configurations.deleted_at IS NULL severity=info
```

2. when running `make repos-import && go run cmd/external-repos/main.go process-repos`, i saw this error right before the listener is shut down: `1:12PM ERR Query args=[] err="conn busy" module=pgx pid=294 severity=error sql="UNLISTEN tasks"`. i haven't been able to reproduce this again and i'm not sure if it's related to the changes in your PR, but wanted to note it just in case :)

   EDIT: oh i can reproduce it, forgot i had set `pgx_logging` to `true` :laughing: 

### Comment by @rverdile on January 28, 2026 at 07:53 PM UTC

@xbhouse updated!

---

## Reviews

### Review by @xbhouse - Commented on January 16, 2026 at 07:46 PM UTC

### Review by @rverdile - Commented on January 19, 2026 at 02:26 PM UTC

### Review by @xbhouse - Approved on January 29, 2026 at 07:37 PM UTC

this looks and works great! nice job :) 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1353*
