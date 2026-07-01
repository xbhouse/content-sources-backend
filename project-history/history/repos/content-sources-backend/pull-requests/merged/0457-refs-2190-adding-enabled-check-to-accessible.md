---
type: pull_request
number: 457
title: "Refs 2190: adding enabled check to accessible"
state: merged
author: xbhouse
created: 2023-11-02T18:17:42Z
updated: 2023-12-13T22:02:53Z
closed: 2023-11-02T23:49:15Z
merged: 2023-11-02T23:49:15Z
base_branch: main
head_branch: update_snapshot_accessible_check
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/457
---

# Pull Request #457: Refs 2190: adding enabled check to accessible

**Author**: @xbhouse
**Created**: November 02, 2023 at 06:17 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `update_snapshot_accessible_check`

## Description

## Summary

Adding check so that if a feature is not enabled, it is also not accessible

## Testing steps

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 02, 2023 at 06:23 PM UTC

and you'll want to adjust the title to  "Refs 2190"  refs = references, but doesn't fix it.  That way it'll be associated to the jira task, but won't close it. 

### Comment by @jlsherrill on November 02, 2023 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-2190

### Comment by @jlsherrill on November 02, 2023 at 06:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @swadeley on November 02, 2023 at 11:47 PM UTC

Hello

see test results in associated frontend PR here:
[ Fixes 2190: added snapshot info to content list table #155 ](https://github.com/content-services/content-sources-frontend/pull/155#issuecomment-1791712258)

---

## Reviews

### Review by @jlsherrill - Commented on November 02, 2023 at 06:20 PM UTC

### Review by @jlsherrill - Approved on November 02, 2023 at 06:27 PM UTC

ACK!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/457*
