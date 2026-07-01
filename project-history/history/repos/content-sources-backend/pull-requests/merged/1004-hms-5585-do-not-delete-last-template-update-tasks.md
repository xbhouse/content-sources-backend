---
type: pull_request
number: 1004
title: "HMS-5585: do not delete last template update tasks"
state: merged
author: jlsherrill
created: 2025-03-03T21:41:15Z
updated: 2025-03-04T19:01:10Z
closed: 2025-03-04T19:00:54Z
merged: 2025-03-04T19:00:54Z
base_branch: main
head_branch: 5585
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1004
---

# Pull Request #1004: HMS-5585: do not delete last template update tasks

**Author**: @jlsherrill
**Created**: March 03, 2025 at 09:41 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5585`

## Description

## Summary

Fixes a regression where by templates last update tasks were being deleted

## Testing steps
import and snapshot rhle9 repos
Create a rhel 9 template 
update the dates:

```
update tasks set queued_at =  now() - interval '30 day';
update tasks set started_at =  now() - interval '30 day';
update tasks set finished_at =  now() - interval '30 day';
```

run process-repos:
```
go run cmd/external-repos/main.go process-repos
```

On main, you'll see the task get deleted and the template will show up in the UI as 'invalid'.  "edit" the templates to get them back to a green status and then repeat on this PR.

---

## Discussion

### Comment by @jlsherrill on March 03, 2025 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-5585

---

## Reviews

### Review by @xbhouse - Approved on March 04, 2025 at 03:28 PM UTC

looks good, nice catch!! 

### Review by @swadeley - Approved on March 04, 2025 at 06:14 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1004*
