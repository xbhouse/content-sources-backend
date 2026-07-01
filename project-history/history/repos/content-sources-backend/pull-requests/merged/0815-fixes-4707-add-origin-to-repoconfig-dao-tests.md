---
type: pull_request
number: 815
title: "Fixes 4707: add origin to RepoConfig dao tests"
state: merged
author: rverdile
created: 2024-09-13T17:17:05Z
updated: 2024-09-16T13:46:49Z
closed: 2024-09-16T08:19:20Z
merged: 2024-09-16T08:19:20Z
base_branch: main
head_branch: repo-config-tests
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/815
---

# Pull Request #815: Fixes 4707: add origin to RepoConfig dao tests

**Author**: @rverdile
**Created**: September 13, 2024 at 05:17 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `repo-config-tests`

## Description

## Summary
adds origin filter to all the RepoConfig listing dao tests and fixes bug that prevented "SeedPreexistingRedHatRepos()" method from running for the RepoConfig dao suite

## Testing steps
1. Import red hat repos locally
2. Run unit tests
3. Without the PR, the tests will fail



---

## Discussion

### Comment by @jlsherrill on September 13, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-4707

---

## Reviews

### Review by @xbhouse - Approved on September 13, 2024 at 05:38 PM UTC

looks good and tests pass! 🚀 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/815*
