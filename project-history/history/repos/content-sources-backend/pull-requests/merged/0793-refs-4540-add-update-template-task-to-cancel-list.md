---
type: pull_request
number: 793
title: "Refs 4540: add update template task to cancel list"
state: merged
author: jlsherrill
created: 2024-08-29T11:47:17Z
updated: 2024-08-29T14:12:50Z
closed: 2024-08-29T14:11:18Z
merged: 2024-08-29T14:11:18Z
base_branch: main
head_branch: 4540_fix
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/793
---

# Pull Request #793: Refs 4540: add update template task to cancel list

**Author**: @jlsherrill
**Created**: August 29, 2024 at 11:47 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4540_fix`

## Description

## Summary

Somehow this didn't make it to the cancel list

## Testing steps

Create a repository and grab the UUID,

run this, replacing 8f84b19e-675a-4af2-83c2-b2597f16ae8c with the repo uuid :
```
UUID=`curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/templates/"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg=="     -H "Content-Type: application/json"     -d "{
          \"name\":\"test-$RANDOM\",
          \"arch\": \"x86_64\",
          \"version\": \"9\",
          \"repository_uuids\": [\"8f84b19e-675a-4af2-83c2-b2597f16ae8c\"],
          \"date\": \"2024-08-09T14:04:05Z\"
        }"  | jq .uuid`


curl -X DELETE --location "http://localhost:8000/api/content-sources/v1.0/templates/$UUID"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg=="
```



---

## Discussion

### Comment by @jlsherrill on August 29, 2024 at 12:00 PM UTC

https://issues.redhat.com/browse/HMS-4540

---

## Reviews

### Review by @rverdile - Approved on August 29, 2024 at 01:35 PM UTC

works!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/793*
