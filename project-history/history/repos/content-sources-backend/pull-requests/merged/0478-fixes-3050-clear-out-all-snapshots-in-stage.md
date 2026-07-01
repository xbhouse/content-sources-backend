---
type: pull_request
number: 478
title: "Fixes 3050: clear out all snapshots in stage"
state: merged
author: jlsherrill
created: 2023-11-20T20:30:44Z
updated: 2023-11-21T19:26:26Z
closed: 2023-11-21T19:25:32Z
merged: 2023-11-21T19:25:32Z
base_branch: main
head_branch: 3050_2
labels: []
url: https://github.com/content-services/content-sources-backend/pull/478
---

# Pull Request #478: Fixes 3050: clear out all snapshots in stage

**Author**: @jlsherrill
**Created**: November 20, 2023 at 08:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `3050_2`

## Description

## Summary
This command will execute in stage and delete all:
* repository versions in pulp
* distributions in pulp
* snapshots in our db

It then will run orphan_cleanup in pulp to remove all the artifacts.

The idea was that this would deploy, would run in stage, and then we would delete the added main.go file and deployment.yaml changes.

## Testing steps

1.  Create and snapshots some repositories
2. run some queries in the pulp db:
```
psql -h localhost -p 5432 -U pulp     #(password 'password')
```

```
select count(*) from core_artifact;    
```

```
select number from core_repositoryversion;   
```
You should see some with a zero version, and some non-zero versions


```
select * from rpm_rpmdistribution;
```
3.  now run the cleanup:  ``` go run cmd/cleanup_versions/main.go --force ```
4. re-run all the queries, most things should be zero, but you still should have repository versions that are only zero
5. Also check the content-sources db:  ```make db-cli-connect```
```
select  count(*) from snapshots;
```
This should now be zero.
6.  re-run a snapshot: ```OPTIONS_ALWAYS_RUN_CRON_TASKS=true go run cmd/external-repos/main.go nightly-jobs```
and see snapshots being created again.  Re-run through 2-6 as much as you like

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 20, 2023 at 08:31 PM UTC

https://issues.redhat.com/browse/HMS-3050

---

## Reviews

### Review by @Andrewgdewar - Approved on November 20, 2023 at 10:44 PM UTC

Was able to check functionality, confirmed expected behaviour!
The code looks fine to me, you may want to get @rverdile to take a quick gander as he seems to find things I would miss.

### Review by @rverdile - Approved on November 21, 2023 at 04:27 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/478*
