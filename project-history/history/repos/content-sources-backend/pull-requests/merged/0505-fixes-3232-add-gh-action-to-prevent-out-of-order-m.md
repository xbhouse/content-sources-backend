---
type: pull_request
number: 505
title: "Fixes 3232: add gh action to prevent out of order migrations"
state: merged
author: xbhouse
created: 2023-12-14T15:44:34Z
updated: 2024-02-13T15:16:08Z
closed: 2024-01-03T15:42:18Z
merged: 2024-01-03T15:42:18Z
base_branch: main
head_branch: gh-action-migrations
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/505
---

# Pull Request #505: Fixes 3232: add gh action to prevent out of order migrations

**Author**: @xbhouse
**Created**: December 14, 2023 at 03:44 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `gh-action-migrations`

## Description

## Summary

Added Github Action to check for out of order database migrations.

## Testing steps

If the content in `db/migrations.latest` does not reflect the timestamp of the newest added set of migration files in `db/migrations`, the job will fail. If these timestamps match or if there are no new migrations, the job will succeed. 

Can test by:
- Adding a set of test migration files with an earlier timestamp than the latest and not updating the `db/migrations.latest`. This should cause the job to fail.
- Removing the test migration files. The `db/migrations.latest` content should reflect the newest added set of migration files. This should cause the job to succeed.
- Adding a set of test migration files with a timestamp that matches the timestamp in `db/migrations.latest`. This should cause the job to succeed.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on December 14, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-3232

---

## Reviews

### Review by @jlsherrill - Approved on December 21, 2023 at 03:53 PM UTC

ACK, nice approach! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/505*
