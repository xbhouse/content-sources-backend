---
type: pull_request
number: 159
title: "Refs 276: Fixes Tag on popular_repositories list api"
state: merged
author: Andrewgdewar
created: 2022-12-20T13:58:28Z
updated: 2022-12-20T15:06:23Z
closed: 2022-12-20T15:06:15Z
merged: 2022-12-20T15:06:15Z
base_branch: main
head_branch: popularReposAreNowPopularRepos
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/159
---

# Pull Request #159: Refs 276: Fixes Tag on popular_repositories list api

**Author**: @Andrewgdewar
**Created**: December 20, 2022 at 01:58 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `popularReposAreNowPopularRepos`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on December 20, 2022 at 02:13 PM UTC

will let mike merge :)

### Comment by @jlsherrill on December 20, 2022 at 02:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-276

### Comment by @jlsherrill on December 20, 2022 at 02:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @mshriver on December 20, 2022 at 02:58 PM UTC

/retest

### Comment by @jlsherrill on December 20, 2022 at 03:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-309

### Comment by @mshriver on December 20, 2022 at 03:06 PM UTC

Thanks @Andrewgdewar !

---

## Reviews

### Review by @jlsherrill - Approved on December 20, 2022 at 02:12 PM UTC

### Review by @mshriver - Approved on December 20, 2022 at 02:38 PM UTC

Generated a new client from this and saw the `list_popular_repositories()` only available for popular_repositories, as intended.  The GET continues to work fine. 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/159*
