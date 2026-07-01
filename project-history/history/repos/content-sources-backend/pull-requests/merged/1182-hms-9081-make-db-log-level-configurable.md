---
type: pull_request
number: 1182
title: "HMS-9081: make db log level configurable"
state: merged
author: rverdile
created: 2025-08-19T18:32:30Z
updated: 2025-08-20T15:54:15Z
closed: 2025-08-20T15:54:12Z
merged: 2025-08-20T15:54:12Z
base_branch: main
head_branch: no-log-ctx-cancel
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1182
---

# Pull Request #1182: HMS-9081: make db log level configurable

**Author**: @rverdile
**Created**: August 19, 2025 at 06:32 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `no-log-ctx-cancel`

## Description

## Summary
- Makes the DB log level optionally configurable because the number of logs is freezing up my laptop
- Gorm now ignores "context canceled" error logs 

## Testing steps
1. Use this to test the context canceling. It lists repositories and the cancels it quickly: https://gist.github.com/rverdile/646a8576d3919f18ab1998f7cb2682cd
2. Without this PR, gorm will log the context is canceled
3. Change the db_level option under logging to change it from the global logging level


---

## Discussion

### Comment by @jlsherrill on August 19, 2025 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-9081

### Comment by @swadeley on August 20, 2025 at 09:08 AM UTC

/retest

### Comment by @rverdile on August 20, 2025 at 01:16 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on August 19, 2025 at 08:45 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1182*
