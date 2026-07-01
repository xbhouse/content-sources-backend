---
type: pull_request
number: 806
title: "Refs 4540: prevent requeue on exit for canceled task"
state: merged
author: rverdile
created: 2024-09-09T17:33:30Z
updated: 2024-09-11T19:06:00Z
closed: 2024-09-11T19:05:57Z
merged: 2024-09-11T19:05:57Z
base_branch: main
head_branch: template-uuid
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/806
---

# Pull Request #806: Refs 4540: prevent requeue on exit for canceled task

**Author**: @rverdile
**Created**: September 09, 2024 at 05:33 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `template-uuid`

## Description

## Summary
It was still possible for a canceled task to get requeued on server exit. This fixes that.

It shouldn't be possible for the HeartbeatListener() method to requeue canceled tasks because heartbeats are removed as part of canceling the task. But if this fix doesn't work, that's the next place to look. 

## Testing steps
1. Create and immediately delete a template. Need to replace repository uuid in create request below.

```
UUID=`curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/templates/"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzc5MTU2MCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJzbmFwVXNlciJ9LCJhY2NvdW50X251bWJlciI6IjE3NzkxNTYwIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTc3OTE1NjAifX19Cg=="     -H "Content-Type: application/json"     -d "{
          \"name\":\"test-$RANDOM\",
          \"arch\": \"x86_64\",
          \"version\": \"9\",
          \"repository_uuids\": [\"2f859dec-184b-48d4-9451-140a2336ce4d\"],
          \"date\": \"2024-08-09T14:04:05Z\"
        }"  | jq .uuid`


curl -X DELETE --location "http://localhost:8000/api/content-sources/v1.0/templates/$UUID"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzc5MTU2MCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJzbmFwVXNlciJ9LCJhY2NvdW50X251bWJlciI6IjE3NzkxNTYwIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTc3OTE1NjAifX19Cg=="
```
2. CTRL+C exit the server. Before these changes, the update-template-content task would get requeued.

---

## Discussion

### Comment by @jlsherrill on September 09, 2024 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-4540

### Comment by @jlsherrill on September 11, 2024 at 02:29 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on September 11, 2024 at 05:37 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/806*
