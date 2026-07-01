---
type: pull_request
number: 827
title: "Fixes 4745: use domain list to detect pulp is up"
state: merged
author: jlsherrill
created: 2024-09-19T17:21:01Z
updated: 2024-09-23T12:07:36Z
closed: 2024-09-23T12:07:33Z
merged: 2024-09-23T12:07:33Z
base_branch: main
head_branch: 4745
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/827
---

# Pull Request #827: Fixes 4745: use domain list to detect pulp is up

**Author**: @jlsherrill
**Created**: September 19, 2024 at 05:21 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4745`

## Description

## Summary

Changes our ping call to pulp to use domain list as we won't have permission for remotes list
The health check (using metrics) already uses this!

## Testing steps


monitor the pulp api:
```
 podman logs -f cs_pulp_api_1 
```

launch the server
```
go run cmd/content-sources/main.go api instrumentation consumer
```

call the ping api:
```
 curl localhost:9000/ping  
````

you should only see calls to /api/pulp/default/api/v3/domains/?name=default
You'll only see ping hit this endpoint once per server start.  You should see it get hit by the metrics go routine every ~30 minutes.



---

## Discussion

### Comment by @jlsherrill on September 19, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-4745

---

## Reviews

### Review by @xbhouse - Approved on September 20, 2024 at 02:42 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/827*
