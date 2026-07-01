---
type: pull_request
number: 629
title: "Fixes 3937: use clowder rds ca for tangy db"
state: merged
author: jlsherrill
created: 2024-04-10T16:57:14Z
updated: 2024-04-10T17:14:39Z
closed: 2024-04-10T17:14:39Z
merged: 2024-04-10T17:14:38Z
base_branch: main
head_branch: 3937
labels: []
url: https://github.com/content-services/content-sources-backend/pull/629
---

# Pull Request #629: Fixes 3937: use clowder rds ca for tangy db

**Author**: @jlsherrill
**Created**: April 10, 2024 at 04:57 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `3937`

## Description

## Summary

Stage/prod dbs need to use ssl, which wasn't being done because we don't have a ca cert.  Lets try using the one clowder is giving us for our db.

## Testing steps

Don't think there is an easy way  to test this


---

## Discussion

### Comment by @jlsherrill on April 10, 2024 at 05:05 PM UTC

yes, and this shouldn't affect integration tests at all

### Comment by @jlsherrill on April 10, 2024 at 05:14 PM UTC

merging

---

## Reviews

### Review by @Andrewgdewar - Approved on April 10, 2024 at 05:01 PM UTC

I take it the pulp failure is expected?

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/629*
