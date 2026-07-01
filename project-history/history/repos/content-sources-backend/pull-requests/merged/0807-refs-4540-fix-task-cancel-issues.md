---
type: pull_request
number: 807
title: "Refs 4540: fix task cancel issues"
state: merged
author: jlsherrill
created: 2024-09-09T23:30:13Z
updated: 2024-09-11T19:43:11Z
closed: 2024-09-11T19:40:04Z
merged: 2024-09-11T19:40:04Z
base_branch: main
head_branch: 4540_3
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/807
---

# Pull Request #807: Refs 4540: fix task cancel issues

**Author**: @jlsherrill
**Created**: September 09, 2024 at 11:30 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4540_3`

## Description

## Summary
* improper context passing
* lack of wrapping DaoError errors
* Check for closed transaction error

## Testing steps


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


Do this about 20 times

---

## Discussion

### Comment by @jlsherrill on September 10, 2024 at 12:00 AM UTC

https://issues.redhat.com/browse/HMS-4540

### Comment by @jlsherrill on September 10, 2024 at 06:47 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on September 11, 2024 at 07:31 PM UTC

looks good, not seeing the error

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/807*
