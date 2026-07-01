---
type: pull_request
number: 1128
title: "Build: bump yummy"
state: merged
author: xbhouse
created: 2025-06-20T14:17:47Z
updated: 2026-05-20T18:37:36Z
closed: 2025-06-20T15:52:54Z
merged: 2025-06-20T15:52:54Z
base_branch: main
head_branch: bump-yummy
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1128
---

# Pull Request #1128: Build: bump yummy

**Author**: @xbhouse
**Created**: June 20, 2025 at 02:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `bump-yummy`

## Description

## Summary

Bumps yummy to the latest version where the DefaultMaxXmlSize was increased to accommodate the increase in size of the RHEL9 arm baseos repo

## Testing steps

Introspecting the RHEL9 arm baseos repo should succeed: `./release/external-repos introspect --url https://cdn.redhat.com/content/dist/rhel9/9/aarch64/baseos/os/`


---

## Discussion

### Comment by @xbhouse on June 20, 2025 at 03:28 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on June 20, 2025 at 02:50 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1128*
