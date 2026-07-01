---
type: pull_request
number: 445
title: "Fixes 2282: reduce number of introspect tasks"
state: merged
author: jlsherrill
created: 2023-10-25T21:19:11Z
updated: 2023-12-05T14:44:40Z
closed: 2023-12-03T23:25:36Z
merged: 2023-12-03T23:25:36Z
base_branch: main
head_branch: 2282
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/445
---

# Pull Request #445: Fixes 2282: reduce number of introspect tasks

**Author**: @jlsherrill
**Created**: October 25, 2023 at 09:19 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2282`

## Description

## Summary

This restructures the introspection code to move all decisions about whether a repo should be introspected into the main query to lookup repos for introspection.  the result is that we do not spawn tasks for repos that won't actually be introspected.  It should reduce the tasks we spawn by about 2/3rds.  

Some considerations:
* Now in pkg/external_repos  any introspect function will introspect without question
* The notifications are being sent one at a time for introspection, but they were before as well.  We likely want to revisit how those are generated.

## Testing steps

1.  Create one or more repos for introspection (but turn off snapshotting for simplicity)
2. run ```go run cmd/external-repos/main.go nightly-jobs```
3. notice that no tasks get run
4.  wait ~24 hours or adjust config.IntrospectTimeInterval in value_constraints.go to `time.Minute * 1`  and wait ~ 1 minute
5. run ```go run cmd/external-repos/main.go nightly-jobs```
6. notice taht tasks now get spawned 


## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on October 25, 2023 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-2282

### Comment by @jlsherrill on October 26, 2023 at 11:48 AM UTC

/retest

### Comment by @xbhouse on November 07, 2023 at 02:58 PM UTC

tested and getting expected results. 

no tasks spawned when running that command (`go run cmd/external-repos/main.go nightly-jobs`) right after adding a repo. changed the value_constraints and waited a minute to run it again and see that tasks now get spawned

![Screenshot 2023-11-07 at 9 55 21 AM](https://github.com/content-services/content-sources-backend/assets/28575816/221bc155-6e8e-4ef7-83e9-6ae979778b75)



### Comment by @xbhouse on November 07, 2023 at 03:03 PM UTC

i'm going to ack but maybe someone else will have more feedback / see something i didn't 

### Comment by @swadeley on November 08, 2023 at 07:20 AM UTC

/retest

### Comment by @swadeley on November 08, 2023 at 12:43 PM UTC

Hi, lgtm, waiting for one more ack.

### Comment by @swadeley on November 21, 2023 at 01:25 AM UTC

@jlsherrill rebase please

### Comment by @swadeley on December 03, 2023 at 11:25 PM UTC

lgtm, thank you.

---

## Reviews

### Review by @jlsherrill - Commented on October 25, 2023 at 09:26 PM UTC

### Review by @jlsherrill - Commented on October 25, 2023 at 09:43 PM UTC

### Review by @xbhouse - Dismissed on November 07, 2023 at 03:03 PM UTC

### Review by @rverdile - Commented on November 08, 2023 at 03:03 PM UTC

If introspection fails, e.g. URL http://example.com, I'm not seeing the task marked as failed. It still says "task completed". There's also no error saved on the task. 

### Review by @rverdile - Commented on November 08, 2023 at 03:17 PM UTC

### Review by @rverdile - Commented on November 08, 2023 at 03:18 PM UTC

### Review by @rverdile - Commented on November 08, 2023 at 03:35 PM UTC

I'm also still seeing errors for public repos being logged at the warn level. The issue says to log them at the error level.

Tested by either running `make repos-import` then `go run cmd/external-repos/main.go nightly-jobs` (without a cdn cert so that red hat ones fail). Or, manually updating the database to set a failing repository to public.

### Review by @jlsherrill - Commented on November 08, 2023 at 04:41 PM UTC

### Review by @jlsherrill - Commented on November 08, 2023 at 04:46 PM UTC

### Review by @rverdile - Commented on November 08, 2023 at 09:45 PM UTC

### Review by @rverdile - Commented on November 08, 2023 at 09:47 PM UTC

### Review by @rverdile - Approved on November 28, 2023 at 03:47 PM UTC

working!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/445*
