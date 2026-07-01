---
type: pull_request
number: 1161
title: "HMS-8936: add userfacing upload script"
state: merged
author: jlsherrill
created: 2025-07-29T20:49:59Z
updated: 2025-08-06T07:23:31Z
closed: 2025-08-06T07:23:31Z
merged: 2025-08-06T07:23:31Z
base_branch: main
head_branch: upload_scripts
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1161
---

# Pull Request #1161: HMS-8936: add userfacing upload script

**Author**: @jlsherrill
**Created**: July 29, 2025 at 08:49 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `upload_scripts`

## Description

## Summary

Userfacing python script for uploading rpms, supports chunk uploads

## Testing steps

Create an upload repository, grab the UUID, upload an rpm:

```
USERNAME=<USERNAME> PASSWORD='<PASSWORD>' python docs/upload-rpm.py  f287a13b-7985-4282-b53e-f9e9b8008b33  ~/rpms/google-chrome-unstable-128.0.6601.2-1.x86_64.rpm
```

---

## Discussion

### Comment by @jlsherrill on July 29, 2025 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-8936

### Comment by @swadeley on July 30, 2025 at 09:48 AM UTC

/retest

### Comment by @jlsherrill on August 01, 2025 at 04:34 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Commented on July 31, 2025 at 07:39 PM UTC

### Review by @xbhouse - Commented on July 31, 2025 at 07:39 PM UTC

### Review by @xbhouse - Commented on July 31, 2025 at 07:40 PM UTC

### Review by @xbhouse - Commented on July 31, 2025 at 07:41 PM UTC

tiny nitpicks just on the usage comments, but this looks and works great! :tada: 

### Review by @jlsherrill - Commented on August 01, 2025 at 01:21 PM UTC

### Review by @xbhouse - Approved on August 04, 2025 at 04:54 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1161*
