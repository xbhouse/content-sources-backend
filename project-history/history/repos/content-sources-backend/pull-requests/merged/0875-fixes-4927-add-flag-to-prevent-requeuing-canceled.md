---
type: pull_request
number: 875
title: "Fixes 4927: add flag to prevent requeuing canceled tasks"
state: merged
author: xbhouse
created: 2024-11-06T17:02:04Z
updated: 2024-11-12T17:18:56Z
closed: 2024-11-12T17:18:56Z
merged: 2024-11-12T17:18:56Z
base_branch: main
head_branch: 4927
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/875
---

# Pull Request #875: Fixes 4927: add flag to prevent requeuing canceled tasks

**Author**: @xbhouse
**Created**: November 06, 2024 at 05:02 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4927`

## Description

## Summary

Adds a `cancel_attempted` field to the tasks table and prevents requeuing a task if it was canceled

## Testing steps

1. On main, create a repo and let it snapshot
2. You can use this script to reproduce the "could not find template with UUID" errors (you'd need to adjust it to use your identity header and change the `repository_uuids` to include the repo you've created). Running this on main should reproduce those errors (you may not need to loop as many times):

```
for i in {1..50}; do
  UUID=$(curl -s -X POST --location "http://localhost:8000/api/content-sources/v1.0/templates/"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzY4NDYzMiIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJicnl0dGFuaWVob3VzZSJ9LCJhY2NvdW50X251bWJlciI6IjExNDkxNzg2IiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTc2ODQ2MzIifX19Cg=="     -H "Content-Type: application/json"     -d "{
          \"name\":\"test-$RANDOM\",
          \"arch\": \"x86_64\",
          \"version\": \"8\",
          \"repository_uuids\": [\"0de58827-5301-44da-a226-85db660d6d1f\"],
          \"date\": \"2024-11-05T14:04:05Z\"
        }"  | jq -r .uuid)

  curl -s -X DELETE --location "http://localhost:8000/api/content-sources/v1.0/templates/$UUID"     -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzY4NDYzMiIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJicnl0dGFuaWVob3VzZSJ9LCJhY2NvdW50X251bWJlciI6IjExNDkxNzg2IiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMTc2ODQ2MzIifX19Cg=="

done
```
3. Checkout this PR, run the migration, and run the script again. You shouldn't see any requeues or "could not find template with UUID" errors

---

## Discussion

### Comment by @jlsherrill on November 06, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-4927

---

## Reviews

### Review by @rverdile - Commented on November 06, 2024 at 05:38 PM UTC

### Review by @xbhouse - Commented on November 06, 2024 at 06:07 PM UTC

### Review by @rverdile - Commented on November 07, 2024 at 06:10 PM UTC

### Review by @rverdile - Commented on November 07, 2024 at 06:11 PM UTC

### Review by @rverdile - Commented on November 07, 2024 at 06:13 PM UTC

### Review by @rverdile - Commented on November 07, 2024 at 06:14 PM UTC

### Review by @rverdile - Approved on November 11, 2024 at 03:21 PM UTC

nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/875*
