---
type: pull_request
number: 769
title: "Fixes 4422: edited repo returns incorrect snapshot task"
state: merged
author: xbhouse
created: 2024-08-08T15:54:25Z
updated: 2024-08-09T18:30:24Z
closed: 2024-08-09T18:00:38Z
merged: 2024-08-09T18:00:38Z
base_branch: main
head_branch: 4422
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/769
---

# Pull Request #769: Fixes 4422: edited repo returns incorrect snapshot task

**Author**: @xbhouse
**Created**: August 08, 2024 at 03:54 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4422`

## Description

## Summary

- Adds a clause to a few queries to check only repo configs that have not been soft-deleted

## Testing steps

- Add a repository, let it snapshot, and grab the snapshot task uuid
- Restart the app with pulp disabled to force the delete-repository-snapshots task to fail
- Delete the repository
- Fetch the task associated with that repository, you should get a 404
- Listing tasks filtered on that repository uuid should return an empty list
- Searching environments, package groups, and rpms should work as before
- For QE, this should fix the issue that caused `test_repo_snapshot_functions` to fail, so we should make sure this test passes

---

## Discussion

### Comment by @jlsherrill on August 08, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-4422

### Comment by @swadeley on August 09, 2024 at 09:22 AM UTC

Hi, API version of `test_repo_snapshot_functions` now passes, I'll revert the changes that disabled the test.

### Comment by @swadeley on August 09, 2024 at 09:40 AM UTC

/retest

### Comment by @swadeley on August 09, 2024 at 10:58 AM UTC

/retest

### Comment by @swadeley on August 09, 2024 at 12:11 PM UTC

Hi

```
11:14:34 tests/test_repositories.py::test_repo_snapshot_functions[ViaREST]
<snip>
11:15:30 PASSED 
```

Also tested locally first of course.

---

## Reviews

### Review by @rverdile - Approved on August 09, 2024 at 05:04 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/769*
