---
type: pull_request
number: 264
title: "Fixes 1735: add rbac caching w/ redis"
state: merged
author: jlsherrill
created: 2023-05-08T13:43:47Z
updated: 2023-05-16T17:27:04Z
closed: 2023-05-16T17:27:00Z
merged: 2023-05-16T17:27:00Z
base_branch: main
head_branch: 1735
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/264
---

# Pull Request #264: Fixes 1735: add rbac caching w/ redis

**Author**: @jlsherrill
**Created**: May 08, 2023 at 01:43 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1735`

## Description

## Summary
This adds caching support for rbac requests, with the goal to reduce overall latency from users using the application.

This uses a 1 minute cache expiration by default.  

## Testing steps
A couple of ways to test
1) spin up an ephemeral env with these changes (using custom deployment.yaml from custom bonfire config and pr built image)
2) curl the repositories api ```time curl https://env-ephemeral-pzuqlw-iiajlc06.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/content-sources/v1/repositories/  -u jdoe:CiUhwxHyzE2K87zz ```
3)  immediately curl again. You should see 100-200ms drop in request time
4) wait 60 seconds and curl again.  You should it revert back initially to the higher request time.

Second way (locally in dev):
1) add new config options from config.example.yaml
2) `make compose-up`  (again to get a new redis locally)
2) Turn on rbac:
```
clients:
  rbac_enabled: True
  rbac_base_url: http://localhost:8800/api/rbac/v1
```
3) add a sleep statement to your rbac client:  clients/rbac.go right above 
```
		acl, err = r.client.GetAccess(ctx, identity.GetIdentityHeader(ctx), "")
```
add:

```
time.sleep(5*time.Second)
```
4) ```curl http://localhost:8000/api/content-sources/v1.0/repositories/ -H "`./scripts/header.sh 1 foo`"```  should take 5 seconds, while immediately running it again should take much less time.


TODO: 
* [x] verify behavior for no access from rbac (error or empty list?)
* [x] Testing around this is currently minimal, and i think the ephemeral tests are probably the best tests rather than trying to mock behavior.  I did add some tests around the rbac client to ensure its respecting the configured cache

---

## Discussion

### Comment by @jlsherrill on May 08, 2023 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-1735

### Comment by @jlsherrill on May 15, 2023 at 12:05 PM UTC

The needed stage changes are now available in app-interface

---

## Reviews

### Review by @jlsherrill - Commented on May 08, 2023 at 08:53 PM UTC

### Review by @Andrewgdewar - Commented on May 11, 2023 at 06:02 PM UTC

### Review by @jlsherrill - Commented on May 11, 2023 at 06:07 PM UTC

### Review by @Andrewgdewar - Approved on May 11, 2023 at 06:38 PM UTC

Tested working. We should likely get the mentioned need changes in appSRE before merging.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/264*
