---
type: pull_request
number: 1449
title: "HMS-5855: handle null distribution_hrefs in trc table"
state: merged
author: xbhouse
created: 2026-04-06T14:01:01Z
updated: 2026-04-06T15:58:13Z
closed: 2026-04-06T15:37:21Z
merged: 2026-04-06T15:37:21Z
base_branch: main
head_branch: fix-for-null-disthrefs
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1449
---

# Pull Request #1449: HMS-5855: handle null distribution_hrefs in trc table

**Author**: @xbhouse
**Created**: April 06, 2026 at 02:01 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-for-null-disthrefs`

## Description

## Summary

Updates `GetDistributionHref` to return a string pointer. We're seeing a null `distribution_href` in the templates_repository_configurations table in stage, possibly from a failed snapshot cleanup, which is preventing cleanup of the custom EPEL repos due to this error: `sql: Scan error on column index 0, name "distribution_href": converting NULL to string is unsupported `

## Testing steps

Tests pass

---

## Discussion

### Comment by @xbhouse on April 06, 2026 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-5855

### Comment by @xbhouse on April 06, 2026 at 02:38 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on April 06, 2026 at 03:17 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1449*
