---
type: pull_request
number: 1041
title: "HMS-5570: delete upload entries on failed upload > artifact conversion"
state: merged
author: xbhouse
created: 2025-03-23T20:31:50Z
updated: 2025-03-31T14:10:57Z
closed: 2025-03-31T14:10:57Z
merged: 2025-03-31T14:10:57Z
base_branch: main
head_branch: 5570
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1041
---

# Pull Request #1041: HMS-5570: delete upload entries on failed upload > artifact conversion

**Author**: @xbhouse
**Created**: March 23, 2025 at 08:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5570`

## Description

## Summary

- Adds a DeleteUpload method
- If an upload fails to be converted to an artifact, the upload record is deleted in our db so the user can try to upload the rpm again
- These [docs](https://docs.google.com/document/d/1eaz3lXwljdjehmu-n9f7kpJjYMRkKzLq4axaC5Lu9A0/edit?tab=t.0) may be helpful
- Bug report in pulp opened [here](https://github.com/pulp/pulpcore/issues/6381)

## Testing steps

1. Checkout the main branch
2. Create an upload, upload the same chunk twice, and add the upload. Here is an example using a small rpm with just one chunk:

```
Create the upload:

curl -X POST --location 'localhost:8000/api/content-sources/v1.0/repositories/uploads/' \
-H 'x-rh-identity: <identity>' \
-H 'Content-Type: application/json' \
-d '{
       "size": 7068,
       "chunk_size": 7068,
       "sha256": "08d4c0799e06278e5c6aba9b2b979c55802c5ec6083906220d712bf0031a8450"
}'

Upload the same chunk twice:

curl -X POST --location 'localhost:8000/api/content-sources/v1.0/repositories/uploads/<upload_uuid>/upload_chunk/' \
-H 'x-rh-identity: <identity>' \
-H 'Content-Range: bytes 0-7067/*' \
--form 'file=@"<chunk_filename>"' \
--form 'sha256="08d4c0799e06278e5c6aba9b2b979c55802c5ec6083906220d712bf0031a8450"'

curl -X POST --location 'localhost:8000/api/content-sources/v1.0/repositories/uploads/<upload_uuid>/upload_chunk/' \
-H 'x-rh-identity: <identity>' \
-H 'Content-Range: bytes 0-7067/*' \
--form 'file=@"<chunk_filename>"' \
--form 'sha256="08d4c0799e06278e5c6aba9b2b979c55802c5ec6083906220d712bf0031a8450"'

Add the upload:

curl -X POST --location 'localhost:8000/api/content-sources/v1.0/repositories/<repository_uuid>/add_uploads/' \
-H 'x-rh-identity: <identity>' \
-H 'Content-Type: application/json' \
-d '{
    "uploads": [
        {
            "uuid": "<upload_uuid>",
            "sha256": "08d4c0799e06278e5c6aba9b2b979c55802c5ec6083906220d712bf0031a8450"
        }
    ],
    "artifacts": []
}'
``` 
    
3. The add_uploads call should fail with a checksum mismatch error and you shouldn't be able to create another upload with the same rpm
4. Checkout this PR, and call add_uploads again. It should still fail but you should be able to create another upload with the same rpm and start over. Uploading the chunk just once should succeed

---

## Discussion

### Comment by @jlsherrill on March 23, 2025 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-5570

---

## Reviews

### Review by @rverdile - Approved on March 26, 2025 at 04:22 PM UTC

nice job!!

### Review by @copilot-pull-request-reviewer - Commented on March 28, 2025 at 08:28 PM UTC

## Pull Request Overview

This PR introduces a mechanism to delete an upload record if its conversion to an artifact fails, ensuring that users can try uploading an RPM again.  
- Added a DeleteUpload method that is called upon conversion failure.  
- Updated the DAO interfaces and implemented the DeleteUpload function.  
- Added tests to verify the deletion of uploads.

### Reviewed Changes

Copilot reviewed 4 out of 4 changed files in this pull request and generated no comments.

| File                         | Description                                                                  |
| ---------------------------- | ---------------------------------------------------------------------------- |
| pkg/tasks/add_uploads.go     | Deletes the upload record on poll task failure to allow retrying uploads.    |
| pkg/dao/uploads.go           | Implements the DeleteUpload method for removing an upload record by UUID.     |
| pkg/dao/interfaces.go        | Adds the DeleteUpload method signature to the DAO interface.               |
| pkg/dao/uploads_test.go       | Introduces tests for verifying that the upload record is correctly deleted.  |


<details>
<summary>Comments suppressed due to low confidence (1)</summary>

**pkg/dao/uploads_test.go:74**
* [nitpick] Consider using errors.Is(err, gorm.ErrRecordNotFound) for checking the 'record not found' error instead of comparing the error message string.
```
First(&found, "upload_uuid = ?", uploadUUID).Error
```
</details>



---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1041*
