---
type: pull_request
number: 1437
title: "HMS-10379: filter template list by stream"
state: merged
author: katarinazaprazna
created: 2026-03-31T23:10:48Z
updated: 2026-04-08T14:36:16Z
closed: 2026-04-08T14:36:16Z
merged: 2026-04-08T14:36:16Z
base_branch: main
head_branch: filter-template-list-by-stream
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1437
---

# Pull Request #1437: HMS-10379: filter template list by stream

**Author**: @katarinazaprazna
**Created**: March 31, 2026 at 11:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `filter-template-list-by-stream`

## Description

## Summary

- Adds multi-value filter for `extended_release`. Supports filtering by "none" to retrieve standard (non-EUS) templates. When "none" is present in the stream list (e.g., "eus,none"), the query returns both templates matching the specified streams (e.g., eus) and standard templates

- Adds multi-value filter for `extended_release_version`
- Adds openapi documentation for both the `extended_release` and `extended_release_version` filters
- Adds handler and dao tests for both extended release and extended release version filter params and list filtering


## Testing steps
All tests should pass

---

## Discussion

### Comment by @xbhouse on March 31, 2026 at 11:30 PM UTC

https://issues.redhat.com/browse/HMS-10379

---

## Reviews

### Review by @rverdile - Commented on April 01, 2026 at 06:44 PM UTC

### Review by @katarinazaprazna - Commented on April 01, 2026 at 08:53 PM UTC

### Review by @xbhouse - Commented on April 02, 2026 at 06:15 PM UTC

### Review by @rverdile - Commented on April 03, 2026 at 01:29 PM UTC

### Review by @katarinazaprazna - Commented on April 06, 2026 at 09:57 PM UTC

### Review by @katarinazaprazna - Commented on April 06, 2026 at 11:18 PM UTC

### Review by @xbhouse - Commented on April 07, 2026 at 02:00 PM UTC

### Review by @rverdile - Commented on April 07, 2026 at 02:34 PM UTC

### Review by @rverdile - Commented on April 07, 2026 at 02:37 PM UTC

### Review by @katarinazaprazna - Commented on April 07, 2026 at 03:03 PM UTC

### Review by @xbhouse - Commented on April 07, 2026 at 03:11 PM UTC

### Review by @xbhouse - Commented on April 07, 2026 at 03:15 PM UTC

### Review by @katarinazaprazna - Commented on April 07, 2026 at 03:24 PM UTC

### Review by @rverdile - Commented on April 07, 2026 at 05:09 PM UTC

### Review by @rverdile - Approved on April 08, 2026 at 02:10 PM UTC

nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1437*
