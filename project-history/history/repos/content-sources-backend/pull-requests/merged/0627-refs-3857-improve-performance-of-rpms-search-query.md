---
type: pull_request
number: 627
title: "Refs 3857: improve performance of rpms search query"
state: merged
author: jlsherrill
created: 2024-04-09T16:54:08Z
updated: 2024-04-09T17:52:19Z
closed: 2024-04-09T17:52:19Z
merged: 2024-04-09T17:52:19Z
base_branch: main
head_branch: 3857
labels: []
url: https://github.com/content-services/content-sources-backend/pull/627
---

# Pull Request #627: Refs 3857: improve performance of rpms search query

**Author**: @jlsherrill
**Created**: April 09, 2024 at 04:54 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `3857`

## Description

## Summary
Improve performance of rpms search query

## Testing steps
I wasn't fully able to reproduce the ~20 seconds we are seeing in stage, but i could reproduce ~200 ms.

1.  I would turn off snapshotting to make it easier, edit config.yaml, turn the snapshotting feature off
2. ```make repos-import```
3. ```go run cmd/external-repos/main.go nightly-jobs```
4.  run this to create 10,000 'instances' of the repo:
```
for i in `seq 1 10000`; do curl -X POST  http://localhost:8000/api/content-sources/v1.0/repositories/ -H "`./scripts/header.sh $RANDOM$RANDOM`" -H "Content-Type: application/json" -d '{"name":"FOO", "url": "https://dl.fedoraproject.org/pub/epel/9/Everything/x86_64/"}'; done
```
5.  Run this request:

```
####
POST http://localhost:8000/api/content-sources/v1.0/rpms/names
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json

{
  "urls": ["https://dl.fedoraproject.org/pub/epel/9/Everything/x86_64/"],
  "search": "kernel"
}
```

On main, i was seeing requests take ~250  ms, with this change, it dropped to ~60-80ms

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 09, 2024 at 05:52 PM UTC

merging to fix stage performance issue

---

## Reviews

### Review by @jlsherrill - Commented on April 09, 2024 at 05:07 PM UTC

### Review by @rverdile - Approved on April 09, 2024 at 05:36 PM UTC

tested and I see the performance improvement

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/627*
