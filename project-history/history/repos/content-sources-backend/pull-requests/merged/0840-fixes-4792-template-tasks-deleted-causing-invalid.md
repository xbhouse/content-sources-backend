---
type: pull_request
number: 840
title: "Fixes 4792: template tasks deleted causing invalid status"
state: merged
author: xbhouse
created: 2024-10-08T15:26:19Z
updated: 2024-10-09T13:55:57Z
closed: 2024-10-09T13:55:57Z
merged: 2024-10-09T13:55:57Z
base_branch: main
head_branch: 4792
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/840
---

# Pull Request #840: Fixes 4792: template tasks deleted causing invalid status

**Author**: @xbhouse
**Created**: October 08, 2024 at 03:26 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4792`

## Description

## Summary

- Changed task cleanup to not delete tasks if they are the last task that updated a template

## Testing steps

1. Create a custom repo, RH repo, let them snapshot, and create a template with those repos
2. Update the task that just updated the template to a date at least 20 days prior
```
UPDATE tasks
SET queued_at = '2024-09-08 15:24:52.793340 +00:00', started_at = '2024-09-09 15:24:52.793340 +00:00', finished_at = '2024-09-10 15:24:52.793340 +00:00'
WHERE id = '<task id>';
```
3. Run `go run cmd/external-repos/main.go nightly-jobs`
4. The `update-template-content` task that just updated the template should not have been cleaned up. And in the UI, the status of the template should still be valid



---

## Discussion

### Comment by @jlsherrill on October 08, 2024 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-4792

---

## Reviews

### Review by @jlsherrill - Approved on October 08, 2024 at 06:27 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/840*
