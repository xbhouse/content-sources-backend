---
type: pull_request
number: 384
title: "update to go 1.20 and dependencies"
state: merged
author: jlsherrill
created: 2023-09-08T13:18:00Z
updated: 2023-09-14T17:34:13Z
closed: 2023-09-14T17:33:51Z
merged: 2023-09-14T17:33:51Z
base_branch: main
head_branch: deps_108
labels: []
url: https://github.com/content-services/content-sources-backend/pull/384
---

# Pull Request #384: update to go 1.20 and dependencies

**Author**: @jlsherrill
**Created**: September 08, 2023 at 01:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `deps_108`

## Description

## Summary

Updates to go 1.20
Action to developers:
* Upgrade to go 1.20 (or 1.21) https://go.dev/doc/install
* Delete and redownload your golangci-lint binary:  
 `rm release/golangci-lint; make install-golangci-lint`
* if you are using GoLand, update to the latest ( 2023.2.1)  You may need to update multiple times to get to the latest.

## Testing steps
tests pass

---

## Discussion

### Comment by @jlsherrill on September 08, 2023 at 03:57 PM UTC

/retest

### Comment by @jlsherrill on September 12, 2023 at 05:02 PM UTC

/retest

### Comment by @swadeley on September 13, 2023 at 04:22 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Commented on September 08, 2023 at 03:45 PM UTC

### Review by @jlsherrill - Commented on September 08, 2023 at 03:45 PM UTC

### Review by @rverdile - Approved on September 14, 2023 at 05:32 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/384*
