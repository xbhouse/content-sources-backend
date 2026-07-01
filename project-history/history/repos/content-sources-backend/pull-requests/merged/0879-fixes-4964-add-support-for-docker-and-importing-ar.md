---
type: pull_request
number: 879
title: "Fixes 4964: add support for docker and importing arm repos"
state: merged
author: xbhouse
created: 2024-11-07T20:15:14Z
updated: 2024-11-14T09:45:21Z
closed: 2024-11-14T09:45:21Z
merged: 2024-11-14T09:45:21Z
base_branch: main
head_branch: 4964
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/879
---

# Pull Request #879: Fixes 4964: add support for docker and importing arm repos

**Author**: @xbhouse
**Created**: November 07, 2024 at 08:15 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4964`

## Description

## Summary

- Adjusts the make command for container volume removal to add support for docker
- Adds a selector to RHEL9 arm repos so it's easier to import those repos

## Testing steps

1. Using docker, `make compose-clean` should clear out all unused volumes
2. Using podman, `make compose-clean` should still work as before
3. Running `OPTIONS_REPOSITORY_IMPORT_FILTER=arm go run cmd/external-repos/main.go import && go run cmd/external-repos/main.go nightly-jobs` should import and snapshot RHEL9 arm repos


---

## Discussion

### Comment by @jlsherrill on November 07, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-4964

### Comment by @swadeley on November 13, 2024 at 09:31 AM UTC

/retest

---

## Reviews

### Review by @dominikvagner - Changes Requested on November 12, 2024 at 11:41 AM UTC

### Review by @xbhouse - Commented on November 12, 2024 at 04:24 PM UTC

### Review by @dominikvagner - Commented on November 13, 2024 at 08:01 AM UTC

### Review by @dominikvagner - Approved on November 13, 2024 at 08:03 AM UTC

everything works now, both the compose and the import of rhel 9 arm repos! 👌🏼 
approved, good job! 👏🏼😄   

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/879*
