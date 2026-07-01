---
type: pull_request
number: 922
title: "HMS-4934: Add modules to introspected data"
state: merged
author: jlsherrill
created: 2024-12-16T20:55:52Z
updated: 2025-01-17T13:19:46Z
closed: 2025-01-17T13:19:46Z
merged: 2025-01-17T13:19:46Z
base_branch: main
head_branch: modstream
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/922
---

# Pull Request #922: HMS-4934: Add modules to introspected data

**Author**: @jlsherrill
**Created**: December 16, 2024 at 08:55 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `modstream`

## Description

## Summary

Adds module streams as searchable for introspected repos

## Testing steps

1.  Adjust go.mod to point to your localy checkout of yummy, which should have https://github.com/content-services/yummy/pull/31

```
replace github.com/content-services/yummy => /home/jlsherri/git/yummy/
```

2. `go get ./...`
3. ` make db-migrate-up`
4. `make repos-import`
5. ```go run cmd/external-repos/main.go introspect go run cmd/external-repos/main.go introspect https://cdn.redhat.com/content/dist/rhel8/8/x86_64/appstream/os```
6. ```go run cmd/external-repos/main.go introspect go run cmd/external-repos/main.go introspect https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os```
7.   play around with the module streams search api:
```
#### module streams
POST http://localhost:8000/api/content-sources/v1.0/module_streams/search
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxIn0sImFjY291bnRfbnVtYmVyIjoiZm9vIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K
Content-Type: application/json

{"urls":["https://cdn.redhat.com/content/dist/rhel8/8/x86_64/appstream/os"],
  "search":"",
  "rpm_names": []
}
```

feel free to play with the search and rpm_names


---

## Discussion

### Comment by @jlsherrill on December 16, 2024 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-4932

### Comment by @jlsherrill on December 17, 2024 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-HMS-4932

### Comment by @jlsherrill on December 18, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-HMS-4934

### Comment by @jlsherrill on December 18, 2024 at 06:12 PM UTC

https://issues.redhat.com/browse/Fixes HMS-4934

### Comment by @jlsherrill on December 18, 2024 at 07:05 PM UTC

https://issues.redhat.com/browse/HMS-4934

### Comment by @swadeley on January 08, 2025 at 04:26 AM UTC

/retest

### Comment by @swadeley on January 08, 2025 at 06:12 AM UTC

@jlsherrill this failed to build

### Comment by @jlsherrill on January 08, 2025 at 12:53 PM UTC

 @swadeley yes, it will fail to build until https://github.com/content-services/yummy/pull/31 is merged and i update it to use the new yummy

### Comment by @jlsherrill on January 13, 2025 at 06:11 PM UTC

good call, fixed

### Comment by @jlsherrill on January 14, 2025 at 09:21 PM UTC

good call, updated!

### Comment by @rverdile on January 15, 2025 at 09:16 PM UTC

acked the yummy PR! Small question unrelated to the yummy changes but looks good

### Comment by @swadeley on January 16, 2025 at 04:27 AM UTC

Hi

testing in ephemeral  first with : https://rverdile.fedorapeople.org/dummy-repos/modules/repo1/

(will test introspected repo in stage)

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

### Comment by @swadeley on January 17, 2025 at 03:42 AM UTC

Hi @jlsherrill , please rebase

---

## Reviews

### Review by @rverdile - Commented on January 10, 2025 at 08:22 PM UTC

If I search for a stream that doesn't exist, I see null as the response. Would empty list be better? I think that's what I would've expected

### Review by @rverdile - Commented on January 13, 2025 at 08:39 PM UTC

I noticed if I create and introspect a repository before checking-out this PR, then checkout this PR and re-introspect, the module streams don't get added because the repository hasn't changed (I suspect)

### Review by @rverdile - Commented on January 15, 2025 at 09:16 PM UTC

### Review by @jlsherrill - Commented on January 15, 2025 at 09:31 PM UTC

### Review by @rverdile - Commented on January 16, 2025 at 01:56 PM UTC

### Review by @rverdile - Approved on January 16, 2025 at 01:56 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/922*
