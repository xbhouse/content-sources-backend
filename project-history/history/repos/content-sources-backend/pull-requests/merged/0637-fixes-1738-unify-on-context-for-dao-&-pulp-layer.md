---
type: pull_request
number: 637
title: "Fixes 1738: Unify on context for dao & pulp layer"
state: merged
author: jlsherrill
created: 2024-04-15T16:48:45Z
updated: 2024-04-26T20:59:07Z
closed: 2024-04-26T20:59:07Z
merged: 2024-04-26T20:59:07Z
base_branch: main
head_branch: request_id
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/637
---

# Pull Request #637: Fixes 1738: Unify on context for dao & pulp layer

**Author**: @jlsherrill
**Created**: April 15, 2024 at 04:48 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `request_id`

## Description

## Summary
Unifies context passing into methods across Dao and PulpClient layers from handler and tasks.  Also sets up the logger to properly pull the request ID from the context. Now the context is passed fully from:
* the handler, to the dao, to the pulp_client
* the handler to the task to the doa & pulp client

All dao db calls should now have a .WithContext() call to add the context to the db call.  

## Testing steps

Run a curl request to create (and snapshot a repo):

```
UUID="1cf4b0b0-fc0a-11ee-8913-201e8814f721"
curl -H "x-rh-insights-request-id:$UUID"  http://localhost:8000/api/content-sources/v1/repositories/  -H "`./scripts/header.sh 1 1`"   -d '{"name":"foo", "url":"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/", "snapshot":true}'  -H "Content-Type: application/json"
```

And watch the server logs:

you should see a lot of:
``` request_id=1cf4b0b0-fc0a-11ee-8913-201e8814f721```
Both at the api level and then as the task runs

Also check  the pulp workers logs:

```podman logs cs_pulp_worker_1 ```
and you should see the same request id there:
```
pulp [1cf4b0b0-fc0a-11ee-8913-201e8814f721]: pulpcore.tasking.tasks:INFO: Task completed 018ee838-baf4-7997-aa91-a4f24bcd82d9
```

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 15, 2024 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-1738

### Comment by @xbhouse on April 23, 2024 at 03:51 PM UTC

tested and i see the request id passed through, just a couple comments so far. @rverdile could you take a look too when you get a chance? i may have missed things

### Comment by @swadeley on April 24, 2024 at 07:26 AM UTC

Hi
I see many test failures in the pr check #1133:
`Bad Request: x-rh-identity header has an invalid or missing org_id`

### Comment by @jlsherrill on April 24, 2024 at 12:01 PM UTC

/retest

### Comment by @swadeley on April 24, 2024 at 07:59 PM UTC

/retest

### Comment by @jlsherrill on April 25, 2024 at 03:02 PM UTC

/retest

### Comment by @swadeley on April 26, 2024 at 06:38 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Commented on April 23, 2024 at 03:40 PM UTC

### Review by @xbhouse - Commented on April 23, 2024 at 03:44 PM UTC

### Review by @xbhouse - Commented on April 23, 2024 at 03:48 PM UTC

### Review by @xbhouse - Commented on April 23, 2024 at 03:56 PM UTC

### Review by @jlsherrill - Commented on April 23, 2024 at 04:00 PM UTC

### Review by @jlsherrill - Commented on April 23, 2024 at 04:01 PM UTC

### Review by @jlsherrill - Commented on April 23, 2024 at 06:44 PM UTC

### Review by @rverdile - Commented on April 23, 2024 at 08:07 PM UTC

ack from me

### Review by @xbhouse - Approved on April 24, 2024 at 01:32 PM UTC

ack! thanks @rverdile 

### Review by @rverdile - Changes Requested on April 24, 2024 at 04:03 PM UTC

Actually I think we need to double check that the request ID is the correct format. Usually they look like this: a9d3e55f2bf040ec995811920e1993d7

I see some errors in the pulp API logs, I'm not sure if they're related: 
```
pulp [None]: django_guid:WARNING: eAQguKUTcRFZtdERBGZnnuLccaFKToTA is not a valid GUID. New GUID is b7e7c7c70fa34845b441d4564553f20e
```

Might just need to update where we generate a request ID locally if one isn't specified.

### Review by @rverdile - Commented on April 24, 2024 at 05:37 PM UTC

### Review by @jlsherrill - Commented on April 24, 2024 at 05:50 PM UTC

### Review by @jlsherrill - Commented on April 24, 2024 at 05:51 PM UTC

### Review by @rverdile - Approved on April 24, 2024 at 06:09 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/637*
