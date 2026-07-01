---
type: pull_request
number: 20
title: "HMS-8902: use content_ids for version queries"
state: merged
author: jlsherrill
created: 2025-09-08T20:35:32Z
updated: 2025-09-17T18:02:02Z
closed: 2025-09-17T18:01:59Z
merged: 2025-09-17T18:01:59Z
base_branch: main
head_branch: 8902
labels: []
url: https://github.com/content-services/tang/pull/20
---

# Pull Request #20: HMS-8902: use content_ids for version queries

**Author**: @jlsherrill
**Created**: September 08, 2025 at 08:35 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `8902`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on September 10, 2025 at 02:34 PM UTC

This should speed up these queries in stage and prod, but its hard to see this locally without syncing a lot of content.

To test locally:

edit go.mod and add (using the path to your tang checkout:

```
replace github.com/content-services/tang => /home/jlsherri/git/tang/
```

run:
```
go get ./...
rm release/content-sources
make run
```

and then:

```
make repos-import-rhel9
go run cmd/external-repos/main.go process-repos
```

wait for it all to be snapshotted.

Then grab all the snapshot uuids:

```
make db-cli-connect

select uuid from snapshots;
```


And use curl for rpm search and listing within snapshots:

```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/snapshots/rpms/names" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==" \
    -H "Content-Type: application/json" \
    -d '{
        	"search": "vim",
        	"uuids": [
        		"f56e9ed0-54f0-431c-afa8-821a458e5ba5",
              "06708418-4d00-482a-867d-193d1349e564",
              "5ccca9ea-67d1-4c4c-a0a8-fa564841bdc5"
        	]
        }'
```

```
curl -X GET --location "http://localhost:8000/api/content-sources/v1.0/snapshots/5ccca9ea-67d1-4c4c-a0a8-fa564841bdc5/rpms" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==" \
    -H "Content-Type: application/json"

```




---

## Reviews

### Review by @xbhouse - Approved on September 17, 2025 at 03:47 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/tang/pull/20*
