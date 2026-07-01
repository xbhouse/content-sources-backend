---
type: pull_request
number: 1390
title: "HMS-10156: override urls for extended release repos"
state: merged
author: rverdile
created: 2026-02-10T21:18:08Z
updated: 2026-02-16T15:22:49Z
closed: 2026-02-16T15:22:43Z
merged: 2026-02-16T15:22:42Z
base_branch: main
head_branch: override-urls
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1390
---

# Pull Request #1390: HMS-10156: override urls for extended release repos

**Author**: @rverdile
**Created**: February 10, 2026 at 09:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `override-urls`

## Description

## Summary
Adds baseurl content overrides for extended release repositories

## Testing steps
1. Create a template with an extended release repository
2. Add that template to a system
3. `dnf install httpd` will be possible without setting a release version for the system



---

## Discussion

### Comment by @xbhouse on February 10, 2026 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-10156

---

## Reviews

### Review by @TenSt - Approved on February 12, 2026 at 11:07 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1390*
