---
type: pull_request
number: 809
title: "Fixes 4665: improve error responses for upload apis"
state: merged
author: xbhouse
created: 2024-09-10T15:03:24Z
updated: 2024-09-12T11:00:30Z
closed: 2024-09-12T10:45:09Z
merged: 2024-09-12T10:45:09Z
base_branch: main
head_branch: 4665
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/809
---

# Pull Request #809: Fixes 4665: improve error responses for upload apis

**Author**: @xbhouse
**Created**: September 10, 2024 at 03:03 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4665`

## Description

## Summary

This fixes the error responses for the upload APIs

## Testing steps

1. Create an upload with a size of 0 or less. You should see this error in the response:
```
{
    "errors": [
        {
            "status": 400,
            "title": "error creating upload",
            "detail": "upload size must be greater than 0"
        }
    ]
}
```
2. Create an upload with a size greater than 0 and upload a chunk with the wrong Content-Range start and end bytes. You should see this error in the response:
```
{
    "errors": [
        {
            "status": 400,
            "title": "error uploading chunk",
            "detail": "uploading file chunk: 400 Bad Request: {\"non_field_errors\":[\"Chunk size does not match content range.\"]}"
        }
    ]
}
```
3. Try to upload a chunk without a file in the request. You should see a 400 response code instead of a 500



---

## Discussion

### Comment by @jlsherrill on September 10, 2024 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-4665

### Comment by @xbhouse on September 10, 2024 at 07:01 PM UTC

/retest

### Comment by @jlsherrill on September 11, 2024 at 08:52 PM UTC

small nitpick, everything else works great and looks great

### Comment by @swadeley on September 12, 2024 at 10:44 AM UTC

Hi

```
In [6]: ls -laZ /tmp/testaa
-rw-r--r--. 1 sjw sjw unconfined_u:object_r:user_tmp_t:s0 7056 Sep 10 13:44 /tmp/testaa

In [7]: app.content_sources.rest_client.repositories_api.create_upload({"size": 7056})
2024-09-12 11:35:53.099 [    INFO] [iqe.base.rest_client] REST: POST https://ee-mnkask2u.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/ with query params [] and x-rh-insights-request-id=a<>9
Out[7]: 
{'created': '2024-09-12T10:35:53.047322Z',
 'last_updated': '2024-09-12T10:35:53.04734Z',
 'size': 7056,
 'upload_uuid': '0191e5cd-c356-7b62-b659-0103ea715995'}

In [8]: with open('/tmp/testaa', 'rb') as file:
   ...:     response = app.content_sources.rest_client.repositories_api.upload_chunk('0191e5cd-c356-7b62-b659-0103ea715995', 'bytes 0-666/*', file, 'f78250d172b47cead1373a996a4cab330fb8d7d2aa59d864355ea7ef5767706
   ...: 5')
   ...: 

<snip traceback>
HTTP response body: {"errors":[{"status":400,"title":"error uploading chunk","detail":"uploading file chunk: 400 Bad Request: {\"non_field_errors\":[\"Chunk size does not match content range.\"]}"}]}

```

```
In [11]: app.content_sources.rest_client.repositories_api.create_upload({"size": 0})
2024-09-12 11:41:34.239 [    INFO] [iqe.base.rest_client] REST: POST https://ee-mnkask2u.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/repositories/uploads/ 
<snip traceback>
HTTP response body: {"errors":[{"status":400,"title":"error creating upload","detail":"upload size must be greater than 0"}]}
```

---

## Reviews

### Review by @xbhouse - Commented on September 10, 2024 at 03:26 PM UTC

### Review by @jlsherrill - Commented on September 11, 2024 at 01:24 PM UTC

### Review by @jlsherrill - Commented on September 11, 2024 at 08:52 PM UTC

### Review by @xbhouse - Commented on September 11, 2024 at 09:02 PM UTC

### Review by @jlsherrill - Approved on September 11, 2024 at 09:07 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/809*
