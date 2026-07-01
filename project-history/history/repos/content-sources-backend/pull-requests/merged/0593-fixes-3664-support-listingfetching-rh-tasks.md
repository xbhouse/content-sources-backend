---
type: pull_request
number: 593
title: "Fixes 3664: support listing/fetching RH tasks"
state: merged
author: jlsherrill
created: 2024-03-04T19:49:36Z
updated: 2024-03-08T10:00:26Z
closed: 2024-03-08T09:53:40Z
merged: 2024-03-08T09:53:40Z
base_branch: main
head_branch: 3664
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/593
---

# Pull Request #593: Fixes 3664: support listing/fetching RH tasks

**Author**: @jlsherrill
**Created**: March 04, 2024 at 07:49 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3664`

## Description

## Summary

allows all orgs to fetch tasks for RH repos

## Testing steps

Ensure a RH repo has been snapshotted, List tasks:

GET /api/content-sources/v1.0/tasks/

and then fetch one's UUID:

GET /api/content-sources/v1.0/tasks/UUID/

Bonus points if you try to cancel a RH task:

POST /api/content-sources/v1.0/tasks/uuid/cancel/

(you shouldn't be able to)

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 04, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-3664

### Comment by @swadeley on March 06, 2024 at 09:36 AM UTC

/retest

### Comment by @swadeley on March 06, 2024 at 01:27 PM UTC

/retest

### Comment by @swadeley on March 06, 2024 at 02:30 PM UTC

@jlsherrill rebase please

### Comment by @swadeley on March 06, 2024 at 04:09 PM UTC

/retest

### Comment by @xbhouse on March 06, 2024 at 04:47 PM UTC

this looks good to me for fetching RH tasks. i haven't been able to test canceling one, what does the curl command look like to do that? i don't see the request in the api docs so i've just tried this which isn't found:

`curl -X POST -H "$( ./scripts/header.sh 9999 1111 )" http://localhost:8000/api/content-sources/v1.0/tasks/76801a1c-ba93-49f5-a0b5-8e1f89369f7a/cancel`

this (with the slash at the end) causes a panic:

`curl -X POST -H "$( ./scripts/header.sh 9999 1111 )" http://localhost:8000/api/content-sources/v1.0/tasks/76801a1c-ba93-49f5-a0b5-8e1f89369f7a/cancel/`

### Comment by @jlsherrill on March 06, 2024 at 07:30 PM UTC

i'll file a bug around the lack of api docs, the correct api is  http://localhost:8000/api/content-sources/v1.0/tasks/76801a1c-ba93-49f5-a0b5-8e1f89369f7a/cancel/

what kind of panic did you get?

### Comment by @xbhouse on March 06, 2024 at 09:05 PM UTC

oh ok. hmm weird, looked like it's panicking due to a null pointer dereference:

`http: panic serving [::1]:52305: runtime error: invalid memory address or nil pointer dereference`

once the task finished the same API call responded with `error canceling task, detail=Could not find task with UUID 76801a1c-ba93-49f5-a0b5-8e1f89369f7a`

if you're not seeing this though it could be something with my config that's wonky

---

## Reviews

### Review by @xbhouse - Approved on March 07, 2024 at 03:38 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/593*
