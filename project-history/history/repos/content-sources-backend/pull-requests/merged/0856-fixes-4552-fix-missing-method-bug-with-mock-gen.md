---
type: pull_request
number: 856
title: "Fixes 4552: fix missing method bug with mock gen"
state: merged
author: dominikvagner
created: 2024-10-21T13:23:15Z
updated: 2024-10-22T15:45:59Z
closed: 2024-10-22T15:45:59Z
merged: 2024-10-22T15:45:59Z
base_branch: main
head_branch: 4552
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/856
---

# Pull Request #856: Fixes 4552: fix missing method bug with mock gen

**Author**: @dominikvagner
**Created**: October 21, 2024 at 01:23 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4552`

## Description

## Summary
This PR addresses the missing method bug (i.e. automates commenting out the `registry_mock.go` file) when generating mocks and updates the way we use the `mockery` tool to a new config as the old way be deprecated.

## Testing steps
- Verify that after adding a new method to a `dao` interface and to its "impl struct", `make mock` runs successfully without errors.
- Also `make mock` should now run without warnings about using an outdated config.


---

## Discussion

### Comment by @jlsherrill on October 21, 2024 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-4552

---

## Reviews

### Review by @xbhouse - Commented on October 21, 2024 at 05:43 PM UTC

### Review by @xbhouse - Commented on October 21, 2024 at 05:46 PM UTC

### Review by @dominikvagner - Commented on October 22, 2024 at 05:35 AM UTC

### Review by @dominikvagner - Commented on October 22, 2024 at 06:30 AM UTC

### Review by @xbhouse - Commented on October 22, 2024 at 01:09 PM UTC

### Review by @xbhouse - Approved on October 22, 2024 at 01:10 PM UTC

lgtm! nice work!! 🚀 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/856*
