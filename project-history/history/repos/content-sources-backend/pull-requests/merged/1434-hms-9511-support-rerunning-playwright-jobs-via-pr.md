---
type: pull_request
number: 1434
title: "HMS-9511: Support rerunning playwright jobs via PR comment"
state: merged
author: swadeley
created: 2026-03-30T15:41:42Z
updated: 2026-04-07T13:08:23Z
closed: 2026-04-07T13:08:23Z
merged: 2026-04-07T13:08:23Z
base_branch: main
head_branch: swadeley/HMS-9511
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1434
---

# Pull Request #1434: HMS-9511: Support rerunning playwright jobs via PR comment

**Author**: @swadeley
**Created**: March 30, 2026 at 03:41 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/HMS-9511`

## Description

## Summary

Add command for PW API & UI tests rerun
Support '/retest-playwright' as PR comment to rerun playwright actions

## Testing steps

'/retest-playwright' as a PR comment reruns playwright actions (API & UI tests).

---

## Discussion

### Comment by @xbhouse on March 30, 2026 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-9511

### Comment by @swadeley on March 31, 2026 at 10:45 AM UTC

/retest-ui

EDIT: This will only work once merged.

### Comment by @rverdile on April 03, 2026 at 01:30 PM UTC

Would we want to include the API tests too? Maybe use the command "/retest-playwright" instead

### Comment by @xbhouse on April 03, 2026 at 01:39 PM UTC

> Would we want to include the API tests too? Maybe use the command "/retest-playwright" instead

oh yes! i agree :) 

### Comment by @swadeley on April 07, 2026 at 11:54 AM UTC

@xbhouse and @rverdile 
Changes made, thank you.
I'll update frontend to match.

### Comment by @xbhouse on April 07, 2026 at 12:42 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Dismissed on April 03, 2026 at 01:00 PM UTC

### Review by @xbhouse - Approved on April 07, 2026 at 12:42 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1434*
