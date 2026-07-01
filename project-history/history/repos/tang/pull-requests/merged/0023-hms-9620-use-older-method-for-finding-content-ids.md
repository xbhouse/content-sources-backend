---
type: pull_request
number: 23
title: "HMS-9620: use older method for finding content ids if they are null"
state: merged
author: jlsherrill
created: 2025-12-04T21:25:17Z
updated: 2026-01-05T17:21:03Z
closed: 2026-01-05T17:20:57Z
merged: 2026-01-05T17:20:57Z
base_branch: main
head_branch: back_compat
labels: []
url: https://github.com/content-services/tang/pull/23
---

# Pull Request #23: HMS-9620: use older method for finding content ids if they are null

**Author**: @jlsherrill
**Created**: December 04, 2025 at 09:25 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `back_compat`

## Description

summary:

Pulp added a new column on repo versions "Content_ids" which should make querying versions for content faster.  They did not, however, update previous versions to include these ids, leaving that column null.

This change checks to see if any of the repo versions wanting to be searched have a null content_ids column.  If so we use the old method for searching.  If not we use the new method.  This does take an extra small query to figure it out, but is much simpler than trying to do it all in a single query.    

The code was also written in a way that hopefully will be easy to clean up after ~August 1st, 2026


To Test:

1.  add this to your go.mod
```
replace github.com/content-services/tang => /path/to/pr/tang/
```
run:
```
go get ./...
make repos-import-rhel9
make process-repos
rm release/content-sources
make run
```

2.  let them snapshot and grab the UUIDs of the snapshots, either from the db or rest api:

```
select uuid from snapshots
```

3.  Query the snapshots
```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/snapshots/rpms/names" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==" \
    -H "Content-Type: application/json" \
    -H "User-Agent: IntelliJ HTTP Client/GoLand 2025.2.2" \
    -H "Accept-Encoding: br, deflate, gzip, x-gzip" \
    -H "Accept: */*" \
    -d '{
        	"search": "vim",
        	"uuids": [
        		"d4ea24b6-54a2-43c1-98e0-a05fd4f5e6ec"
        	]
        }'
```

4.  connect to pulp and clear the content_ids (password is password):
```
psql --host localhost --port 5432  -U pulp

# update core_repositoryversion set content_ids = null;
```

5.  re-run the curl command to see the same results.


You can also turn on debug logging to see the queries being run and confirm they are different.


---

## Reviews

### Review by @xbhouse - Approved on December 09, 2025 at 10:16 PM UTC

tested and works well! ack!

---

*Archived from: https://github.com/content-services/tang/pull/23*
