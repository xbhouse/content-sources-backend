---
type: pull_request
number: 555
title: "Fixes 3346: new pattern for pulp response errors"
state: merged
author: rverdile
created: 2024-01-31T22:20:34Z
updated: 2024-02-05T21:28:58Z
closed: 2024-02-05T21:28:51Z
merged: 2024-02-05T21:28:51Z
base_branch: main
head_branch: 3346
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/555
---

# Pull Request #555: Fixes 3346: new pattern for pulp response errors

**Author**: @rverdile
**Created**: January 31, 2024 at 10:20 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3346`

## Description

## Summary
Establishes new pattern for handling pulp api response errors to ensure nil check and include the response body in the error. 

## Testing steps
1. Cause a pulp error somehow. One way is to do `podman stop cs_pulp_api_1`, then try to sync a repository.
2. You should see the error message in the logs includes the body response.
3. Also look through the pulp client package and see if I missed adding this response pattern to the responses.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 31, 2024 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-3346

---

## Reviews

### Review by @jlsherrill - Approved on February 05, 2024 at 03:52 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/555*
