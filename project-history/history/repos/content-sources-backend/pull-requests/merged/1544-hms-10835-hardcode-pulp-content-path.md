---
type: pull_request
number: 1544
title: "HMS-10835: Hardcode pulp content path"
state: merged
author: rverdile
created: 2026-06-18T20:25:54Z
updated: 2026-06-23T14:09:39Z
closed: 2026-06-23T14:09:35Z
merged: 2026-06-23T14:09:35Z
base_branch: main
head_branch: hardcode-content-path
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1544
---

# Pull Request #1544: HMS-10835: Hardcode pulp content path

**Author**: @rverdile
**Created**: June 18, 2026 at 08:25 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `hardcode-content-path`

## Description

## Summary
The pulp content path should never change. This makes is configurable, but hardcodes it with the expected default.

## Testing steps
1. Integration tests are probably enough
2. You could also create a snapshot and then make sure the content URL is correct


---

## Discussion

### Comment by @xbhouse on June 18, 2026 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-10835

### Comment by @rverdile on June 22, 2026 at 04:07 PM UTC

thanks @swadeley, updated!

---

## Reviews

### Review by @swadeley - Commented on June 19, 2026 at 10:39 AM UTC

### Review by @swadeley - Commented on June 19, 2026 at 10:50 AM UTC

### Review by @swadeley - Approved on June 22, 2026 at 04:15 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1544*
