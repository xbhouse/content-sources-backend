---
type: pull_request
number: 720
title: "Build: restart pulp migration on error"
state: merged
author: jlsherrill
created: 2024-06-25T20:46:32Z
updated: 2024-06-27T13:29:27Z
closed: 2024-06-27T13:29:24Z
merged: 2024-06-27T13:29:24Z
base_branch: main
head_branch: pulp_fix
labels: []
url: https://github.com/content-services/content-sources-backend/pull/720
---

# Pull Request #720: Build: restart pulp migration on error

**Author**: @jlsherrill
**Created**: June 25, 2024 at 08:46 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `pulp_fix`

## Description

## Summary

There seems to be a race condition where podman thinks postgres is 'up' but its not.  So restart the pulp migration if it errors.

## Testing steps

```
make compose-clean compose-up
```

try to create a repo for snapshotting

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @swadeley on June 26, 2024 at 08:44 AM UTC

/retest

### Comment by @swadeley on June 27, 2024 at 08:32 AM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on June 26, 2024 at 06:29 PM UTC

I can `make compose-clean compose-up` again!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/720*
