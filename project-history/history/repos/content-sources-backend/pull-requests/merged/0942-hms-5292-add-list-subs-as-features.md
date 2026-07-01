---
type: pull_request
number: 942
title: "HMS-5292: Add list subs-as-features"
state: merged
author: rverdile
created: 2025-01-16T14:15:21Z
updated: 2025-01-23T21:33:54Z
closed: 2025-01-23T21:33:50Z
merged: 2025-01-23T21:33:50Z
base_branch: main
head_branch: subs-as-feat
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/942
---

# Pull Request #942: HMS-5292: Add list subs-as-features

**Author**: @rverdile
**Created**: January 16, 2025 at 02:15 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `subs-as-feat`

## Description

## Summary
Adds `GET  /admin/features`, which lists the features from the subs-as-features API.

Returns `{ features: ["feature1", "feature2", ...] }`

## Testing steps
1. Grab the certs you need
2. Make a GET request to `/admin/features` and you should see the features listed


---

## Discussion

### Comment by @jlsherrill on January 16, 2025 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-5292

---

## Reviews

### Review by @rverdile - Commented on January 16, 2025 at 02:21 PM UTC

### Review by @xbhouse - Commented on January 20, 2025 at 09:32 PM UTC

### Review by @xbhouse - Commented on January 20, 2025 at 09:32 PM UTC

### Review by @xbhouse - Commented on January 20, 2025 at 09:38 PM UTC

### Review by @rverdile - Commented on January 21, 2025 at 07:12 PM UTC

### Review by @rverdile - Commented on January 21, 2025 at 07:12 PM UTC

### Review by @xbhouse - Commented on January 21, 2025 at 08:48 PM UTC

### Review by @xbhouse - Commented on January 21, 2025 at 09:06 PM UTC

### Review by @rverdile - Commented on January 22, 2025 at 06:29 PM UTC

### Review by @rverdile - Commented on January 22, 2025 at 06:29 PM UTC

### Review by @xbhouse - Commented on January 22, 2025 at 06:50 PM UTC

### Review by @rverdile - Commented on January 22, 2025 at 08:43 PM UTC

### Review by @xbhouse - Approved on January 22, 2025 at 09:24 PM UTC

nice job! lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/942*
