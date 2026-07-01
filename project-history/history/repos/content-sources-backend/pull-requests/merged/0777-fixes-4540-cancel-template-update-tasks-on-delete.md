---
type: pull_request
number: 777
title: "Fixes 4540: cancel template update tasks on delete"
state: merged
author: jlsherrill
created: 2024-08-15T19:58:59Z
updated: 2024-08-28T18:49:52Z
closed: 2024-08-28T18:49:49Z
merged: 2024-08-28T18:49:49Z
base_branch: main
head_branch: cancel_update_template
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/777
---

# Pull Request #777: Fixes 4540: cancel template update tasks on delete

**Author**: @jlsherrill
**Created**: August 15, 2024 at 07:58 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `cancel_update_template`

## Description

## Summary

## Testing steps



---

## Discussion

### Comment by @jlsherrill on August 15, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-4540

### Comment by @jlsherrill on August 15, 2024 at 08:31 PM UTC

Other things this PR needs:
* [x]  https://github.com/content-services/content-sources-backend/pull/774 merged
* [x] https://issues.redhat.com/browse/HMS-4478 completed so that this will cancel pending tasks and not just ones already running.  This will make this Pr actually useful 

### Comment by @jlsherrill on August 28, 2024 at 03:00 PM UTC

[test]

### Comment by @jlsherrill on August 28, 2024 at 05:05 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on August 28, 2024 at 06:19 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/777*
