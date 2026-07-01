---
type: pull_request
number: 638
title: "Fixes 3766: add check for missing repos/snapshots"
state: merged
author: xbhouse
created: 2024-04-15T20:19:52Z
updated: 2024-04-19T09:30:23Z
closed: 2024-04-19T09:01:33Z
merged: 2024-04-19T09:01:33Z
base_branch: main
head_branch: check-missing-repos-snaps
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/638
---

# Pull Request #638: Fixes 3766: add check for missing repos/snapshots

**Author**: @xbhouse
**Created**: April 15, 2024 at 08:19 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `check-missing-repos-snaps`

## Description

## Summary

- Returns a 400 response code if a repository / snapshot is not found for these search endpoints:
```
/rpms/names
/package_groups/names
/environments/names
/snapshots/rpms/names
/snapshots/package_groups/names
/snapshots/environments/names 
```
- Updates the error response for requests with missing urls / uuids to return a 400 instead of a 500
- Also adds a handler test for DetectRpms because I realized I forgot to do that :) 

## Testing steps

- Make a request to each endpoint with an incorrect uuid / url, each should return 400
- Requests with no uuid / url to these endpoints should also return 400

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 15, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-3766

### Comment by @mayurilahane on April 15, 2024 at 09:14 PM UTC

/retest

### Comment by @swadeley on April 17, 2024 at 08:35 AM UTC

/retest

### Comment by @swadeley on April 17, 2024 at 01:09 PM UTC

Hi

for rpm names:

```In [6]:  app.content_sources.rest_client.repositories_api.search_rpm(
   ...:         dict(
   ...:             search="bear",
   ...:             urls=[
   ...:                 'https://jlsherrill.fedorapeople.org/fake-repos/revision/two/',
   ...:             ],
   ...:         )
   ...:     )
   ...: 
   ...: 
2024-04-17 14:05:29.539 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[6]: [{'package_name': 'bear', 'summary': 'A dummy package of bear'}]

In [7]:  app.content_sources.rest_client.repositories_api.search_rpm(
   ...:         dict(
   ...:             search="bear",
   ...:             urls=[
   ...:                 'https://jlsherrill.fedorapeople.org/fake-repos/revision/none_such/',
   ...:             ],
   ...:         )
   ...:     )
<snip>
HTTP response body: {"errors":[{"status":400,"title":"Error searching RPMs","detail":"Could not find repository with URL: https://jlsherrill.fedorapeople.org/fake-repos/revision/none_such/"}]}
```

### Comment by @swadeley on April 17, 2024 at 02:15 PM UTC

Hi

for /package_groups/names:

```In [1]: app.content_sources.rest_client.repositories_api.search_package_group({"search": "Xfce", "urls": ["https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/"]})
2024-04-17 15:13:58.606 [    INFO] [root] Using <function client_obj_maker.<locals>.make_obj at 0x7f8aac2f0790> object....with url https://ee-u2i7kj4d.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1 and verify_ssl set to True
2024-04-17 15:13:58.606 [    INFO] [iqe.base.auth] Setting auth_type to jwt
2024-04-17 15:13:58.606 [    INFO] [iqe.base.auth] Setting jwt_grant_type to password
2024-04-17 15:13:59.223 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[1]: 
[{'description': 'A lightweight desktop environment that works well on low end '
                 'machines.',
  'id': 'xfce-desktop',
  'package_group_name': 'Xfce',
  'package_list': ['xfce4-screensaver',
                   'tumbler',
                   'NetworkManager-gnome',
                   'openssh-askpass',
                   'gdm',
                   'xfce4-terminal',
                   'thunar-volman',
                   'xfce4-panel',
                   'xfce4-appfinder',
                   'xfwm4',
                   'xfce4-settings',
                   'xfce4-pulseaudio-plugin',
                   'pinentry-gtk',
                   'mousepad',
                   'xfdesktop',
                   'xfconf',
                   'Thunar',
                   'xfce-polkit',
                   'thunar-archive-plugin',
                   'xfce4-session',
                   'xfce4-power-manager']}]

In [2]: app.content_sources.rest_client.repositories_api.search_package_group({"search": "Xfce", "urls": ["https://dl.fedoraproject.org/pub/epel/8/Nothing/x86_64/"]})
2024-04-17 15:14:17.424 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]

HTTP response body: {"errors":[{"status":400,"title":"Error searching package groups","detail":"Could not find repository with URL: https://dl.fedoraproject.org/pub/epel/8/Nothing/x86_64/"}]}

```

### Comment by @swadeley on April 17, 2024 at 02:18 PM UTC

Hi
for /environments/names:

```
In [3]: app.content_sources.rest_client.repositories_api.search_environments(dict(urls=[ "https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/" ], search="kde" ))
2024-04-17 15:17:15.913 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=IDKVpZcrjmcEBcuxduYshjHlDnaIbltl, params=[]
Out[3]: 
[{'description': 'The KDE Plasma Workspaces, a highly-configurable graphical '
                 'user interface which includes a panel, desktop, system icons '
                 'and desktop widgets, and many powerful KDE applications.',
  'environment_name': 'KDE Plasma Workspaces',
  'id': 'kde-desktop-environment'}]

In [4]: app.content_sources.rest_client.repositories_api.search_environments(dict(urls=[ "https://dl.fedoraproject.org/pub/epel/8/Noththing/x86_64/" ], search="kde" ))
2024-04-17 15:17:28.673 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]

HTTP response body: {"errors":[{"status":400,"title":"Error searching environments","detail":"Could not find repository with URL: https://dl.fedoraproject.org/pub/epel/8/Noththing/x86_64/"}]}



### Comment by @swadeley on April 17, 2024 at 02:25 PM UTC

Hi
For /snapshots/rpms/names:

```In [6]:  app.content_sources.rest_client.snapshots_api.search_snapshot_rpms(dict(search='bear', uuids=['cabb9874-9c0c-4d23-84f3-1f5712abdf73'],))
2024-04-17 15:23:40.994 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=RwZnoDFeTgKUTnAwskEGkhxazqrawUBN, params=[]
Out[6]: [{'package_name': 'bear', 'summary': 'A dummy package of bear'}]

In [7]:  app.content_sources.rest_client.snapshots_api.search_snapshot_rpms(dict(search='bear', uuids=['zzzzzzzz-9c0c-4d23-84f3-1f5712abdf73'],))
2024-04-17 15:24:33.095 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]

HTTP response body: {"errors":[{"status":400,"title":"Error searching RPMs","detail":"Could not find snapshot with UUID: zzzzzzzz-9c0c-4d23-84f3-1f5712abdf73"}]}


```

### Comment by @swadeley on April 17, 2024 at 02:37 PM UTC

Hi

for /snapshots/package_groups/names:

```
In [10]: app.content_sources.rest_client.snapshots_api.search_snapshot_package_groups(dict(search='birds', uuids=['48bb5e11-af91-4fa7-be24-000b233969a5']))
2024-04-17 15:35:36.561 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=ErCjMRaEFhzkDkaiJaUSrOpaNIqgNfeG, params=[]
Out[10]: 
[{'description': 'birds',
  'id': 'birds',
  'package_group_name': 'birds',
  'package_list': ['cockateel', 'penguin', 'stork']}]

In [11]: app.content_sources.rest_client.snapshots_api.search_snapshot_package_groups(dict(search='birds', uuids=['zzzzzzzz-af91-4fa7-be24-000b233969a5']))
2024-04-17 15:36:03.386 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]

HTTP response body: {"errors":[{"status":400,"title":"Error searching package groups","detail":"Could not find snapshot with UUID: zzzzzzzz-af91-4fa7-be24-000b233969a5"}]}

```

### Comment by @swadeley on April 17, 2024 at 02:44 PM UTC

Hi

for /snapshots/environments/names:

```
In [12]: app.content_sources.rest_client.snapshots_api.search_snapshot_environments(dict(uuids=['48bb5e11-af91-4fa7-be24-000b233969a5'],search='avians'))
2024-04-17 15:43:55.961 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=lQJAQzBuRTLHxubZPNJEUKSCfvsLhmPj, params=[]
Out[12]: [{'description': 'avians', 'environment_name': 'avians', 'id': 'avians'}]

In [13]: app.content_sources.rest_client.snapshots_api.search_snapshot_environments(dict(uuids=['zzzzzzzz-af91-4fa7-be24-000b233969a5'],search='avians'))
2024-04-17 15:44:14.387 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]


HTTP response body: {"errors":[{"status":400,"title":"Error searching environments","detail":"Could not find snapshot with UUID: zzzzzzzz-af91-4fa7-be24-000b233969a5"}]}

```

### Comment by @swadeley on April 17, 2024 at 02:45 PM UTC

all good!

---

## Reviews

### Review by @rverdile - Commented on April 17, 2024 at 08:23 PM UTC

### Review by @xbhouse - Commented on April 17, 2024 at 08:27 PM UTC

### Review by @xbhouse - Commented on April 18, 2024 at 03:06 PM UTC

### Review by @rverdile - Commented on April 18, 2024 at 06:41 PM UTC

### Review by @rverdile - Approved on April 18, 2024 at 06:42 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/638*
