---
type: pull_request
number: 350
title: "Fixes 2270: prevent out of order migrations"
state: merged
author: jlsherrill
created: 2023-08-03T18:54:03Z
updated: 2023-08-04T15:09:21Z
closed: 2023-08-04T13:25:49Z
merged: 2023-08-04T13:25:49Z
base_branch: main
head_branch: 2270
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/350
---

# Pull Request #350: Fixes 2270: prevent out of order migrations

**Author**: @jlsherrill
**Created**: August 03, 2023 at 06:54 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `2270`

## Description

## Summary

1.  Creates a new file ./db/migrations.latest which holds the timestamp of the latest migration
2. on db-migrate-up, the command errors out if the value in   ./db/migrations.latest doesn't match the latest in migration in ./db/migrations
3.  `go run ./cmd/dbmigrate/go.mod  new my_migration'  will auto-populate migrations.latest with the newest timestamp
4. The result is that prs should create conflicts on merge if they both contain migrations, ensuring that we never commit any that are out of order.

## Testing steps

Try generating a new migration without updating migrations.latest.  Try to migrate up (you should get an error).



---

## Discussion

### Comment by @jlsherrill on August 03, 2023 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-2270

---

## Reviews

### Review by @rverdile - Commented on August 03, 2023 at 10:33 PM UTC

### Review by @rverdile - Approved on August 03, 2023 at 10:34 PM UTC

found 1 typo but otherwise lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/350*
