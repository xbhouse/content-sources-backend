---
type: pull_request
number: 583
title: "Fixes 3629: add org id to features config"
state: merged
author: jlsherrill
created: 2024-02-19T16:18:26Z
updated: 2024-02-26T16:50:08Z
closed: 2024-02-21T14:30:08Z
merged: 2024-02-21T14:30:08Z
base_branch: main
head_branch: 3629
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/583
---

# Pull Request #583: Fixes 3629: add org id to features config

**Author**: @jlsherrill
**Created**: February 19, 2024 at 04:18 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3629`

## Description

## Summary

Adds org ids as a way to make features accessible.

## Testing steps

edit your config.yaml and edit/add:
```
features:
  snapshots:
    enabled: true #true
    accounts: 
    users: 
    organizations: ["FOO"]
```

use header.sh to test the org_id.  this should work:

```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/repositories/"       -H "Content-Type: application/json"     -d "{
          \"name\": \"needed errata\",
          \"url\": \"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/\",
          \"gpg_key\": \"\",
          \"snapshot\": true
        }"   -H "`./scripts/header.sh FOO 1 1`"
```
This should fail:
```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/repositories/"       -H "Content-Type: application/json"     -d "{
          \"name\": \"needed errata\",
          \"url\": \"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/\",
          \"gpg_key\": \"\",
          \"snapshot\": true
        }"   -H "`./scripts/header.sh NOT_FOO 1 1`"
```


## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 19, 2024 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-3629

### Comment by @swadeley on February 21, 2024 at 02:30 PM UTC

tested, lgtm, thank you.

---

## Reviews

### Review by @xbhouse - Approved on February 19, 2024 at 04:38 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/583*
