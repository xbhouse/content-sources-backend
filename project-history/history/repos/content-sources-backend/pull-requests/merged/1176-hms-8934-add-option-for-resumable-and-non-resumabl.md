---
type: pull_request
number: 1176
title: "HMS-8934: Add option for resumable and non-resumable uploads"
state: merged
author: Starle21
created: 2025-08-14T08:39:45Z
updated: 2025-08-21T14:30:08Z
closed: 2025-08-21T14:30:01Z
merged: 2025-08-21T14:30:01Z
base_branch: main
head_branch: starle/HMS-8934
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1176
---

# Pull Request #1176: HMS-8934: Add option for resumable and non-resumable uploads

**Author**: @Starle21
**Created**: August 14, 2025 at 08:39 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `starle/HMS-8934`

## Description

## Summary
In the /repositories/uploads/ api if there has been already a previous
    upload matching the sha256 and chunk_size created, it was used on
    subsequent uploads (for instance when the full upload process had been
    interrupted and was resumed later) and the existing UUID was returned.
    There was no option to resume with a fresh upload, returning a new UUID.
    This adds the new option to allow non-resumable uploads, passing
    resumable=false parameter in the post request body for the api. It makes it
    easier for users that use this api with their custom scripts, not having
    to write resumable logic. Passing resumable=true will behave as it did,
    returning the existing UUID.

## Testing steps
Use your favorite REST client to hit the endpoint `/repositories/uploads/`


related changes on the frontend: https://github.com/content-services/content-sources-frontend/pull/614


---

## Discussion

### Comment by @jlsherrill on August 14, 2025 at 09:00 AM UTC

https://issues.redhat.com/browse/HMS-8934

### Comment by @rverdile on August 20, 2025 at 07:27 PM UTC

/retest

### Comment by @swadeley on August 21, 2025 at 09:25 AM UTC

Hi, just one test in PR check `FAILED tests/test_repository_api_only.py::test_bulk_import`  and that is unrelated to this PR.

---

## Reviews

### Review by @rverdile - Commented on August 18, 2025 at 04:10 PM UTC

### Review by @Starle21 - Commented on August 19, 2025 at 01:25 PM UTC

### Review by @rverdile - Commented on August 19, 2025 at 08:21 PM UTC

### Review by @rverdile - Commented on August 19, 2025 at 08:21 PM UTC

### Review by @rverdile - Commented on August 19, 2025 at 08:30 PM UTC

### Review by @rverdile - Commented on August 19, 2025 at 08:31 PM UTC

### Review by @rverdile - Commented on August 19, 2025 at 08:31 PM UTC

### Review by @jlsherrill - Commented on August 19, 2025 at 08:39 PM UTC

### Review by @Starle21 - Commented on August 20, 2025 at 11:03 AM UTC

### Review by @Starle21 - Commented on August 20, 2025 at 11:05 AM UTC

### Review by @rverdile - Commented on August 20, 2025 at 07:45 PM UTC

Looks good! One more thought

To make the API behavior more predictable, I would add an `Order` clause to the `GetExistingUploadIDAndCompletedChunks` query. Otherwise, it orders by uuid. 

e.g.
```
SELECT * FROM "uploads" WHERE org_id = '17791560' AND chunk_size = 2476 AND sha256 = '9b3d22d05187810d8521d99ca2483232e7da80087691e5c1f8fa106a25118bdf' AND size = 2476 ORDER BY "uploads"."upload_uuid" LIMIT 1
```

For example, let's say I do the following (assuming the sha256 and chunk_size are the same). For the resumable upload, the API could return either the first or second upload.
1. Create an upload with resumable=false
2. Create an upload with resumable=false
3. Create an upload with resumable=true


And I would say let's get the latest upload ("created_at desc")

### Review by @rverdile - Approved on August 21, 2025 at 01:32 PM UTC

nice job!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1176*
