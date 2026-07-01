---
type: pull_request
number: 1225
title: "HMS-9391: remove skipping RBAC for users with admin privileges"
state: merged
author: Starle21
created: 2025-10-02T16:35:34Z
updated: 2025-10-07T14:16:26Z
closed: 2025-10-07T14:16:25Z
merged: 2025-10-07T14:16:25Z
base_branch: main
head_branch: starle/HMS-9391
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1225
---

# Pull Request #1225: HMS-9391: remove skipping RBAC for users with admin privileges

**Author**: @Starle21
**Created**: October 02, 2025 at 04:35 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `starle/HMS-9391`

## Description

## Summary

Previously, if a user was an admin, RBAC was skipped.  
It was added as a work around for something that is no longer needed.  
The commit removes this behavior.

## Testing steps

- Test locally, enable RBAC in config.yaml.
- If you want for a user to get authorized, add their username under mocks.rbac in config.yaml under appropriate role.
- Perform a request to an endpoint.
- A user should not get authorized, even if they are an admin, if their username is not listed in mocks.rbac.

Note that redis is set up to keep identity headers for a certain period of time.
You might want to set it up in config.yaml under redis.expiration.rbac to use a short time frame.
I used a 1s setting (default is 1min).



---

## Discussion

### Comment by @jlsherrill on October 02, 2025 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-9391

### Comment by @katarinazaprazna on October 06, 2025 at 12:37 PM UTC

LGTM, nice work!

### Comment by @rverdile on October 07, 2025 at 02:16 PM UTC

going to merge so we can keep testing kessel

---

## Reviews

### Review by @katarinazaprazna - Commented on October 06, 2025 at 11:40 AM UTC

### Review by @rverdile - Commented on October 06, 2025 at 01:44 PM UTC

### Review by @Starle21 - Commented on October 06, 2025 at 02:46 PM UTC

### Review by @rverdile - Approved on October 06, 2025 at 05:27 PM UTC

looks good! nice job

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1225*
