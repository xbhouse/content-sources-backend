---
type: pull_request
number: 550
title: "Fixes 3278: migrate tasking to pgx"
state: merged
author: jlsherrill
created: 2024-01-26T15:37:10Z
updated: 2024-01-30T17:39:27Z
closed: 2024-01-30T17:39:24Z
merged: 2024-01-30T17:39:23Z
base_branch: main
head_branch: 3278
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/550
---

# Pull Request #550: Fixes 3278: migrate tasking to pgx

**Author**: @jlsherrill
**Created**: January 26, 2024 at 03:37 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3278`

## Description

## Summary

migrates our tasking system to pgxv5

## Testing steps

1) test the app normally, triggering a snapshot and watch it complete without error
2) Turn on pgx logging and set log level to trace:

```
logging:
  level: trace
  console: true

tasking:
  pgx_logging: true
  heartbeat: 1m
  worker_count: 3
```

verify that you see tasking logs

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 26, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-3278

### Comment by @jlsherrill on January 29, 2024 at 12:53 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on January 29, 2024 at 09:12 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/550*
