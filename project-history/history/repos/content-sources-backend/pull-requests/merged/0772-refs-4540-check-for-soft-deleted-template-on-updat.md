---
type: pull_request
number: 772
title: "Refs 4540: check for soft deleted template on update"
state: merged
author: jlsherrill
created: 2024-08-12T21:07:41Z
updated: 2024-08-27T17:22:29Z
closed: 2024-08-27T17:22:25Z
merged: 2024-08-27T17:22:25Z
base_branch: main
head_branch: 4540
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/772
---

# Pull Request #772: Refs 4540: check for soft deleted template on update

**Author**: @jlsherrill
**Created**: August 12, 2024 at 09:07 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4540`

## Description

## Summary
On the template update task, check to see if the template has been softdeleted, and simple return from the task if so

## Testing steps

With your server running
Create a repo with snapshotting

Run the server with just the api:
`go run ./cmd/content-sources/main.go api`

Create a template with the repo:
```
####
POST http://localhost:8000/api/content-sources/v1.0/templates/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json

{
  "name":"test5",
  "arch": "x86_64",
  "version": "9",
  "repository_uuids": ["cd79da0f-dea2-46d4-82c4-56d3012db8bb"],
  "date": "2024-08-09T14:04:05Z"
}
```

delete the repo:
```
####
DELETE http://localhost:8000/api/content-sources/v1.0/templates/08273910-70b5-45e0-99a6-57706f61efe5
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json
```

Now stop the server and restart it with the worker:
`go run ./cmd/content-sources/main.go api consumer`

This will cause the update and delete task to run at the same time.  Monitor the server logs and look for `Could not find template with UUID d95ee14e-00a1-40fa-be8f-27edfb5e043e`, you shouldn't see it.



---

## Discussion

### Comment by @jlsherrill on August 12, 2024 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-4540

### Comment by @swadeley on August 27, 2024 at 12:21 PM UTC

Hi @jlsherrill 

please rebase to pick up change in https://github.com/content-services/content-sources-backend/pull/788

Thank you

### Comment by @jlsherrill on August 27, 2024 at 12:45 PM UTC

[test]

### Comment by @jlsherrill on August 27, 2024 at 01:15 PM UTC

[test]

### Comment by @jlsherrill on August 27, 2024 at 03:19 PM UTC

[test]

---

## Reviews

### Review by @rverdile - Commented on August 13, 2024 at 05:17 PM UTC

### Review by @jlsherrill - Commented on August 13, 2024 at 05:40 PM UTC

### Review by @jlsherrill - Commented on August 13, 2024 at 05:41 PM UTC

### Review by @rverdile - Commented on August 13, 2024 at 06:38 PM UTC

### Review by @jlsherrill - Commented on August 15, 2024 at 08:29 PM UTC

### Review by @rverdile - Commented on August 19, 2024 at 02:50 PM UTC

### Review by @rverdile - Commented on August 19, 2024 at 02:50 PM UTC

looks good. one comment and needs rebase

### Review by @jlsherrill - Commented on August 26, 2024 at 02:46 PM UTC

### Review by @rverdile - Approved on August 26, 2024 at 06:37 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/772*
