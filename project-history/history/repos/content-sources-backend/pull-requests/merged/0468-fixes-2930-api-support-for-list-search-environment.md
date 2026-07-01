---
type: pull_request
number: 468
title: "Fixes 2930: API support for list / search environments"
state: merged
author: xbhouse
created: 2023-11-14T15:50:48Z
updated: 2024-02-13T15:16:02Z
closed: 2024-01-10T13:36:03Z
merged: 2024-01-10T13:36:03Z
base_branch: main
head_branch: envs-api-support
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/468
---

# Pull Request #468: Fixes 2930: API support for list / search environments

**Author**: @xbhouse
**Created**: November 14, 2023 at 03:50 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `envs-api-support`

## Description

## Summary

This adds:

- `repositories/{uuid}/environments` endpoint for listing, returns ID, Name, and Description metadata
- `environments/names` endpoint for searching, returns Name and Description of all environments found across the repositories requested
- database support to store environment metadata
- environments are de-duplicated based on combination of name/ID

## Testing steps

Endpoint: `$BASE_URL/environments/names`

Example request: `{ "urls": [ "https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/" ], "search": "kde" }`

- add a repository or 2 that has a comps.xml file (e.g. https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/) and make a request to `repositories/{uuid}/environments`, should list environments associated with the repo UUID
- make a request to `environments/names` with at least one valid repo URL or repo UUID and a search term, should return the environments with that search term in the name

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 14, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-2930

### Comment by @rverdile on December 05, 2023 at 06:59 PM UTC

also needs a big rebase. good luck :)

### Comment by @rverdile on December 14, 2023 at 03:01 PM UTC

/retest

### Comment by @jlsherrill on January 03, 2024 at 12:34 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @swadeley on January 08, 2024 at 02:25 PM UTC

/retest

### Comment by @swadeley on January 09, 2024 at 09:09 AM UTC

/retest

### Comment by @swadeley on January 09, 2024 at 07:40 PM UTC

/retest

### Comment by @swadeley on January 10, 2024 at 01:23 PM UTC

Hi

looks good:
```
In [1]: app.content_sources.rest_client.repositories_api.search_environments(dict(urls=[ "https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/" ], search="kde" ))
<snip>
2024-01-10 13:15:40.526 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=ACYDIjNNlsifsOGzTIMSJcZuHXfoIcJe, params=[]
Out[1]: 
[{'description': 'The KDE Plasma Workspaces, a highly-configurable graphical '
                 'user interface which includes a panel, desktop, system icons '
                 'and desktop widgets, and many powerful KDE applications.',
  'environment_name': 'KDE Plasma Workspaces'}]

In [2]: 

```

"uuids" works too:

```
In [7]: app.content_sources.rest_client.repositories_api.search_environments(dict(uuids=[ "a93fa910-6838-4b11-b5e9-6dddb9633bf0" ], search="kde" ))
2024-01-10 13:22:13.441 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=QxVJKNUKWtemVyOEyDkBNduWEVFytpIj, params=[]
Out[7]: 
[{'description': 'The KDE Plasma Workspaces, a highly-configurable graphical '
                 'user interface which includes a panel, desktop, system icons '
                 'and desktop widgets, and many powerful KDE applications.',
  'environment_name': 'KDE Plasma Workspaces'}]
```

---

## Reviews

### Review by @rverdile - Commented on December 05, 2023 at 06:42 PM UTC

I might find more things as a I test, but overall looks pretty good

### Review by @xbhouse - Commented on December 05, 2023 at 07:18 PM UTC

### Review by @xbhouse - Commented on December 05, 2023 at 07:23 PM UTC

### Review by @xbhouse - Commented on December 05, 2023 at 07:23 PM UTC

### Review by @rverdile - Commented on December 05, 2023 at 10:04 PM UTC

### Review by @xbhouse - Commented on December 06, 2023 at 05:16 PM UTC

### Review by @xbhouse - Commented on December 06, 2023 at 05:17 PM UTC

### Review by @xbhouse - Commented on December 06, 2023 at 05:18 PM UTC

### Review by @rverdile - Commented on December 06, 2023 at 09:17 PM UTC

### Review by @rverdile - Commented on December 06, 2023 at 09:17 PM UTC

### Review by @xbhouse - Commented on December 06, 2023 at 09:58 PM UTC

### Review by @rverdile - Commented on December 07, 2023 at 02:49 PM UTC

### Review by @xbhouse - Commented on December 07, 2023 at 02:58 PM UTC

### Review by @rverdile - Approved on December 07, 2023 at 04:35 PM UTC

tested with a few red hat repos and looks good!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/468*
