---
type: pull_request
number: 407
title: "Fixes 2245: add red hat repo for nightly snapshotting"
state: merged
author: jlsherrill
created: 2023-09-29T17:20:10Z
updated: 2023-10-10T18:49:55Z
closed: 2023-10-10T13:55:51Z
merged: 2023-10-10T13:55:51Z
base_branch: main
head_branch: rh_repo
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/407
---

# Pull Request #407: Fixes 2245: add red hat repo for nightly snapshotting

**Author**: @jlsherrill
**Created**: September 29, 2023 at 05:20 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `rh_repo`

## Description

## Summary

Adds an intial set of Red hat repositories that is loaded from a json file.  
Adds a single gpg key loaded for all red hat repos
These repos get created into the repository_configs table with an org_id of -1

## Testing steps

load in ephemeral, or run `make repos-import`

1. curl  http://localhost:8000/api/content-sources/v1.0/repositories/  for some new org
  By default, nothing should be returned
2.  curl http://localhost:8000/api/content-sources/v1.0/repositories/?origin=red_hat
  you should get all the red hat repos
3.  create some repo
4. curl http://localhost:8000/api/content-sources/v1.0/repositories/?origin=red_hat,external
  you  should get both custom and redhat repos


---

## Discussion

### Comment by @jlsherrill on September 29, 2023 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-2245

### Comment by @jlsherrill on September 29, 2023 at 09:12 PM UTC

Realized i need to use the red hat cdn cert for syncing, will work on adding that

### Comment by @jlsherrill on October 03, 2023 at 01:34 PM UTC

/retest

### Comment by @jlsherrill on October 03, 2023 at 03:05 PM UTC

this seems to be failing with: 

{"level":"panic","error":"ERROR: invalid reference to FROM-clause entry for table \"repository_configurations\" (SQLSTATE 42P01)","time":"2023-10-03T15:05:35Z","message":"Failed to save repositories"}
panic: Failed to save repositories

will investigate

Fixed!

### Comment by @jlsherrill on October 03, 2023 at 06:14 PM UTC

/retest

### Comment by @jlsherrill on October 03, 2023 at 06:38 PM UTC

nothing should be returned for a 'new org'.  do you have the NewRepositoryFiltering feature turned on? because that will affect the behavior.  With it turned on, you'll get both.  WIth it turned off (default), you'll get only get external (which for a new org would have none).

### Comment by @jlsherrill on October 03, 2023 at 06:38 PM UTC

updated the testing steps to try to make that more clear

### Comment by @jlsherrill on October 03, 2023 at 06:39 PM UTC

/retest

### Comment by @rverdile on October 03, 2023 at 07:04 PM UTC

> nothing should be returned for a 'new org'. do you have the NewRepositoryFiltering feature turned on? because that will affect the behavior. With it turned on, you'll get both. WIth it turned off (default), you'll get only get external (which for a new org would have none).

Got it. I did have the feature turned on.

### Comment by @swadeley on October 09, 2023 at 12:53 PM UTC

Hi

I can't get this to deploy:

content-sources-backend-kafka-consumer-fc68c66b7-tfqct           0/1     Init:ImagePullBackOff
content-sources-backend-service-897844f57-jl4zg                  0/2     Init:ImagePullBackOff

both fail with `waiting to start: PodInitializing`



### Comment by @jlsherrill on October 10, 2023 at 12:00 PM UTC

/retest

### Comment by @swadeley on October 10, 2023 at 01:55 PM UTC

Hi

deployed new images

```
In [5]: app.content_sources.rest_client.repositories_api.list_repositories(origin='red_hat')
2023-10-10 14:48:00.411 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=shDW6AGqfvP9rRlB0lTYr8SNkVyK1aTf, params=[('origin', 'red_hat')]
Out[5]: 
{'data': [{'account_id': '',
           'content_type': 'rpm',
           'distribution_arch': 'x86_64',
           'distribution_versions': ['8'],
           'failed_introspections_count': 0,
           'gpg_key': '-----BEGIN PGP PUBLIC KEY BLOCK-----\n'
                      '\n'
                      'mQINBErgSTsBEACh2A4b0O9t+vzC9VrVtL1AKvUWi9OPCjkvR7Xd8DtJxeeMZ5eF\n
<big-snip>
           'name': 'Red Hat Enterprise Linux 9 for x86_64 - AppStream (RPMs)',
           'org_id': '-1',
           'origin': 'red_hat','


```

`In [6]: app.content_sources.rest_client.repositories_api.list_repositories(origin='external')
2023-10-10 14:52:22.096 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=9Yc3132QeURHAUyhKdnlIQXZYqAZvf1h, params=[('origin', 'external')]
Out[6]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'any',
           'distribution_versions': ['any'],
           'failed_introspections_count': 0,
           'gpg_key': '',`

and this gets all repos:
`In [9]: app.content_sources.rest_client.repositories_api.list_repositories(origin='external,red_hat')
`

Thank you

---

## Reviews

### Review by @jlsherrill - Commented on September 29, 2023 at 05:28 PM UTC

### Review by @jlsherrill - Commented on September 29, 2023 at 05:31 PM UTC

### Review by @rverdile - Commented on September 29, 2023 at 06:14 PM UTC

### Review by @jlsherrill - Commented on September 29, 2023 at 06:24 PM UTC

### Review by @jlsherrill - Commented on September 29, 2023 at 06:38 PM UTC

### Review by @jlsherrill - Commented on September 29, 2023 at 06:38 PM UTC

### Review by @rverdile - Commented on October 03, 2023 at 06:26 PM UTC

> curl http://localhost:8000/api/content-sources/v1.0/repositories/ for some org
By default, nothing should be returned

Here it says nothing should be returned when I curl this. If I curl this, I get both red hat repos and whatever repos have been created for the org. I would expect to get both (or only custom), but which is the behavior you intended?

### Review by @rverdile - Commented on October 03, 2023 at 06:41 PM UTC

### Review by @jlsherrill - Commented on October 03, 2023 at 06:47 PM UTC

### Review by @rverdile - Commented on October 03, 2023 at 07:01 PM UTC

### Review by @rverdile - Commented on October 03, 2023 at 07:08 PM UTC

### Review by @jlsherrill - Commented on October 03, 2023 at 07:13 PM UTC

### Review by @rverdile - Commented on October 03, 2023 at 07:18 PM UTC

### Review by @jlsherrill - Commented on October 03, 2023 at 07:22 PM UTC

### Review by @rverdile - Commented on October 04, 2023 at 07:32 PM UTC

### Review by @jlsherrill - Commented on October 04, 2023 at 08:09 PM UTC

### Review by @rverdile - Approved on October 04, 2023 at 08:26 PM UTC

tested and looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/407*
