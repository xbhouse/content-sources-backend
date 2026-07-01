---
type: pull_request
number: 446
title: "Fixes 2474: Recover from tasks panics"
state: merged
author: jlsherrill
created: 2023-10-26T19:55:13Z
updated: 2023-11-03T19:51:30Z
closed: 2023-11-03T19:50:46Z
merged: 2023-11-03T19:50:46Z
base_branch: main
head_branch: 2474
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/446
---

# Pull Request #446: Fixes 2474: Recover from tasks panics

**Author**: @jlsherrill
**Created**: October 26, 2023 at 07:55 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `2474`

## Description

## Summary
Previously if a task paniced for some reason, the worker routine would get 'stuck' and refuse to process anything else.  Now the task is marked as failed and the worker can continue to pick up another task.

## Testing steps

modify the pkg/taskss/introspection.go  to panic sometimes or all the time:
```
	if rand.Int()%2 == 0 {
		panic(fmt.Errorf("You are unlucky!!"))
	}
```

Create some repository with with snapshotting turned off (for simplicity)
```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/repositories/" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJqdXN0aW4ifSwiYWNjb3VudF9udW1iZXIiOiIxMTQ4NjU4OSIsImludGVybmFsIjp7Im9yZ19pZCI6IjE2Nzg4NjE2In19fQo=" \
    -H "Content-Type: application/json" \
    -d "{
          \"name\": \"errata\",
          \"url\": \"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/\",
          \"snapshot\": false
        }"
```

Now trigger introspection: 
```
 OPTIONS_ALWAYS_RUN_CRON_TASKS=true go run cmd/external-repos/main.go nightly-jobs
```

Run this a few times.  eventually your workers will stop processing tasks.  With this PR, they should always recover and mark the task as failed.   You can check the tasks table' status column to confirm they are marked as failed.

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on October 26, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-2474

### Comment by @jlsherrill on October 27, 2023 at 05:58 PM UTC

/retest

### Comment by @jlsherrill on October 27, 2023 at 07:37 PM UTC

/retest

### Comment by @jlsherrill on November 03, 2023 at 07:50 PM UTC

merging as this does not require qe and the failing backend test is unrelated

---

## Reviews

### Review by @rverdile - Commented on October 30, 2023 at 05:56 PM UTC

### Review by @jlsherrill - Commented on October 30, 2023 at 06:55 PM UTC

### Review by @rverdile - Approved on November 03, 2023 at 07:27 PM UTC

nice!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/446*
