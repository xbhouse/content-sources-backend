---
type: pull_request
number: 897
title: "HMS-4932: Add search module streams"
state: merged
author: Andrewgdewar
created: 2024-11-18T22:27:35Z
updated: 2025-01-16T03:58:01Z
closed: 2025-01-07T01:58:11Z
merged: 2025-01-07T01:58:11Z
base_branch: main
head_branch: HMS-4932
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/897
---

# Pull Request #897: HMS-4932: Add search module streams

**Author**: @Andrewgdewar
**Created**: November 18, 2024 at 10:27 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `HMS-4932`

## Description

## Summary

Adds the POST "snapshots/module_streams/search" endpoint.
That takes the following inputs
```
{
  "uuids": [
    "9d164ac6-b866-4808-af1f-847fbf588658",
    "36e449b9-3d61-43e8-a495-6ac332fdd3a8"
  ],
  "rpm_names": [],
  "search": ""
}

```
and returns a list of modules with their streams:

```
{
 [
  {
    "module_name": "mariadb",
    "streams": [
      {
        "name": "mariadb",
        "stream": "10.11",
        "context": "rhel9",
        "arch": "x86_64",
        "cersion": "9040020240126110506",
        "description": "MariaDB is a....",
        "profiles": {
          "client": [
            "mariadb"
          ],
          "galera": [
            "mariadb-server",
            "mariadb-server-galera"
          ],
          "server": [
            "mariadb-server"
          ]
        }
      }
    ]
  },
]

}
```

- [ ] Needs to [update tangy ](https://github.com/content-services/tang/pull/14)and tangy reference prior to review

## Testing steps



---

## Discussion

### Comment by @jlsherrill on November 18, 2024 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-4932

### Comment by @jlsherrill on November 27, 2024 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-4932

EDIT: remove duplicated HMS

### Comment by @swadeley on December 02, 2024 at 01:26 AM UTC

/retest

### Comment by @jlsherrill on December 02, 2024 at 01:30 AM UTC

https://issues.redhat.com/browse/HMS-HMS-4932

### Comment by @swadeley on December 03, 2024 at 07:14 AM UTC

Warning: Unexpected input(s) 'skip-go-installation', 

### Comment by @Andrewgdewar on December 03, 2024 at 04:16 PM UTC

> Warning: Unexpected input(s) 'skip-go-installation',

This requires the tangy PR to be integrated, don't test this yet.

### Comment by @jlsherrill on December 12, 2024 at 05:41 PM UTC

seems to be some linting issue

### Comment by @swadeley on December 13, 2024 at 02:20 AM UTC

/retest

### Comment by @swadeley on December 13, 2024 at 02:56 AM UTC

> > Warning: Unexpected input(s) 'skip-go-installation',
> 
> This requires the tangy PR to be integrated, don't test this yet.

OK but winter is coming :)

`02:31:28 pkg/dao/module_streams.go:65:37: (*config.Tang).RpmRepositoryVersionModuleStreamsList undefined (type tangy.Tangy has no field or method RpmRepositoryVersionModuleStreamsList)`

### Comment by @Andrewgdewar on December 16, 2024 at 05:18 PM UTC

> seems to be some linting issue

tangy..

### Comment by @jlsherrill on December 17, 2024 at 06:21 PM UTC

few small comment, but overall looks good

### Comment by @jlsherrill on December 19, 2024 at 01:46 PM UTC

lets merge https://github.com/content-services/tang/pull/14 and tag a new tang version, then pull that update in here?

### Comment by @jlsherrill on December 20, 2024 at 08:41 PM UTC

note this require iqe update to merge

### Comment by @swadeley on January 02, 2025 at 01:51 AM UTC

@Andrewgdewar please rebase

### Comment by @swadeley on January 02, 2025 at 01:51 AM UTC

/retest

### Comment by @swadeley on January 03, 2025 at 06:26 AM UTC

Hi

I built the API docs for IQE but its taking forever to get a suitable RHEL repo in ephemeral snapshotted. Will test in stage.

Format for IQE:
```
In [9]: app.content_sources.rest_client.module_streams_api.search_snapshot_module_streams(dict(search='mariadb', uuids=['776a0172-5855-4734-b382-8e4d60a44a84'], rpm_names=[]))

```

### Comment by @swadeley on January 03, 2025 at 07:31 AM UTC

Hi

testing with: `https://rverdile.fedorapeople.org/dummy-repos/modules/repo1/`

```
In [9]: app.content_sources.rest_client.module_streams_api.search_snapshot_module_streams(dict(search='', uuids=['6d1da6ea-6603-4cb9-bec4-f803dcaee830'], rpm_names=['kangaroo']))
2025-01-03 15:29:40.761 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bt0c0eqb.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/module_streams/search with query params [] and x-rh-insights-request-id=<>
Out[9]: 
[{'module_name': 'kangaroo',
  'streams': [{'arch': 'noarch',
               'context': 'deadbeef',
               'description': 'A module for the kangaroo 0.2 package',
               'name': 'kangaroo',
               'profiles': {'default': ['kangaroo']},
               'stream': '0',
               'version': '20180704111719'}]}]
```

but nothing if blank search:

```
In [12]: app.content_sources.rest_client.module_streams_api.search_snapshot_module_streams(dict(search="", uuids=['6d1da6ea-6603-4cb9-bec4-f803dcaee830'], rpm_names=[""]))
2025-01-03 15:33:05.554 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bt0c0eqb.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/module_streams/search with query params [] and x-rh-insights-request-id=<>
Out[12]: []
```

### Comment by @Andrewgdewar on January 06, 2025 at 06:41 PM UTC

> Hi
> 
> testing with: `https://rverdile.fedorapeople.org/dummy-repos/modules/repo1/`
> 
> ```
> In [9]: app.content_sources.rest_client.module_streams_api.search_snapshot_module_streams(dict(search='', uuids=['6d1da6ea-6603-4cb9-bec4-f803dcaee830'], rpm_names=['kangaroo']))
> 2025-01-03 15:29:40.761 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bt0c0eqb.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/module_streams/search with query params [] and x-rh-insights-request-id=<>
> Out[9]: 
> [{'module_name': 'kangaroo',
>   'streams': [{'arch': 'noarch',
>                'context': 'deadbeef',
>                'description': 'A module for the kangaroo 0.2 package',
>                'name': 'kangaroo',
>                'profiles': {'default': ['kangaroo']},
>                'stream': '0',
>                'version': '20180704111719'}]}]
> ```
> 
> but nothing if blank search:
> 
> ```
> In [12]: app.content_sources.rest_client.module_streams_api.search_snapshot_module_streams(dict(search="", uuids=['6d1da6ea-6603-4cb9-bec4-f803dcaee830'], rpm_names=[""]))
> 2025-01-03 15:33:05.554 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bt0c0eqb.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/module_streams/search with query params [] and x-rh-insights-request-id=<>
> Out[12]: []
> ```

This is expected. You need to give the api at least one uuid.

### Comment by @swadeley on January 07, 2025 at 01:57 AM UTC

> > Hi
> > testing with: `https://rverdile.fedorapeople.org/dummy-repos/modules/repo1/`
> > ```
> > In [9]: app.content_sources.rest_client.module_streams_api.search_snapshot_module_streams(dict(search='', uuids=['6d1da6ea-6603-4cb9-bec4-f803dcaee830'], rpm_names=['kangaroo']))
> > 2025-01-03 15:29:40.761 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bt0c0eqb.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/module_streams/search with query params [] and x-rh-insights-request-id=<>
> > Out[9]: 
> > [{'module_name': 'kangaroo',
> >   'streams': [{'arch': 'noarch',
> >                'context': 'deadbeef',
> >                'description': 'A module for the kangaroo 0.2 package',
> >                'name': 'kangaroo',
> >                'profiles': {'default': ['kangaroo']},
> >                'stream': '0',
> >                'version': '20180704111719'}]}]
> > ```
> > 
> > 
> >     
> >       
> >     
> > 
> >       
> >     
> > 
> >     
> >   
> > but nothing if blank search:
> > ```
> > In [12]: app.content_sources.rest_client.module_streams_api.search_snapshot_module_streams(dict(search="", uuids=['6d1da6ea-6603-4cb9-bec4-f803dcaee830'], rpm_names=[""]))
> > 2025-01-03 15:33:05.554 [    INFO] [iqe.base.rest_client] REST: POST https://ee-bt0c0eqb.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/module_streams/search with query params [] and x-rh-insights-request-id=<>
> > Out[12]: []
> > ```
> 
> This is expected. You need to give the api at least one uuid.

OK, thank you

### Comment by @swadeley on January 16, 2025 at 03:57 AM UTC

forgot to add this better example output:

```
In [7]: app.content_sources.rest_client.module_streams_api.search_snapshot_module_streams(dict(search='walrus', uuids=['177e9b6f-0125-4628-9953-40195e368822'], rpm_names=[]))
2025-01-16 11:52:41.209 [    INFO] [iqe.base.rest_client] REST: POST https://ee-fryb0ek4.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/snapshots/module_streams/search with query params [] and x-rh-insights-request-id=<>
Out[7]: 
[{'module_name': 'walrus',
  'streams': [{'arch': 'x86_64',
               'context': 'c0ffee42',
               'description': 'A module for the walrus 0.71 package',
               'name': 'walrus',
               'profiles': {'default': ['walrus'], 'flipper': ['walrus']},
               'stream': '0.71',
               'version': '20180707144203'},
              {'arch': 'x86_64',
               'context': 'deadbeef',
               'description': 'A module for the walrus 5.21 package',
               'name': 'walrus',
               'profiles': {'default': ['walrus']},
               'stream': '5.21',
               'version': '20180704144203'}]}]
```

---

## Reviews

### Review by @jlsherrill - Commented on December 02, 2024 at 01:41 PM UTC

### Review by @jlsherrill - Commented on December 02, 2024 at 01:41 PM UTC

### Review by @jlsherrill - Commented on December 02, 2024 at 01:44 PM UTC

### Review by @Andrewgdewar - Commented on December 02, 2024 at 05:58 PM UTC

### Review by @Andrewgdewar - Commented on December 02, 2024 at 08:22 PM UTC

### Review by @jlsherrill - Commented on December 02, 2024 at 08:26 PM UTC

### Review by @jlsherrill - Commented on December 02, 2024 at 08:27 PM UTC

### Review by @jlsherrill - Commented on December 02, 2024 at 08:30 PM UTC

### Review by @Andrewgdewar - Commented on December 03, 2024 at 04:04 PM UTC

### Review by @rverdile - Commented on December 04, 2024 at 03:06 PM UTC

### Review by @rverdile - Commented on December 04, 2024 at 03:08 PM UTC

### Review by @jlsherrill - Commented on December 04, 2024 at 09:47 PM UTC

### Review by @rverdile - Commented on December 05, 2024 at 02:53 PM UTC

### Review by @Andrewgdewar - Commented on December 05, 2024 at 10:14 PM UTC

### Review by @jlsherrill - Commented on December 09, 2024 at 04:47 PM UTC

### Review by @jlsherrill - Commented on December 09, 2024 at 04:47 PM UTC

### Review by @jlsherrill - Commented on December 09, 2024 at 04:47 PM UTC

### Review by @jlsherrill - Commented on December 09, 2024 at 04:58 PM UTC

### Review by @rverdile - Commented on December 09, 2024 at 07:06 PM UTC

### Review by @Andrewgdewar - Commented on December 09, 2024 at 08:03 PM UTC

### Review by @rverdile - Commented on December 10, 2024 at 03:01 PM UTC

### Review by @Andrewgdewar - Commented on December 11, 2024 at 09:30 PM UTC

### Review by @Andrewgdewar - Commented on December 16, 2024 at 05:16 PM UTC

### Review by @jlsherrill - Commented on December 16, 2024 at 06:40 PM UTC

### Review by @jlsherrill - Commented on December 16, 2024 at 08:26 PM UTC

### Review by @jlsherrill - Commented on December 17, 2024 at 06:20 PM UTC

### Review by @jlsherrill - Commented on December 17, 2024 at 06:20 PM UTC

### Review by @Andrewgdewar - Commented on December 18, 2024 at 05:23 PM UTC

### Review by @jlsherrill - Approved on December 20, 2024 at 08:41 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/897*
