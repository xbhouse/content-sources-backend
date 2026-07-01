---
type: pull_request
number: 29
title: "CONTENT-56: Add list/delete dao layer and honor org id"
state: merged
author: rverdile
created: 2022-06-10T15:36:38Z
updated: 2022-06-14T19:54:41Z
closed: 2022-06-14T19:54:30Z
merged: 2022-06-14T19:54:30Z
base_branch: main
head_branch: list-delete-dao
labels: []
url: https://github.com/content-services/content-sources-backend/pull/29
---

# Pull Request #29: CONTENT-56: Add list/delete dao layer and honor org id

**Author**: @rverdile
**Created**: June 10, 2022 at 03:36 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `list-delete-dao`

## Description

#### Adds
- Delete and List actions to the DAO layer.
- Additional error checking to Fetch/Update in the DAO layer.
- Rewrites the repository handler tests to use serveHTTP instead of context (because it feels less cumbersome)
- Org ID filters to all the queries that were missing them.
- Tests for verifying Org ID is correct when returning results from database and API
- IdentityHeader constant to api/api.go
- Rename collectionResponse -> setCollectionResponseMetadata

The TestList pagination tests in the handler tests are a bit tricky now because of the mocking. I'm not 100% sure they're still testing what they need to test.

---

## Discussion

### Comment by @jlsherrill on June 10, 2022 at 04:34 PM UTC

couple of comments, and failing lint, otherwise looks 🚀 

### Comment by @jlsherrill on June 13, 2022 at 05:35 PM UTC

This needs a rebase

### Comment by @rverdile on June 14, 2022 at 07:49 PM UTC

it's green

### Comment by @jlsherrill on June 14, 2022 at 07:54 PM UTC

went ahead and merged so @Andrewgdewar can rebase

---

## Reviews

### Review by @jlsherrill - Commented on June 10, 2022 at 04:13 PM UTC

### Review by @jlsherrill - Commented on June 10, 2022 at 04:17 PM UTC

### Review by @jlsherrill - Commented on June 10, 2022 at 04:34 PM UTC

### Review by @rverdile - Commented on June 10, 2022 at 04:57 PM UTC

### Review by @rverdile - Commented on June 10, 2022 at 04:58 PM UTC

### Review by @jlsherrill - Commented on June 10, 2022 at 05:24 PM UTC

### Review by @rverdile - Commented on June 10, 2022 at 06:41 PM UTC

### Review by @jlsherrill - Approved on June 14, 2022 at 07:54 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/29*
