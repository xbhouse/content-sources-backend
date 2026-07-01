---
type: pull_request
number: 885
title: "Fixes 4887: use livez for checking pulp connectivity"
state: merged
author: dominikvagner
created: 2024-11-12T09:40:07Z
updated: 2024-11-13T16:13:26Z
closed: 2024-11-13T16:13:26Z
merged: 2024-11-13T16:13:26Z
base_branch: main
head_branch: 4887
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/885
---

# Pull Request #885: Fixes 4887: use livez for checking pulp connectivity

**Author**: @dominikvagner
**Created**: November 12, 2024 at 09:40 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4887`

## Description

## Summary
This PR changes the pulp connectivity check in the metrics collector to use the `livez` endpoint instead of listing the domains for checking this.

## Testing steps
- Start all the necessary containers and the server, and verify that metrics report the pulp connectivity as '1'.
- Stop the `cs_pulp_api_1` and `cs_pulp_content_1` containers.
- Verify the pulp connectivity metric is now '0'.


---

## Discussion

### Comment by @jlsherrill on November 12, 2024 at 10:00 AM UTC

https://issues.redhat.com/browse/HMS-4887

---

## Reviews

### Review by @xbhouse - Approved on November 13, 2024 at 03:33 PM UTC

working as expected! mock fix looks good too 🎉 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/885*
