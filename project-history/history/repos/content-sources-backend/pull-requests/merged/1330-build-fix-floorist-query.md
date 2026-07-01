---
type: pull_request
number: 1330
title: "Build: fix floorist query"
state: merged
author: jlsherrill
created: 2025-12-09T14:47:55Z
updated: 2025-12-10T12:09:40Z
closed: 2025-12-10T12:09:40Z
merged: 2025-12-10T12:09:40Z
base_branch: main
head_branch: query_fix
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1330
---

# Pull Request #1330: Build: fix floorist query

**Author**: @jlsherrill
**Created**: December 09, 2025 at 02:47 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `query_fix`

## Description

## Summary

mistakenly used wrong table alias for summary:

```
sqlalchemy.exc.ProgrammingError: (psycopg2.errors.UndefinedColumn) column rc.summary does not exist
LINE 1: ...pms.version, rpms.release, rpms.arch, rpms.epoch, rc.summary...
^
HINT: Perhaps you meant to reference the column "rpms.summary".
```

## Testing steps



---

## Reviews

### Review by @TenSt - Approved on December 09, 2025 at 03:41 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1330*
