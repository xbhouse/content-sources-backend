---
type: pull_request
number: 1299
title: "HMS-8948: Fix module stream version parsing"
state: merged
author: rverdile
created: 2025-11-13T15:26:47Z
updated: 2025-12-01T18:39:26Z
closed: 2025-12-01T18:39:22Z
merged: 2025-12-01T18:39:22Z
base_branch: main
head_branch: perl
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1299
---

# Pull Request #1299: HMS-8948: Fix module stream version parsing

**Author**: @rverdile
**Created**: November 13, 2025 at 03:26 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `perl`

## Description

## Summary
A bug in the yummy dependency led to incorrect module stream version parsing. This bumps the yummy version to include the fix and fixes a bug that prevented the incorrect stream from being removed

You can pull in these changes locally and edit the go.mod to point to them:
https://github.com/content-services/yummy/pull/39

## Testing steps
1. Introspect RHEL 8 Appstream
2. Search module streams for perl. Without this PR, it would return perl "5.3" (and other versions). With this PR, "5.3" should be replaced by "5.30"
3. Search the RHEL 8 Appstream repository for the perl RPM using `/rpms/names/`. Include package sources.
4. The package sources should not include perl "5.3", only perl "5.30"


---

## Discussion

### Comment by @xbhouse on November 13, 2025 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-8948

---

## Reviews

### Review by @TenSt - Approved on November 27, 2025 at 01:12 PM UTC

LGTM! I've tested it locally with latest yammy version, all looks good

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1299*
