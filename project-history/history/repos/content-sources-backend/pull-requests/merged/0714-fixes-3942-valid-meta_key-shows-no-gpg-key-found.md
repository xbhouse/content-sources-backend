---
type: pull_request
number: 714
title: "Fixes 3942: valid meta_key shows no gpg key found"
state: merged
author: xbhouse
created: 2024-06-20T18:31:24Z
updated: 2024-06-21T15:30:29Z
closed: 2024-06-21T15:08:36Z
merged: 2024-06-21T15:08:36Z
base_branch: main
head_branch: 3942
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/714
---

# Pull Request #714: Fixes 3942: valid meta_key shows no gpg key found

**Author**: @xbhouse
**Created**: June 20, 2024 at 06:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3942`

## Description

## Summary

Adds the request body to the api spec for `/repository_parameters/external_gpg_key/`. This was missing and is needed for QE to test validation of gpg keys when fetching from a remote repo.

## Testing steps

- Add a repository (QE example [here](https://issues.redhat.com/browse/HMS-3942) is using this one: https://jlsherrill.fedorapeople.org/fake-repos/signed/) 
- Fetch the gpg key from this URL: https://jlsherrill.fedorapeople.org/fake-repos/signed/GPG-KEY.gpg
```
POST /repository_parameters/external_gpg_key/

Request:
{
     "url": "https://jlsherrill.fedorapeople.org/fake-repos/signed/GPG-KEY.gpg"
}

Response:
{
    "gpg_key": "-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmQINBGPhWqgBEACnFSBzQlJ1ilLhS/toT46iJqInMQFtQz...-----END PGP PUBLIC KEY BLOCK-----\n"
}    
```
- Using the gpg key from the previous response, validate the gpg key for the repository
```
POST /repository_parameters/validate/

Request:
[
  {
    "gpg_key": "-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmQINBGPhWqgBEACnFS...",
    "metadata_verification": true,
    "name": "test1",
    "url": "https://jlsherrill.fedorapeople.org/fake-repos/signed/",
    "uuid": "1c90d2e1-642b-409f-a3ee-256e1a3e6d78"
  }
]

Response:
[
    {
        "name": {
            "skipped": false,
            "valid": true,
            "error": ""
        },
        "url": {
            "skipped": false,
            "valid": true,
            "error": "",
            "http_code": 200,
            "metadata_present": true,
            "metadata_signature_present": false
        },
        "gpg_key": {
            "skipped": false,
            "valid": true,
            "error": ""
        }
    }
]
```
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 20, 2024 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-3942

### Comment by @swadeley on June 21, 2024 at 11:55 AM UTC

Hi

I tried (with http and https):

`app.content_sources.rest_client.gpg_key_api.fetch_gpg_key(dict(url="https://jlsherrill.fedorapeople.org/fake-repos/signed/GPG-KEY.gpg"))
`
and

`app.content_sources.rest_client.gpg_key_api.fetch_gpg_key({"url": "https://jlsherrill.fedorapeople.org/fake-repos/signed/GPG-KEY.gpg"})`


but get

`HTTP response body: {"errors":[{"status":400,"detail":"Bad Request"}]}
`

### Comment by @swadeley on June 21, 2024 at 12:35 PM UTC

Hi

If I create a repo with a key I can download that:
```
In [18]: app.content_sources.rest_client.repositories_api.get_gpg_key_file('bf5b5ddf-6cdf-454b-95c9-a37f25b031c3')
2024-06-21 13:30:07.510 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[18]: '-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmQINBGP<snip><snip>uQy8Gk0M=\n=Avdh\n-----END PGP PUBLIC KEY BLOCK-----\n'
```

Confusingly the attribute_map says:

```
In [19]: app.content_sources.rest_client.repositories_api.get_gpg_key_file.attribute_map
Out[19]: {'uuid': 'uuid'}
```

### Comment by @xbhouse on June 21, 2024 at 02:28 PM UTC

hi @swadeley can you try with the new openapi spec? the endpoint was missing a slash, fixed that and this request works for me:

```
In [1]: app.content_sources.rest_client.gpg_key_api.fetch_gpg_key(dict(url="https://jlsherrill.fedorapeople.org/fake-repos/signed/GPG-KEY.gpg"))
2024-06-21 09:59:41.031 [    INFO] [root] Using <function client_obj_maker.<locals>.make_obj at 0x10e023430> object....with url https://ee-6kkc3juv.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1 and verify_ssl set to True
2024-06-21 09:59:41.031 [    INFO] [iqe.base.auth] Setting auth_type to jwt
2024-06-21 09:59:41.031 [    INFO] [iqe.base.auth] Setting jwt_grant_type to password
2024-06-21 09:59:41.590 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=4ad1caab-5bd4-4476-8e48-a6539865c333, params=[]
Out[1]: 
{'gpg_key': '-----BEGIN PGP PUBLIC KEY BLOCK-----\n'
            '\n'
            'mQINBGPhWqgBEACnFSBzQlJ1ilLhS/toT46iJqInMQFtQzQgexR6/XKNfvovqwzH\n'
            'O5KlPFc3zfl1PWfUqhhH4o4YcxXpRa9pv7MiJEc+AMLC9ziNutw+eAdYpNU8l29k\n'
            'ZgoMQWrT5zKB5o48He7cgZigGBoHTL2MVoMQ5bPZDu1mh3ERVeefL2bEc5D00NPD\n'
            'QMDca7LAoy5V131sm4zuVWG8KMomCvB9Jp5B3IraZTdZMveV3sFhQRui7nIJTHkk\n'
            ...
            '-----END PGP PUBLIC KEY BLOCK-----\n'}
```

### Comment by @swadeley on June 21, 2024 at 02:45 PM UTC

Hi @xbhouse , yes, that fixed it. Thank you.

```
In [1]: app.content_sources.rest_client.gpg_key_api.fetch_gpg_key({'url':'http://jlsherrill.fedorapeople.org/fake-repos/signed/GPG-KEY.gpg'})

2024-06-21 15:41:19.119 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[1]: 
{'gpg_key': '-----BEGIN PGP PUBLIC KEY BLOCK-----\n'
            '\n'
            'mQINBGPhWqgBEACnFSBzQlJ1ilLhS/toT46iJqInMQFtQzQgexR6/XKNfvovqwzH\n'

```

---

## Reviews

### Review by @jlsherrill - Approved on June 20, 2024 at 07:10 PM UTC

good catch!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/714*
