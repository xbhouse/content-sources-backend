---
type: pull_request
number: 934
title: "HMS-5127: force task requeue on server exit"
state: merged
author: rverdile
created: 2025-01-07T21:02:53Z
updated: 2025-01-10T19:25:01Z
closed: 2025-01-10T19:24:54Z
merged: 2025-01-10T19:24:54Z
base_branch: main
head_branch: req-server-exit
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/934
---

# Pull Request #934: HMS-5127: force task requeue on server exit

**Author**: @rverdile
**Created**: January 07, 2025 at 09:02 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `req-server-exit`

## Description

## Summary
This will make sure a running task requeues  on server exit, rather than possibly failing before having the chance to requeue

## Testing steps
1. You can apply this patch to help recreate the conditions of the bug
```
diff --git a/pkg/tasks/queue/queue.go b/pkg/tasks/queue/queue.go
index cbf0105..b97c5bd 100644
--- a/pkg/tasks/queue/queue.go
+++ b/pkg/tasks/queue/queue.go
@@ -11,7 +11,7 @@ import (
 	"github.com/google/uuid"
 )
 
-const MaxTaskRetries = 3 // Maximum number of times a task can be retried before failing
+const MaxTaskRetries = 100 // Maximum number of times a task can be retried before failing
 
 type Task struct {
 	Typename     string
diff --git a/pkg/tasks/worker/worker.go b/pkg/tasks/worker/worker.go
index 3fac3e8..f8754f4 100644
--- a/pkg/tasks/worker/worker.go
+++ b/pkg/tasks/worker/worker.go
@@ -84,6 +84,7 @@ func (w *worker) start(ctx context.Context) {
 	defer beat.Stop()
 
 	for {
+		time.Sleep(time.Second)
 		select {
 		case <-w.stopChan:
 			if w.runningTask.id != uuid.Nil {
```
2. With this PR, closing the server (CTRL+C) should always allow the running tasks to requeue
3. Without this PR (and with this patch), you should see the tasks fail and therefore fail to requeue as well



---

## Discussion

### Comment by @jlsherrill on January 07, 2025 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-5127

---

## Reviews

### Review by @xbhouse - Approved on January 10, 2025 at 03:25 PM UTC

looks good and the fix works well! ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/934*
