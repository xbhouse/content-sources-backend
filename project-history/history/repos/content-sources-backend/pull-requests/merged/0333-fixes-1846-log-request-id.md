---
type: pull_request
number: 333
title: "Fixes 1846: log request ID"
state: merged
author: rverdile
created: 2023-07-13T18:47:25Z
updated: 2023-08-02T16:07:20Z
closed: 2023-08-02T16:07:16Z
merged: 2023-08-02T16:07:16Z
base_branch: main
head_branch: request-id
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/333
---

# Pull Request #333: Fixes 1846: log request ID

**Author**: @rverdile
**Created**: July 13, 2023 at 06:47 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `request-id`

## Description

## Summary
- Adds support for the request ID, `x-rh-insights-request-id`, sent in the header and propagates that request ID to logging in the API, tasking, and redis
- Also propagates TaskID/Type logging throughout tasking system and pulp. 

The goal here is to more easily track which events correspond to which request. For example, if I get an error during repository snapshot, it is nice to know which request or task triggered that failed snapshot.

TODO
- [x] I need to double check redis logging. Will add testing step for that as well.


## Testing steps
1. Make an API request, for example, one to create a repository. Add the `x-rh-insights-request-id` header.
```
POST http://localhost:8000/api/content-sources/v1.0/repositories/
x-rh-identity: zEzcvAe....
x-Rh-Insights-Request-Id: 12345
Content-Type: application/json
```
3. You should see the request ID propagated through the to the logs for the API, tasking, and introspection. Set your log level to  `debug` to make sure you see all the logs.
4. Try variations of snapshotting, introspecting, CTRL+C exiting the server while it's in the middle of a task, and make sure the logs are all showing the request ID (or task ID if applicable).
5. Provide feedback if you think something should or shouldn't be logged.
6. A quick test for RBAC/Redis logging is to hard-code an error in `client_wrapper.go` and see the error get logged.


---

## Discussion

### Comment by @jlsherrill on July 13, 2023 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-1846

### Comment by @rverdile on July 21, 2023 at 07:07 PM UTC

This is ready for another review. 

Propagating context through admin_tasks handler to pulp_client is a bigger change than I anticipated because of how it affects tests. Since is has no impact on the logging (since nothing is logged there at the moment), it might make sense to leave it out from here.

---

## Reviews

### Review by @dpang314 - Commented on July 19, 2023 at 03:46 PM UTC

Testing logging for snapshots, introspection, and interrupted tasks works for me. 

I'm not sure whether this is the intended behavior, but when a request is made without the `x-Rh-Insights-Request-Id` header, `echo_middleware.RequestIDWithConfig` automatically generates one and attaches it to the HTTP response header. However, that automatically generated id is not used by the logger.

### Review by @rverdile - Commented on July 20, 2023 at 01:51 PM UTC

### Review by @rverdile - Commented on July 20, 2023 at 01:51 PM UTC

### Review by @rverdile - Commented on July 20, 2023 at 07:56 PM UTC

### Review by @rverdile - Commented on July 20, 2023 at 08:01 PM UTC

### Review by @dpang314 - Commented on July 24, 2023 at 07:41 PM UTC

### Review by @rverdile - Commented on July 25, 2023 at 12:38 PM UTC

### Review by @dpang314 - Approved on July 25, 2023 at 01:23 PM UTC

Looks good to me!

### Review by @jlsherrill - Commented on July 31, 2023 at 06:46 PM UTC

### Review by @jlsherrill - Commented on July 31, 2023 at 06:48 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/333*
