---
type: pull_request
number: 373
title: "Refs 645: use only runAsUser"
state: closed
author: jlsherrill
created: 2023-08-24T15:52:52Z
updated: 2023-08-24T16:41:21Z
closed: 2023-08-24T16:41:21Z
base_branch: main
head_branch: removeNon
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/373
---

# Pull Request #373: Refs 645: use only runAsUser

**Author**: @jlsherrill
**Created**: August 24, 2023 at 03:52 PM UTC
**Status**: Closed
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `removeNon`

## Description

## Summary
I think that runAsNonroot is not needed (and may conflict)  if we specify runAsUser

## Testing steps


---

## Discussion

### Comment by @jlsherrill on August 24, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-645

### Comment by @jlsherrill on August 24, 2023 at 04:41 PM UTC

closing in favor of https://github.com/content-services/content-sources-backend/pull/374

---

## Reviews

### Review by @mshriver - Approved on August 24, 2023 at 04:31 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/373*
