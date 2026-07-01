---
type: pull_request
number: 234
title: "Fixes 1440: update dependencies"
state: merged
author: rverdile
created: 2023-03-22T20:36:43Z
updated: 2023-03-23T20:13:13Z
closed: 2023-03-23T20:13:07Z
merged: 2023-03-23T20:13:07Z
base_branch: main
head_branch: up-deps
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/234
---

# Pull Request #234: Fixes 1440: update dependencies

**Author**: @rverdile
**Created**: March 22, 2023 at 08:36 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `up-deps`

## Description

## Summary
There is a change here: https://github.com/go-gorm/postgres (which was causing our tests to fail), so I had to pin it to the previous version (v1.4.8). There's an issue open for it, so hopefully it gets resolved.

It's causing a bug for us because gorm's error was no longer returning as type `pgconn.PgError`, but `errors.errorString` instead, which is a problem here: https://github.com/content-services/content-sources-backend/blob/main/pkg/dao/repository_configs.go#L31
## Testing steps
none


---

## Discussion

### Comment by @jlsherrill on March 22, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-1440

---

## Reviews

### Review by @rverdile - Commented on March 22, 2023 at 08:39 PM UTC

### Review by @jlsherrill - Approved on March 23, 2023 at 03:43 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/234*
