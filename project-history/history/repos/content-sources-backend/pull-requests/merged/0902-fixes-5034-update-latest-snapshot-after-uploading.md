---
type: pull_request
number: 902
title: "Fixes 5034: update latest snapshot after uploading content"
state: merged
author: xbhouse
created: 2024-11-21T16:35:50Z
updated: 2024-11-22T13:00:29Z
closed: 2024-11-22T12:50:05Z
merged: 2024-11-22T12:50:05Z
base_branch: main
head_branch: 5034
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/902
---

# Pull Request #902: Fixes 5034: update latest snapshot after uploading content

**Author**: @xbhouse
**Created**: November 21, 2024 at 04:35 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `5034`

## Description

## Summary

Triggers the update-latest-snapshot task after uploading content to a repo. Without this, use_latest templates with upload repos wouldn't get updated with the latest content.

## Testing steps

(It's probably easiest to do most of this in the UI)

1. Add a RH repo and an upload repo with 1 rpm, let them snapshot
2. Create a template using those repos and set it to use latest content
3. Upload another rpm to the upload repo
4. Check the template's packages via the UI or the API (`/templates/:uuid/rpms?search=<uploaded_rpm_name>`). They should include all the uploaded packages


---

## Discussion

### Comment by @jlsherrill on November 21, 2024 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-5034

### Comment by @swadeley on November 22, 2024 at 12:49 PM UTC

Hi, looks good:

![image](https://github.com/user-attachments/assets/b5d136fb-79d0-45ec-b31c-73bc6a0927cc)

![image](https://github.com/user-attachments/assets/408be0d1-b9a5-4241-bad3-d742c1c00506)



![image](https://github.com/user-attachments/assets/4e98b27b-f553-4789-8b49-fa5b1d9571db)


---

## Reviews

### Review by @rverdile - Commented on November 21, 2024 at 06:06 PM UTC

### Review by @xbhouse - Commented on November 21, 2024 at 06:38 PM UTC

### Review by @rverdile - Commented on November 21, 2024 at 07:22 PM UTC

### Review by @rverdile - Commented on November 21, 2024 at 07:22 PM UTC

### Review by @xbhouse - Commented on November 21, 2024 at 07:34 PM UTC

### Review by @rverdile - Approved on November 21, 2024 at 07:35 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/902*
