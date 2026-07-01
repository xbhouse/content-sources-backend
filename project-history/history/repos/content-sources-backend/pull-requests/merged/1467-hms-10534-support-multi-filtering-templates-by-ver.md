---
type: pull_request
number: 1467
title: "HMS-10534: support multi-filtering templates by version"
state: merged
author: xbhouse
created: 2026-04-15T21:56:02Z
updated: 2026-04-16T16:04:19Z
closed: 2026-04-16T16:04:19Z
merged: 2026-04-16T16:04:19Z
base_branch: main
head_branch: 10534
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1467
---

# Pull Request #1467: HMS-10534: support multi-filtering templates by version

**Author**: @xbhouse
**Created**: April 15, 2026 at 09:56 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `10534`

## Description

## Summary

Adds support for filtering templates by multiple major versions

## Testing steps

1. Create templates with different versions
2. Filtering by multiple versions should work (e.g. `GET /templates/?version=8,9&extended_release_version=none`)

---

## Discussion

### Comment by @xbhouse on April 15, 2026 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-10534

---

## Reviews

### Review by @TenSt - Approved on April 16, 2026 at 01:58 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1467*
