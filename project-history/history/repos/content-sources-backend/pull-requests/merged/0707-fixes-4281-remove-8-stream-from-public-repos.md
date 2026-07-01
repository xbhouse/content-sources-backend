---
type: pull_request
number: 707
title: "Fixes 4281: remove 8-stream from public repos"
state: merged
author: jlsherrill
created: 2024-06-14T19:21:07Z
updated: 2024-06-17T18:30:49Z
closed: 2024-06-17T18:14:40Z
merged: 2024-06-17T18:14:40Z
base_branch: main
head_branch: 4281
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/707
---

# Pull Request #707: Fixes 4281: remove 8-stream from public repos

**Author**: @jlsherrill
**Created**: June 14, 2024 at 07:21 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4281`

## Description

## Summary

and make repos removed from the external_repos.json file no longer public

## Testing steps

1. on main run `make repos-import`
2. connect to the db and query 8-stream public repos:
```
make db-cli-connect
# select public, url from repositories  where url ilike '%8-stream%';
```
3.  Check out this PR, re-run `make repos-import`
4. re-query the query and make sure the 8-stream repos are no longer public
5. query the 9-stream repos and make sure they are still public
```
# select public, url from repositories  where url ilike '%9-stream%';
```

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 14, 2024 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-4281

### Comment by @jlsherrill on June 14, 2024 at 07:51 PM UTC

/retest

### Comment by @swadeley on June 17, 2024 at 03:59 PM UTC

QE ACK, tested OK

---

## Reviews

### Review by @xbhouse - Approved on June 17, 2024 at 05:20 PM UTC

works and looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/707*
