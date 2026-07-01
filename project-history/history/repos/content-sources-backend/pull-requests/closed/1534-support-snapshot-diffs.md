---
type: pull_request
number: 1534
title: "Support snapshot diffs"
state: closed
author: katarinazaprazna
created: 2026-06-10T14:14:51Z
updated: 2026-06-23T20:07:14Z
closed: 2026-06-23T20:07:14Z
base_branch: main
head_branch: support-snapshot-diffs
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1534
---

# Pull Request #1534: Support snapshot diffs

**Author**: @katarinazaprazna
**Created**: June 10, 2026 at 02:14 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `support-snapshot-diffs`

## Description

## Summary

- Adds a new endpoint that returns RPM packages added and removed in a snapshot (diff), merged into a single alphabetically-sorted paginated list
- Each entry is tagged with "added" or "removed" status; same-name packages are grouped together (removed before added) so the frontend can render update pairs side-by-side


## Testing steps



---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1534*
