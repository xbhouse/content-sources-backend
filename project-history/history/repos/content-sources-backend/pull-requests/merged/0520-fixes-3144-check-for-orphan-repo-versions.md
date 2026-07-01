---
type: pull_request
number: 520
title: "Fixes 3144: check for orphan repo versions"
state: merged
author: jlsherrill
created: 2024-01-05T13:04:24Z
updated: 2024-01-19T13:55:43Z
closed: 2024-01-19T13:55:43Z
merged: 2024-01-19T13:55:43Z
base_branch: main
head_branch: 3144
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/520
---

# Pull Request #520: Fixes 3144: check for orphan repo versions

**Author**: @jlsherrill
**Created**: January 05, 2024 at 01:04 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3144`

## Description

## Summary

during snapshot, and if one is found use that to continue.  

This helps the situation where some issue occurs and the pulp repo sync is successful, but the task dies, panics, or the publication creation fails.  If the repo hasn't changed at the next snapshot time, no new snapshot gets created.  With this change, the previously created repo version is re-used if there isn't a new one during the sync.

## Testing steps

I couldn't figure out a way to simulate this, so instead we can manually clear the db.  On a fresh empty db:

1)  create  a repo with snapshotting enabled.  Wait for it to finish snapshotting
2) within the database delete the snapshot:    `make db-cli-connect`
```
 delete * from snapshots;
```
3) trigger the snapshot manually:

```
POST http://localhost:8000/api/content-sources/v1.0/repositories/659bc4b3-b1a8-4d02-8d73-f36fe2b47d00/snapshot/
```

4) Check the db and notice a new snapshot was created, pointing to the old repo version (1)
```
# select version_href from snapshots;
```
```
                                        version_href                                         
---------------------------------------------------------------------------------------------
 /pulp/399c928e/api/v3/repositories/rpm/rpm/018cd9dd-d4bf-7683-9610-8299e4eccfd6/versions/1/

```
## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 08, 2024 at 05:23 PM UTC

https://issues.redhat.com/browse/HMS-3144

### Comment by @jlsherrill on January 18, 2024 at 03:08 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on January 16, 2024 at 04:28 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/520*
