---
type: pull_request
number: 69
title: "Fixes 124: Search filter error on URL column"
state: merged
author: rverdile
created: 2022-08-02T21:11:24Z
updated: 2022-08-09T16:15:28Z
closed: 2022-08-09T16:15:28Z
merged: 2022-08-09T16:15:28Z
base_branch: main
head_branch: search-error
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/69
---

# Pull Request #69: Fixes 124: Search filter error on URL column

**Author**: @rverdile
**Created**: August 02, 2022 at 09:11 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `search-error`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on August 02, 2022 at 09:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-124

### Comment by @jlsherrill on August 04, 2022 at 06:17 PM UTC

I noticed that on 'main' if you search with a URL you get no actual error, it just returns zero items.  I think this type of error should throw a 500. Would you mind fixing it as part of this?

### Comment by @rverdile on August 04, 2022 at 08:58 PM UTC

good find. updated.

---

## Reviews

### Review by @jlsherrill - Approved on August 08, 2022 at 05:17 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/69*
