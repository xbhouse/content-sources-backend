---
type: pull_request
number: 760
title: "Fixes 4416: listing tasks returns 500 for bad repo uuid"
state: merged
author: xbhouse
created: 2024-08-01T17:15:11Z
updated: 2024-08-05T09:30:52Z
closed: 2024-08-05T09:21:30Z
merged: 2024-08-05T09:21:30Z
base_branch: main
head_branch: 4416
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/760
---

# Pull Request #760: Fixes 4416: listing tasks returns 500 for bad repo uuid

**Author**: @xbhouse
**Created**: August 01, 2024 at 05:15 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4416`

## Description

## Summary

Fixes 500 error on invalid repo uuids when listing tasks

## Testing steps

- List tasks filtered by an invalid repository uuid (for devs: `/tasks/?repository_uuid=repository_uuid`, for QE: `app.content_sources.rest_client.tasks_api.list_tasks(repository_uuid="repository_uuid")`)
- Should return 200 and an empty list

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on August 01, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-4416

### Comment by @swadeley on August 05, 2024 at 08:34 AM UTC

/retest

### Comment by @swadeley on August 05, 2024 at 09:20 AM UTC

Hi

in a shell to stage_proxy:

`ServiceException: (500)
Reason: Internal Server Error`

`HTTP response body: {"errors":[{"status":500,"title":"Error listing tasks","detail":"Database Error: ERROR: invalid input syntax for type uuid: \"repository_uuid\" (SQLSTATE 22P02)"}]}
`


in a shell to ephemeral built form this branch:

```
In [1]: app.content_sources.rest_client.tasks_api.list_tasks(repository_uuid="repository_uuid")
<snip>
Out[1]: 
{'data': [],
 'links': {'first': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=repository_uuid',
           'last': '/api/content-sources/v1/tasks/?limit=100&offset=0&repository_uuid=repository_uuid'},
 'meta': {'count': 0, 'limit': 100, 'offset': 0}}
```



---

## Reviews

### Review by @rverdile - Approved on August 02, 2024 at 01:33 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/760*
