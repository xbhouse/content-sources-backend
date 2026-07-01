---
type: pull_request
number: 713
title: "Refs 4256: change snapshot interval to 24 hours"
state: merged
author: jlsherrill
created: 2024-06-20T14:32:12Z
updated: 2024-06-20T17:23:42Z
closed: 2024-06-20T17:23:16Z
merged: 2024-06-20T17:23:16Z
base_branch: main
head_branch: 4256_24
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/713
---

# Pull Request #713: Refs 4256: change snapshot interval to 24 hours

**Author**: @jlsherrill
**Created**: June 20, 2024 at 02:32 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4256_24`

## Description

## Summary

By changing it to 24 hours, we are forcing snapshotting to happen only if they haven't been snapshotted in 24 hours, not 20.  Since the job is running hourly, it won't now grab ~3 hours worth of repos

## Testing steps

None right now, I think we should merge and monitor the behavior

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 20, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4256

### Comment by @jlsherrill on June 20, 2024 at 05:23 PM UTC

due to the simplicity i'm gonna go ahead and merge

---

## Reviews

### Review by @swadeley - Approved on June 20, 2024 at 02:39 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/713*
