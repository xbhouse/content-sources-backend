---
type: pull_request
number: 656
title: "Fixes 4055: treat filter/page param parse as info"
state: merged
author: jlsherrill
created: 2024-05-01T14:38:18Z
updated: 2024-05-02T18:50:16Z
closed: 2024-05-02T13:52:26Z
merged: 2024-05-02T13:52:25Z
base_branch: main
head_branch: 4055
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/656
---

# Pull Request #656: Fixes 4055: treat filter/page param parse as info

**Author**: @jlsherrill
**Created**: May 01, 2024 at 02:38 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4055`

## Description

## Summary

Changes filter/pagination parsing errors as info level logging instead of error

## Testing steps

perform a request:
GET http://localhost:8000/api/content-sources/v1.0/templates/?limit=a

watch your server, it should show logging at INFO

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 01, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4055

### Comment by @xbhouse on May 01, 2024 at 03:50 PM UTC

i see `11:40AM ERR ../../go/pkg/mod/github.com/content-services/lecho/v3@v3.5.2/middleware.go:179 > error="error: code=500, title=Error listing templates, detail=ERROR: invalid byte sequence for encoding \"UTF8\": 0x00 (SQLSTATE 22021) \n"` when sending a request like `/templates/?name=%00`

do you need to make these updates to the templates handler as well? and maybe the task_info handler, though invalid syntax (`/tasks/?status=%00`) logs as DBG there so i'm not sure

### Comment by @jlsherrill on May 01, 2024 at 07:14 PM UTC

Good catch!  The template actually had a function to detect this, but 1) it wasn't being called & 2) it wasn't checking for a query error properly, this should be fixed!

---

## Reviews

### Review by @xbhouse - Approved on May 02, 2024 at 01:12 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/656*
