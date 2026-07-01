---
type: pull_request
number: 432
title: "Fixes 2832: trim introspection error"
state: merged
author: jlsherrill
created: 2023-10-17T19:43:19Z
updated: 2023-10-19T18:40:23Z
closed: 2023-10-19T18:05:33Z
merged: 2023-10-19T18:05:33Z
base_branch: main
head_branch: 2832
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/432
---

# Pull Request #432: Fixes 2832: trim introspection error

**Author**: @jlsherrill
**Created**: October 17, 2023 at 07:43 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2832`

## Description

## Summary
This trims the introspection error to 255 characters

## Testing steps
I'm not sure how to reproduce the error, but i've added a unit test to cover it.


---

## Discussion

### Comment by @jlsherrill on October 17, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-2832

### Comment by @jlsherrill on October 17, 2023 at 10:16 PM UTC

oh good idea!  So to reproduce, you would create a repo with a url of "http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

With this fix, you will get a proper error for introspection

### Comment by @swadeley on October 19, 2023 at 10:33 AM UTC

Hi

I used API to create repo:
```

Out[2]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 0,
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '',
 'last_snapshot_task_uuid': '7fd3afa7-3678-43db-8fab-dbe2c6af44ad',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'name': 'pr-432-test1',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 0,
 'snapshot': True,
 'status': 'Pending',
 'url': 'http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/',
 'uuid': 'df1e94ae-b5ec-457e-95ec-be22cb503400'}
```

On the Admin > Snapshots > Task Details tab I see a traceback followed by "[(UnicodeError: label too long). ](description: encoding with 'idna' codec failed (UnicodeError: label too long).)"

Using API:

```
In [3]: app.content_sources.rest_client.repositories_api.get_repository('df1e94ae-b5ec-457e-95ec-be22cb503400')
2023-10-19 11:31:35.571 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=VHwnPZcOZLKHNvssvSOFkWIOoFnhGCSK, params=[]
Out[3]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 1,
 'gpg_key': '',
 'last_introspection_error': 'GET error for file '
                             'http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/repodata/repomd.xml: '
                             'Get '
                             '"http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
 'last_introspection_time': '2023-10-19T10:27:03Z',
 'last_snapshot_task_uuid': '7fd3afa7-3678-43db-8fab-dbe2c6af44ad',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'name': 'pr-432-test1',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 0,
 'snapshot': True,
 'status': 'Invalid',
 'url': 'http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/',
 'uuid': 'df1e94ae-b5ec-457e-95ec-be22cb503400'}
```

### Comment by @swadeley on October 19, 2023 at 10:42 AM UTC

I created another repo with URL of 254 characters.

```
In [4]:      repo = dict(
   ...:            snapshot=True,
   ...:            name='pr-432-test2',
   ...:            url='http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbb
   ...: bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb'
   ...:       )

In [5]: app.content_sources.rest_client.repositories_api.create_repository(repo)
<snip>
ServiceException: (500)
Reason: Internal Server Error
HTTP response headers: HTTPHeaderDict({'content-length': '146', 'content-type': 'application/json; charset=UTF-8', 'date': 'Thu, 19 Oct 2023 10:39:21 GMT', 
<snip>
HTTP response body: {"errors":[{"status":500,"title":"Error creating repository","detail":"ERROR: value too long for type character varying(255) (SQLSTATE 22001)"}]}

```


### Comment by @swadeley on October 19, 2023 at 11:06 AM UTC

Hi @jlsherrill 

When you say trim introspection error, is this what you expected:

```
'last_introspection_error': 'GET error for file '
                             'http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/repodata/repomd.xml: '
                             'Get '
                             '"http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
```

In the web UI, Admin tab, introspection is marked as "Completed" and Task Details shows:

```
Type
    introspect
Status
    completed
Queued At
    19 Oct 2023 11:27 UTC+01:00
Started At
    19 Oct 2023 11:27 UTC+01:00
Finished At
    19 Oct 2023 11:27 UTC+01:00
Error
    None
```

"Finished" and no error suggests introspection was completed,
but API shows no introspection time, suggesting it was not completed:

```
(In [6]: app.content_sources.rest_client.repositories_api.get_repository('df1e94ae-b5ec-457e-95ec-be22cb503400')
2023-10-19 11:56:11.290 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, 
<snip>
Out[6]: 
{'account_id': '0369233',
 'content_type': 'rpm',
 'distribution_arch': 'any',
 'distribution_versions': ['any'],
 'failed_introspections_count': 1,
 'gpg_key': '',
 'last_introspection_error': 'GET error for file '
                             'http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/repodata/repomd.xml: '
                             'Get '
                             '"http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
 'last_introspection_time': '2023-10-19T10:27:03Z',
 'last_snapshot_task_uuid': '7fd3afa7-3678-43db-8fab-dbe2c6af44ad',
 'last_success_introspection_time': '',
 'last_update_introspection_time': '',
 'metadata_verification': False,
 'name': 'pr-432-test1',
 'org_id': '3340851',
 'origin': 'external',
 'package_count': 0,
 'snapshot': True,
 'status': 'Invalid',
 'url': 'http://aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/',
 'uuid': 'df1e94ae-b5ec-457e-95ec-be22cb503400'}
```

### Comment by @jlsherrill on October 19, 2023 at 05:48 PM UTC

Yes, without this change the 'last_introspection_error' would be blank.  Failed introspections are still marked as 'successful' from a task level, but this may change soon

### Comment by @swadeley on October 19, 2023 at 06:05 PM UTC

OK, introspection task completed not equal to successful introspection. Thank you.

---

## Reviews

### Review by @rverdile - Commented on October 17, 2023 at 08:26 PM UTC

lgtm!

tested by creating a repository with a long URL, but still under the URL character limit. Then the error (which contains the URL) was over the character limit.


### Review by @rverdile - Approved on October 18, 2023 at 01:19 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/432*
