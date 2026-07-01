---
type: pull_request
number: 1495
title: "HMS-10580: Set limit on pulp orphan cleanup tasks"
state: merged
author: swadeley
created: 2026-05-14T13:09:41Z
updated: 2026-05-20T13:41:04Z
closed: 2026-05-20T13:41:04Z
merged: 2026-05-20T13:41:03Z
base_branch: main
head_branch: swadeley/HMS-10580
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1495
---

# Pull Request #1495: HMS-10580: Set limit on pulp orphan cleanup tasks

**Author**: @swadeley
**Created**: May 14, 2026 at 01:09 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/HMS-10580`

## Description



## Summary
Set limit on pulp orphan cleanup tasks
Sets a global cap and adds new constant:
 maxGetTaskAttempts (defaults to 5 when unset) 

HMS-10580 API - pulp orphan cleanup runs too many tasks
## Testing steps

tests pass

---

## Discussion

### Comment by @xbhouse on May 14, 2026 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-10580

### Comment by @swadeley on May 14, 2026 at 05:46 PM UTC

Testing with "maxPollFailures = 2":

```
:~/content-sources-backend$ go test -count=1 -v ./pkg/external_repos/commands/ -run TestRunPulpOrphanCleanupForDomains
=== RUN   TestRunPulpOrphanCleanupForDomains_StopsAfterMaxPollFailures
=== PAUSE TestRunPulpOrphanCleanupForDomains_StopsAfterMaxPollFailures
=== CONT  TestRunPulpOrphanCleanupForDomains_StopsAfterMaxPollFailures
{"level":"info","org_id":"1","pulp_domain_name":"d1","task_href":"https://pulp.example/api/v3/tasks/fake-href/","time":"2026-05-14T18:44:21+01:00","message":"running orphan cleanup for org: 1"}
{"level":"error","org_id":"1","pulp_domain_name":"d1","error":"injected poll failure","time":"2026-05-14T18:44:21+01:00","message":"error polling pulp task for orphan cleanup"}
{"level":"info","org_id":"2","pulp_domain_name":"d2","task_href":"https://pulp.example/api/v3/tasks/fake-href/","time":"2026-05-14T18:44:21+01:00","message":"running orphan cleanup for org: 2"}
{"level":"error","org_id":"2","pulp_domain_name":"d2","error":"injected poll failure","time":"2026-05-14T18:44:21+01:00","message":"error polling pulp task for orphan cleanup"}
{"level":"error","max_poll_failures":2,"poll_failures":2,"time":"2026-05-14T18:44:21+01:00","message":"stopping pulp orphan cleanup: too many failed task polls; remaining domains skipped (orphan cleanup tasks might still be running on Pulp)"}
--- PASS: TestRunPulpOrphanCleanupForDomains_StopsAfterMaxPollFailures (0.00s)
PASS
ok  	github.com/content-services/content-sources-backend/pkg/external_repos/commands	0.006s
```


### Comment by @swadeley on May 20, 2026 at 01:40 PM UTC

> looks good, nice job!

Thank you.

---

## Reviews

### Review by @rverdile - Commented on May 14, 2026 at 04:05 PM UTC

Let's just use a constant here. I don't think this is a value that needs to be configurable.

Also this is a behavior that can be tested locally by manually writing an error after the return for the polling task

### Review by @rverdile - Commented on May 15, 2026 at 12:26 PM UTC

### Review by @swadeley - Commented on May 18, 2026 at 10:12 AM UTC

### Review by @swadeley - Commented on May 18, 2026 at 10:13 AM UTC

### Review by @TenSt - Commented on May 18, 2026 at 02:05 PM UTC

### Review by @swadeley - Commented on May 18, 2026 at 02:42 PM UTC

### Review by @rverdile - Commented on May 18, 2026 at 08:44 PM UTC

### Review by @rverdile - Commented on May 18, 2026 at 08:55 PM UTC

### Review by @rverdile - Commented on May 18, 2026 at 08:57 PM UTC

### Review by @rverdile - Commented on May 18, 2026 at 08:57 PM UTC

### Review by @swadeley - Commented on May 19, 2026 at 09:08 AM UTC

### Review by @swadeley - Commented on May 19, 2026 at 09:08 AM UTC

### Review by @rverdile - Commented on May 19, 2026 at 03:03 PM UTC

### Review by @rverdile - Commented on May 19, 2026 at 03:04 PM UTC

### Review by @rverdile - Commented on May 19, 2026 at 03:04 PM UTC

### Review by @rverdile - Commented on May 19, 2026 at 03:11 PM UTC

### Review by @swadeley - Commented on May 19, 2026 at 04:15 PM UTC

### Review by @swadeley - Commented on May 19, 2026 at 04:43 PM UTC

### Review by @swadeley - Commented on May 19, 2026 at 04:57 PM UTC

### Review by @rverdile - Approved on May 20, 2026 at 01:28 PM UTC

looks good, nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1495*
