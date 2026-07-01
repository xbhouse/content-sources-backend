---
type: pull_request
number: 344
title: "Fixes 2180: add index for tasks query"
state: merged
author: rverdile
created: 2023-08-02T16:05:09Z
updated: 2023-08-07T20:30:32Z
closed: 2023-08-07T20:00:46Z
merged: 2023-08-07T20:00:46Z
base_branch: main
head_branch: tasks-idx
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/344
---

# Pull Request #344: Fixes 2180: add index for tasks query

**Author**: @rverdile
**Created**: August 02, 2023 at 04:05 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `tasks-idx`

## Description

## Summary
A query of tasks such as `SELECT * FROM tasks WHERE org_id = '5894300' and repository_uuid = '7300282a-dc65-432b-bf20-52dc958b72f6' and status = 'running' and type = 'snapshot' ORDER BY tasks.id LIMIT 1"` is taking a while in stage.

## Testing steps
1. Run the db migration
2. Seed something like 100,000 tasks with a few different orgIDs and statuses.
3. Run a similar query to the one above. Add `explain` to the beginning i.e. `EXPLAIN SELECT...`
4. You should see in the output that the index is being used to make the query. 
5. You may not notice improved performance locally because the seeded data is already pretty well sorted.

---

## Discussion

### Comment by @jlsherrill on August 02, 2023 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-2180

### Comment by @jlsherrill on August 04, 2023 at 02:02 PM UTC

You'll need to rebase and update your migrations to a later timestamp.  The easiest way is:
```
go run cmd/dbmigrate/main.go new my_migration_name
mv db/migrations/20230717141727_my_migration_name.down.sql db/migrations/20230804092757_my_migration_name.down.sql 
mv db/migrations/20230717141727_my_migration_name.up.sql db/migrations/20230804092757_my_migration_name.up.sql 
```

### Comment by @jlsherrill on August 07, 2023 at 07:59 PM UTC

another thing to note, is that the 'cost' goes way down:

before:
```
 Limit  (cost=117843.06..117843.07 rows=1 width=161)
```

after:
```
 Limit  (cost=8.46..8.47 rows=1 width=161)
```



### Comment by @jlsherrill on August 07, 2023 at 08:00 PM UTC

going ahead and merging so i can rebase my other PR :)

### Comment by @jlsherrill on August 07, 2023 at 08:01 PM UTC

Apologies this had qe-testing-needed, but i don't think it actually needed it

---

## Reviews

### Review by @jlsherrill - Approved on August 07, 2023 at 07:59 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/344*
