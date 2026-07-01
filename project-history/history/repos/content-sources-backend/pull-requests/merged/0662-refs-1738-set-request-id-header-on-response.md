---
type: pull_request
number: 662
title: "Refs 1738: set request id header on response"
state: merged
author: jlsherrill
created: 2024-05-08T17:48:56Z
updated: 2024-05-09T15:28:23Z
closed: 2024-05-09T15:28:23Z
merged: 2024-05-09T15:28:23Z
base_branch: main
head_branch: 1738
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/662
---

# Pull Request #662: Refs 1738: set request id header on response

**Author**: @jlsherrill
**Created**: May 08, 2024 at 05:48 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1738`

## Description

## Summary

The previous pr: https://github.com/content-services/content-sources-backend/pull/637  had a bug where the request_id wasn't being set on the response in all cases (when one was specified). The result is that tasks didn't get their proper request_id

## Testing steps

Run a curl request to create (and snapshot a repo):

UUID="1cf4b0b0-fc0a-11ee-8913-201e8814f721"
curl -H "x-rh-insights-request-id:$UUID"  http://localhost:8000/api/content-sources/v1/repositories/  -H "`./scripts/header.sh 1 1`"   -d '{"name":"foo", "url":"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/", "snapshot":true}'  -H "Content-Type: application/json"

And watch the server logs:

you should see a lot of:
```
 request_id=1cf4b0b0-fc0a-11ee-8913-201e8814f721
```
Both at the api level and then as the task runs

Also check the pulp workers logs:

podman logs cs_pulp_worker_1 
and you should see the same request id there:
```
pulp [1cf4b0b0-fc0a-11ee-8913-201e8814f721]: pulpcore.tasking.tasks:INFO: Task completed 018ee838-baf4
```
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 08, 2024 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-1738

---

## Reviews

### Review by @rverdile - Approved on May 09, 2024 at 03:19 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/662*
