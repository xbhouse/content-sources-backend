---
type: pull_request
number: 1489
title: "HMS-10637: Include EUS in template filter by major version"
state: merged
author: rverdile
created: 2026-05-04T19:33:52Z
updated: 2026-05-05T18:14:59Z
closed: 2026-05-04T21:09:31Z
merged: 2026-05-04T21:09:31Z
base_branch: main
head_branch: fix-filter-by-major-version
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1489
---

# Pull Request #1489: HMS-10637: Include EUS in template filter by major version

**Author**: @rverdile
**Created**: May 04, 2026 at 07:33 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-filter-by-major-version`

## Description

## Summary

Template filtering on the major version only needs to include EUS templates to compatible with image builder. But this adds an option to exclude it so it is still compatible with our UI.

## Testing steps
1. Import and snapshot the extended release repos
2. Create templates of varying versions (9, 9.4, 9.6, 10, 10.0, etc)
3. Checkout the frontend PR as well: https://github.com/content-services/content-sources-frontend/pull/993
4. Use the version filtering in various combinations and make sure it works as you'd expect



---

## Discussion

### Comment by @xbhouse on May 04, 2026 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-10637

---

## Reviews

### Review by @xbhouse - Approved on May 04, 2026 at 08:25 PM UTC

looks good! tested with IB too, templates are showing up :) thank you!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1489*
