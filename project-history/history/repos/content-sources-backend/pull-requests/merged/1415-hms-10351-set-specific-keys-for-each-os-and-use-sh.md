---
type: pull_request
number: 1415
title: "HMS-10351: set specific keys for each os and use sha256 signed"
state: merged
author: jlsherrill
created: 2026-03-18T19:53:31Z
updated: 2026-03-19T15:50:54Z
closed: 2026-03-19T15:50:54Z
merged: 2026-03-19T15:50:54Z
base_branch: main
head_branch: HMS-10351_2
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1415
---

# Pull Request #1415: HMS-10351: set specific keys for each os and use sha256 signed

**Author**: @jlsherrill
**Created**: March 18, 2026 at 07:53 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `HMS-10351_2`

## Description

## Summary

there was a bit more required than previous thought, here i split up the gpg key files into all 3 os versions, even though 8 and 9 are the same.  This removes previous keys and adds keys that image builder is using to build 8,9, & 10 on RHEL 10.  

## Testing steps



---

## Discussion

### Comment by @xbhouse on March 18, 2026 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-10351

---

## Reviews

### Review by @croissanne - Commented on March 19, 2026 at 09:18 AM UTC

### Review by @croissanne - Commented on March 19, 2026 at 09:19 AM UTC

### Review by @croissanne - Commented on March 19, 2026 at 09:25 AM UTC

### Review by @croissanne - Commented on March 19, 2026 at 09:25 AM UTC

This will work, some small comments but i don't think they're required.

### Review by @croissanne - Approved on March 19, 2026 at 03:17 PM UTC

perf! thank you 💯 

### Review by @TenSt - Approved on March 19, 2026 at 03:21 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1415*
