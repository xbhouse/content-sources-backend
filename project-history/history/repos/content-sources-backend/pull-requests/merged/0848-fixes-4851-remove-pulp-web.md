---
type: pull_request
number: 848
title: "Fixes 4851: remove pulp-web"
state: merged
author: jlsherrill
created: 2024-10-15T12:21:33Z
updated: 2024-10-16T13:45:30Z
closed: 2024-10-16T13:43:04Z
merged: 2024-10-16T13:43:04Z
base_branch: main
head_branch: compose
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/848
---

# Pull Request #848: Fixes 4851: remove pulp-web

**Author**: @jlsherrill
**Created**: October 15, 2024 at 12:21 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `compose`

## Description

## Summary

Removes pulp-web from docker-compose  as its not used in our real deployments and we no longer need it.

as part of this, the pulp-api is still available on :8080 but pulp-content is now available on :8081
This caused some tests to need updates that were assuming api and content were on the same port


## Testing steps



---

## Discussion

### Comment by @jlsherrill on October 16, 2024 at 01:00 AM UTC

https://issues.redhat.com/browse/HMS-4851

---

## Reviews

### Review by @dominikvagner - Commented on October 16, 2024 at 12:38 PM UTC

looks good! 👍🏼 just one small nit 😅

### Review by @jlsherrill - Commented on October 16, 2024 at 12:42 PM UTC

### Review by @dominikvagner - Approved on October 16, 2024 at 01:15 PM UTC

good job! approved!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/848*
