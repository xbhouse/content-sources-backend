---
type: pull_request
number: 674
title: "Refs 4160: update label constraint migration to change label"
state: merged
author: rverdile
created: 2024-05-20T19:07:43Z
updated: 2024-05-20T20:17:37Z
closed: 2024-05-20T20:17:28Z
merged: 2024-05-20T20:17:28Z
base_branch: main
head_branch: fix-migration
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/674
---

# Pull Request #674: Refs 4160: update label constraint migration to change label

**Author**: @rverdile
**Created**: May 20, 2024 at 07:07 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `fix-migration`

## Description

## Summary
This label was duplicated before being fixed in #671 

The label needs to be updated as part of the migration, or the migration will fail.

## Testing steps
1. Assuming your local DB has the latest migration (label uniqueness constraint), run `go run cmd/dbmigrate/main.go down --steps 1`
2. Update the redhat_repos.json to revert the change made in #671.
3. Run `make repos-import`
4. Run `go run cmd/dbmigration/main.go up` to run the uniqueness constraint migration.
5. The migration should not fail, and the label for this repository in repository_configurations should be correct.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 20, 2024 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-4160

### Comment by @jlsherrill on May 20, 2024 at 07:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

---

## Reviews

### Review by @xbhouse - Approved on May 20, 2024 at 07:14 PM UTC

tested and works well :) thank you!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/674*
