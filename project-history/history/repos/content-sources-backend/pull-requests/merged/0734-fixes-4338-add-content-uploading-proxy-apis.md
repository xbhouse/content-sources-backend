---
type: pull_request
number: 734
title: "Fixes 4338: add content uploading proxy apis"
state: merged
author: xbhouse
created: 2024-07-08T20:34:27Z
updated: 2024-07-15T16:33:52Z
closed: 2024-07-15T16:33:51Z
merged: 2024-07-15T16:33:51Z
base_branch: main
head_branch: 4338
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/734
---

# Pull Request #734: Fixes 4338: add content uploading proxy apis

**Author**: @xbhouse
**Created**: July 08, 2024 at 08:34 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4338`

## Description

## Summary

Adds the following proxy APIs for uploading content to pulp:
- Create an upload
- Upload a file chunk
- Finish an upload
- Fetch a task

## Testing steps

There are a lot of steps to test this flow. 

- Pulp's API docs for uploads can be found [here](https://pulpproject.org/pulpcore/restapi/#tag/Uploads)
- Usage docs can be found [here](https://github.com/pulp/pulpcore/blob/b2e4e22ffff3b2fc72c8eaf566e0408a5c60d947/docs/workflows/upload-publish.rst#chunked-uploads)
- You can also use this [script](https://github.com/xbhouse/testing-uploads/blob/main/test_upload.py) which may be helpful

Otherwise you can test with the following steps:

1. Download and extract the [LibreOffice rpms](https://www.libreoffice.org/download/download-libreoffice/?type=rpm-x86_64&version=7.6.6&lang=en-US) to test with. After extracting, the rpms will be in the RPMS directory.

2. Get the size in bytes of the rpm you plan to test with (there are lots of ways to do this):
`ls -l <rpm_file> | awk '{print $5}'`

3. Split one of the rpms into chunks (this is splitting into 6MB chunks and each chunk will be named like test_aa, test_ab, etc):
`split -b 6M <rpm_file> test`

4. Generate the checksum for the rpm:
Mac: `shasum -a 256 <rpm_file>`
Linux: `sha256sum <rpm_file>`

5. Optionally generate checksums for each rpm chunk. The chunk's checksum is an optional parameter for the PUT request.

6. Create the upload:
```
curl -X POST localhost:8000/api/content-sources/v1.0/pulp/uploads/ \
  -H 'x-rh-identity: <identity_header>' \
  -H 'Content-Type: application/json' \
  -d '{"size": <rpm_size>}'
```

Response should look like this:
```
{
    "pulp_created": "2024-07-01T14:04:44.220526Z",
    "pulp_href": "/api/pulp/dd51c811/api/v3/uploads/01906e9c-bd3b-74b2-a7b3-e122e9ca88d5/",
    "pulp_last_updated": "2024-07-01T14:04:44.220541Z",
    "size": 7056
}
```
The `pulp_href` is the `upload_href` you will use in the next 2 requests.

7. Upload a chunk:
```
Content-Range header values:

start = chunk number * chunk size in bytes (chunk number starts at 0)
end = start + chunk size in bytes - 1

curl -X PUT localhost:8000/api/content-sources/v1.0/pulp/uploads/<upload_href> \
  -H 'x-rh-identity: <identity_header>' \
  -H 'Content-Range: bytes <start>-<end>/*' \
  --form 'file=@"<chunk_file>"' \
  --form 'sha256="<rpm_chunk_checksum>"'
```

Response should look like this:
```
{
    "pulp_created": "2024-07-01T14:44:58.435098Z",
    "pulp_href": "/api/pulp/dd51c811/api/v3/uploads/01906ec1-93c2-7af5-99b5-12d880a1e485/",
    "pulp_last_updated": "2024-07-01T14:44:58.435113Z",
    "size": 7056
}
```

8. Finish the upload:
```
curl -X POST localhost:8000/api/content-sources/v1.0/pulp/uploads/<upload_href> \
  -H 'x-rh-identity: <identity_header>' \
  -H 'Content-Type: application/json' \
  -d '{"sha256": "<rpm_checksum>"}'
```

Response returns the task href that is queued to create the artifact and should look like this:
```
{
    "task": "/api/pulp/dd51c811/api/v3/tasks/01906eb8-d423-7323-9123-03a775bb3500/"
}
```

9. Get the artifact href from the task:
```
curl -X GET localhost:8000/api/content-sources/v1.0/pulp/tasks/<task_href> \
  -H 'x-rh-identity: <identity_header>' \
  -H 'Content-Type: application/json'
```

Response should look like this, the artifact_href is under created_resources:
```
{
    "child_tasks": [],
    "created_by": "/api/pulp/dd51c811/api/v3/users/1/",
    "created_resources": [
        "/api/pulp/dd51c811/api/v3/artifacts/01906ea1-4343-7a12-a71a-180bd6faa122/"
    ],
    "finished_at": "2024-07-01T14:09:40.68993Z",
    "logging_cid": "9ad506d1-9f18-44f0-abe4-615056e51817",
    "name": "pulpcore.app.tasks.upload.commit",
    "progress_reports": [],
    "pulp_created": "2024-07-01T14:09:40.609565Z",
    "pulp_href": "/api/pulp/dd51c811/api/v3/tasks/01906ea1-4301-7315-8cc9-f517f5fa79d7/",
    "pulp_last_updated": "2024-07-01T14:09:40.609584Z",
    "reserved_resources_record": [
        "/api/pulp/dd51c811/api/v3/uploads/01906e9c-bd3b-74b2-a7b3-e122e9ca88d5/",
        "prn:core.upload:01906e9c-bd3b-74b2-a7b3-e122e9ca88d5",
        "shared:prn:core.domain:019027bb-562c-7ab6-8636-50d1e0134ea8",
        "shared:/api/pulp/default/api/v3/domains/019027bb-562c-7ab6-8636-50d1e0134ea8/"
    ],
    "started_at": "2024-07-01T14:09:40.657947Z",
    "state": "completed",
    "unblocked_at": "2024-07-01T14:09:40.619289Z",
    "worker": "/api/pulp/default/api/v3/workers/01902733-d5d1-7de0-a98e-df71415e451b/"
}
```



## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 08, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-4338

---

## Reviews

### Review by @jlsherrill - Commented on July 12, 2024 at 05:50 PM UTC

### Review by @xbhouse - Commented on July 12, 2024 at 05:52 PM UTC

### Review by @jlsherrill - Approved on July 12, 2024 at 09:15 PM UTC

Tested with different chunk sizes and it all worked really well.

As mentioned, i  might would check in the python upload script, it may be useful at least in the short term going forward (and we may want to enhance it as we add other apis).

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/734*
