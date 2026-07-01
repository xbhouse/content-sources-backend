---
type: pull_request
number: 1321
title: "HMS-9332: support for import upload repos"
state: merged
author: TenSt
created: 2025-12-01T12:15:08Z
updated: 2025-12-10T08:44:06Z
closed: 2025-12-10T08:44:06Z
merged: 2025-12-10T08:44:06Z
base_branch: main
head_branch: stepan/HMS-9332-import-export-upload-repos
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1321
---

# Pull Request #1321: HMS-9332: support for import upload repos

**Author**: @TenSt
**Created**: December 01, 2025 at 12:15 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `stepan/HMS-9332-import-export-upload-repos`

## Description

## Summary
This PR:
- adds support for import upload repos according to requirements
- refactors some common used code
- updates related tests

## Testing steps
1. Run the backend locally
2. Execute API request to import some new upload repo
```bash
curl -i --request POST -H "$( ./scripts/header.sh 9999 1111 )" \
  --url http://localhost:8000/api/content-sources/v1.0/repositories/bulk_import/ \
  --header 'Content-Type: application/json' \
  --data '[
  {
    "name": "sm-test",
    "url": "",
    "origin": "upload",
    "distribution_versions": [
      "any"
    ],
    "distribution_arch": "any",
    "gpg_key": "",
    "metadata_verification": false,
    "module_hotfixes": false,
    "snapshot": true
  }
]'
```
3. You should receive 201 and a warning: `"upload repository was exported from a different organization, make sure to add content to it manually"`
4. Execute the same API request again, you should receive 201 and a warning that about "name" field that it exists.

Note: export is already supported for upload, but in case you need it, here is the API call:
```bash
curl --request POST -H "$( ./scripts/header.sh 9999 1111 )" \
  --url http://localhost:8000/api/content-sources/v1.0/repositories/bulk_export/ \
  --header 'Content-Type: application/json' \
  --data '{
  "repository_uuids": [
    "fd06edc4-12bf-4019-8c4a-552e7ff7b335"
  ]
}' | jq
```

---

## Discussion

### Comment by @xbhouse on December 01, 2025 at 12:30 PM UTC

https://issues.redhat.com/browse/HMS-9332

---

## Reviews

### Review by @rverdile - Commented on December 03, 2025 at 07:58 PM UTC

looks good overall, two comments

### Review by @TenSt - Commented on December 04, 2025 at 11:27 AM UTC

### Review by @TenSt - Commented on December 04, 2025 at 12:54 PM UTC

### Review by @rverdile - Commented on December 04, 2025 at 08:34 PM UTC

### Review by @rverdile - Commented on December 04, 2025 at 09:34 PM UTC

### Review by @TenSt - Commented on December 05, 2025 at 12:51 PM UTC

### Review by @rverdile - Commented on December 05, 2025 at 03:05 PM UTC

### Review by @TenSt - Commented on December 09, 2025 at 01:23 PM UTC

### Review by @rverdile - Approved on December 09, 2025 at 08:36 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1321*
