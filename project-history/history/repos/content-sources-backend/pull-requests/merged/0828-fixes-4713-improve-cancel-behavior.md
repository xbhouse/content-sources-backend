---
type: pull_request
number: 828
title: "Fixes 4713: improve cancel behavior"
state: merged
author: rverdile
created: 2024-09-19T19:49:58Z
updated: 2024-09-25T19:56:33Z
closed: 2024-09-25T19:56:30Z
merged: 2024-09-25T19:56:30Z
base_branch: main
head_branch: template-not-found
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/828
---

# Pull Request #828: Fixes 4713: improve cancel behavior

**Author**: @rverdile
**Created**: September 19, 2024 at 07:49 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `template-not-found`

## Description

## Summary
Now cancels failed tasks, multiple tasks, and fixes bug in worker's cancel check.

## Testing steps
1. Start the server
2. Run this in a script, replacing the headers and repository uuid in the template create request 
```
UUID=`curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/templates/"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzc5MTU2MCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJzbmFwVXNlciJ9LCJhY2NvdW50X251bWJlciI6IjE3NzkxNTYwIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTc3OTE1NjAifX19Cg=="     -H "Content-Type: application/json"     -d "{
          \"name\":\"test-$RANDOM\",
          \"arch\": \"x86_64\",
          \"version\": \"9\",
          \"repository_uuids\": [\"acdb2f91-9666-4600-8662-45bab8b7bbbe\"],
          \"date\": \"2024-08-09T14:04:05Z\"
        }"  | jq .uuid`


echo $UUID

curl -X PATCH --location "http://localhost:8000/api/content-sources/v1.0/templates/$UUID"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzc5MTU2MCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJzbmFwVXNlciJ9LCJhY2NvdW50X251bWJlciI6IjE3NzkxNTYwIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTc3OTE1NjAifX19Cg=="     -H "Content-Type: application/json"     -d "{
          \"name\":\"test-$RANDOM\",
          \"arch\": \"x86_64\",
          \"version\": \"9\",
          \"date\": \"2024-08-09T14:04:05Z\"
        }"

sleep .025

curl -X DELETE --location "http://localhost:8000/api/content-sources/v1.0/templates/$UUID"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzc5MTU2MCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJzbmFwVXNlciJ9LCJhY2NvdW50X251bWJlciI6IjE3NzkxNTYwIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTc3OTE1NjAifX19Cg=="
```
3. After running this several times, you should not see the "Could not find template with UUID 5c14723e-709f-447d-ba23-eff2ea4d154b" error logged anymore.


---

## Discussion

### Comment by @jlsherrill on September 19, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-4713

---

## Reviews

### Review by @rverdile - Commented on September 19, 2024 at 07:50 PM UTC

### Review by @jlsherrill - Approved on September 25, 2024 at 06:43 PM UTC

ACK, just needs a rebase



---

*Archived from: https://github.com/content-services/content-sources-backend/pull/828*
