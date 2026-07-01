---
type: pull_request
number: 1345
title: "HMS-9924: turn off epel for most orgs"
state: merged
author: jlsherrill
created: 2025-12-18T07:03:48Z
updated: 2025-12-19T14:54:52Z
closed: 2025-12-19T14:54:52Z
merged: 2025-12-19T14:54:51Z
base_branch: main
head_branch: epel_snapshot
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1345
---

# Pull Request #1345: HMS-9924: turn off epel for most orgs

**Author**: @jlsherrill
**Created**: December 18, 2025 at 07:03 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `epel_snapshot`

## Description

## Summary

Adds a job to run in stage and production that sets all non-community epel repositories' snapshot attribute to false unless they fall into two categories:
* the org they belong to has a content template
* the org has built an image in the last 6 weeks. (Passed via the EPEL_ORG_ID_SKIP env id)

From some poking in production, this should reduce the number of snaphotted epel repositories from 4379 to 158

Note that the tests were written by cursor

## Testing steps

make sure you have this in your config.yaml:
```
features:
  allow_custom_epel_creation:
    enabled: true
```

create an epel repository:

```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/repositories/" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==" \
    -H "Content-Type: application/json" \
    -d '{
          "name": "epel",
          "url": "https://dl.fedoraproject.org/pub/epel/9/Everything/x86_64/",
          "snapshot": true
        }'
```

run the job telling it to skip the org we just created:
```
$ EPEL_ORG_ID_SKIP="9" go run cmd/jobs/main.go  disable-snapshot-for-epel-repos
```
you should see no updates


now run it with some other org_id in the skip list:
```
EPEL_ORG_ID_SKIP="-123" go run cmd/jobs/main.go  disable-snapshot-for-epel-repos
```

You should see the one repo that is updated get printed and then updated:

```
2:02AM WRN Updating org: 9, repository: e04d9a0b-fc33-486b-9d9c-a470a52c41e1 severity=warn
2:02AM WRN Updated repos: 1 severity=warn
```

---

## Discussion

### Comment by @xbhouse on December 18, 2025 at 07:30 AM UTC

https://issues.redhat.com/browse/HMS-9924

### Comment by @marusak on December 18, 2025 at 11:18 AM UTC

> this should reduce the number of snaphotted epel repositories from 4379 to 4221

Sadly, that is not significant :/ 

> the org they belong to has a content template

I assume, content template that is using this repo?

### Comment by @jlsherrill on December 18, 2025 at 12:21 PM UTC

@marusak sorry thats a typo, it should reduce it by 4221, leaving only 158

### Comment by @xbhouse on December 18, 2025 at 02:48 PM UTC

tested and this looks good, just one question - it looks like this wouldn't cancel pending or running custom epel snapshot tasks, but it would prevent queuing new custom epel snapshot tasks? am i understanding that correctly?

### Comment by @jlsherrill on December 18, 2025 at 07:13 PM UTC

@xbhouse yes exactly.  I figured it was good enough to stop future ones from being created rather than stop ones already in progress or queued

### Comment by @jlsherrill on December 18, 2025 at 07:21 PM UTC

I'll go ahead and set the EPEL_ORG_ID_SKIP in stage (to a dummy value) and prod (to the actual list of IB ids, before i merge

---

## Reviews

### Review by @xbhouse - Approved on December 18, 2025 at 07:18 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1345*
