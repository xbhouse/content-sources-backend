---
type: pull_request
number: 353
title: "Fixes 2310: fix 404 err on delete task"
state: merged
author: jlsherrill
created: 2023-08-09T15:25:13Z
updated: 2023-08-09T17:25:36Z
closed: 2023-08-09T17:25:33Z
merged: 2023-08-09T17:25:33Z
base_branch: main
head_branch: 2310
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/353
---

# Pull Request #353: Fixes 2310: fix 404 err on delete task

**Author**: @jlsherrill
**Created**: August 09, 2023 at 03:25 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `2310`

## Description

## Summary
if no domain was ever created

## Testing steps

In an org that has never created any repos (or at least none that were snapshotted).  Create a repo:

```
POST http://localhost:8000/api/content-sources/v1.0/repositories/
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJmb28ifSwiYWNjb3VudF9udW1iZXIiOiIzMCIsImludGVybmFsIjp7Im9yZ19pZCI6IjMwIn19fQo=
x-Rh-Insights-Request-Id: 9876


{"name":"pulp 3.16","url":"http://yum.theforeman.org/pulpcore/3.16/el8/x86_64/","snapshot": false}
```

get the UUID of the created repo and delete it:

```

DELETE http://localhost:8000/api/content-sources/v1.0/repositories/14c3a83d-7574-4654-9a2d-8d4c92254b39
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJmb28ifSwiYWNjb3VudF9udW1iZXIiOiIzMCIsImludGVybmFsIjp7Im9yZ19pZCI6IjMwIn19fQo=

```

Now check the tasks:

```
GET http://localhost:8000/api/content-sources/v1.0/tasks/
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJmb28ifSwiYWNjb3VudF9udW1iZXIiOiIzMCIsImludGVybmFsIjp7Im9yZ19pZCI6IjMwIn19fQo=
```

the delete task should not be in error state


---

## Discussion

### Comment by @jlsherrill on August 09, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-2310

---

## Reviews

### Review by @rverdile - Commented on August 09, 2023 at 04:24 PM UTC

### Review by @rverdile - Commented on August 09, 2023 at 04:25 PM UTC

### Review by @rverdile - Commented on August 09, 2023 at 04:27 PM UTC

### Review by @jlsherrill - Commented on August 09, 2023 at 04:54 PM UTC

### Review by @jlsherrill - Commented on August 09, 2023 at 04:58 PM UTC

### Review by @rverdile - Commented on August 09, 2023 at 05:13 PM UTC

### Review by @rverdile - Commented on August 09, 2023 at 05:13 PM UTC

### Review by @rverdile - Approved on August 09, 2023 at 05:14 PM UTC

tested and looks good, but there's a linting error

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/353*
