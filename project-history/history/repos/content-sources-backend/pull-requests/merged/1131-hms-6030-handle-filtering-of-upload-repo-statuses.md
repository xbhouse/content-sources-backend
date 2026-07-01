---
type: pull_request
number: 1131
title: "HMS-6030: handle filtering of upload repo statuses"
state: merged
author: xbhouse
created: 2025-06-23T19:10:31Z
updated: 2025-07-02T12:13:08Z
closed: 2025-07-02T12:13:08Z
merged: 2025-07-02T12:13:08Z
base_branch: main
head_branch: 6030
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1131
---

# Pull Request #1131: HMS-6030: handle filtering of upload repo statuses

**Author**: @xbhouse
**Created**: June 23, 2025 at 07:10 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `6030`

## Description

## Summary

Fixes logic for status filtering to also handle upload repos

## Testing steps

1. Create an upload repo and upload an rpm (either via API or UI)
2. When listing the repositories with a filter on the status, the upload repo should be included in the response or displayed in the UI

---

## Discussion

### Comment by @jlsherrill on June 23, 2025 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-6030

### Comment by @swadeley on July 02, 2025 at 12:13 PM UTC

Hi

works as described:

![image](https://github.com/user-attachments/assets/b71b9680-0142-4c52-9340-acbfe65a7498)


---

## Reviews

### Review by @rverdile - Approved on June 27, 2025 at 08:25 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1131*
