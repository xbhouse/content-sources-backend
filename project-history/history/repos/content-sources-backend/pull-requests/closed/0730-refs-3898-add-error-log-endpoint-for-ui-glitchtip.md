---
type: pull_request
number: 730
title: "Refs 3898: add error log endpoint for UI glitchtip"
state: closed
author: rverdile
created: 2024-07-05T20:19:45Z
updated: 2024-07-08T13:53:43Z
closed: 2024-07-08T13:53:43Z
base_branch: main
head_branch: ui-glitchtip
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/730
---

# Pull Request #730: Refs 3898: add error log endpoint for UI glitchtip

**Author**: @rverdile
**Created**: July 05, 2024 at 08:19 PM UTC
**Status**: Closed
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `ui-glitchtip`

## Description

## Summary
Adds an endpoint that logs the error sent to it, so that the frontend can send errors to it, and we can get glitchtip notifications for those errors

See https://github.com/content-services/content-sources-frontend/pull/296

POST request to `/content-sources/v1/log_error`

with the body 
```
{
    "error_title": "UI error",
    "error_details": "an error has occurred [might include stack trace]"
}
```

## Testing steps
1. Send a POST request with the body above to the new endpoint
2. In the terminal, you should see an error log

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 05, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-3898

### Comment by @jlsherrill on July 05, 2024 at 08:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @rverdile on July 08, 2024 at 01:53 PM UTC

closing in favor of waiting for official support

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/730*
