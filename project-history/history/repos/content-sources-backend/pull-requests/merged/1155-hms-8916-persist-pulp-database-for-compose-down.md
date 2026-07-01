---
type: pull_request
number: 1155
title: "HMS-8916: Persist pulp database for compose-down"
state: merged
author: rverdile
created: 2025-07-24T18:03:14Z
updated: 2025-07-25T14:37:23Z
closed: 2025-07-24T20:03:57Z
merged: 2025-07-24T20:03:57Z
base_branch: main
head_branch: pulp-db-persistence
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1155
---

# Pull Request #1155: HMS-8916: Persist pulp database for compose-down

**Author**: @rverdile
**Created**: July 24, 2025 at 06:03 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `pulp-db-persistence`

## Description

## Summary
The pulp database was not persisting between compose-down and compose-up

## Testing steps
1. `make compose-clean compose-up`
2. Create a repository with snapshot
3. List the repositories in pulp. e.g. `get http://localhost:8080/api/pulp/cs-62edc243ec/api/v3/repositories/?offset=0&limit=25`
4. The repository will be listed
5. `make compose-down compose-up`
6. Again, list the repositories in pulp
7. The repository should still be listed. Without this PR, no repositories would be found


---

## Discussion

### Comment by @jlsherrill on July 24, 2025 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-8916

---

## Reviews

### Review by @xbhouse - Approved on July 24, 2025 at 07:42 PM UTC

lgtm! thank you! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1155*
