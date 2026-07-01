---
type: pull_request
number: 1462
title: "HMS-10537: add none filter for extended release version"
state: merged
author: rverdile
created: 2026-04-14T15:05:32Z
updated: 2026-04-15T17:02:09Z
closed: 2026-04-15T17:02:06Z
merged: 2026-04-15T17:02:06Z
base_branch: main
head_branch: none-filter
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1462
---

# Pull Request #1462: HMS-10537: add none filter for extended release version

**Author**: @rverdile
**Created**: April 14, 2026 at 03:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `none-filter`

## Description

## Summary
Adds a "none" filter to extended release version so it is possible to filter by major version only

## Testing steps
1. Create a template with an extended release version and import an extended support repository
2. Filtering by `/templates?version=9&extended_release_version=none`. Should return only templates with major version 9 that are not extended support templates.
3. Filtering by `/repositories?distribution_version=9&extended_release_version=none` should return only repositories with major version 9 that are not extended release repositories.



---

## Discussion

### Comment by @xbhouse on April 14, 2026 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-10537

### Comment by @xbhouse on April 14, 2026 at 03:44 PM UTC

/retest

### Comment by @rverdile on April 14, 2026 at 04:34 PM UTC

/retest

### Comment by @rverdile on April 14, 2026 at 05:03 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on April 14, 2026 at 03:34 PM UTC

thank you! tested and looks good :) ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1462*
