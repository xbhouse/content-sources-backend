---
type: pull_request
number: 949
title: "HMS-5380: create_upload fixes"
state: merged
author: xbhouse
created: 2025-01-22T16:38:16Z
updated: 2025-01-24T21:01:11Z
closed: 2025-01-22T19:19:51Z
merged: 2025-01-22T19:19:51Z
base_branch: main
head_branch: fix-uploads
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/949
---

# Pull Request #949: HMS-5380: create_upload fixes

**Author**: @xbhouse
**Created**: January 22, 2025 at 04:38 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `fix-uploads`

## Description

## Summary

Fixes some issues when creating an upload:
- changes artifact href to be a string instead of null if empty (this is blocking IQE from calling this api)
- rejects request if chunk_size is 0 

## Testing steps

1. Artifact href should be set to an empty string if it doesn't exist
2. Request should be rejected if chunk size is 0 or less



---

## Discussion

### Comment by @jlsherrill on January 22, 2025 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-5380

### Comment by @mayurilahane on January 22, 2025 at 05:14 PM UTC

not facing the issue now and was able to create the upload 

![image](https://github.com/user-attachments/assets/a6b53379-7a38-4cdd-b1e2-3915450f0a9a)


### Comment by @mayurilahane on January 22, 2025 at 05:33 PM UTC

When chunk size is "0" it is giving error as expected 

![image](https://github.com/user-attachments/assets/7eefccf5-cc26-403c-a473-eadea55f483e)


### Comment by @mayurilahane on January 22, 2025 at 05:33 PM UTC

ACK

---

## Reviews

### Review by @Andrewgdewar - Approved on January 22, 2025 at 05:47 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/949*
