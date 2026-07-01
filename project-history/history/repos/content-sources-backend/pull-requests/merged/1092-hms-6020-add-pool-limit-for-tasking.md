---
type: pull_request
number: 1092
title: "HMS-6020: add pool limit for tasking"
state: merged
author: jlsherrill
created: 2025-04-25T17:09:39Z
updated: 2025-04-26T17:03:49Z
closed: 2025-04-26T17:03:38Z
merged: 2025-04-26T17:03:38Z
base_branch: main
head_branch: 6020
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1092
---

# Pull Request #1092: HMS-6020: add pool limit for tasking

**Author**: @jlsherrill
**Created**: April 25, 2025 at 05:09 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `6020`

## Description

## Summary

and add a warning if its too small

## Testing steps

in config.yaml set: 
```
  worker_count: 25
  pool_limit: 30
```

then import and sync all repos:

```
make repos-import
go run cmd/external-repos/main.go process-repos
```


see that they all complete



---

## Discussion

### Comment by @jlsherrill on April 25, 2025 at 05:34 PM UTC

https://issues.redhat.com/browse/HMS-6020

### Comment by @jlsherrill on April 26, 2025 at 12:36 AM UTC

/retest

### Comment by @jlsherrill on April 26, 2025 at 05:03 PM UTC

this is actually no-qe-needed (just updated card)

---

## Reviews

### Review by @rverdile - Approved on April 25, 2025 at 06:04 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1092*
