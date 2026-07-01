---
type: pull_request
number: 908
title: "Fixes 5058: separate set cancel from set cancel_attempted"
state: merged
author: rverdile
created: 2024-11-25T20:10:25Z
updated: 2024-11-27T18:25:20Z
closed: 2024-11-27T18:25:17Z
merged: 2024-11-27T18:25:17Z
base_branch: main
head_branch: cancel-query
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/908
---

# Pull Request #908: Fixes 5058: separate set cancel from set cancel_attempted

**Author**: @rverdile
**Created**: November 25, 2024 at 08:10 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `cancel-query`

## Description

## Summary
The "template not found" error is still showing up because of an edge case where a task is failed before it is canceled

I missed this case in the discussion [here](https://github.com/content-services/content-sources-backend/pull/875#discussion_r1831462633
). Bryttanie had it right. :)


## Testing steps
I could not reproduce this locally, but the unit test is able to test this scenario

To reproduce this locally, you could try running the following as a script, while also modifying that update-template-content task to fail. But this did not work for me.

```
UUID=$(curl -s -X POST --location "http://localhost:8000/api/content-sources/v1.0/templates/"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzc5MTU2MCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJzbmFwVXNlciJ9LCJhY2NvdW50X251bWJlciI6IjE3NzkxNTYwIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTc3OTE1NjAifX19Cg=="     -H "Content-Type: application/json"     -d "{
          \"name\":\"test-$RANDOM\",
          \"arch\": \"x86_64\",
          \"version\": \"8\",
          \"repository_uuids\": [\"917abddc-0745-4752-a100-ec147c7ee34b\"],
          \"date\": \"2024-08-09T14:04:05Z\"
        }"  | jq -r .uuid)

curl -s -X DELETE --location "http://localhost:8000/api/content-sources/v1.0/templates/$UUID"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzc5MTU2MCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJzbmFwVXNlciJ9LCJhY2NvdW50X251bWJlciI6IjE3NzkxNTYwIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTc3OTE1NjAifX19Cg=="
```

---

## Discussion

### Comment by @jlsherrill on November 25, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-5058

---

## Reviews

### Review by @xbhouse - Approved on November 27, 2024 at 04:08 PM UTC

ack! 🤞 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/908*
