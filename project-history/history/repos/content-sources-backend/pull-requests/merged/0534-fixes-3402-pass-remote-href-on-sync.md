---
type: pull_request
number: 534
title: "Fixes 3402: pass remote href on sync"
state: merged
author: jlsherrill
created: 2024-01-15T20:31:08Z
updated: 2024-01-16T21:30:37Z
closed: 2024-01-16T21:21:26Z
merged: 2024-01-16T21:21:26Z
base_branch: main
head_branch: 3402
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/534
---

# Pull Request #534: Fixes 3402: pass remote href on sync

**Author**: @jlsherrill
**Created**: January 15, 2024 at 08:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3402`

## Description

## Summary

pass the remote href when doing a sync, as pulp has an rbac bug preventing us from using the remote on the repo

## Testing steps

integration tests pass

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 15, 2024 at 08:59 PM UTC

/retest

### Comment by @jlsherrill on January 15, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-3402

### Comment by @jlsherrill on January 15, 2024 at 09:35 PM UTC

/retest

### Comment by @jlsherrill on January 16, 2024 at 06:23 PM UTC

/retest

### Comment by @jlsherrill on January 16, 2024 at 08:39 PM UTC

/retest

### Comment by @jlsherrill on January 16, 2024 at 09:21 PM UTC

merging to fix ephemeral and stage

---

## Reviews

### Review by @rverdile - Approved on January 15, 2024 at 09:38 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/534*
