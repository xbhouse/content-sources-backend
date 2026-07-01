---
type: pull_request
number: 1251
title: "HMS-9203: properly detect admin urls"
state: merged
author: jlsherrill
created: 2025-10-24T15:18:45Z
updated: 2025-10-24T17:07:26Z
closed: 2025-10-24T17:07:23Z
merged: 2025-10-24T17:07:23Z
base_branch: main
head_branch: 9203
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1251
---

# Pull Request #1251: HMS-9203: properly detect admin urls

**Author**: @jlsherrill
**Created**: October 24, 2025 at 03:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `9203`

## Description

## Summary

The previous change didn't fully fix the problem because /admin/tasks/:task_uuid wasn't a valid path, instead it was /admin/tasks/:uuid

Also adding the features paths

## Testing steps

on a server with some tasks already created , find a task uuid
```
make db-cli-connect
# select id from tasks;
```

Then curl the admin api:
```
curl -X GET --location "http://localhost:8000/api/content-sources/v1.0/admin/tasks/d38ee12e-b589-42a5-b844-0de6c03344e9" \
    -H "Content-Type: application/json" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg=="
```


---

## Discussion

### Comment by @xbhouse on October 24, 2025 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-9203

---

## Reviews

### Review by @xbhouse - Approved on October 24, 2025 at 03:59 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1251*
