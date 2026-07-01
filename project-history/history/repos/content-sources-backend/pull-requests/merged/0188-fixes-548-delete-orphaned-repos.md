---
type: pull_request
number: 188
title: "Fixes 548: delete orphaned repos"
state: merged
author: jlsherrill
created: 2023-01-31T21:31:11Z
updated: 2023-02-06T19:03:27Z
closed: 2023-02-06T19:03:26Z
merged: 2023-02-06T19:03:26Z
base_branch: main
head_branch: 548
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/188
---

# Pull Request #188: Fixes 548: delete orphaned repos

**Author**: @jlsherrill
**Created**: January 31, 2023 at 09:31 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `548`

## Description

## Summary

as part of introspect-all.  This deletes any non-public repositories that were created more than a week ago and do not have any repo_configs defined


## Testing steps
Create a repo:
```
curl -X POST http://localhost:8000/api/content-sources/v1/repositories/bulk_create/  -d '[{"name":"foobarz", "url":"https://jlsherrill.fedorapeople.org/fake-repos/needed-errataz"}]'  -H "`./scripts/header.sh 3 1`"   -H "Content-Type: application/json"   
```
update the created_at time:
```
make db-cli-connect
```
```
# update repositories  set created_at = LOCALTIMESTAMP - INTERVAL '1 week' where url = 'https://jlsherrill.fedorapeople.org/fake-repos/needed-errataz/';
UPDATE 1
```

Delete the repo (using UUID from create output):
```
curl -X DELETE http://localhost:8000/api/content-sources/v1/repositories/f3430e8e-38b1-4563-84c1-4b93c3e63585  -H "`./scripts/header.sh 3 1`"   -H "Content-Type: application/json"
```

introspect-all
```
go run cmd/external-repos/main.go introspect-all
```

notice:

```
{"level":"debug","time":"2023-01-31T16:30:43-05:00","message":"Cleaned up 1 orphaned repositories"}
```


and in the db, it should be gone:
```
select url from repositories;
```

---

## Discussion

### Comment by @jlsherrill on January 31, 2023 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-548

---

## Reviews

### Review by @rverdile - Commented on February 01, 2023 at 04:49 PM UTC

### Review by @rverdile - Commented on February 01, 2023 at 05:12 PM UTC

### Review by @rverdile - Commented on February 01, 2023 at 06:41 PM UTC

played around with this and looks good!

I left a couple of comments. I don't think they'll lead to any big changes.  but I guess I'll wait and see before acking.

### Review by @rverdile - Commented on February 01, 2023 at 06:47 PM UTC

### Review by @jlsherrill - Commented on February 01, 2023 at 08:32 PM UTC

### Review by @jlsherrill - Commented on February 01, 2023 at 08:32 PM UTC

### Review by @jlsherrill - Commented on February 01, 2023 at 08:34 PM UTC

### Review by @rverdile - Approved on February 01, 2023 at 09:57 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/188*
