---
type: pull_request
number: 996
title: "HMS-5551: use prefix match for rpm search"
state: merged
author: jlsherrill
created: 2025-02-27T21:53:07Z
updated: 2025-03-12T15:25:29Z
closed: 2025-03-12T15:25:29Z
merged: 2025-03-12T15:25:29Z
base_branch: main
head_branch: 5551
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/996
---

# Pull Request #996: HMS-5551: use prefix match for rpm search

**Author**: @jlsherrill
**Created**: February 27, 2025 at 09:53 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5551`

## Description

## Summary

Changes the rpm name search to only search using prefixes.  This has a couple advantages:

1.  it is much much faster
2. The exact match will appear in the list first

Snapshot search is also switched here: https://github.com/content-services/tang/pull/18

## Testing steps

create a repo, then search for rpms:

```
POST http://localhost:8000/api/content-sources/v1.0/rpms/names
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxIn0sImFjY291bnRfbnVtYmVyIjoiZm9vIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K
Content-Type: application/json

{"urls":["https://dl.fedoraproject.org/pub/epel/9/Everything/x86_64/",
  "http://yum.theforeman.org/pulpcore/3.16/el8/x86_64/"],
  "search":"p"}
```

from the results, try search for a prefix, or a middle substring.  Only prefix's should return a match


Do the same for the created snapshot:

```
GET http://localhost:8000/api/content-sources/v1.0/snapshots/b158a14a-50d9-470a-9d3f-e0e596d859d9/rpms?search=a&limit=10
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json
```


---

## Discussion

### Comment by @jlsherrill on February 27, 2025 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-5551

### Comment by @swadeley on March 12, 2025 at 03:25 PM UTC

All API test pass, no API docs changes, lgtm.

---

## Reviews

### Review by @rverdile - Approved on March 10, 2025 at 02:43 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/996*
