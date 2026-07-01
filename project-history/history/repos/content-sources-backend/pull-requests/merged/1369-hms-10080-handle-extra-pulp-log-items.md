---
type: pull_request
number: 1369
title: "HMS-10080: handle extra pulp log items"
state: merged
author: jlsherrill
created: 2026-01-20T15:42:32Z
updated: 2026-01-20T16:05:49Z
closed: 2026-01-20T16:01:34Z
merged: 2026-01-20T16:01:34Z
base_branch: main
head_branch: 10080
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1369
---

# Pull Request #1369: HMS-10080: handle extra pulp log items

**Author**: @jlsherrill
**Created**: January 20, 2026 at 03:42 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `10080`

## Description

## Summary

When parsing pulp log entries, currently if the pulp team adds more items to the log entry, we will fail to parse it.  This make this more robust by ignoring anything after request_org_id.  

## Testing steps

Tests should pass, i added a test to cover this situation.



---

## Discussion

### Comment by @xbhouse on January 20, 2026 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-10080

---

## Reviews

### Review by @rverdile - Approved on January 20, 2026 at 03:55 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1369*
