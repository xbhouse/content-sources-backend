---
type: pull_request
number: 1102
title: "HMS-5790: cannot change size of upload"
state: merged
author: xbhouse
created: 2025-05-02T15:45:48Z
updated: 2025-05-13T13:41:46Z
closed: 2025-05-13T13:41:46Z
merged: 2025-05-13T13:41:46Z
base_branch: main
head_branch: 5790
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1102
---

# Pull Request #1102: HMS-5790: cannot change size of upload

**Author**: @xbhouse
**Created**: May 02, 2025 at 03:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5790`

## Description

## Summary

Adds the size of the upload to the uploads table and takes the size into account when creating an upload

## Testing steps

1. Create an upload via the API and specify a size that's smaller than the chunk size: 
```
curl --location 'localhost:8000/api/content-sources/v1.0/repositories/uploads/' \
--header 'x-rh-identity: <identity>' \
--header 'Content-Type: application/json' \
--data '{
    "size": 7050,
    "chunk_size": 7068,
    "sha256": "08d4c0799e06278e5c6aba9b2b979c55802c5ec6083906220d712bf0031a8450"
}'

{
  "upload_uuid":"01969199-f854-712c-a1a4-6497c73a6d03",
  "created":"2025-05-02T15:25:09.077058Z",
  "last_updated":"2025-05-02T15:25:09.077071Z",
  "size":7050
}
```

You should see an upload created in the uploads table in the content sources db and when listing uploads (`pulp upload list`) via the pulp CLI.

2. Recreate the upload with the same sha and chunk size, but with the correct size of the rpm. You should see a new upload created with a different UUID than the first one.

3. Uploading a chunk to the original upload (the one with the wrong size) should fail, this is expected: 
```
curl --location 'localhost:8000/api/content-sources/v1.0/repositories/uploads/01969199-f854-712c-a1a4-6497c73a6d03/upload_chunk/' \
--header 'x-rh-identity: <identity>' \
--header 'Content-Range: bytes 0-7067/*' \ 
--form 'file=@"/Users/bhouse/rpms/testfileaa"' \      
--form 'sha256="08d4c0799e06278e5c6aba9b2b979c55802c5ec6083906220d712bf0031a8450"'  

{"errors":[{"status":400,"title":"error uploading chunk","detail":"uploading file chunk: 400 Bad Request: {\"non_field_errors\":[\"End byte is greater than upload size.\"]}"}]}
```

Uploading a chunk to the new upload (the one with the correct size) should succeed.

4. Uploading rpms in the UI and upload cleanup (`go run cmd/external-repos/main.go upload-cleanup`) should still work as before.


---

## Discussion

### Comment by @jlsherrill on May 02, 2025 at 05:02 PM UTC

https://issues.redhat.com/browse/HMS-5790

---

## Reviews

### Review by @rverdile - Approved on May 06, 2025 at 07:40 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1102*
