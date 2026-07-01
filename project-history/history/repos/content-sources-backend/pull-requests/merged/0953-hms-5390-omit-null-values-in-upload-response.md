---
type: pull_request
number: 953
title: "HMS-5390: omit null values in upload response"
state: merged
author: xbhouse
created: 2025-01-24T21:08:03Z
updated: 2025-02-03T15:43:22Z
closed: 2025-02-03T15:43:22Z
merged: 2025-02-03T15:43:22Z
base_branch: main
head_branch: fix-uploads
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/953
---

# Pull Request #953: HMS-5390: omit null values in upload response

**Author**: @xbhouse
**Created**: January 24, 2025 at 09:08 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `fix-uploads`

## Description

## Summary

- Omits null values for `created`, `last_updated`, `artifact_href`, and `completed_checksums` from the upload response (used when creating an upload or uploading a chunk). These values are not always needed and null values here prevent the generated IQE client from calling these APIs
- Adds the upload uuid to the response when an artifact href is found (i'm not sure this should be null or omitted)

## Testing steps

1. Responses from creating an upload and uploading a chunk shouldn't have any null values
2. Uploading rpms in the UI should still work the same (existing uploads should be reused)
3. IQE: Go through the uploading flow using the IQE environment, let me know if there are any other API type errors etc and I'll make the changes here



---

## Discussion

### Comment by @jlsherrill on January 24, 2025 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-5390

### Comment by @mayurilahane on January 27, 2025 at 04:13 PM UTC

/retest

### Comment by @mayurilahane on January 27, 2025 at 08:06 PM UTC

/retest

### Comment by @mayurilahane on January 29, 2025 at 04:27 PM UTC

/retest

### Comment by @mayurilahane on January 29, 2025 at 05:06 PM UTC

null values in upload response is fixed and able to upload the chunk of rpm file 

### Comment by @mayurilahane on January 29, 2025 at 05:06 PM UTC

ACK

### Comment by @dominikvagner on January 31, 2025 at 08:08 AM UTC

/retest

---

## Reviews

### Review by @dominikvagner - Commented on January 30, 2025 at 09:50 AM UTC

one small comment, works and looks good! 👏🏼😄 

### Review by @xbhouse - Commented on January 30, 2025 at 01:52 PM UTC

### Review by @xbhouse - Commented on January 30, 2025 at 04:22 PM UTC

### Review by @dominikvagner - Commented on January 31, 2025 at 08:09 AM UTC

### Review by @dominikvagner - Approved on January 31, 2025 at 08:10 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/953*
