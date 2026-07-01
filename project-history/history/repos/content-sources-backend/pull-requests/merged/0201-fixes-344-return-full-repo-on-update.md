---
type: pull_request
number: 201
title: "Fixes 344: return full repo on update"
state: merged
author: jlsherrill
created: 2023-02-08T15:14:21Z
updated: 2023-02-14T16:30:37Z
closed: 2023-02-14T16:00:30Z
merged: 2023-02-14T16:00:30Z
base_branch: main
head_branch: 344
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/201
---

# Pull Request #201: Fixes 344: return full repo on update

**Author**: @jlsherrill
**Created**: February 08, 2023 at 03:14 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `344`

## Description

## Summary
previously PUT/PATCH /repositories/UUID/  would return a string, this modifies it to to return a repository object.
## Testing steps

create a repo:
```
$ curl -X POST    -H 'Content-Type: application/json'  -H "`./scripts/header.sh 1`"  http://localhost:8000/api/content-sources/v1/repositories/  -d '{"name":"foo", "url":"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/"}'


{"uuid":"49bbc072-53cc-4fb7-9f8d-3dbda9d4c553","name":"foo","url":"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/","distribution_versions":["any"],"distribution_arch":"any","account_id":"","org_id":"1","last_introspection_time":"","last_success_introspection_time":"","last_update_introspection_time":"","last_introspection_error":"","package_count":0,"status":"Valid","gpg_key":"","metadata_verification":false}
```
update it:

```
$ curl -X PATCH -H 'Content-Type: application/json'  -H "`./scripts/header.sh 1`"  http://localhost:8000/api/content-sources/v1/repositories/49bbc072-53cc-4fb7-9f8d-3dbda9d4c553  -d '{"name":"foo2"}'
{"uuid":"49bbc072-53cc-4fb7-9f8d-3dbda9d4c553","name":"foo2","url":"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/","distribution_versions":["any"],"distribution_arch":"any","account_id":"","org_id":"1","last_introspection_time":"2023-02-08 10:12:51.659822 -0500 EST","last_success_introspection_time":"2023-02-08 10:12:51.659822 -0500 EST","last_update_introspection_time":"2023-02-01 11:27:46.514447 -0500 EST","last_introspection_error":"","package_count":32,"status":"Valid","gpg_key":"","metadata_verification":false}
```

notice the response


---

## Discussion

### Comment by @jlsherrill on February 08, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-344

### Comment by @swadeley on February 13, 2023 at 08:32 PM UTC

/retest

### Comment by @mayurilahane on February 14, 2023 at 03:55 PM UTC

varified !
lgtm !

`ipdb> response = app.content_sources.rest_client.repositories_api.full_update_repository(
        repo.uuid,
        {
            "distribution_arch": "x86_64",
            "distribution_versions": ["7"],
            "name": "updated",
            "url": "http://jlsherrill.fedorapeople.org/fake-repos/needed-errata/",
        },
    )
`

----------------------------------------
`ipdb> response
{'account_id': '12345',
 'distribution_arch': 'x86_64',
 'distribution_versions': ['7'],
 'gpg_key': '',
 'last_introspection_error': '',
 'last_introspection_time': '2023-02-14 15:51:32.543639 +0000 UTC',
 'last_success_introspection_time': '2023-02-14 15:51:32.543639 +0000 UTC',
 'last_update_introspection_time': '2023-02-13 20:56:01.085204 +0000 UTC',
 'metadata_verification': False,
 'name': 'updated',
 'org_id': '12345',
 'package_count': 32,
 'status': 'Valid',
 'url': 'http://jlsherrill.fedorapeople.org/fake-repos/needed-errata/',
 'uuid': '2833acff-7768-4413-9f06-df7b9ae7c878'} `

---

## Reviews

### Review by @rverdile - Commented on February 09, 2023 at 03:44 PM UTC

### Review by @rverdile - Commented on February 09, 2023 at 03:59 PM UTC

### Review by @rverdile - Approved on February 09, 2023 at 03:59 PM UTC

lgtm!

### Review by @jlsherrill - Commented on February 09, 2023 at 05:58 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/201*
