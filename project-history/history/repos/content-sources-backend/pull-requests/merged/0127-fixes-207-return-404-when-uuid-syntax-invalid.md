---
type: pull_request
number: 127
title: "Fixes 207: return 404 when uuid syntax invalid"
state: merged
author: rverdile
created: 2022-10-21T15:28:43Z
updated: 2022-10-24T11:30:42Z
closed: 2022-10-24T11:08:09Z
merged: 2022-10-24T11:08:09Z
base_branch: main
head_branch: bad-uuid
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/127
---

# Pull Request #127: Fixes 207: return 404 when uuid syntax invalid

**Author**: @rverdile
**Created**: October 21, 2022 at 03:28 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `bad-uuid`

## Description

To test:
Make a PUT/PATCH/DELETE request where the UUID has invalid syntax i.e. /repostories/bad-uuid

Previously, the API would return 500 error and complain about invalid UUID syntax. Now, it returns a 404.

---

## Discussion

### Comment by @jlsherrill on October 21, 2022 at 03:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-207

---

## Reviews

### Review by @Andrewgdewar - Approved on October 21, 2022 at 04:09 PM UTC

Make it so!

### Review by @swadeley - Commented on October 24, 2022 at 11:08 AM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/127*
