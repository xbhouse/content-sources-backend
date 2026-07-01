---
type: pull_request
number: 402
title: "Fixes 2490: make migrations idempotent"
state: merged
author: rverdile
created: 2023-09-22T17:29:18Z
updated: 2023-10-02T13:50:29Z
closed: 2023-10-02T13:50:24Z
merged: 2023-10-02T13:50:24Z
base_branch: main
head_branch: if-not-exists
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/402
---

# Pull Request #402: Fixes 2490: make migrations idempotent

**Author**: @rverdile
**Created**: September 22, 2023 at 05:29 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `if-not-exists`

## Description

## Summary
Makes it so migrations be be re-run on an already migrated database. Also fixed a few down migrations that were not working. Adds a CI test for migrations.

## Testing steps
1. Migrate up
2. Migrate down
3. Migrate up
4. Drop the `schema_migrations` table
5. Migrate up again

---

## Discussion

### Comment by @jlsherrill on September 22, 2023 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-2490

---

## Reviews

### Review by @Andrewgdewar - Dismissed on September 26, 2023 at 05:38 PM UTC

Tested, no errors, migrated successfully.

### Review by @Andrewgdewar - Approved on September 28, 2023 at 07:58 PM UTC

Reviewed test locally, passed without issues. 
CI is passing as well. 
Ack from me!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/402*
