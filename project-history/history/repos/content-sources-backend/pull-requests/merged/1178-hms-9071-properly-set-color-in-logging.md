---
type: pull_request
number: 1178
title: "HMS-9071: Properly set color in logging"
state: merged
author: jlsherrill
created: 2025-08-15T15:19:45Z
updated: 2025-08-19T11:14:26Z
closed: 2025-08-19T11:14:26Z
merged: 2025-08-19T11:14:26Z
base_branch: main
head_branch: 9071
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1178
---

# Pull Request #1178: HMS-9071: Properly set color in logging

**Author**: @jlsherrill
**Created**: August 15, 2025 at 03:19 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `9071`

## Description

## Summary

makes logger coloring optional and defaults to off. Sets this in both zerolog and in  gorm

## Testing steps

set color to true in logging
```
logging:
  level: info
  color: true
```
start server, see 🌈 .  Set it to false and try again, see ⚫ ⚪ 





---

## Discussion

### Comment by @jlsherrill on August 15, 2025 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-9071

### Comment by @swadeley on August 18, 2025 at 07:46 AM UTC

/retest

### Comment by @jlsherrill on August 18, 2025 at 12:54 PM UTC

/retest

### Comment by @jlsherrill on August 19, 2025 at 01:33 AM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on August 18, 2025 at 01:53 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1178*
