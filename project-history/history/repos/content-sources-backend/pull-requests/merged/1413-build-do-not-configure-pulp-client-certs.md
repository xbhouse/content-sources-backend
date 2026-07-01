---
type: pull_request
number: 1413
title: "Build: do not configure pulp client certs"
state: merged
author: jlsherrill
created: 2026-03-16T21:51:59Z
updated: 2026-03-17T08:47:23Z
closed: 2026-03-17T08:47:23Z
merged: 2026-03-17T08:47:23Z
base_branch: main
head_branch: useClientCerts
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1413
---

# Pull Request #1413: Build: do not configure pulp client certs

**Author**: @jlsherrill
**Created**: March 16, 2026 at 09:51 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `useClientCerts`

## Description



## Summary

do not configure pulp client certs if basic auth is provided. 

Currently if both are provided, the client is configured to use both, which results in both being sent.  Apparently the gateway cannot handle this and seems to reject some requests while allowing some others?

## Testing steps



---

## Discussion

### Comment by @jlsherrill on March 16, 2026 at 09:59 PM UTC

alternatively we could have removed the pulp certs from deployment.yaml, but that would have meant we cannot control this via configuration, which i'd rather be able to do to be able to revert quickly.

---

## Reviews

### Review by @dominikvagner - Approved on March 17, 2026 at 08:46 AM UTC

ack! ✅ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1413*
