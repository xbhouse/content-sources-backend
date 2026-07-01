---
type: pull_request
number: 869
title: "Fixes 4610: don't trigger snapshot on adding slash to url"
state: merged
author: xbhouse
created: 2024-10-31T13:15:57Z
updated: 2024-11-04T12:30:33Z
closed: 2024-11-04T12:23:23Z
merged: 2024-11-04T12:23:23Z
base_branch: main
head_branch: 4610
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/869
---

# Pull Request #869: Fixes 4610: don't trigger snapshot on adding slash to url

**Author**: @xbhouse
**Created**: October 31, 2024 at 01:15 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4610`

## Description

## Summary

Adjusts comparison to check formatted url to keep from triggering snapshots if only a slash was added or removed.

## Testing steps

1. Create a repository and let it snapshot
2. Update the repository URL by adding an extra slash or removing a slash. Snapshots should not be triggered with either PUT or PATCH

---

## Discussion

### Comment by @jlsherrill on October 31, 2024 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-4610

### Comment by @swadeley on November 04, 2024 at 10:06 AM UTC

/retest

### Comment by @swadeley on November 04, 2024 at 12:23 PM UTC

Hi

works as described.

Thank you

---

## Reviews

### Review by @jlsherrill - Approved on November 01, 2024 at 04:30 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/869*
