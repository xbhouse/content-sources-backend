---
type: pull_request
number: 1363
title: "HMS-9958: add minor versions to repository parameters"
state: merged
author: rverdile
created: 2026-01-16T21:06:15Z
updated: 2026-02-05T14:34:01Z
closed: 2026-02-05T14:33:58Z
merged: 2026-02-05T14:33:58Z
base_branch: main
head_branch: repo-params-new
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1363
---

# Pull Request #1363: HMS-9958: add minor versions to repository parameters

**Author**: @rverdile
**Created**: January 16, 2026 at 09:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `repo-params-new`

## Description

## Summary
Updates the repository parameters to include minor repository versions and extended release features. Filters them out if they are not entitled in the feature service API

## Testing steps
1. Set the "RHEL-EUS-x86_64" and "RHEL-E4S-x86_64" features in the feature_filter option in the config.yaml
2. call the `/repository_parameters/` API
3. All the minor version and extended release feature names should be listed
4. If you remove the features from the feature filter, they should no longer be listed



---

## Discussion

### Comment by @xbhouse on January 16, 2026 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-9958

---

## Reviews

### Review by @dominikvagner - Commented on January 21, 2026 at 10:04 AM UTC

### Review by @rverdile - Commented on January 21, 2026 at 03:20 PM UTC

### Review by @katarinazaprazna - Commented on January 21, 2026 at 03:46 PM UTC

### Review by @rverdile - Commented on January 22, 2026 at 09:19 PM UTC

### Review by @dominikvagner - Approved on February 05, 2026 at 01:29 PM UTC

works great! looks nice! 👏🏼 
thanks! 🚀 ack! ✅ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1363*
