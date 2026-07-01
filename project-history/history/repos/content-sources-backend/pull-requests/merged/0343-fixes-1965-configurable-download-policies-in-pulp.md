---
type: pull_request
number: 343
title: "Fixes 1965: configurable download policies in pulp"
state: merged
author: jlsherrill
created: 2023-08-01T19:20:14Z
updated: 2023-08-15T19:47:52Z
closed: 2023-08-15T19:47:48Z
merged: 2023-08-15T19:47:48Z
base_branch: main
head_branch: 1965
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/343
---

# Pull Request #343: Fixes 1965: configurable download policies in pulp

**Author**: @jlsherrill
**Created**: August 01, 2023 at 07:20 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1965`

## Description

## Summary

Allows setting the download policy via an environment variable.  Currently only on_demand and immediate are supported.

## Testing steps

for local dev testing run:
```
CLIENTS_PULP_DOWNLOAD_POLICY=immediate  make run
```

create a repo with snapshotting turned on, then check pulp (use password 'password'):
```
$ psql -h localhost -p 5432  pulp pulp 

pulp=# select policy from core_remote;
```

There's not really any way to test this from a users perspective. right now

---

## Discussion

### Comment by @jlsherrill on August 01, 2023 at 07:20 PM UTC

also, this will set the policies in the various environments:  https://gitlab.cee.redhat.com/service/app-interface/-/merge_requests/75987

### Comment by @jlsherrill on August 01, 2023 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-1965

---

## Reviews

### Review by @rverdile - Approved on August 15, 2023 at 06:09 PM UTC

tested and working!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/343*
