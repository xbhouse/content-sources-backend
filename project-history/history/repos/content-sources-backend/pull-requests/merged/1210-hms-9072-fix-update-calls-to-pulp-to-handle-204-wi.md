---
type: pull_request
number: 1210
title: "HMS-9072: fix update calls to pulp to handle 204 with no task"
state: merged
author: Starle21
created: 2025-09-19T09:55:45Z
updated: 2025-10-01T07:50:08Z
closed: 2025-10-01T07:49:59Z
merged: 2025-10-01T07:49:58Z
base_branch: main
head_branch: starle/HMS-9072
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1210
---

# Pull Request #1210: HMS-9072: fix update calls to pulp to handle 204 with no task

**Author**: @Starle21
**Created**: September 19, 2025 at 09:55 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `starle/HMS-9072`

## Description

## Summary

**Pulp update calls now return 200 with no task**

Previously pulp always created an internal pulp task and returned the task href,
when there had been triggered an update call to domains or remote, irrespective if
the update was changing anything or not. 

From now on, when an update call to
- domain (UpdateDomainIfNeeded) or 
- remote (UpdateRmpRemote) does not change anything,

pulp does not create a task. 
This is indicated by 200 status code. 

If there is an actual change, pulp sends 202 with taskHref back. 

This commit updates the parts of code
that deal with these pulp functions to take this pulp api change into account.

Note that Pulp team has been waiting on content-sources to be updated first,
before they update Pulp. This means that zest client is still expecting to receive
a task href always, failing when it is empty. There is a temporary workaround
in the code, which will be removed after Pulp will get updated.

## Testing steps
Run the integration tests.



---

## Discussion

### Comment by @jlsherrill on September 19, 2025 at 10:00 AM UTC

https://issues.redhat.com/browse/HMS-9072

### Comment by @jlsherrill on September 23, 2025 at 06:59 PM UTC

@Starle21 the pulp team is has prepared the change to use this and actually have a container here that we could test with:  http://quay.io/redhat-user-workloads/pulp-services-tenant/pulp/pulp@sha256:9f6d0a08cf63cda6476021385a652bf16b5a79300fadfa61234ba96e1f941f56



### Comment by @jlsherrill on September 24, 2025 at 06:59 PM UTC

apologies, the image i mentioned couldn't be used directly, but:

`quay.io/redhat-user-workloads/pulp-services-tenant/pulp/pulp:on-pr-682`

can.  After `podman login quay.io`  and replacing the pulp images in compose_files/pulp/docker-compose.yml
  a `make compose-up` seems to use the right images

and you can verify the pulp version is correct:

```
 curl localhost:8080/api/pulp/api/v3/status/ | jq
```

```
  "versions": [
    {
      "component": "core",
      "version": "3.90.0",
```

### Comment by @jlsherrill on September 30, 2025 at 03:18 PM UTC

/retest

### Comment by @jlsherrill on September 30, 2025 at 06:47 PM UTC

/retest


---

## Reviews

### Review by @jlsherrill - Approved on September 30, 2025 at 05:39 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1210*
